package jetpack_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJetpack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jetpack Suite")
}
