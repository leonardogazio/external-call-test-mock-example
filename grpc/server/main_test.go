package server

import (
	"os"
	"testing"

	foo "github.com/leonardogazio/external-call-test-mock-example/client/foo-service-api"
)

var testInstance = &apiServer{}

func TestMain(m *testing.M) {
	foo.Client = &foo.Mock{}
	os.Exit(m.Run())
}
