// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &ToCidrFunction{}

type ToCidrFunction struct{}

func NewToCidrFunction() function.Function {
	return &ToCidrFunction{}
}

func (f *ToCidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "tocidr"
}

func (f *ToCidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "`tocidr` converts subnet and netmask to the cidr format.",
		Description: "Given subnet and netmask, return subnet and netmask in the cidr format. Fails for invalid subnet or netmasks.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "subnet",
				Description: "Network subnet",
			},
			function.StringParameter{
				Name:        "netmask",
				Description: "Network mask",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *ToCidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var subnet string
	var netmask string
	var response string

	// Read Terraform argument data into the variables
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &subnet, &netmask))

	ip := net.ParseIP(subnet)
	ip4 := ip.To4()
	if ip == nil || ip4 == nil {
		resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("not an IPv4 address/invalid address: %s", subnet))
		return
	}

	maskIP := net.ParseIP(netmask)
	mask4 := maskIP.To4()
	if maskIP == nil || mask4 == nil {
		resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("not an IPv4 netmask/invalid netmask: %q", netmask))
		return
	}

	mask := net.IPMask(mask4)
	ones, bits := mask.Size()
	if bits != 32 || ones < 0 {
		resp.Error = function.NewArgumentFuncError(1, "invalid netmask (non-contiguous or unsupported")
		return
	}

	network := net.IP(make([]byte, 4))
	for i := 0; i < 4; i++ {
		network[i] = ip4[i] & mask[i]
	}

	response = fmt.Sprintf("%s/%d", network.String(), ones)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, response))
}
