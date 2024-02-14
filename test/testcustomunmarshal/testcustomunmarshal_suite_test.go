package testcustomunmarshal_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestcustomunmarshal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testcustomunmarshal Suite")
}
