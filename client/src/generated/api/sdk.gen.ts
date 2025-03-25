// This file is auto-generated by @hey-api/openapi-ts

import type { Options as ClientOptions, TDataShape, Client } from '@hey-api/client-next';
import type { GetArtToysData, GetArtToysResponse, GetArtToysError, CreateArtToyData, CreateArtToyResponse, CreateArtToyError, DeleteReviewData, DeleteReviewResponse, DeleteReviewError, GetReviewData, GetReviewResponse, GetReviewError, UpdateReviewData, UpdateReviewResponse, UpdateReviewError, CreateReviewData, CreateReviewResponse, CreateReviewError, SearchArtToysData, SearchArtToysResponse, SearchArtToysError, DeleteArtToyData, DeleteArtToyResponse, DeleteArtToyError, GetArtToyByIdData, GetArtToyByIdResponse, GetArtToyByIdError, UpdateArtToyData, UpdateArtToyResponse, UpdateArtToyError, ChangePasswordData, ChangePasswordResponse, ChangePasswordError, LoginData, LoginResponse, LoginError, LogoutData, LogoutResponse, LogoutError, MeData, MeResponse, MeError, PasswordResetRequestsData, PasswordResetRequestsResponse, PasswordResetRequestsError, RegisterData, RegisterResponse, RegisterError, ResetPasswordData, ResetPasswordResponse, ResetPasswordError, CompleteOrderData, CompleteOrderResponse, CompleteOrderError, GetOrderOfBuyerData, GetOrderOfBuyerResponse, GetOrderOfBuyerError, GetOrdersOfBuyerData, GetOrdersOfBuyerResponse, GetOrdersOfBuyerError, GetCartData, GetCartResponse, GetCartError, AddItemToCartData, AddItemToCartResponse, AddItemToCartError, CheckoutData, CheckoutResponse, CheckoutError, ClearItemsFromCartData, ClearItemsFromCartResponse, ClearItemsFromCartError, PaymentSuccessCallbackData, PaymentSuccessCallbackResponse, PaymentSuccessCallbackError, RemoveItemFromCartData, RemoveItemFromCartResponse, RemoveItemFromCartError, GetMyArtToysData, GetMyArtToysResponse, GetMyArtToysError, GetOrdersByStatusData, GetOrdersByStatusResponse, GetOrdersByStatusError, UpdateOrderData, UpdateOrderResponse, UpdateOrderError, GetOrderOfSellerData, GetOrderOfSellerResponse, GetOrderOfSellerError, GetOrdersOfSellerData, GetOrdersOfSellerResponse, GetOrdersOfSellerError, GetArtToysOfSellerData, GetArtToysOfSellerResponse, GetArtToysOfSellerError, UpdateUserData, UpdateUserResponse, UpdateUserError } from './types.gen';
import { client as _heyApiClient } from './client.gen';

export type Options<TData extends TDataShape = TDataShape, ThrowOnError extends boolean = boolean> = ClientOptions<TData, ThrowOnError> & {
    /**
     * You can provide a client instance returned by `createClient()` instead of
     * individual options. This might be also useful if you want to implement a
     * custom client.
     */
    client?: Client;
    /**
     * You can pass arbitrary values through the `meta` object. This can be
     * used to access values that aren't defined as part of the SDK function.
     */
    meta?: Record<string, unknown>;
};

/**
 * Get Art Toys
 * Get art toys
 */
export const getArtToys = <ThrowOnError extends boolean = false>(options?: Options<GetArtToysData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetArtToysResponse, GetArtToysError, ThrowOnError>({
        url: '/v1/art-toy',
        ...options
    });
};

/**
 * Create Art toy
 * Create a new art toy record
 */
