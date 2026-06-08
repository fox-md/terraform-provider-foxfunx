// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ function.Function = &DirEmptyFunction{}

type DirEmptyFunction struct{}

func NewDirEmptyFunction() function.Function {
	return &DirEmptyFunction{}
}

func (f *DirEmptyFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "dirempty"
}

func (f *DirEmptyFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "`dirempty` determines whether a directory is empty.",
		Description: "Given a path, return boolean depending on directory content. Fails for files and non-existing directories.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "path",
				Description: "Path to directory.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *DirEmptyFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var path string
	var response bool

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &path))

	_, err := isDir(ctx, path)
	if err != nil {
		resp.Error = function.NewArgumentFuncError(0, err.Error())
		return
	}

	response, err = isDirEmpty(path)
	if err != nil {
		tflog.Error(ctx, err.Error())
		resp.Error = function.NewArgumentFuncError(0, err.Error())
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, response))
}
