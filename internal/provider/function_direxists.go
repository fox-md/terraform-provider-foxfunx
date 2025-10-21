// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &DirExistsFunction{}

type DirExistsFunction struct{}

func NewDirExistsFunction() function.Function {
	return &DirExistsFunction{}
}

func (f *DirExistsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "direxists"
}

func (f *DirExistsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "`direxists` determines whether a directory exists at a given path.",
		Description: "Given a path, return boolean depending on directory existence. Fails for files.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "path",
				Description: "Path to directory.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *DirExistsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var path string
	var response bool

	// Read Terraform argument data into the variables
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &path))

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		response = false
	} else if err != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to verify path. Error: %s", err.Error()))
		resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("Failed to check '%s' path. Error: %s", path, err.Error()))
		return
	} else if info.IsDir() {
		response = true
	} else {
		resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("'%s' is a file, not a directory", path))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, response))
}
