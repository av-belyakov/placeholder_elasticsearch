package testnatsinteraction_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestnatsinteraction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testnatsinteraction Suite")
}
