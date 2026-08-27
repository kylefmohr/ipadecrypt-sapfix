// Package mescal signs Apple Store action requests with the SAP signing
// service that Apple exposes through macOS' private CommerceKit framework.
// The private auth endpoint rejects login requests that don't carry a valid
// X-Apple-ActionSignature header.
package mescal

import "errors"

// ErrUnavailable indicates that the host cannot create Apple's SAP action
// signatures. Apple currently provides the required signing service through a
// private macOS framework, so it is only available on macOS builds compiled
// with cgo.
var ErrUnavailable = errors.New("Apple SAP signing is unavailable")
