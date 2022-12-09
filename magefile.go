//go:build mage
// +build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"

	// mage:import
	"github.com/jmcconnell26/nolte-plumbing/cmd/golang"
	"github.com/jmcconnell26/nolte-plumbing/pkg"
)

// GitHubWorkflow Mage Command Namespace.
type GitHubWorkflow mg.Namespace

// All Targets in a pipeline.
func All(ctx context.Context) {
	mg.CtxDeps(ctx, golang.Golang.StaticTests)
	mg.SerialCtxDeps(ctx, GitHubWorkflow.GH)
}

// GH will be run all GitHub Actions Locally.
func (GitHubWorkflow) GH(ctx context.Context) {
	mg.SerialCtxDeps(ctx, GitHubWorkflow.GHBuild)
	mg.SerialCtxDeps(ctx, GitHubWorkflow.GHLint)
}

// GHLint start the Github lint Workflow.
func (GitHubWorkflow) GHLint(ctx context.Context) error {
	job := pkg.ActJob{
		Name: "lint",
	}
	return job.Execute()
}

// GHBuild start the Github build Workflow.
func (GitHubWorkflow) GHBuild(ctx context.Context) error {
	job := pkg.ActJob{
		Name: "build",
	}
	return job.Execute()
}
