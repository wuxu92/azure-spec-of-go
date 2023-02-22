package specs_test

import (
	"azure-spec-of-go/specs"
	"azure-spec-of-go/utils"
	"azure-spec-of-go/utils/logs"

	"github.com/go-openapi/spec"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Load Spec", func() {

	testSpecStruct := func(swag *spec.Swagger) *specs.Spec {
		spec := specs.NewSpec(swag)
		logs.Info(utils.JSONIndent(spec))
		return spec
	}
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
		When("Load Expanded Spec", func() {
			doc, _ := specs.LoadExpanded(fileName)
			//Expect(err).To(BeNil())
			//Expect(doc.Spec()).ToNot(BeNil())
			spec := testSpecStruct(doc.Spec())

			FIt("mock json", func() {
				str := spec.RenderDefinitions()
				logs.Info("%s", str)
			})
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
