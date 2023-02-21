package specs_test

import (
	"os"
	"path"
	"testing"

	"azure-spec-of-go/specs"
	"azure-spec-of-go/utils"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestSpecs(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Specs")
}

var _ = ginkgo.BeforeSuite(func() {
	if env := os.Getenv(specs.SwagBaseKey); env == "" {
		err := os.Setenv(specs.SwagBaseKey, path.Join(utils.ProjectDir(), "testdata", "specs"))
		if err != nil {
			ginkgo.Fail("set env err")
		}
	}
})
