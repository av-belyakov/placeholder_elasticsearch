package testhandlerttp_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTesthandlerttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testhandlerttp Suite")
}
