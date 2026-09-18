//go:build linux

package fc

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetCPUConfigForwardsTemplateVerbatim(t *testing.T) {
	t.Parallel()

	const template = `{"cpuid_modifiers":[{"leaf":"0x1","subleaf":"0x0","flags":0,"modifiers":[{"register":"ecx","bitmap":"0b0"}]}],"kvm_capabilities":[],"msr_modifiers":[]}`

	path := filepath.Join(t.TempDir(), "cpu-template.json")
	require.NoError(t, os.WriteFile(path, []byte(template), 0o600))

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		require.Equal(t, http.MethodPut, request.Method)
		require.Equal(t, "/cpu-config", request.URL.Path)
		require.Equal(t, "application/json", request.Header.Get("Content-Type"))

		body, err := io.ReadAll(request.Body)
		require.NoError(t, err)
		require.Equal(t, template, string(body))
		require.NotContains(t, string(body), "reg_modifiers")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := &apiClient{rawClient: server.Client(), cpuConfigURL: server.URL}
	require.NoError(t, client.setCPUConfig(context.Background(), path))
}
