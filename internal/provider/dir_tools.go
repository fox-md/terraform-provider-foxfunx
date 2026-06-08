// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func isDir(ctx context.Context, path string) (bool, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("'%s' path does not exist", path)
	} else if err != nil {
		tflog.Error(ctx, fmt.Sprintf("failed to verify path. Error: %s", err.Error()))
		return false, fmt.Errorf("failed to check '%s' path. Error: %s", path, err.Error())
	} else if info.IsDir() {
		return true, nil
	} else {
		return false, fmt.Errorf("'%s' is not a directory", path)
	}
}

func isDirEmpty(path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}
	return len(entries) == 0, nil
}
