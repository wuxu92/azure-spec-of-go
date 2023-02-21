package specs_test

import (
	"azure-spec-of-go/specs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load Spec", func() {
	Context("load foo.json", func() {
		var fileName = "foo.json"
		It("loads.Spec", func() {
			doc, err := specs.LoadSpec(fileName)
			Expect(err).To(BeNil())
			GinkgoT().Logf("got origin spec: %+v", doc)
			Expect(doc.Spec()).NotTo(BeNil())
		})
	})

	Context("Load cycle.json", func() {
		var fileName = "cycle.json"
		FIt("Load Expanded Spec", func() {
			doc, err := specs.LoadExpanded(fileName)
			Expect(err).To(BeNil())
			Expect(doc.Spec()).ToNot(BeNil())
		})
	})

	Context("Load keys.json", func() {
		var fileName = "kv/stable/2023-02-21/keys.json"
		It("Load Expanded Spec", func() {
			doc, err := specs.LoadExpanded(fileName)
			Expect(err).To(BeNil())
			Expect(doc.Spec()).ToNot(BeNil())
		})
	})
})
