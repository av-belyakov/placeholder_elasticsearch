package testdatetime_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestdatetime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testdatetime Suite")
}
