*** Settings ***
Library    SeleniumLibrary
Library    String

*** Variables ***
${BROWSER}    chrome
${URL}    http://localhost:3000
${VALID_EMAIL}    p.kunlayanathee@gmail.com
${VALID_PASSWORD}    password
${VALID_ART_TOY_NAME}    Test Art Toy
${VALID_ART_TOY_DESCRIPTION}    This is a valid art toy description
${VALID_ART_TOY_PRICE}    100.00
${VALID_IMAGE_PATH}    ${CURDIR}/test.jpg

*** Keywords ***
Open Browser To Home Page
    Open Browser    ${URL}    ${BROWSER}
    Maximize Browser Window

Login As Seller
    [Arguments]    ${email}    ${password}
    Go To    ${URL}/login
    Input Text    xpath=//input[@name='email']    ${email}
    Input Text    xpath=//input[@name='password']    ${password}
    Click Button    xpath=//button[contains(text(), 'เข้าสู่ระบบ')]
    Wait Until Element Is Visible    xpath=//p[contains(text(), 'สวัสดี,')]

Navigate To Create Art Toy Page
    Go To    ${URL}/product/create
    Wait Until Element Is Visible    xpath=//h4[contains(text(), 'วางจำหน่าย Art Toy ใหม่')]

Clear Database
    # This would need to be implemented based on your test database setup
    Log    Clearing test database
