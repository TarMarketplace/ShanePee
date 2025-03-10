Feature: Login user
    In order to use the application
    As a user
    I want to login to my account

    Background: Clean database
      Given SQL command
    """
    DELETE FROM myschema.users;
    """
      And reset mock server

    Scenario: Login with valid credentials
        Given a mock server is running
        And a user with an email "example@email.com" already exists with password "password"
        When a mock server request with method: "POST" to "/v1/auth/login" and body
        """json
        {
          "email" : "example@email.com",
          "password" : "password"
        }
        """ 
        Then the mock server response with status 200 and response payload should match json:
        """json
        {
            "$schema": "http://.../schemas/User.json",
            "id": 4109519149,
            "email": "johndoe@example.com",
            "address": {...},
            "payment_method": {...},
            ...
        }
        """
  
    Scenario: Login with invalid credentials
        Given a mock server is running
        And a user with an email "example@email.com" already exists with password "password"
        When a mock server request with method: "POST" to "/v1/auth/login" and body
        """json
        {
          "email" : "example@email.com",
          "password" : "ppp"
        }
        """ 
        Then the mock server response with status 401 and response payload should match json:
        """json
        {
          "title": "Forbidden",
          "status": 403,
          "detail": "Incorrect email or password",
          ...          
        }
        """