Feature: User login (EPIC-1-7)

  Scenario: User login with valid credentials
    Given a registered user with a valid email and password
    When they enter their credentials correctly
    Then the system logs them in

  Scenario: User login with invalid credentials
    Given a registered user with a valid email and password
    When they enter their credentials incorrectly
    Then the system does not log them in
