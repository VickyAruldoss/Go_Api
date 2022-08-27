package jetpack

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJetpack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jetpack Suite")
}

var _ = Describe("Jetpack test", func() {
	var (
		jetpack Jetpack
	)
	BeforeEach(func() {
		jetpack = NewJetpack()
	})

	AfterEach(func() {

	})

	Context("Print the Jetpack", func() {
		It("Should print the values", func() {
			val := jetpack.Print()
			Expect(val).To(Equal("OMG it works"))
		})
	})

})
