package export_artifacts

import (
	"errors"
	"testing"
)

type exporterStub struct {
	err    error
	called bool
}

func (e *exporterStub) FindAndExport() error {
	e.called = true
	return e.err
}

func TestRun_DelegatesToExporter(t *testing.T) {
	// GIVEN an exporter runner
	stub := &exporterStub{}

	// WHEN running export
	err := Run(stub)

	// THEN it delegates to the exporter
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !stub.called {
		t.Fatal("expected exporter to be called")
	}
}

func TestRun_ReturnsExporterError(t *testing.T) {
	// GIVEN an exporter that fails
	stub := &exporterStub{err: errors.New("export failed")}

	// WHEN running export
	err := Run(stub)

	// THEN it returns the error
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "export failed" {
		t.Fatalf("expected error to be propagated, got %v", err)
	}
}
