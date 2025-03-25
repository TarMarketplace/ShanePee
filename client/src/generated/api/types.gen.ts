// This file is auto-generated by @hey-api/openapi-ts

export type Address = {
    district?: string;
    house_no?: string;
    postcode?: string;
    province?: string;
};

export type ArrayResponseArtToy = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    data: Array<ArtToy> | null;
};

export type ArrayResponseCartItem = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    data: Array<CartItem> | null;
};

export type ArrayResponseOrder = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    data: Array<Order> | null;
};

export type ArrayResponseReview = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    data: Array<Review> | null;
};

export type ArtToy = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    availability: boolean;
    average_rating: number;
    description: string;
    id: number;
    name: string;
    owner_id: number;
    photo?: string;
    price: number;
    release_date: string;
};

export type ArtToyCreateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    description: string;
    name: string;
    photo?: string | null;
    price: number;
};

export type ArtToyUpdateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    availability?: boolean;
    description?: string;
    name?: string;
    photo?: string;
    price?: number;
};

export type CartItem = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    art_toy?: ArtToy;
    art_toy_id: number;
    id: number;
    owner?: User;
    owner_id: number;
};

export type CartItemCreateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    art_toy_id: number;
};

export type ChangePasswordBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    new_password: string;
    old_password: string;
};

export type CheckoutOutputBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    url: string;
};

export type ErrorDetail = {
    /**
     * Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id'
     */
    location?: string;
    /**
     * Error message text
     */
    message?: string;
    /**
     * The value at the given location
     */
    value?: unknown;
};

export type ErrorModel = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    /**
     * A human-readable explanation specific to this occurrence of the problem.
     */
    detail?: string;
    /**
     * Optional list of individual error details
     */
    errors?: Array<ErrorDetail> | null;
    /**
     * A URI reference that identifies the specific occurrence of the problem.
     */
    instance?: string;
    /**
     * HTTP status code
     */
    status?: number;
    /**
     * A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
     */
    title?: string;
    /**
     * A URI reference to human-readable documentation for the error.
     */
    type?: string;
};

export type LoginBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    email: string;
    password: string;
};

export type Order = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    buyer_id: number;
    created_at: string;
    delivery_service?: string;
    id: number;
    order_items?: Array<OrderItem> | null;
    seller_id: number;
    status: 'PREPARING' | 'DELIVERING' | 'COMPLETED';
    tracking_number?: string;
};

export type OrderItem = {
    art_toy?: ArtToy;
    art_toy_id: number;
    id: number;
    order_id: number;
};

export type OrderUpdateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    delivery_service?: string;
    status?: 'PREPARING' | 'DELIVERING' | 'COMPLETED';
    tracking_number?: string;
};

export type PartialAddress = {
    district?: string;
    house_no?: string;
    postcode?: string;
    province?: string;
};

export type PartialPaymentMethod = {
    card_number?: string;
    card_owner?: string;
    cvv?: string;
    expire_date?: string;
};

export type PaymentMethod = {
    card_number?: string;
    card_owner?: string;
    cvv?: string;
    expire_date?: string;
};

export type RegisterBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    email: string;
    password: string;
};

export type RequestPasswordResetBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    email: string;
};

export type ResetPasswordBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    new_password: string;
    request_id: number;
    token: string;
};

export type Review = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    art_toy_id: number;
    comment: string;
    id: number;
    rating: number;
};

export type ReviewCreateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    comment: string;
    rating: number;
};

export type ReviewUpdateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    comment?: string;
    rating?: number;
};

export type User = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    address: Address;
    email: string;
    first_name?: string;
    gender?: string;
    id: number;
    last_name?: string;
    payment_method: PaymentMethod;
    photo?: string;
    tel?: string;
};

export type UserUpdateBody = {
    /**
     * A URL to the JSON Schema for this object.
     */
    readonly $schema?: string;
    address?: PartialAddress;
    first_name?: string;
    gender?: 'MALE' | 'FEMALE' | 'OTHER';
    last_name?: string;
    payment_method?: PartialPaymentMethod;
    photo?: string;
    tel?: string;
};

