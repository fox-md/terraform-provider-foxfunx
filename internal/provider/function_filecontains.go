// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ function.Function = &FileContainsFunction{}

type FileContainsFunction struct{}

func NewFileContainsFunction() function.Function {
	return &FileContainsFunction{}
}

func (f *FileContainsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "filecontains"
}

func (f *FileContainsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "`filecontains` determines whether search string exists in the file content.",
		Description: "Given a path, return boolean depending on whether the file content matches the search string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "file_path",
				Description: "Path to the file.",
				Validators: []function.StringParameterValidator{
					stringvalidator.LengthAtLeast(1),
				},
			},
			function.StringParameter{
				Name:        "search_text",
				Description: "Text to search in the file content.",
				Validators: []function.StringParameterValidator{
					stringvalidator.LengthAtLeast(1),
				},
			},
		},
		VariadicParameter: function.BoolParameter{
			Name:        "match_case",
			Description: "Match search text case. Default to `true`. No need to pass, unless you want to set it to `false`.",
		},
		Return: function.BoolReturn{},
	}
}

func (f *FileContainsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var file_path string
	var search_text string
	var match_case_ref []types.Bool

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &file_path, &search_text, &match_case_ref))

	match_case := true

	if len(match_case_ref) > 0 {
		if !match_case_ref[0].IsNull() && !match_case_ref[0].IsUnknown() {
			match_case = match_case_ref[0].ValueBool()
		}
	}

	response, err := fileContains(file_path, search_text, match_case)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to search for text in the '%s' file. Error: %s", file_path, err.Error()))
		resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("Failed to search for text in the '%s' path. Error: %s", file_path, err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, response))
}

func fileContains(filePath, target string, match_case bool) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	switch match_case {
	case true:
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), target) {
				return true, nil
			}
		}
	case false:
		for scanner.Scan() {
			if strings.EqualFold(scanner.Text(), target) {
				return true, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	return false, nil
}
