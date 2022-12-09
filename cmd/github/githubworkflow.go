// Package github for gh tasks
package github

import (
	"context"

	"github.com/jmcconnell26/nolte-plumbing/pkg"
	"github.com/magefile/mage/mg"
)

// GitHubWorkflow Mage Command Namespace.
type GitHubWorkflow mg.Namespace

// StartJob Github Workflow
func (GitHubWorkflow) StartJob(ctx context.Context) error {
	jobName, ok := ctx.Value("jobName").(string)
	if !ok {
		jobName = "build"
	}
	job := pkg.ActJob{
		Name: jobName,
	}
	return job.Execute()
}