export type GetArtToysData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/art-toy';
};

export type GetArtToysErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetArtToysError = GetArtToysErrors[keyof GetArtToysErrors];

export type GetArtToysResponses = {
    /**
     * OK
     */
    200: ArrayResponseArtToy;
};

export type GetArtToysResponse = GetArtToysResponses[keyof GetArtToysResponses];

export type CreateArtToyData = {
    body: ArtToyCreateBody;
    path?: never;
    query?: never;
    url: '/v1/art-toy';
};

export type CreateArtToyErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type CreateArtToyError = CreateArtToyErrors[keyof CreateArtToyErrors];

export type CreateArtToyResponses = {
    /**
     * OK
     */
    200: ArtToy;
};

export type CreateArtToyResponse = CreateArtToyResponses[keyof CreateArtToyResponses];

export type DeleteReviewData = {
    body?: never;
    path: {
        artToyID: number;
    };
    query?: never;
    url: '/v1/art-toy/review/{artToyID}';
};

export type DeleteReviewErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type DeleteReviewError = DeleteReviewErrors[keyof DeleteReviewErrors];

export type DeleteReviewResponses = {
    /**
     * No Content
     */
    204: void;
};

export type DeleteReviewResponse = DeleteReviewResponses[keyof DeleteReviewResponses];

export type GetReviewData = {
    body?: never;
    path: {
        artToyID: number;
    };
    query?: never;
    url: '/v1/art-toy/review/{artToyID}';
};

export type GetReviewErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetReviewError = GetReviewErrors[keyof GetReviewErrors];

export type GetReviewResponses = {
    /**
     * OK
     */
    200: Review;
};

export type GetReviewResponse = GetReviewResponses[keyof GetReviewResponses];

export type UpdateReviewData = {
    body: ReviewUpdateBody;
    path: {
        artToyID: number;
    };
    query?: never;
    url: '/v1/art-toy/review/{artToyID}';
};

export type UpdateReviewErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type UpdateReviewError = UpdateReviewErrors[keyof UpdateReviewErrors];

export type UpdateReviewResponses = {
    /**
     * OK
     */
    200: Review;
};

export type UpdateReviewResponse = UpdateReviewResponses[keyof UpdateReviewResponses];

export type CreateReviewData = {
    body: ReviewCreateBody;
    path: {
        artToyID: number;
    };
    query?: never;
    url: '/v1/art-toy/review/{artToyID}';
};

export type CreateReviewErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type CreateReviewError = CreateReviewErrors[keyof CreateReviewErrors];

export type CreateReviewResponses = {
    /**
     * OK
     */
    200: Review;
};

export type CreateReviewResponse = CreateReviewResponses[keyof CreateReviewResponses];

export type SearchArtToysData = {
    body?: never;
    path?: never;
    query?: {
        keyword?: string;
        /**
         * Sorting key. Available values: 'release_date', 'price'.
         */
        sort_key?: 'release_date' | 'price';
        /**
         * If true, sorting is in descending order. Sorting is applied only if 'sort_key' is defined.
         */
        reverse?: boolean;
    };
    url: '/v1/art-toy/search';
};

export type SearchArtToysErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type SearchArtToysError = SearchArtToysErrors[keyof SearchArtToysErrors];

export type SearchArtToysResponses = {
    /**
     * OK
     */
    200: ArrayResponseArtToy;
};

export type SearchArtToysResponse = SearchArtToysResponses[keyof SearchArtToysResponses];

export type DeleteArtToyData = {
    body?: never;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/art-toy/{id}';
};

export type DeleteArtToyErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type DeleteArtToyError = DeleteArtToyErrors[keyof DeleteArtToyErrors];

export type DeleteArtToyResponses = {
    /**
     * No Content
     */
    204: void;
};

export type DeleteArtToyResponse = DeleteArtToyResponses[keyof DeleteArtToyResponses];

export type GetArtToyByIdData = {
    body?: never;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/art-toy/{id}';
};

