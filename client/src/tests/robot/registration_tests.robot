*** Settings ***
Documentation    Test suite for user registration functionality
Resource    resources/common.resource
Test Setup    Open Browser To Home Page
Test Teardown    Close Browser

*** Test Cases ***
Invalid Registration - Email Without @ Symbol
    [Documentation]    EC1: Email without @ symbol should be invalid
    Go To Registration Page
    Input Registration Data    invalid.email    ${VALID_PASSWORD}    ${VALID_PASSWORD}
    Submit Registration
    Form Message Should appear    Invalid email

Invalid Registration - Empty Local Part Before @
    [Documentation]    EC2: Email with empty part before @ should be invalid
    Go To Registration Page
    Input Registration Data    @domain.com    ${VALID_PASSWORD}    ${VALID_PASSWORD}
    Submit Registration
    Form Message Should appear    Invalid email

Invalid Registration - Invalid Domain Format
    [Documentation]    EC3: Email without valid domain should be invalid
    Go To Registration Page
    Input Registration Data    test@invalid    ${VALID_PASSWORD}    ${VALID_PASSWORD}
    Submit Registration
    Form Message Should appear    Invalid email 

Invalid Registration - Existing Email
    [Documentation]    EC5: Email that already exists should be invalid
    Create Existing User
    Go To Registration Page
    Input Registration Data    ${VALID_EMAIL}    ${VALID_PASSWORD}    ${VALID_PASSWORD}
    Submit Registration
    Toast Message Should appear    Something went wrong

Invalid Registration - Short Password
    [Documentation]    EC7: Password with less than 8 characters should be invalid
    Go To Registration Page
    Input Registration Data    ${VALID_EMAIL}    short    short
    Submit Registration
    Form Message Should appear    Password must be at least 8 characters

Invalid Registration - Password Mismatch
    [Documentation]    EC9: Password and confirm password must match
    Go To Registration Page
    Input Registration Data    ${VALID_EMAIL}    ${VALID_PASSWORD}    differentpass
    Submit Registration
    Form Message Should appear    Password does not match

Valid Registration
    [Documentation]    EC4, EC6, EC8, EC10: Valid registration with proper email and matching passwords
    Go To Registration Page
		${RANDOM_EMAIL}=    Evaluate    'test_{}@example.com'.format(datetime.datetime.now().strftime('%Y%m%d%H%M%S'))    datetime
		Input Registration Data    ${RANDOM_EMAIL}    ${VALID_PASSWORD}    ${VALID_PASSWORD}
    Submit Registration
    Registration Should Succeed

*** Keywords ***
Go To Registration Page
    Go To    ${URL}/login?mode=register
    Wait Until Element Is Visible    xpath=//h5[contains(text(), 'สมัครใช้งาน')]

Input Registration Data
    [Arguments]    ${email}    ${password}    ${confirm_password}
		Input Text    xpath=//input[@name='name']    Something
		Input Text    xpath=//input[@name='surname']    Random
		Input Text    xpath=//input[@type='tel']    0123456789
		Click Button    xpath=//button[@data-testid='gender-select']
		Click Element    xpath=//span[contains(text(), 'ชาย')]
		Next Step
    Input Text    xpath=//input[@name='email']    ${email}
    Input Text    xpath=//input[@name='password']    ${password}
    Input Text    xpath=//input[@name='confirmPassword']    ${confirm_password}

Next Step
		Click Button    xpath=//button[contains(text(), 'ถัดไป')]
		Wait Until Element Is Visible    xpath=//input[@name='email']

Submit Registration
    Click Button    xpath=//button[contains(text(), 'สมัครใช้งาน')]

Form Message Should appear
    [Arguments]    ${error_message}
		Wait Until Element Is Visible    xpath=//p[substring(@id, string-length(@id) - string-length('-form-item-message') + 1) = '-form-item-message']
    Element Should Contain    xpath=//p[substring(@id, string-length(@id) - string-length('-form-item-message') + 1) = '-form-item-message']    ${error_message}

Toast Message Should appear
		[Arguments]    ${error_message}
		Wait Until Element Is Visible    xpath=//div[@data-title]
    Element Should Contain    xpath=//div[@data-title]    ${error_message}

Registration Should Succeed
		Wait Until Element Is Visible    xpath=//h5[contains(text(), 'Art Toys แนะนำสำหรับคุณ')]

Create Existing User
    # This would need to be implemented based on your test database setup
    Log    Creating existing user in test database
