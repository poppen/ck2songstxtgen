package main

import (
	"testing"
)

type musicInfo struct {
	file  string
	title string
}

var oggtests = []struct {
	in  musicInfo
	out musicInfo
}{
	{
		musicInfo{
			"internal/ogg_test/example.ogg",
			"",
		},
		musicInfo{
			"",
			"OGG Test File",
		},
	},
	{
		musicInfo{
			"internal/ogg_test/example_without_title.ogg",
			"",
		},
		musicInfo{
			"",
			"example_without_title",
		},
	},
}

func TestTitle(t *testing.T) {
	for _, tt := range oggtests {
		if s := title(tt.in.file); s != tt.out.title {
			t.Errorf("title(%q) => %q, want: %q", tt.in.file, s, tt.out.title)
		}
	}
}
