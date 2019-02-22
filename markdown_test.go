package main

import "testing"

func Test_markdownToHTML(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := markdownToHTML(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("markdownToHTML() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
