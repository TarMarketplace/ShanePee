Feature: Calculator
    Scenario: Add two numbers
        Given a Calculator
        When I add 1 and 2
        Then the result should be 3