export type GetArtToyByIdErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetArtToyByIdError = GetArtToyByIdErrors[keyof GetArtToyByIdErrors];

export type GetArtToyByIdResponses = {
    /**
     * OK
     */
    200: ArtToy;
};

export type GetArtToyByIdResponse = GetArtToyByIdResponses[keyof GetArtToyByIdResponses];

export type UpdateArtToyData = {
    body: ArtToyUpdateBody;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/art-toy/{id}';
};

export type UpdateArtToyErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type UpdateArtToyError = UpdateArtToyErrors[keyof UpdateArtToyErrors];

export type UpdateArtToyResponses = {
    /**
     * OK
     */
    200: ArtToy;
};

export type UpdateArtToyResponse = UpdateArtToyResponses[keyof UpdateArtToyResponses];

export type ChangePasswordData = {
    body: ChangePasswordBody;
    path?: never;
    query?: never;
    url: '/v1/auth/change-password';
};

export type ChangePasswordErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type ChangePasswordError = ChangePasswordErrors[keyof ChangePasswordErrors];

export type ChangePasswordResponses = {
    /**
     * No Content
     */
    204: void;
};

export type ChangePasswordResponse = ChangePasswordResponses[keyof ChangePasswordResponses];

export type LoginData = {
    body: LoginBody;
    path?: never;
    query?: never;
    url: '/v1/auth/login';
};

export type LoginErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type LoginError = LoginErrors[keyof LoginErrors];

export type LoginResponses = {
    /**
     * OK
     */
    200: User;
};

export type LoginResponse = LoginResponses[keyof LoginResponses];

export type LogoutData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/auth/logout';
};

export type LogoutErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type LogoutError = LogoutErrors[keyof LogoutErrors];

export type LogoutResponses = {
    /**
     * No Content
     */
    204: void;
};

export type LogoutResponse = LogoutResponses[keyof LogoutResponses];

export type MeData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/auth/me';
};

export type MeErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type MeError = MeErrors[keyof MeErrors];

export type MeResponses = {
    /**
     * OK
     */
    200: User;
};

export type MeResponse = MeResponses[keyof MeResponses];

export type PasswordResetRequestsData = {
    body: RequestPasswordResetBody;
    path?: never;
    query?: never;
    url: '/v1/auth/password-reset-requests';
};

export type PasswordResetRequestsErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type PasswordResetRequestsError = PasswordResetRequestsErrors[keyof PasswordResetRequestsErrors];

export type PasswordResetRequestsResponses = {
    /**
     * No Content
     */
    204: void;
};

export type PasswordResetRequestsResponse = PasswordResetRequestsResponses[keyof PasswordResetRequestsResponses];

export type RegisterData = {
    body: RegisterBody;
    path?: never;
    query?: never;
    url: '/v1/auth/register';
};

export type RegisterErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type RegisterError = RegisterErrors[keyof RegisterErrors];

export type RegisterResponses = {
    /**
     * OK
     */
    200: User;
};

export type RegisterResponse = RegisterResponses[keyof RegisterResponses];

export type ResetPasswordData = {
    body: ResetPasswordBody;
    path?: never;
    query?: never;
    url: '/v1/auth/reset-password';
};

export type ResetPasswordErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type ResetPasswordError = ResetPasswordErrors[keyof ResetPasswordErrors];

export type ResetPasswordResponses = {
    /**
     * No Content
     */
    204: void;
};

export type ResetPasswordResponse = ResetPasswordResponses[keyof ResetPasswordResponses];

export type CompleteOrderData = {
    body?: never;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/buyer/order/{id}/complete';
};

export type CompleteOrderErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type CompleteOrderError = CompleteOrderErrors[keyof CompleteOrderErrors];

export type CompleteOrderResponses = {
    /**
     * OK
     */
    200: Order;
};

export type CompleteOrderResponse = CompleteOrderResponses[keyof CompleteOrderResponses];

export type GetOrderOfBuyerData = {
    body?: never;
    path: {
        orderID: number;
    };
    query?: never;
    url: '/v1/buyer/order/{orderID}';
};

