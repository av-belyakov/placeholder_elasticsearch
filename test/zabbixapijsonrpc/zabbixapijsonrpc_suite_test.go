package testzabbixapijsonrpc_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestzabbixapijsonrpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testzabbixapijsonrpc Suite")
}
