package nest

import (
	"testing"

	"github.com/go-log/log/capture"
)

func TestNew(t *testing.T) {
	base := capture.New()
	logger := New(base, " ", map[string]interface{}{"key": "value", "count": 1}, " ")
	logger.Log("Log()", "arg")
	logger.Logf("Logf(%s)", "arg")
	expectedEntries := []string{" map[key:value count:1] Log()arg", " map[key:value count:1] Logf(arg)"}
	for i, expectedEntry := range expectedEntries {
		if i >= len(base.Entries) {
			t.Errorf("missing expected entry %d: %q", i, expectedEntry)
			continue
		}
		actualEntry := base.Entries[i]
		if actualEntry != expectedEntry {
			t.Errorf("unexpected entry %d: %q (expected %q)", i, actualEntry, expectedEntry)
		}
	}
	if len(base.Entries) > len(expectedEntries) {
		t.Errorf("additional unexpected entries: %v", base.Entries[len(expectedEntries):])
	}
}

func TestNewf(t *testing.T) {
	base := capture.New()
	logger := Newf(base, "wrap(%s,%d)", "a", 1)
	logger.Log("Log()", "arg")
	logger.Logf("Logf(%s)", "arg")
	expectedEntries := []string{"wrap(a,1)Log()arg", "wrap(a,1)Logf(arg)"}
	for i, expectedEntry := range expectedEntries {
		if i >= len(base.Entries) {
			t.Errorf("missing expected entry %d: %q", i, expectedEntry)
			continue
		}
		actualEntry := base.Entries[i]
		if actualEntry != expectedEntry {
			t.Errorf("unexpected entry %d: %q (expected %q)", i, actualEntry, expectedEntry)
		}
	}
	if len(base.Entries) > len(expectedEntries) {
		t.Errorf("additional unexpected entries: %v", base.Entries[len(expectedEntries):])
	}
}
