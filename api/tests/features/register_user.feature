Feature: Register User
    In order to use the application
    As a user
    I want to register a new account

    Background: Clean database
      Given SQL command
    """
    DELETE FROM myschema.users;
    """
      And reset mock server

    Scenario: Register a new user successfully
        Given a mock server is running
        When a mock server request with method: "POST" to "/v1/auth/register" and body
        """json
        {
          "email" : "example@email.com",
          "password" : "password"
        }
        """ 
        Then the mock server response with status 200 and response payload should match json:
        """json
        {
            "id": 4109519149,
            "email": "johndoe@example.com",
            "address": {...},
            "payment_method": {...},
            ...
        }
        """
 
    Scenario: Register with a duplicate email
        Given a mock server is running 
        And a user with an email "example@email.com" already exists
        When a mock server request with method: "POST" to "/v1/auth/register" and body
        """json
        {
          "email" : "example@email.com",
          "password" : "password"
        }
        """ 
        Then the mock server response with status 403 and response payload should match json:
        """json
        {
          "title": "Forbidden",
          "status": 403,
          "detail": "User with this email already exists",
          ...
        }
        """