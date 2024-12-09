package testhandlerobservables_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTesthandlerobservables(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testhandlerobservables Suite")
}
