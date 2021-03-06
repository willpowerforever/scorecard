package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ossf/scorecard/checker"
	"github.com/ossf/scorecard/checks"
)

var _ = Describe("E2E TEST:CITests", func() {
	Context("E2E TEST:Validating use of CI tests", func() {
		It("Should return use of CI tests", func() {
			l := log{}
			checker := checker.Checker{
				Ctx:         context.Background(),
				Client:      ghClient,
				HttpClient:  client,
				Owner:       "apache",
				Repo:        "airflow",
				GraphClient: graphClient,
				Logf:        l.Logf,
			}
			result := checks.CITests(checker)
			Expect(result.Error).Should(BeNil())
			Expect(result.Pass).Should(BeTrue())
		})
	})
})
