// Package types provides tests for custom types
package types

import (
	"encoding/json"
	"testing"
)

func TestFlexInt_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int64
		wantErr  bool
	}{
		{
			name:     "Integer value",
			input:    `{"uid": 12345}`,
			expected: 12345,
			wantErr:  false,
		},
		{
			name:     "String value",
			input:    `{"uid": "67890"}`,
			expected: 67890,
			wantErr:  false,
		},
		{
			name:     "Empty string",
			input:    `{"uid": ""}`,
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Zero integer",
			input:    `{"uid": 0}`,
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Zero string",
			input:    `{"uid": "0"}`,
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "Large integer",
			input:    `{"uid": 9223372036854775807}`,
			expected: 9223372036854775807,
			wantErr:  false,
		},
		{
			name:     "Large string",
			input:    `{"uid": "9223372036854775807"}`,
			expected: 9223372036854775807,
			wantErr:  false,
		},
		{
			name:     "Negative integer",
			input:    `{"uid": -100}`,
			expected: -100,
			wantErr:  false,
		},
		{
			name:     "Negative string",
			input:    `{"uid": "-100"}`,
			expected: -100,
			wantErr:  false,
		},
		{
			name:     "Invalid string",
			input:    `{"uid": "abc"}`,
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Float string should fail",
			input:    `{"uid": "123.45"}`,
			expected: 0,
			wantErr:  true,
		},
	}

	type testStruct struct {
		UID FlexInt `json:"uid"`
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result testStruct
			err := json.Unmarshal([]byte(tc.input), &result)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if result.UID.Int64() != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result.UID.Int64())
			}
		})
	}
}

func TestFlexInt_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		value    FlexInt
		expected string
	}{
		{
			name:     "Positive value",
			value:    FlexInt(12345),
			expected: `{"uid":12345}`,
		},
		{
			name:     "Zero value",
			value:    FlexInt(0),
			expected: `{"uid":0}`,
		},
		{
			name:     "Negative value",
			value:    FlexInt(-100),
			expected: `{"uid":-100}`,
		},
	}

	type testStruct struct {
		UID FlexInt `json:"uid"`
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := testStruct{UID: tc.value}
			data, err := json.Marshal(s)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if string(data) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, string(data))
			}
		})
	}
}

func TestFlexInt_String(t *testing.T) {
	f := FlexInt(12345)
	if f.String() != "12345" {
		t.Errorf("Expected '12345', got '%s'", f.String())
	}
}

func TestFlexInt_UserInfo(t *testing.T) {
	// Test with UserInfo struct
	testCases := []struct {
		name     string
		input    string
		expected int64
	}{
		{
			name:     "UID as integer",
			input:    `{"uid": 15036904, "nickname": "test_user"}`,
			expected: 15036904,
		},
		{
			name:     "UID as string",
			input:    `{"uid": "15036904", "nickname": "test_user"}`,
			expected: 15036904,
		},
		{
			name:     "UID as empty string",
			input:    `{"uid": "", "nickname": ""}`,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var user UserInfo
			err := json.Unmarshal([]byte(tc.input), &user)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if user.UID.Int64() != tc.expected {
				t.Errorf("Expected UID %d, got %d", tc.expected, user.UID.Int64())
			}
		})
	}
}
