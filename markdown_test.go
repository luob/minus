package main

import "testing"

func Test_markdownToHTML(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "strongAndEm",
			args: args{
				input: "***dsadasdad***",
			},
			want: "<strong><em>***dsadasdad***</em></strong>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := markdownToHTML(tt.args.input); got != tt.want {
				t.Errorf("markdownToHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
