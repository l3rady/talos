// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package components implements specific widgets for the dashboard.
package components

import (
	"fmt"
	"strings"

	"github.com/siderolabs/gen/slices"
)

const (
	noData       = "..."
	notAvailable = "n/a"
	none         = "<none>"
	maxLogLines  = 1000
)

// field represents a field in a widget consist of a name and a value, rendered next to each other.
type field struct {
	Name  string
	Value string
}

func (f *field) render(nameWidth int) string {
	return fmt.Sprintf("[::b]%s[::-] %s", padRight(f.Name, nameWidth), f.Value)
}

type fieldGroup struct {
	fields []field
}

// String implements the Stringer interface.
func (fg *fieldGroup) String() string {
	width := fg.maxFieldNameLength()

	return strings.Join(
		slices.Map(fg.fields, func(t field) string {
			return t.render(width)
		}),
		"\n",
	)
}

func (fg *fieldGroup) maxFieldNameLength() int {
	max := 0

	for _, f := range fg.fields {
		if len(f.Name) > max {
			max = len(f.Name)
		}
	}

	return max
}

// padRight pads a string to the specified width by appending spaces to the end.
func padRight(s string, width int) string {
	return fmt.Sprintf("%-*s", width, s)
}

func toHealthStatus(healthy bool) string {
	if healthy {
		return formatStatus("Healthy")
	}

	return formatStatus("Unhealthy")
}

func formatStatus(status any) string {
	statusStr := capitalizeFirst(fmt.Sprintf("%v", status))

	switch strings.ToLower(statusStr) {
	case "running", "healthy", "true":
		return fmt.Sprintf("[green]%s[-]", statusStr)
	case "stopped", "unhealthy", "false":
		return fmt.Sprintf("[red]%s[-]", statusStr)
	default:
		return statusStr
	}
}

// capitalizeFirst capitalizes the first character of string.
func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}

	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
