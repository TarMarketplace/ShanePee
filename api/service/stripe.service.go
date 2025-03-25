package service

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"shanepee.com/api/config"
	"shanepee.com/api/domain"
)

type StripeService interface {
	Checkout(ctx context.Context, userID int64) (url string, err error)
	PaymentSuccessCallback(ctx context.Context, body []byte, signature string) error
}

type stripeServiceImpl struct {
	orderRepo                       domain.OrderRepository
	reviewRepo                      domain.ReviewRepository
	userRepo                        domain.UserRepository
	cartRepo                        domain.CartRepository
	cartSvc                         CartService
	stripeKey                       string
	stripePaymentSuccessRedirectURL string
}

func NewStripeService(orderRepo domain.OrderRepository, reviewRepo domain.ReviewRepository, userRepo domain.UserRepository, cartRepo domain.CartRepository, cartSvc CartService, cfg config.Config) StripeService {
	return &stripeServiceImpl{
		orderRepo:                       orderRepo,
		reviewRepo:                      reviewRepo,
		userRepo:                        userRepo,
		cartRepo:                        cartRepo,
		cartSvc:                         cartSvc,
		stripeKey:                       cfg.StripeKey,
		stripePaymentSuccessRedirectURL: cfg.StripePaymentSuccessRedirectURL,
	}
}

var _ StripeService = &stripeServiceImpl{}

func (s *stripeServiceImpl) Checkout(ctx context.Context, userID int64) (string, error) {
	_, err := s.userRepo.FindUserByID(ctx, userID)

	if err != nil {
		return "", err
	}

	cartItems, err := s.cartRepo.GetCartWithItemByOwnerID(ctx, userID)
	if err != nil {
		return "", err
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams

	for _, cartItem := range cartItems {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				Currency: stripe.String("thb"),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name: stripe.String(cartItem.ArtToy.Name),
				},
				UnitAmount: stripe.Int64(int64(cartItem.ArtToy.Price * 100)),
			},
			Quantity: stripe.Int64(1),
		})
	}

	stripe.Key = s.stripeKey
	params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(s.stripePaymentSuccessRedirectURL),
		LineItems:  lineItems,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
	}
	params.AddMetadata("user_id", strconv.Itoa(int(userID)))
	result, err := session.New(params)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

func (s *stripeServiceImpl) PaymentSuccessCallback(ctx context.Context, body []byte, signature string) error {
	// Pass the request body & Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// You can find your endpoint's secret in your webhook settings
	// event, _ := webhook.ConstructEvent(body, signature, endpointSecret)

	var event stripe.Event
	err := json.Unmarshal(body, &event)
	if err != nil {
		return errors.New("Unable to decode JSON")
	}

	// if err != nil {
	// 	logrus.Info(err)
	// 	return errors.New("Invalid signature")
	// }

	if event.Type == "checkout.session.completed" {
		var checkoutSession stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &checkoutSession)
		if err != nil {
			logrus.Warnf("Error parsing checkout session: %v", err)
			return nil
		}

		// Access metadata directly from the Checkout Session
		userID := checkoutSession.Metadata["user_id"]
		logrus.Printf("Checkout session completed for user: %s", userID)

		logrus.Printf("Payment status: %s", checkoutSession.PaymentStatus)
		logrus.Printf("Customer email: %s", checkoutSession.CustomerDetails.Email)

		v, err := strconv.Atoi(userID)
		if err != nil {
			logrus.Warnf("Error parsing user ID: %v", err)
			return nil
		}
		err = s.cartSvc.Checkout(ctx, int64(v))
		return err
	}

	logrus.Warn("Invalid event type")
	return nil
}
