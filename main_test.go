package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstDay(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		cutOff   int
		schedule string
		expected int
	}{
		{3, "Tuesday", 10},
		{3, "Thursday", 19},
		{3, "Monday", 16},
		{3, "Saturday", 14},
		{3, "Sunday", 15},
	}

	for _, test := range tests {
		scheduled := []string{test.schedule}
		operation := int(time.Date(2021, 8, getFirstDay(test.cutOff, scheduled).Day(), 0, 0, 0, 0, time.UTC).Weekday())
		result := int(time.Date(2021, 8, test.expected, 0, 0, 0, 0, time.UTC).Weekday())
		assert.Equal(operation, result)
	}
}

func TestGetFirstBiWeekly(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		cutOff   int
		schedule string
		expected int
	}{
		{3, "Tuesday", 10},
		{3, "Thursday", 19},
		{3, "Monday", 16},
		{3, "Saturday", 14},
		{3, "Sunday", 15},
	}

	for _, test := range tests {
		scheduled := []string{test.schedule}
		operation := int(time.Date(2021, 8, getFirstBiWeekly(test.cutOff, scheduled).Day(), 0, 0, 0, 0, time.UTC).Weekday())
		result := int(time.Date(2021, 8, test.expected, 0, 0, 0, 0, time.UTC).Weekday())
		assert.Equal(operation, result)
	}
}

func TestGetSecondBiWeekly(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		cutOff   int
		schedule string
		expected int
	}{
		{3, "Tuesday", 10},
		{3, "Thursday", 19},
		{3, "Monday", 16},
		{3, "Saturday", 14},
		{3, "Sunday", 15},
	}

	for _, test := range tests {
		scheduled := []string{"0", test.schedule}
		operation := int(time.Date(2021, 8, getSecondBiWeekly(test.cutOff, scheduled).Day(), 0, 0, 0, 0, time.UTC).Weekday())
		result := int(time.Date(2021, 8, test.expected, 0, 0, 0, 0, time.UTC).Weekday())
		assert.Equal(operation, result)
	}
}
