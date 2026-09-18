//go:build linux

package machineinfo

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetectSnapshotCPUProfile(t *testing.T) {
	info, err := Detect(" amd-zenplus-v1 ")
	require.NoError(t, err)
	assert.Equal(t, runtime.GOARCH, info.Arch)
	assert.Equal(t, "snapshot-profile", info.Family)
	assert.Equal(t, "amd-zenplus-v1", info.Model)
	assert.Equal(t, "Firecracker snapshot CPU profile amd-zenplus-v1", info.ModelName)
	assert.Empty(t, info.Flags)
}
