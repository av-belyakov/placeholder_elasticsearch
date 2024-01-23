package testruleshandler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestruleshandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testruleshandler Suite")
}
