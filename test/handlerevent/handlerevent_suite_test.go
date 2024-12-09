package testhandlerevent_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTesthandlerevent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testhandlerevent Suite")
}
