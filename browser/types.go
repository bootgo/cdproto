package browser

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// WindowID [no description].
type WindowID int64

// Int64 returns the WindowID as int64 value.
func (t WindowID) Int64() int64 {
	return int64(t)
}

// WindowState the state of the browser window.
type WindowState string

// String returns the WindowState as string value.
func (t WindowState) String() string {
	return string(t)
}

// WindowState values.
const (
	WindowStateNormal     WindowState = "normal"
	WindowStateMinimized  WindowState = "minimized"
	WindowStateMaximized  WindowState = "maximized"
	WindowStateFullscreen WindowState = "fullscreen"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t WindowState) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t WindowState) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *WindowState) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch WindowState(in.String()) {
	case WindowStateNormal:
		*t = WindowStateNormal
	case WindowStateMinimized:
		*t = WindowStateMinimized
	case WindowStateMaximized:
		*t = WindowStateMaximized
	case WindowStateFullscreen:
		*t = WindowStateFullscreen

	default:
		in.AddError(errors.New("unknown WindowState value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *WindowState) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Bounds browser window bounds information.
type Bounds struct {
	Left        int64       `json:"left,omitempty"`        // The offset from the left edge of the screen to the window in pixels.
	Top         int64       `json:"top,omitempty"`         // The offset from the top edge of the screen to the window in pixels.
	Width       int64       `json:"width,omitempty"`       // The window width in pixels.
	Height      int64       `json:"height,omitempty"`      // The window height in pixels.
	WindowState WindowState `json:"windowState,omitempty"` // The window state. Default to normal.
}

// PermissionType [no description].
type PermissionType string

// String returns the PermissionType as string value.
func (t PermissionType) String() string {
	return string(t)
}

// PermissionType values.
const (
	PermissionTypeAccessibilityEvents      PermissionType = "accessibilityEvents"
	PermissionTypeAudioCapture             PermissionType = "audioCapture"
	PermissionTypeBackgroundSync           PermissionType = "backgroundSync"
	PermissionTypeClipboardRead            PermissionType = "clipboardRead"
	PermissionTypeClipboardWrite           PermissionType = "clipboardWrite"
	PermissionTypeDurableStorage           PermissionType = "durableStorage"
	PermissionTypeFlash                    PermissionType = "flash"
	PermissionTypeGeolocation              PermissionType = "geolocation"
	PermissionTypeMidi                     PermissionType = "midi"
	PermissionTypeMidiSysex                PermissionType = "midiSysex"
	PermissionTypeNotifications            PermissionType = "notifications"
	PermissionTypePaymentHandler           PermissionType = "paymentHandler"
	PermissionTypeProtectedMediaIdentifier PermissionType = "protectedMediaIdentifier"
	PermissionTypeSensors                  PermissionType = "sensors"
	PermissionTypeVideoCapture             PermissionType = "videoCapture"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t PermissionType) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t PermissionType) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *PermissionType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch PermissionType(in.String()) {
	case PermissionTypeAccessibilityEvents:
		*t = PermissionTypeAccessibilityEvents
	case PermissionTypeAudioCapture:
		*t = PermissionTypeAudioCapture
	case PermissionTypeBackgroundSync:
		*t = PermissionTypeBackgroundSync
	case PermissionTypeClipboardRead:
		*t = PermissionTypeClipboardRead
	case PermissionTypeClipboardWrite:
		*t = PermissionTypeClipboardWrite
	case PermissionTypeDurableStorage:
		*t = PermissionTypeDurableStorage
	case PermissionTypeFlash:
		*t = PermissionTypeFlash
	case PermissionTypeGeolocation:
		*t = PermissionTypeGeolocation
	case PermissionTypeMidi:
		*t = PermissionTypeMidi
	case PermissionTypeMidiSysex:
		*t = PermissionTypeMidiSysex
	case PermissionTypeNotifications:
		*t = PermissionTypeNotifications
	case PermissionTypePaymentHandler:
		*t = PermissionTypePaymentHandler
	case PermissionTypeProtectedMediaIdentifier:
		*t = PermissionTypeProtectedMediaIdentifier
	case PermissionTypeSensors:
		*t = PermissionTypeSensors
	case PermissionTypeVideoCapture:
		*t = PermissionTypeVideoCapture

	default:
		in.AddError(errors.New("unknown PermissionType value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *PermissionType) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// Bucket chrome histogram bucket.
type Bucket struct {
	Low   int64 `json:"low"`   // Minimum value (inclusive).
	High  int64 `json:"high"`  // Maximum value (exclusive).
	Count int64 `json:"count"` // Number of samples.
}

// Histogram chrome histogram.
type Histogram struct {
	Name    string    `json:"name"`    // Name.
	Sum     int64     `json:"sum"`     // Sum of sample values.
	Count   int64     `json:"count"`   // Total number of samples.
	Buckets []*Bucket `json:"buckets"` // Buckets.
}