export type GetOrderOfBuyerErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetOrderOfBuyerError = GetOrderOfBuyerErrors[keyof GetOrderOfBuyerErrors];

export type GetOrderOfBuyerResponses = {
    /**
     * OK
     */
    200: Order;
};

export type GetOrderOfBuyerResponse = GetOrderOfBuyerResponses[keyof GetOrderOfBuyerResponses];

export type GetOrdersOfBuyerData = {
    body?: never;
    path?: never;
    query?: {
        status?: 'PREPARING' | 'DELIVERING' | 'COMPLETED';
    };
    url: '/v1/buyer/orders';
};

export type GetOrdersOfBuyerErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetOrdersOfBuyerError = GetOrdersOfBuyerErrors[keyof GetOrdersOfBuyerErrors];

export type GetOrdersOfBuyerResponses = {
    /**
     * OK
     */
    200: ArrayResponseOrder;
};

export type GetOrdersOfBuyerResponse = GetOrdersOfBuyerResponses[keyof GetOrdersOfBuyerResponses];

export type GetCartData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/cart';
};

export type GetCartErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetCartError = GetCartErrors[keyof GetCartErrors];

export type GetCartResponses = {
    /**
     * OK
     */
    200: ArrayResponseCartItem;
};

export type GetCartResponse = GetCartResponses[keyof GetCartResponses];

export type AddItemToCartData = {
    body: CartItemCreateBody;
    path?: never;
    query?: never;
    url: '/v1/cart/add-item';
};

export type AddItemToCartErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type AddItemToCartError = AddItemToCartErrors[keyof AddItemToCartErrors];

export type AddItemToCartResponses = {
    /**
     * OK
     */
    200: CartItem;
};

export type AddItemToCartResponse = AddItemToCartResponses[keyof AddItemToCartResponses];

export type CheckoutData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/cart/checkout';
};

export type CheckoutErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type CheckoutError = CheckoutErrors[keyof CheckoutErrors];

export type CheckoutResponses = {
    /**
     * OK
     */
    200: CheckoutOutputBody;
};

export type CheckoutResponse = CheckoutResponses[keyof CheckoutResponses];

export type ClearItemsFromCartData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/cart/clear-items';
};

export type ClearItemsFromCartErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type ClearItemsFromCartError = ClearItemsFromCartErrors[keyof ClearItemsFromCartErrors];

export type ClearItemsFromCartResponses = {
    /**
     * No Content
     */
    204: void;
};

export type ClearItemsFromCartResponse = ClearItemsFromCartResponses[keyof ClearItemsFromCartResponses];

export type PaymentSuccessCallbackData = {
    body: Blob | File;
    headers?: {
        'Stripe-Signature'?: string;
    };
    path?: never;
    query?: never;
    url: '/v1/cart/payment-success-callback';
};

export type PaymentSuccessCallbackErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type PaymentSuccessCallbackError = PaymentSuccessCallbackErrors[keyof PaymentSuccessCallbackErrors];

export type PaymentSuccessCallbackResponses = {
    /**
     * No Content
     */
    204: void;
};

export type PaymentSuccessCallbackResponse = PaymentSuccessCallbackResponses[keyof PaymentSuccessCallbackResponses];

export type RemoveItemFromCartData = {
    body?: never;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/cart/remove-item/{id}';
};

export type RemoveItemFromCartErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type RemoveItemFromCartError = RemoveItemFromCartErrors[keyof RemoveItemFromCartErrors];

export type RemoveItemFromCartResponses = {
    /**
     * No Content
     */
    204: void;
};

export type RemoveItemFromCartResponse = RemoveItemFromCartResponses[keyof RemoveItemFromCartResponses];

export type GetMyArtToysData = {
    body?: never;
    path?: never;
    query?: never;
    url: '/v1/my-art-toy';
};

export type GetMyArtToysErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetMyArtToysError = GetMyArtToysErrors[keyof GetMyArtToysErrors];

