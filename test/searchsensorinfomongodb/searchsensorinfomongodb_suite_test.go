package testsearchsensorinfomongodb_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestsearchsensorinfomongodb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsearchsensorinfomongodb Suite")
}
