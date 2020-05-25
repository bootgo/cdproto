// Package memory provides the Chrome DevTools Protocol
// commands, types, and events for the Memory domain.
//
// Generated by the cdproto-gen command.
package memory

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/bootgo/cdproto/cdp"
)

// GetDOMCountersParams [no description].
type GetDOMCountersParams struct{}

// GetDOMCounters [no description].
func GetDOMCounters() *GetDOMCountersParams {
	return &GetDOMCountersParams{}
}

// GetDOMCountersReturns return values.
type GetDOMCountersReturns struct {
	Documents        int64 `json:"documents,omitempty"`
	Nodes            int64 `json:"nodes,omitempty"`
	JsEventListeners int64 `json:"jsEventListeners,omitempty"`
}

// Do executes Memory.getDOMCounters against the provided context.
//
// returns:
//   documents
//   nodes
//   jsEventListeners
func (p *GetDOMCountersParams) Do(ctxt context.Context, h cdp.Executor) (documents int64, nodes int64, jsEventListeners int64, err error) {
	// execute
	var res GetDOMCountersReturns
	err = h.Execute(ctxt, CommandGetDOMCounters, nil, &res)
	if err != nil {
		return 0, 0, 0, err
	}

	return res.Documents, res.Nodes, res.JsEventListeners, nil
}

// PrepareForLeakDetectionParams [no description].
type PrepareForLeakDetectionParams struct{}

// PrepareForLeakDetection [no description].
func PrepareForLeakDetection() *PrepareForLeakDetectionParams {
	return &PrepareForLeakDetectionParams{}
}

// Do executes Memory.prepareForLeakDetection against the provided context.
func (p *PrepareForLeakDetectionParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandPrepareForLeakDetection, nil, nil)
}

// SetPressureNotificationsSuppressedParams enable/disable suppressing memory
// pressure notifications in all processes.
type SetPressureNotificationsSuppressedParams struct {
	Suppressed bool `json:"suppressed"` // If true, memory pressure notifications will be suppressed.
}

// SetPressureNotificationsSuppressed enable/disable suppressing memory
// pressure notifications in all processes.
//
// parameters:
//   suppressed - If true, memory pressure notifications will be suppressed.
func SetPressureNotificationsSuppressed(suppressed bool) *SetPressureNotificationsSuppressedParams {
	return &SetPressureNotificationsSuppressedParams{
		Suppressed: suppressed,
	}
}

// Do executes Memory.setPressureNotificationsSuppressed against the provided context.
func (p *SetPressureNotificationsSuppressedParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetPressureNotificationsSuppressed, p, nil)
}

// SimulatePressureNotificationParams simulate a memory pressure notification
// in all processes.
type SimulatePressureNotificationParams struct {
	Level PressureLevel `json:"level"` // Memory pressure level of the notification.
}

// SimulatePressureNotification simulate a memory pressure notification in
// all processes.
//
// parameters:
//   level - Memory pressure level of the notification.
func SimulatePressureNotification(level PressureLevel) *SimulatePressureNotificationParams {
	return &SimulatePressureNotificationParams{
		Level: level,
	}
}

// Do executes Memory.simulatePressureNotification against the provided context.
func (p *SimulatePressureNotificationParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSimulatePressureNotification, p, nil)
}

// StartSamplingParams start collecting native memory profile.
type StartSamplingParams struct {
	SamplingInterval   int64 `json:"samplingInterval,omitempty"`   // Average number of bytes between samples.
	SuppressRandomness bool  `json:"suppressRandomness,omitempty"` // Do not randomize intervals between samples.
}

// StartSampling start collecting native memory profile.
//
// parameters:
func StartSampling() *StartSamplingParams {
	return &StartSamplingParams{}
}

// WithSamplingInterval average number of bytes between samples.
func (p StartSamplingParams) WithSamplingInterval(samplingInterval int64) *StartSamplingParams {
	p.SamplingInterval = samplingInterval
	return &p
}

// WithSuppressRandomness do not randomize intervals between samples.
func (p StartSamplingParams) WithSuppressRandomness(suppressRandomness bool) *StartSamplingParams {
	p.SuppressRandomness = suppressRandomness
	return &p
}

