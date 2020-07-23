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

package opsmanager

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

type LogsJobsListOpts struct {
	cli.GlobalOpts
	verbose bool
	store   store.LogJobLister
}

func (opts *LogsJobsListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *LogsJobsListOpts) Run() error {
	r, err := opts.store.LogCollectionJobs(opts.ConfigProjectID(), opts.newLogListOptions())
	if err != nil {
		return err
	}
	return output.Print(config.Default(), "", r)
}

func (opts *LogsJobsListOpts) newLogListOptions() *opsmngr.LogListOptions {
	return &opsmngr.LogListOptions{Verbose: opts.verbose}
}

// mongocli om logs jobs list --verbose verbose [--projectId projectId]
func LogsJobsListOptsBuilder() *cobra.Command {
	opts := &LogsJobsListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   description.ListLogCollectionJobs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.verbose, flag.Verbose, false, usage.Verbose)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
