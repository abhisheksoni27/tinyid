package tinyid

import (
	"crypto/rand"
	"errors"
	"strings"
	"testing"
)

func TestNew_NoLength(t *testing.T) {
	id := New(0)

	if id == "" {
		t.Errorf("could not generate tinyid %+v", id)
	}
}

func TestNew_ValidLength(t *testing.T) {
	id := New(10)

	if id == "" {
		t.Errorf("could not generate tinyid %+v", id)
	}
}

func TestNew_GeneratedIdLengthEqualPassedLength(t *testing.T) {
	specifiedLen := 10
	id := New(specifiedLen)

	if len(id) != specifiedLen {
		t.Errorf("tinyid's length is not equal to what was requested: want = %+v, got = %+v(%+v)", specifiedLen, id, len(id))
	}
}

func TestNew_IsUrlSafe(t *testing.T) {
	id := New(10)

	for _, char := range []string{"+", "/", "="} {
		if strings.Index(id, char) != -1 {
			t.Errorf("invalid url unsafe character found: id = %+v, char = %+v", id, char)
		}
	}

}

// MockErrorRand is a custom implementation of rand.Reader for testing.
type MockErrorRand struct{}

// Read generates deterministic random bytes for testing.
func (r *MockErrorRand) Read(p []byte) (n int, err error) {
	// Replace this with the desired deterministic byte sequence for your test.
	return 0, errors.New("invalid byte")
}

func TestNew_MockRandError(t *testing.T) {
	// Replace the existing rand.Reader with the MockErrorRand implementation
	// during the test by swapping it out temporarily and restoring it afterward.
	originalRandReader := rand.Reader
	rand.Reader = &MockErrorRand{}
	defer func() {
		rand.Reader = originalRandReader
	}()

	// Call your New function with the desired length
	result := New(8)

	expected := ""
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
