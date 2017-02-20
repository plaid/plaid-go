package plaid

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInstitutions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "institutions tests")
}

var _ = Describe("institutions", func() {

	Describe("GetInstitutionsSearch", func() {

		It("returns non-empty array", func() {
			institutions, err := GetInstitutionsSearch(Tartan, "redwood", "auth", "ins_100042")
			Expect(err).To(BeNil(), "err should be nil")
			Expect(institutions).ToNot(BeEmpty())
		})

		It("returns non-empty array", func() {
			institutions, err := GetInstitutionsSearch(Tartan, "chase", "connect", "")
			Expect(err).To(BeNil(), "err should be nil")
			Expect(institutions).ToNot(BeEmpty())
		})

		It("returns an error", func() {
			institutions, err := GetInstitutionsSearch(Tartan, "", "connect", "")
			Expect(err).ToNot(BeNil(), "err should not be nil")
			Expect(err.Error()).To(Equal("/institutions/all/ - query or institution id must be specified"))
			Expect(institutions).To(BeEmpty())
		})

	})

	Describe("GetInstitution", func() {

		It("returns proper fields", func() {
			i, err := GetInstitution(Tartan, "5301a9d704977c52b60000db")
			Expect(err).To(BeNil(), "err should be nil")
			Expect(i.HasMFA).To(BeFalse())
			Expect(i.ID).To(Equal("5301a9d704977c52b60000db"))
			Expect(i.MFA).To(BeEmpty())
			Expect(i.Name).To(Equal("American Express"))
			Expect(i.Type).To(Equal("amex"))
			Expect(i.Products).ToNot(BeEmpty())
			Expect(i.Products).To(ContainElement("balance"))
			Expect(i.Products).To(ContainElement("connect"))
		})

	})

	Describe("GetInstitution", func() {

		It("returns an error", func() {
			_, err := GetInstitution(Tartan, "")
			Expect(err).ToNot(BeNil(), "err should not be nil")
			Expect(err.Error()).To(Equal("/institutions/all/:id - institution id must be specified"))
		})

	})

	Describe("GetInstitutions", func() {
		It("returns non-empty array", func() {
			c := NewClient("test_id", "test_secret", Tartan)
			institutions, err := c.GetInstitutions(Tartan, []string{"connect", "auth"}, 20, 0)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(institutions).ToNot(BeEmpty())
		})
	})

	Describe("GetInstitutions", func() {
		It("returns non-empty array", func() {
			c := NewClient("test_id", "test_secret", Tartan)
			institutions, err := c.GetInstitutions(Tartan, []string{}, 0, 0)
			Expect(err).To(BeNil(), "err should be nil")
			Expect(institutions).ToNot(BeEmpty())
		})
	})

})

func ExampleGetInstitution() {
	institution, err := GetInstitution(Tartan, "5301a9d704977c52b60000db")
	fmt.Println(err)
	fmt.Println(institution.Name)
	fmt.Println(institution.Type)
	// Output: <nil>
	// American Express
	// amex
}
