package tests

import "github.com/cucumber/godog"

func aUserWithAnEmailAlreadyExists(arg1 string) error {
	return godog.ErrPending
}

func registerUserSuccessfullyScenario(ctx *godog.ScenarioContext) {
	// Background
	ctx.Step(`^SQL command$`, sQLCommand)
	ctx.Step(`^reset mock server$`, resetMockServer)

	// Given
	ctx.Step(`^a mock server is running$`, aMockServerIsRunning)

	// When
	ctx.Step(`^a mock server request with method: "([^"]*)" to "([^"]*)" and body$`, aMockServerRequestWithMethodToAndBody)

	// Then
	ctx.Step(`^the mock server response with status (\d+) and response payload should match json:$`, theMockServerResponseWithStatusAndResponsePayloadShouldMatchJson)
}

func registerWithDuplicateEmailScenario(ctx *godog.ScenarioContext) {
	// Background
	ctx.Step(`^SQL command$`, sQLCommand)
	ctx.Step(`^reset mock server$`, resetMockServer)

	// Given
	ctx.Step(`^a mock server is running$`, aMockServerIsRunning)
	ctx.Step(`^a user with an email "([^"]*)" already exists$`, aUserWithAnEmailAlreadyExists)

	// When
	ctx.Step(`^a mock server request with method: "([^"]*)" to "([^"]*)" and body$`, aMockServerRequestWithMethodToAndBody)

	// Then
	ctx.Step(`^the mock server response with status (\d+)$`, theMockServerResponseWithStatus)
}

func InitializeRegisterUserFeature(ctx *godog.ScenarioContext) {
	registerUserSuccessfullyScenario(ctx)
	registerWithDuplicateEmailScenario(ctx)
}
