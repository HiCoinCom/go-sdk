// Package utils provides utility functions and constants for ChainUp Custody SDK.
package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// -----------------------------------------------------------------------------
// Common Types (shared between Custody and MPC)
// -----------------------------------------------------------------------------

// FlexInt is a custom type that can unmarshal from both int and string in JSON.
// When unmarshaling from string, empty string becomes 0, otherwise it must be a valid integer.
type FlexInt int64

// MarshalJSON implements json.Marshaler.
func (f FlexInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(f))
}

// UnmarshalJSON implements json.Unmarshaler.
// Supports both int and string formats. Empty string becomes 0.
func (f *FlexInt) UnmarshalJSON(data []byte) error {
	// Try int first
	var intVal int64
	if err := json.Unmarshal(data, &intVal); err == nil {
		*f = FlexInt(intVal)
		return nil
	}

	// Try string
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return fmt.Errorf("FlexInt: cannot unmarshal %s into int64 or string", string(data))
	}

	// Empty string becomes 0
	if strVal == "" {
		*f = 0
		return nil
	}

	// Parse string as int
	intVal, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		return fmt.Errorf("FlexInt: cannot parse string %q as int64: %w", strVal, err)
	}
	*f = FlexInt(intVal)
	return nil
}

// Int64 returns the value as int64.
func (f FlexInt) Int64() int64 {
	return int64(f)
}

// String returns the value as string.
func (f FlexInt) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// Timestamp is a custom type for handling Unix timestamp JSON serialization.
type Timestamp struct {
	time.Time
}

// MarshalJSON implements json.Marshaler.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("0"), nil
	}
	return json.Marshal(t.UnixMilli())
}

// UnmarshalJSON implements json.Unmarshaler.
// Supports both int64 (milliseconds) and string formats.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	// Try int64 first (milliseconds timestamp)
	var ms int64
	if err := json.Unmarshal(data, &ms); err == nil {
		if ms > 0 {
			t.Time = time.UnixMilli(ms)
		}
		return nil
	}

	// Try string format
	var strVal string
	if err := json.Unmarshal(data, &strVal); err != nil {
		return fmt.Errorf("Timestamp: cannot unmarshal %s into int64 or string", string(data))
	}

	// Empty string becomes zero time
	if strVal == "" {
		t.Time = time.Time{}
		return nil
	}

	// Try to parse as numeric string (milliseconds)
	if ms, err := strconv.ParseInt(strVal, 10, 64); err == nil {
		if ms > 0 {
			t.Time = time.UnixMilli(ms)
		}
		return nil
	}

	// Try common date formats
	formats := []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}
	for _, format := range formats {
		if parsed, err := time.Parse(format, strVal); err == nil {
			t.Time = parsed
			return nil
		}
	}

	return fmt.Errorf("Timestamp: cannot parse %q as timestamp", strVal)
}
