package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeSelector(t *testing.T) {
	selector, err := NormalizeSelector("https://github.com/octo/demo/pull/7", 7)
	require.NoError(t, err)
	assert.Equal(t, "https://github.com/octo/demo/pull/7", selector)

	selector, err = NormalizeSelector("7", 7)
	require.NoError(t, err)
	assert.Equal(t, "7", selector)

	selector, err = NormalizeSelector("", 42)
	require.NoError(t, err)
	assert.Equal(t, "42", selector)

	_, err = NormalizeSelector("https://github.com/octo/demo/pull/7", 8)
	require.Error(t, err)

	_, err = NormalizeSelector("octo/demo#7", 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "URL or number")
}

func TestResolveURL(t *testing.T) {
	id, err := Resolve("https://github.com/octo/demo/pull/9", "", "")
	require.NoError(t, err)
	assert.Equal(t, Identity{Owner: "octo", Repo: "demo", Host: "github.com", Number: 9}, id)
}

func TestResolveHostSanitization(t *testing.T) {
	id, err := Resolve("11", "octo/demo", "HTTPS://GHE.EXAMPLE.COM:8443/")
	require.NoError(t, err)
	assert.Equal(t, "ghe.example.com", id.Host)
}

func TestResolveURLHostPrecedence(t *testing.T) {
	id, err := Resolve("https://git.enterprise.local:8443/octo/demo/pull/13", "", "github.com")
	require.NoError(t, err)
	assert.Equal(t, Identity{Owner: "octo", Repo: "demo", Host: "git.enterprise.local", Number: 13}, id)
}

func TestResolveNumberRequiresRepo(t *testing.T) {
	_, err := Resolve("7", "", "")
	require.Error(t, err)

	id, err := Resolve("7", "octo/demo", "github.com")
	require.NoError(t, err)
	assert.Equal(t, Identity{Owner: "octo", Repo: "demo", Host: "github.com", Number: 7}, id)
}

func TestResolveNumberMissingRepoPrefix(t *testing.T) {
	_, err := Resolve("7", "", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "--repo:")
}

func TestResolveNumberInvalidRepoPrefix(t *testing.T) {
	_, err := Resolve("7", "not-a-repo", "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "--repo:")
}

func TestSanitizeHost(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{"empty", "", "github.com"},
		{"bare", "github.com", "github.com"},
		{"https scheme", "https://github.com", "github.com"},
		{"http scheme", "http://github.com", "github.com"},
		{"scheme with path", "https://github.com/some/path", "github.com"},
		{"with port", "ghe.example.com:8443", "ghe.example.com"},
		{"scheme with port", "https://ghe.example.com:8443", "ghe.example.com"},
		{"mixed case", "HTTPS://GitHub.Com", "github.com"},
		{"trailing slash", "https://github.com/", "github.com"},
		{"whitespace", "  github.com  ", "github.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeHost(tt.raw)
			assert.Equal(t, tt.want, got)
		})
	}
}
