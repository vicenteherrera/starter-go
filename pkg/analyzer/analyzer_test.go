package analyzer_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	analyzer "github.com/vicenteherrera/starter-go/pkg/analyzer/containerfile"
)

var client analyzer.Client

var _ = BeforeSuite(func() {
	client = analyzer.NewClient("")
})

var _ = Describe("PsaEvaluator", func() {

	Context("When I analize without a filename", func() {
		It("It returns an error", func() {
			// _, _ = fmt.Fprintf(GinkgoWriter, "Executing analyzer:\n%s", client)
			_, err := client.AnalyzeFile()
			Expect(err).Should(HaveOccurred())
		})
	})

})
