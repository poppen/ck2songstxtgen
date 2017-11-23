package main

import (
	"testing"
)

var oggtests = []struct {
	in  song
	out song
}{
	{
		song{
			"internal/ogg_test/example.ogg",
			"",
			1,
			0.48,
		},
		song{
			"",
			"OGG Test File",
			1,
			0.48,
		},
	},
	{
		song{
			"internal/ogg_test/example_without_title.ogg",
			"",
			1,
			0.48,
		},
		song{
			"",
			"example_without_title",
			1,
			0.48,
		},
	},
}

func TestTitle(t *testing.T) {
	for _, tt := range oggtests {
		if s := title(tt.in.name); s != tt.out.songName {
			t.Errorf("songName(%q) => %q, want: %q", tt.in.name, s, tt.out.songName)
		}
	}
}

func TestEntry(t *testing.T) {
}

func TestPrintEntry(t *testing.T) {
}