export type GetMyArtToysResponses = {
    /**
     * OK
     */
    200: ArrayResponseArtToy;
};

export type GetMyArtToysResponse = GetMyArtToysResponses[keyof GetMyArtToysResponses];

export type GetOrdersByStatusData = {
    body?: never;
    path: {
        status: 'PREPARING' | 'DELIVERING' | 'COMPLETED';
    };
    query?: never;
    url: '/v1/order/{status}';
};

export type GetOrdersByStatusErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetOrdersByStatusError = GetOrdersByStatusErrors[keyof GetOrdersByStatusErrors];

export type GetOrdersByStatusResponses = {
    /**
     * OK
     */
    200: ArrayResponseOrder;
};

export type GetOrdersByStatusResponse = GetOrdersByStatusResponses[keyof GetOrdersByStatusResponses];

export type GetReviewData = {
    body?: never;
    path: {
        sellerID: number;
    };
    query?: never;
    url: '/v1/seller/art-toy/review/{sellerID}';
};

export type GetReviewErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetReviewError = GetReviewErrors[keyof GetReviewErrors];

export type GetReviewResponses = {
    /**
     * OK
     */
    200: ArrayResponseReview;
};

export type GetReviewResponse = GetReviewResponses[keyof GetReviewResponses];

export type UpdateOrderData = {
    body: OrderUpdateBody;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/seller/order/{id}';
};

export type UpdateOrderErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type UpdateOrderError = UpdateOrderErrors[keyof UpdateOrderErrors];

export type UpdateOrderResponses = {
    /**
     * OK
     */
    200: Order;
};

export type UpdateOrderResponse = UpdateOrderResponses[keyof UpdateOrderResponses];

export type GetOrderOfSellerData = {
    body?: never;
    path: {
        orderID: number;
    };
    query?: never;
    url: '/v1/seller/order/{orderID}';
};

export type GetOrderOfSellerErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetOrderOfSellerError = GetOrderOfSellerErrors[keyof GetOrderOfSellerErrors];

export type GetOrderOfSellerResponses = {
    /**
     * OK
     */
    200: Order;
};

export type GetOrderOfSellerResponse = GetOrderOfSellerResponses[keyof GetOrderOfSellerResponses];

export type GetOrdersOfSellerData = {
    body?: never;
    path?: never;
    query?: {
        status?: 'PREPARING' | 'DELIVERING' | 'COMPLETED';
    };
    url: '/v1/seller/orders';
};

export type GetOrdersOfSellerErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetOrdersOfSellerError = GetOrdersOfSellerErrors[keyof GetOrdersOfSellerErrors];

export type GetOrdersOfSellerResponses = {
    /**
     * OK
     */
    200: ArrayResponseOrder;
};

export type GetOrdersOfSellerResponse = GetOrdersOfSellerResponses[keyof GetOrdersOfSellerResponses];

export type GetArtToysOfSellerData = {
    body?: never;
    path: {
        id: number;
    };
    query?: never;
    url: '/v1/seller/{id}/art-toy';
};

export type GetArtToysOfSellerErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type GetArtToysOfSellerError = GetArtToysOfSellerErrors[keyof GetArtToysOfSellerErrors];

export type GetArtToysOfSellerResponses = {
    /**
     * OK
     */
    200: ArrayResponseArtToy;
};

export type GetArtToysOfSellerResponse = GetArtToysOfSellerResponses[keyof GetArtToysOfSellerResponses];

export type UpdateUserData = {
    body: UserUpdateBody;
    path?: never;
    query?: never;
    url: '/v1/user';
};

export type UpdateUserErrors = {
    /**
     * Error
     */
    default: ErrorModel;
};

export type UpdateUserError = UpdateUserErrors[keyof UpdateUserErrors];

export type UpdateUserResponses = {
    /**
     * No Content
     */
    204: void;
};

export type UpdateUserResponse = UpdateUserResponses[keyof UpdateUserResponses];

export type ClientOptions = {
    baseUrl: `${string}://${string}` | (string & {});
};