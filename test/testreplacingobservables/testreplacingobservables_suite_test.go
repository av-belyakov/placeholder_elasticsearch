package testreplacingobservables_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestreplacingobservables(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testreplacingobservables Suite")
}
