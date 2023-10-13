package carve

import "testing"

func TestCanonicalPath(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			name:   "",
			input:  "./example",
			expect: "example",
		},
		{
			name:   "",
			input:  "example",
			expect: "example",
		},
		{
			name:   "",
			input:  "./example/1",
			expect: "example/1",
		},
		{
			name:   "",
			input:  "/example",
			expect: "/example",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canonicalPath(tt.input)
			if got != tt.expect {
				t.Errorf("got %s want %s", got, tt.expect)
			}
		})
	}
}
