package tests

import "github.com/cucumber/godog"

func aMockServerIsRunning() error {
	return godog.ErrPending
}

func resetMockServer() error {
	return godog.ErrPending
}

func sQLCommand(arg1 *godog.DocString) error {
	return godog.ErrPending
}

func aMockServerRequestWithMethodToAndBody(arg1, arg2 string, arg3 *godog.DocString) error {
	return godog.ErrPending
}

func theMockServerResponseWithStatus(arg1 int) error {
	return godog.ErrPending
}

func theMockServerResponseWithStatusAndResponsePayloadShouldMatchJson(arg1 int, arg2 *godog.DocString) error {
	return godog.ErrPending
}
