package azure_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"go.mondoo.io/mondoo/motor/discovery/azure"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAzureInstanceFetch(t *testing.T) {
//  subscriptionid := "/subscriptions/abc/resourceGroups/name"
// 	client, err := azure.NewCompute(subscriptionid)
// 	require.NoError(t, err)

// 	ctx := context.Background()
// 	instances, err := client.ListInstances(ctx)
// 	require.NoError(t, err)

// 	assert.Equal(t, 1, len(instances))
// }
