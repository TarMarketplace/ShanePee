*** Settings ***
Documentation    Test suite for art toy creation functionality
Resource    resources/common.resource
Test Setup    Run Keywords    Open Browser To Home Page    AND    Login As Seller    ${VALID_EMAIL}    ${VALID_PASSWORD}    AND    Navigate To Create Art Toy Page
Test Teardown    Close Browser

*** Test Cases ***
Invalid Art Toy Creation - Null Name
    [Documentation]    EC1: Art toy name is null
    Input Art Toy Data    ${EMPTY}    ${VALID_ART_TOY_DESCRIPTION}    ${VALID_ART_TOY_PRICE}    ${VALID_IMAGE_PATH}
    Submit Art Toy Form
    Form Message Should appear    Name is required

Invalid Art Toy Creation - Null Description
    [Documentation]    EC4: Art toy description is null
    Input Art Toy Data    ${VALID_ART_TOY_NAME}    ${EMPTY}    ${VALID_ART_TOY_PRICE}    ${VALID_IMAGE_PATH}
    Submit Art Toy Form
    Form Message Should appear    Description is required

Invalid Art Toy Creation - Null Price
    [Documentation]    EC7: Art toy price is null
    Input Art Toy Data    ${VALID_ART_TOY_NAME}    ${VALID_ART_TOY_DESCRIPTION}    ${EMPTY}    ${VALID_IMAGE_PATH}
    Submit Art Toy Form
    Form Message Should appear    Price is required

Invalid Art Toy Creation - Zero Price
    [Documentation]    EC10: Art toy price is exactly 0
    Input Art Toy Data    ${VALID_ART_TOY_NAME}    ${VALID_ART_TOY_DESCRIPTION}    0    ${VALID_IMAGE_PATH}
    Submit Art Toy Form
    Form Message Should appear    Price is required

Invalid Art Toy Creation - No Image
    [Documentation]    EC12: Art toy image is null
    Input Art Toy Data    ${VALID_ART_TOY_NAME}    ${VALID_ART_TOY_DESCRIPTION}    ${VALID_ART_TOY_PRICE}    ${EMPTY}
    Submit Art Toy Form
    Form Message Should appear    Please upload an image

Valid Art Toy Creation
    [Documentation]    EC3, EC6, EC11, EC14: Valid art toy creation with all valid inputs
    Input Art Toy Data    ${VALID_ART_TOY_NAME}    ${VALID_ART_TOY_DESCRIPTION}    ${VALID_ART_TOY_PRICE}    ${VALID_IMAGE_PATH}
    Submit Art Toy Form
    Art Toy Creation Should Succeed

*** Keywords ***
Input Art Toy Data
    [Arguments]    ${name}    ${description}    ${price}    ${image_path}
    Input Text    xpath=//input[@name='name']    ${name}
    Input Text    xpath=//textarea[@name='description']    ${description}
    Input Text    xpath=//input[@name='price']    ${price}
    Run Keyword If    '${image_path}' != '${EMPTY}'    Choose File    xpath=//input[@type='file']    ${image_path}

Submit Art Toy Form
    Click Button    xpath=//button[contains(text(), 'วางจำหน่าย')]

Form Message Should appear
    [Arguments]    ${error_message}
		Wait Until Element Is Visible    xpath=//p[substring(@id, string-length(@id) - string-length('-form-item-message') + 1) = '-form-item-message']
    Element Should Contain    xpath=//p[substring(@id, string-length(@id) - string-length('-form-item-message') + 1) = '-form-item-message']    ${error_message}

Art Toy Creation Should Succeed
    Wait Until Element Is Visible    xpath=//h4[contains(text(), '${VALID_ART_TOY_NAME}')]
