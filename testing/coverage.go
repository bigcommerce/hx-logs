package testing

import (
	"fmt"
	"os"
	t "testing"
)

func RunAndRequireCoverage(m *t.M, minimumCoverage float64) {
	result := m.Run()
	if result == 0 && t.CoverMode() != "" {
		coverage := t.Coverage()
		if coverage < minimumCoverage {
			fmt.Printf("Only %0.1f%% of lines were covered out of the required %0.1f%%\n", coverage*100, minimumCoverage*100)
			result = -1
		}
	}
	os.Exit(result)
}
