package markdown

import (
	"html/template"
	"reflect"
	"testing"
)

func TestDefaultProcessor_Markdown(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		p    *DefaultProcessor
		args args
		want template.HTML
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DefaultProcessor{}
			if got := p.Markdown(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProcessor.Markdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
