// This file is auto-generated by @hey-api/openapi-ts
import type {
  Client,
  Options as ClientOptions,
  TDataShape,
} from '@hey-api/client-next'

import { client as _heyApiClient } from './client.gen'
import type {
  ChangePasswordData,
  ChangePasswordError,
  ChangePasswordResponse,
  CreateArtToyData,
  CreateArtToyError,
  CreateArtToyResponse,
  GetArtToyByIdData,
  GetArtToyByIdError,
  GetArtToyByIdResponse,
  GetArtToysData,
  GetArtToysError,
  GetArtToysResponse,
  LoginData,
  LoginError,
  LoginResponse,
  LogoutData,
  LogoutError,
  LogoutResponse,
  MeData,
  MeError,
  MeResponse,
  PasswordChangeRequestsData,
  PasswordChangeRequestsError,
  PasswordChangeRequestsResponse,
  RegisterData,
  RegisterError,
  RegisterResponse,
  UpdateArtToyData,
  UpdateArtToyError,
  UpdateArtToyResponse,
  UpdateUserData,
  UpdateUserError,
  UpdateUserResponse,
} from './types.gen'

export type Options<
  TData extends TDataShape = TDataShape,
  ThrowOnError extends boolean = boolean,
> = ClientOptions<TData, ThrowOnError> & {
  /**
   * You can provide a client instance returned by `createClient()` instead of
   * individual options. This might be also useful if you want to implement a
   * custom client.
   */
  client?: Client
  /**
   * You can pass arbitrary values through the `meta` object. This can be
   * used to access values that aren't defined as part of the SDK function.
   */
  meta?: Record<string, unknown>
}

/**
 * Get Art Toys
 * Get art toys
 */
export const getArtToys = <ThrowOnError extends boolean = false>(
  options?: Options<GetArtToysData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    GetArtToysResponse,
    GetArtToysError,
    ThrowOnError
  >({
    url: '/v1/art-toy',
    ...options,
  })
}

/**
 * Create Art toy
 * Create a new art toy record
 */
export const createArtToy = <ThrowOnError extends boolean = false>(
  options: Options<CreateArtToyData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    CreateArtToyResponse,
    CreateArtToyError,
    ThrowOnError
  >({
    url: '/v1/art-toy',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Get Art Toy by ID
 * Get art toy by id
 */
export const getArtToyById = <ThrowOnError extends boolean = false>(
  options: Options<GetArtToyByIdData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).get<
    GetArtToyByIdResponse,
    GetArtToyByIdError,
    ThrowOnError
  >({
    url: '/v1/art-toy/{id}',
    ...options,
  })
}

/**
 * Update Art toy
 * Update an existing art toy by ID
 */
export const updateArtToy = <ThrowOnError extends boolean = false>(
  options: Options<UpdateArtToyData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).put<
    UpdateArtToyResponse,
    UpdateArtToyError,
    ThrowOnError
  >({
    url: '/v1/art-toy/{id}',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Change password
 * Change password of a user using token and request id
 */
export const changePassword = <ThrowOnError extends boolean = false>(
  options: Options<ChangePasswordData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    ChangePasswordResponse,
    ChangePasswordError,
    ThrowOnError
  >({
    url: '/v1/auth/change-password',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Login User
 * Login
 */
export const login = <ThrowOnError extends boolean = false>(
  options: Options<LoginData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    LoginResponse,
    LoginError,
    ThrowOnError
  >({
    url: '/v1/auth/login',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Logout User
 * Logout
 */
export const logout = <ThrowOnError extends boolean = false>(
  options?: Options<LogoutData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).post<
    LogoutResponse,
    LogoutError,
    ThrowOnError
  >({
    url: '/v1/auth/logout',
    ...options,
  })
}

/**
 * Get current authenticated user
 * Get authenticated user from the session
 */
export const me = <ThrowOnError extends boolean = false>(
  options?: Options<MeData, ThrowOnError>
) => {
  return (options?.client ?? _heyApiClient).get<
    MeResponse,
    MeError,
    ThrowOnError
  >({
    url: '/v1/auth/me',
    ...options,
  })
}

/**
 * Request a password reset
 * Initiates a password reset process by sending an email with reset instructions
 */
export const passwordChangeRequests = <ThrowOnError extends boolean = false>(
  options: Options<PasswordChangeRequestsData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    PasswordChangeRequestsResponse,
    PasswordChangeRequestsError,
    ThrowOnError
  >({
    url: '/v1/auth/password-change-requests',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Register User
 * Register
 */
export const register = <ThrowOnError extends boolean = false>(
  options: Options<RegisterData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).post<
    RegisterResponse,
    RegisterError,
    ThrowOnError
  >({
    url: '/v1/auth/register',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}

/**
 * Update User
 * Update user by id
 */
export const updateUser = <ThrowOnError extends boolean = false>(
  options: Options<UpdateUserData, ThrowOnError>
) => {
  return (options.client ?? _heyApiClient).patch<
    UpdateUserResponse,
    UpdateUserError,
    ThrowOnError
  >({
    url: '/v1/user',
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  })
}
