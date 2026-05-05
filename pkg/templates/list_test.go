package templates

import "testing"

func TestListEmbeddedTemplates(t *testing.T) {
	m, err := ListEmbeddedTemplates()
	if err != nil {
		t.Fatalf("ListEmbeddedTemplates returned error: %v", err)
	}
	if len(m) == 0 {
		t.Fatal("expected at least one architecture, got none")
	}
	if variants, ok := m["default"]; !ok || len(variants) == 0 {
		t.Fatalf("expected 'default' to exist with at least one variant, got: %v", variants)
	}
	// Warn if mvc/ddd missing but don't fail to avoid brittle tests
	if _, ok := m["mvc"]; !ok {
		t.Log("warning: 'mvc' not present in embedded templates")
	}
	if _, ok := m["ddd"]; !ok {
		t.Log("warning: 'ddd' not present in embedded templates")
	}
}
