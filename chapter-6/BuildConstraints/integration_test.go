//go:build integration
// +build integration

package integration

import (
	"testing"
)

// go test では TestSomethinsは実行されず、 go test -tags=integrationとした場合のみテストが実行される
func TestSomething(t *testing.T) {
	t.Log("Integration test")
}
