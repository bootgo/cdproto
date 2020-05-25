// Package testing provides the Chrome DevTools Protocol
// commands, types, and events for the Testing domain.
//
// Testing domain is a dumping ground for the capabilities requires for
// browser or app testing that do not fit other domains.
//
// Generated by the cdproto-gen command.
package testing

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/bootgo/cdproto/cdp"
)

// GenerateTestReportParams generates a report for testing.
type GenerateTestReportParams struct {
	Message string `json:"message"`         // Message to be displayed in the report.
	Group   string `json:"group,omitempty"` // Specifies the endpoint group to deliver the report to.
}

// GenerateTestReport generates a report for testing.
//
// parameters:
//   message - Message to be displayed in the report.
func GenerateTestReport(message string) *GenerateTestReportParams {
	return &GenerateTestReportParams{
		Message: message,
	}
}

// WithGroup specifies the endpoint group to deliver the report to.
func (p GenerateTestReportParams) WithGroup(group string) *GenerateTestReportParams {
	p.Group = group
	return &p
}

// Do executes Testing.generateTestReport against the provided context.
func (p *GenerateTestReportParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandGenerateTestReport, p, nil)
}

// Command names.
const (
	CommandGenerateTestReport = "Testing.generateTestReport"
)
