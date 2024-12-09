package testmongodbconnect_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestmongodbconnect(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testmongodbconnect Suite")
}