// Do executes Memory.startSampling against the provided context.
func (p *StartSamplingParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStartSampling, p, nil)
}

// StopSamplingParams stop collecting native memory profile.
type StopSamplingParams struct{}

// StopSampling stop collecting native memory profile.
func StopSampling() *StopSamplingParams {
	return &StopSamplingParams{}
}

// Do executes Memory.stopSampling against the provided context.
func (p *StopSamplingParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStopSampling, nil, nil)
}

// GetAllTimeSamplingProfileParams retrieve native memory allocations profile
// collected since renderer process startup.
type GetAllTimeSamplingProfileParams struct{}

// GetAllTimeSamplingProfile retrieve native memory allocations profile
// collected since renderer process startup.
func GetAllTimeSamplingProfile() *GetAllTimeSamplingProfileParams {
	return &GetAllTimeSamplingProfileParams{}
}

// GetAllTimeSamplingProfileReturns return values.
type GetAllTimeSamplingProfileReturns struct {
	Profile *SamplingProfile `json:"profile,omitempty"`
}

// Do executes Memory.getAllTimeSamplingProfile against the provided context.
//
// returns:
//   profile
func (p *GetAllTimeSamplingProfileParams) Do(ctxt context.Context, h cdp.Executor) (profile *SamplingProfile, err error) {
	// execute
	var res GetAllTimeSamplingProfileReturns
	err = h.Execute(ctxt, CommandGetAllTimeSamplingProfile, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Profile, nil
}

// GetBrowserSamplingProfileParams retrieve native memory allocations profile
// collected since browser process startup.
type GetBrowserSamplingProfileParams struct{}

// GetBrowserSamplingProfile retrieve native memory allocations profile
// collected since browser process startup.
func GetBrowserSamplingProfile() *GetBrowserSamplingProfileParams {
	return &GetBrowserSamplingProfileParams{}
}

// GetBrowserSamplingProfileReturns return values.
type GetBrowserSamplingProfileReturns struct {
	Profile *SamplingProfile `json:"profile,omitempty"`
}

// Do executes Memory.getBrowserSamplingProfile against the provided context.
//
// returns:
//   profile
func (p *GetBrowserSamplingProfileParams) Do(ctxt context.Context, h cdp.Executor) (profile *SamplingProfile, err error) {
	// execute
	var res GetBrowserSamplingProfileReturns
	err = h.Execute(ctxt, CommandGetBrowserSamplingProfile, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Profile, nil
}

// GetSamplingProfileParams retrieve native memory allocations profile
// collected since last startSampling call.
type GetSamplingProfileParams struct{}

// GetSamplingProfile retrieve native memory allocations profile collected
// since last startSampling call.
func GetSamplingProfile() *GetSamplingProfileParams {
	return &GetSamplingProfileParams{}
}

// GetSamplingProfileReturns return values.
type GetSamplingProfileReturns struct {
	Profile *SamplingProfile `json:"profile,omitempty"`
}

// Do executes Memory.getSamplingProfile against the provided context.
//
// returns:
//   profile
func (p *GetSamplingProfileParams) Do(ctxt context.Context, h cdp.Executor) (profile *SamplingProfile, err error) {
	// execute
	var res GetSamplingProfileReturns
	err = h.Execute(ctxt, CommandGetSamplingProfile, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Profile, nil
}

// Command names.
const (
	CommandGetDOMCounters                     = "Memory.getDOMCounters"
	CommandPrepareForLeakDetection            = "Memory.prepareForLeakDetection"
	CommandSetPressureNotificationsSuppressed = "Memory.setPressureNotificationsSuppressed"
	CommandSimulatePressureNotification       = "Memory.simulatePressureNotification"
	CommandStartSampling                      = "Memory.startSampling"
	CommandStopSampling                       = "Memory.stopSampling"
	CommandGetAllTimeSamplingProfile          = "Memory.getAllTimeSamplingProfile"
	CommandGetBrowserSamplingProfile          = "Memory.getBrowserSamplingProfile"
	CommandGetSamplingProfile                 = "Memory.getSamplingProfile"
)
