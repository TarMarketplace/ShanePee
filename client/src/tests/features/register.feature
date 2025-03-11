Feature: Register User (EPIC-1-1)

    Scenario: Register with valid credentials
        Given a user with a unique email
        When they submit their email and password
        Then the system creates their account successfully

    Scenario: Register with invalid credentials
        Given a user with an already registered email
        When they attempt to register with the same email
        Then the system shows an error message