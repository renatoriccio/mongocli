// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iam

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type OrganizationsDeleteOpts struct {
	*cli.DeleteOpts
	store store.OrganizationDeleter
}

func (opts *OrganizationsDeleteOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *OrganizationsDeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteOrganization)
}

// mongocli iam organization(s) delete [id] [--force]
func OrganizationsDeleteBuilder() *cobra.Command {
	opts := &OrganizationsDeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Organization '%s' deleted\n", "Organization not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete [ID]",
		Short: description.DeleteOrganization,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.init(); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}