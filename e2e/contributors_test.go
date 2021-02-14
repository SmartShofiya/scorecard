package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ossf/scorecard/checker"
	"github.com/ossf/scorecard/checks"
)

var _ = Describe("E2E TEST:CodeReview", func() {
	Context("E2E TEST:Validating project contributors", func() {
		It("Should return valid project contributors", func() {
			l := log{}
			checker := checker.Checker{
				Ctx:         context.Background(),
				Client:      ghClient,
				HttpClient:  client,
				Owner:       "ossf",
				Repo:        "scorecard",
				GraphClient: graphClient,
				Logf:        l.Logf,
			}
			result := checks.Contributors(checker)
			Expect(result.Error).Should(BeNil())
			Expect(result.Pass).Should(BeTrue())
		})
		It("Should return valid project contributors", func() {
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
			result := checks.Contributors(checker)
			Expect(result.Error).Should(BeNil())
			Expect(result.Pass).Should(BeTrue())
		})
	})
})
