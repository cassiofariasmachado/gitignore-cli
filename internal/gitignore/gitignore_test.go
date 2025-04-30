package gitignore

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func CreateMockServer(handlerFunc http.HandlerFunc) (*httptest.Server, string) {
	mockServer := httptest.NewServer(http.HandlerFunc(handlerFunc))

	return mockServer, mockServer.URL + "/%s.gitignore"
}

func TestConfigure(t *testing.T) {
	url := "https://example.com"
	Configure(&url)

	if baseUrl == nil || *baseUrl != url {
		t.Errorf("Expected baseUrl to be %s, got %v", url, baseUrl)
	}
}

func TestGetGitIgnoreFoundWithSuccess(t *testing.T) {
	// Mock HTTP server
	mockServer, mockServerUrl := CreateMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mock content"))
	})

	Configure(&mockServerUrl)

	gitignore, err := GetGitIgnore("test", "test.gitignore")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	content, _ := io.ReadAll(gitignore.Content)

	defer mockServer.Close()

	if string(content) != "mock content" {
		t.Errorf("Expected content to be 'mock content', got %s", string(content))
	}
}

func TestGetGitIgnoreNotFound(t *testing.T) {
	// Mock HTTP server
	mockServer, mockServerUrl := CreateMockServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	Configure(&mockServerUrl)

	_, err := GetGitIgnore("test", "test.gitignore")

	defer mockServer.Close()

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestSaveFile(t *testing.T) {
	content := "mock content"
	gitignore := GitIgnore{
		Name:    "test",
		Path:    "test.gitignore",
		Content: io.NopCloser(strings.NewReader(content)),
	}

	gitignore.SaveFile()
	defer os.Remove("test.gitignore")

	data, err := os.ReadFile("test.gitignore")
	if err != nil {
		t.Fatalf("Expected no error reading file, got %v", err)
	}

	if string(data) != content {
		t.Errorf("Expected file content to be '%s', got '%s'", content, string(data))
	}
}