export const createArtToy = <ThrowOnError extends boolean = false>(options: Options<CreateArtToyData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateArtToyResponse, CreateArtToyError, ThrowOnError>({
        url: '/v1/art-toy',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Delete Art Toy Review
 * Delete an existing art toy review by ID
 */
export const deleteReview = <ThrowOnError extends boolean = false>(options: Options<DeleteReviewData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<DeleteReviewResponse, DeleteReviewError, ThrowOnError>({
        url: '/v1/art-toy/review/{artToyID}',
        ...options
    });
};

/**
 * Get Art Toy Review
 * Get art toy review by art toy ID
 */
export const getReview = <ThrowOnError extends boolean = false>(options: Options<GetReviewData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetReviewResponse, GetReviewError, ThrowOnError>({
        url: '/v1/art-toy/review/{artToyID}',
        ...options
    });
};

/**
 * Update Art Toy Review
 * Update an existing art toy review by ID
 */
export const updateReview = <ThrowOnError extends boolean = false>(options: Options<UpdateReviewData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateReviewResponse, UpdateReviewError, ThrowOnError>({
        url: '/v1/art-toy/review/{artToyID}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Create Art Toy Review
 * Create a new art toy review record
 */
export const createReview = <ThrowOnError extends boolean = false>(options: Options<CreateReviewData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<CreateReviewResponse, CreateReviewError, ThrowOnError>({
        url: '/v1/art-toy/review/{artToyID}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Search Art Toys
 * Search art toys
 */
export const searchArtToys = <ThrowOnError extends boolean = false>(options?: Options<SearchArtToysData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<SearchArtToysResponse, SearchArtToysError, ThrowOnError>({
        url: '/v1/art-toy/search',
        ...options
    });
};

/**
 * Delete Art Toy
 * Delete an art toy by ID
 */
export const deleteArtToy = <ThrowOnError extends boolean = false>(options: Options<DeleteArtToyData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<DeleteArtToyResponse, DeleteArtToyError, ThrowOnError>({
        url: '/v1/art-toy/{id}',
        ...options
    });
};

/**
 * Get Art Toy by ID
 * Get art toy by id
 */
export const getArtToyById = <ThrowOnError extends boolean = false>(options: Options<GetArtToyByIdData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetArtToyByIdResponse, GetArtToyByIdError, ThrowOnError>({
        url: '/v1/art-toy/{id}',
        ...options
    });
};

/**
 * Update Art Toy
 * Update an existing art toy by ID
 */
export const updateArtToy = <ThrowOnError extends boolean = false>(options: Options<UpdateArtToyData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateArtToyResponse, UpdateArtToyError, ThrowOnError>({
        url: '/v1/art-toy/{id}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Change password
 * Change password for authenticated user
 */
export const changePassword = <ThrowOnError extends boolean = false>(options: Options<ChangePasswordData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<ChangePasswordResponse, ChangePasswordError, ThrowOnError>({
        url: '/v1/auth/change-password',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Login User
 * Login
 */
export const login = <ThrowOnError extends boolean = false>(options: Options<LoginData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<LoginResponse, LoginError, ThrowOnError>({
        url: '/v1/auth/login',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Logout User
 * Logout
 */
export const logout = <ThrowOnError extends boolean = false>(options?: Options<LogoutData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).post<LogoutResponse, LogoutError, ThrowOnError>({
        url: '/v1/auth/logout',
        ...options
    });
};

/**
 * Get current authenticated user
 * Get authenticated user from the session
 */
export const me = <ThrowOnError extends boolean = false>(options?: Options<MeData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<MeResponse, MeError, ThrowOnError>({
        url: '/v1/auth/me',
        ...options
    });
};

/**
 * Request a password reset
 * Initiates a password reset process by sending an email with reset instructions
 */
export const passwordResetRequests = <ThrowOnError extends boolean = false>(options: Options<PasswordResetRequestsData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<PasswordResetRequestsResponse, PasswordResetRequestsError, ThrowOnError>({
        url: '/v1/auth/password-reset-requests',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Register User
 * Register
 */
export const register = <ThrowOnError extends boolean = false>(options: Options<RegisterData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<RegisterResponse, RegisterError, ThrowOnError>({
        url: '/v1/auth/register',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Reset password
 * Reset password of a user using token and request id
 */
export const resetPassword = <ThrowOnError extends boolean = false>(options: Options<ResetPasswordData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<ResetPasswordResponse, ResetPasswordError, ThrowOnError>({
        url: '/v1/auth/reset-password',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Complete Order by Buyer
 * Update status to completed of an order by buyer
 */
export const completeOrder = <ThrowOnError extends boolean = false>(options: Options<CompleteOrderData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<CompleteOrderResponse, CompleteOrderError, ThrowOnError>({
        url: '/v1/buyer/order/{id}/complete',
        ...options
    });
};

/**
 * Get order detail of buyer
 * Get order detail of buyer
 */
export const getOrderOfBuyer = <ThrowOnError extends boolean = false>(options: Options<GetOrderOfBuyerData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetOrderOfBuyerResponse, GetOrderOfBuyerError, ThrowOnError>({
        url: '/v1/buyer/order/{orderID}',
        ...options
    });
};

/**
 * Get orders of buyer
 * Get orders of buyer
 */
export const getOrdersOfBuyer = <ThrowOnError extends boolean = false>(options?: Options<GetOrdersOfBuyerData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetOrdersOfBuyerResponse, GetOrdersOfBuyerError, ThrowOnError>({
        url: '/v1/buyer/orders',
        ...options
    });
};

/**
 * Get Cart
 * Retrieve the user's cart
 */
export const getCart = <ThrowOnError extends boolean = false>(options?: Options<GetCartData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetCartResponse, GetCartError, ThrowOnError>({
        url: '/v1/cart',
        ...options
    });
};

/**
 * Add Item To Cart
 * Add an item to the cart
 */
export const addItemToCart = <ThrowOnError extends boolean = false>(options: Options<AddItemToCartData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<AddItemToCartResponse, AddItemToCartError, ThrowOnError>({
        url: '/v1/cart/add-item',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Checkout Items In Cart
 * Place a new order from items in the cart
 */
export const checkout = <ThrowOnError extends boolean = false>(options?: Options<CheckoutData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).post<CheckoutResponse, CheckoutError, ThrowOnError>({
        url: '/v1/cart/checkout',
        ...options
    });
};

/**
 * Clear Items From Cart
 * Clear items from the cart
 */
export const clearItemsFromCart = <ThrowOnError extends boolean = false>(options?: Options<ClearItemsFromCartData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).delete<ClearItemsFromCartResponse, ClearItemsFromCartError, ThrowOnError>({
        url: '/v1/cart/clear-items',
        ...options
    });
};

/**
 * Callback after stripe payment success
 * Callback after stripe payment success
 */
export const paymentSuccessCallback = <ThrowOnError extends boolean = false>(options: Options<PaymentSuccessCallbackData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).post<PaymentSuccessCallbackResponse, PaymentSuccessCallbackError, ThrowOnError>({
        url: '/v1/cart/payment-success-callback',
        ...options,
        headers: {
            'Content-Type': 'application/octet-stream',
            ...options?.headers
        }
    });
};

/**
 * Remove Item From Cart
 * Remove an item from the cart
 */
export const removeItemFromCart = <ThrowOnError extends boolean = false>(options: Options<RemoveItemFromCartData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).delete<RemoveItemFromCartResponse, RemoveItemFromCartError, ThrowOnError>({
        url: '/v1/cart/remove-item/{id}',
        ...options
    });
};

/**
 * Get My Art Toys
 * Get my art toys
 */
export const getMyArtToys = <ThrowOnError extends boolean = false>(options?: Options<GetMyArtToysData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetMyArtToysResponse, GetMyArtToysError, ThrowOnError>({
        url: '/v1/my-art-toy',
        ...options
    });
};

/**
 * Get Orders by Status
 * Get orders by status
 */
export const getOrdersByStatus = <ThrowOnError extends boolean = false>(options: Options<GetOrdersByStatusData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetOrdersByStatusResponse, GetOrdersByStatusError, ThrowOnError>({
        url: '/v1/order/{status}',
        ...options
    });
};

/**
 * Get Art Toy Reviews of seller
 * Get art toy reviews of seller
 */
export const getSellerReviews = <ThrowOnError extends boolean = false>(options: Options<GetReviewData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetReviewResponse, GetReviewError, ThrowOnError>({
        url: '/v1/seller/art-toy/review/{sellerID}',
        ...options
    });
};

/**
 * Update Order by Seller
 * Update tracking number, delivery service, status of an order by seller
 */
export const updateOrder = <ThrowOnError extends boolean = false>(options: Options<UpdateOrderData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateOrderResponse, UpdateOrderError, ThrowOnError>({
        url: '/v1/seller/order/{id}',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};

/**
 * Get order detail of seller
 * Get order detail of seller
 */
export const getOrderOfSeller = <ThrowOnError extends boolean = false>(options: Options<GetOrderOfSellerData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetOrderOfSellerResponse, GetOrderOfSellerError, ThrowOnError>({
        url: '/v1/seller/order/{orderID}',
        ...options
    });
};

/**
 * Get orders of seller
 * Get orders of seller
 */
export const getOrdersOfSeller = <ThrowOnError extends boolean = false>(options?: Options<GetOrdersOfSellerData, ThrowOnError>) => {
    return (options?.client ?? _heyApiClient).get<GetOrdersOfSellerResponse, GetOrdersOfSellerError, ThrowOnError>({
        url: '/v1/seller/orders',
        ...options
    });
};

/**
 * Get Art Toys Of Seller
 * Get art toys of seller
 */
export const getArtToysOfSeller = <ThrowOnError extends boolean = false>(options: Options<GetArtToysOfSellerData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).get<GetArtToysOfSellerResponse, GetArtToysOfSellerError, ThrowOnError>({
        url: '/v1/seller/{id}/art-toy',
        ...options
    });
};

/**
 * Update User
 * Update user by id
 */
export const updateUser = <ThrowOnError extends boolean = false>(options: Options<UpdateUserData, ThrowOnError>) => {
    return (options.client ?? _heyApiClient).patch<UpdateUserResponse, UpdateUserError, ThrowOnError>({
        url: '/v1/user',
        ...options,
        headers: {
            'Content-Type': 'application/json',
            ...options?.headers
        }
    });
};
