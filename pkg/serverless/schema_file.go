package serverless

import (
	"fmt"
	"io"

	"github.com/kyma-project/hydroform/function/pkg/workspace"
)

type Schema struct{}

func (s Schema) Write(w io.Writer, _ interface{}) error {
	b, err := workspace.ReflectSchema()
	if err != nil {
		return fmt.Errorf("failed to reflect schema: %w", err)
	}

	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("failed to write schema: %w", err)
	}

	return nil
}

func (s Schema) FileName() string {
	return "schema.json"
}
