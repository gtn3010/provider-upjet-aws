// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package kms

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the kms group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_kms_alias", func(r *config.Resource) {
		r.References["target_key_id"] = config.Reference{
			Type: "Key",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_kms_ciphertext", func(r *config.Resource) {
		r.References["key_id"] = config.Reference{
			Type: "Key",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_kms_grant", func(r *config.Resource) {
		r.References["key_id"] = config.Reference{
			Type:      "Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_kms_replica_key", func(r *config.Resource) {
		r.References["primary_key_arn"] = config.Reference{
			Type:      "Key",
			Extractor: common.PathARNExtractor,
		}
	})

	p.AddResourceConfigurator("aws_kms_replica_external_key", func(r *config.Resource) {
		r.References["primary_key_arn"] = config.Reference{
			Type:      "ExternalKey",
			Extractor: common.PathARNExtractor,
		}
	})
}
