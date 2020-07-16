package vsphere

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestVSphere(t *testing.T) {
	cfg := &Config{
		VSphereServerHost: "127.0.0.1:8990",
		User:              "user",
		Password:          "pass",
	}

	client, err := New(cfg)
	require.NoError(t, err)

	// fetch datacenters
	dcs, err := client.ListDatacenters()
	require.NoError(t, err)
	assert.Equal(t, 1, len(dcs))

	// fetch license
	lcs, err := client.ListLicenses()
	require.NoError(t, err)
	assert.Equal(t, 1, len(lcs))

	for _, dc := range dcs {
		// list hosts
		hosts, err := client.ListHosts(dc)
		require.NoError(t, err)
		assert.Equal(t, 3, len(hosts))

		// list vms
		vms, err := client.ListVirtualMachines(dc)
		require.NoError(t, err)
		assert.Equal(t, 3, len(vms))
	}
}

func TestESXi(t *testing.T) {
	cfg := &Config{
		VSphereServerHost: "192.168.56.102",
		User:              "root",
		Password:          "password1!",
	}

	client, err := New(cfg)
	require.NoError(t, err)

	// fetch datacenters
	dcs, err := client.ListDatacenters()
	require.NoError(t, err)
	assert.Equal(t, 1, len(dcs))

	// // fetch license
	// lcs, err := client.ListLicenses()
	// require.NoError(t, err)
	// assert.Equal(t, 1, len(lcs))

	// list hosts
	for _, dc := range dcs {
		hosts, err := client.ListHosts(dc)
		require.NoError(t, err)
		assert.Equal(t, 1, len(hosts))

		switches, err := EsxiVswitchStandard(client.Client, hosts[0])
		require.NoError(t, err)
		assert.Equal(t, 2, len(switches))

		switches, err = EsxiVswitchDvs(client.Client, hosts[0])
		require.NoError(t, err)
		assert.Equal(t, 0, len(switches))

		nics, err := EsxiVmknics(client.Client, hosts[0])
		require.NoError(t, err)
		assert.Equal(t, 1, len(nics))

		// list vms
		vms, err := client.ListVirtualMachines(dc)
		require.NoError(t, err)
		assert.Equal(t, 0, len(vms))
	}
}
