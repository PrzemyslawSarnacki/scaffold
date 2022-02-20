package scaffold

import "testing"

func Test_isStatic(t *testing.T) {
	tests := []struct {
		name string
		paths []string
		want bool
	}{
		{
			name: "testing positive",
			paths: []string{"path.yml", "path.gif", "path.png"},
			want: true,
		},
		{
			name: "testing negative",
			paths: []string{"path.go", "README.md", "go.mod"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, path := range tt.paths {
				if got := isStatic(path); got != tt.want {
					t.Errorf("isStatic() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
