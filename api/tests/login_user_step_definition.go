package tests

import "github.com/cucumber/godog"

func aUserWithAnEmailAlreadyExistsWithPassword(arg1, arg2 string) error {
	return godog.ErrPending
}

func InitalizeLoginUserFeature(ctx *godog.ScenarioContext) {
	// Background
	ctx.Step(`^SQL command$`, sQLCommand)
	ctx.Step(`^reset mock server$`, resetMockServer)

	// Given
	ctx.Step(`^a mock server is running$`, aMockServerIsRunning)
	ctx.Step(`^a user with an email "([^"]*)" already exists with password "([^"]*)"$`, aUserWithAnEmailAlreadyExistsWithPassword)

	// When
	ctx.Step(`^a mock server request with method: "([^"]*)" to "([^"]*)" and body$`, aMockServerRequestWithMethodToAndBody)

	// Then
	ctx.Step(`^the mock server response with status (\d+) and response payload should match json:$`, theMockServerResponseWithStatusAndResponsePayloadShouldMatchJson)
}
