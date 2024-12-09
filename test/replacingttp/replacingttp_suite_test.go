package testreplacingttp_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestreplacingttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testreplacingttp Suite")
}
