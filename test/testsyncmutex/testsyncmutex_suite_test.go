package testsyncmutex_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestsyncmutex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsyncmutex Suite")
}
