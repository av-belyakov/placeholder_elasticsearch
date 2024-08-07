package casecomparation_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCasecomparation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Casecomparation Suite")
}
