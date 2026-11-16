package registry

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		got  string
		want string
	}{
		{NamespaceListPath(), "/namespaces"},
		{NamespacePath("ns"), "/namespaces/ns"},
		{ResourceListPath("ns"), "/namespaces/ns/resources"},
		{ResourcePath("ns", "res"), "/namespaces/ns/resources/res"},
		{VersionListPath("ns", "res"), "/namespaces/ns/resources/res/versions"},
		{VersionPath("ns", "res", "1.0.0"), "/namespaces/ns/resources/res/versions/1.0.0"},
		{ArchivePath("ns", "res", "1.0.0"), "/namespaces/ns/resources/res/versions/1.0.0/archive"},
		{ChannelListPath("ns", "res"), "/namespaces/ns/resources/res/channels"},
		{ChannelPath("ns", "res", "stable"), "/namespaces/ns/resources/res/channels/stable"},
		{ChannelArchivePath("ns", "res", "stable"), "/namespaces/ns/resources/res/channels/stable/archive"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("path = %q, want %q", tt.got, tt.want)
		}
	}
}

func TestPathEscapesSegments(t *testing.T) {
	if got := NamespacePath("a/b"); got != "/namespaces/a%2Fb" {
		t.Errorf("NamespacePath(a/b) = %q, want %q", got, "/namespaces/a%2Fb")
	}
}

func TestPathsMatchPatterns(t *testing.T) {
	var got map[string]string
	mux := http.NewServeMux()
	mux.HandleFunc(PatternArchive, func(w http.ResponseWriter, r *http.Request) {
		got = map[string]string{
			ParamNamespace: r.PathValue(ParamNamespace),
			ParamResource:  r.PathValue(ParamResource),
			ParamVersion:   r.PathValue(ParamVersion),
		}
	})

	mux.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, ArchivePath("ns", "res", "1.0.0"), nil))
	if got[ParamNamespace] != "ns" || got[ParamResource] != "res" || got[ParamVersion] != "1.0.0" {
		t.Errorf("path values = %v, want ns, res, 1.0.0", got)
	}
}
