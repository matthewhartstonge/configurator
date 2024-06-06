package hcl

import (
	"os"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/matthewhartstonge/configurator"
)

var _ configurator.ConfigTypeable = (*HCL)(nil)

func New(config configurator.ConfigImplementer) *HCL {
	h := &HCL{}
	h.ConfigFileType = configurator.NewConfigFileType(
		config,
		[]string{"hcl"},
		unmarshal(h),
	)

	return h
}

type HCL struct {
	configurator.ConfigFileType
}

func (h HCL) Type() string {
	return "HCL Configurator"
}

// unmarshal is a helper function that returns a Unmarshaler for HCL files.
func unmarshal(h *HCL) configurator.Unmarshaler {
	return func(data []byte, v interface{}) error {
		src, err := os.ReadFile(h.Path)
		if err != nil {
			return err
		}

		file, diags := hclsyntax.ParseConfig(src, h.Path, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			return diags
		}

		diags = gohcl.DecodeBody(file.Body, nil, v)
		if diags.HasErrors() {
			return diags
		}

		return nil
	}
}
