package reporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mondoo.com/cnquery/explorer"
	"go.mondoo.com/cnquery/shared"
	"sigs.k8s.io/yaml"
)

func TestCSVExport(t *testing.T) {
	data, err := os.ReadFile("testdata/kubernetes_report.yaml")
	require.NoError(t, err)

	var report *explorer.ReportCollection
	err = yaml.Unmarshal(data, &report)
	require.NoError(t, err)
	w := shared.IOWriter{Writer: os.Stdout}
	err = ReportCollectionToCSV(report, &w)
	require.NoError(t, err)
}