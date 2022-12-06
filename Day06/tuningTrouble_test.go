package main

import "testing"

func Test_dupsFound(t *testing.T) {
	type args struct {
		chars []string
	}
	tests := []struct {
		name     string
		args     args
		wantDups bool
	}{
		{
			"No dups",
			args{[]string{"a", "b", "c", "d"}},
			false,
		},
		{
			"Dups",
			args{[]string{"a", "b", "c", "a"}},
			true,
		},
		{
			"Dups immediately",
			args{[]string{"a", "a", "c", "a"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDups := dupsFound(tt.args.chars); gotDups != tt.wantDups {
				t.Errorf("dupsFound() = %v, want %v", gotDups, tt.wantDups)
			}
		})
	}
}

func Test_start(t *testing.T) {
	tests := []struct {
		name         string
		signal       string
		wantPosition int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", "nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPosition := start(packetLength, tt.signal); gotPosition != tt.wantPosition {
				t.Errorf("start() = %v, want %v", gotPosition, tt.wantPosition)
			}
		})
	}
}
