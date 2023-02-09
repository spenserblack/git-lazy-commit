package fileutils

import "testing"

func TestSharedDirectory(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				paths: []string{},
			},
			want: "",
		},
		{
			name: "single",
			args: args{
				paths: []string{"foo/bar.txt"},
			},
			want: "foo",
		},
		{
			name: "multiple",
			args: args{
				paths: []string{"foo/bar.txt", "foo/baz.txt"},
			},
			want: "foo",
		},
		{
			name: "multiple different",
			args: args{
				paths: []string{"foo/bar.txt", "foo/baz.txt", "bar/baz.txt"},
			},
			want: "/",
		},
		{
			name: "multiple different with nested",
			args: args{
				paths: []string{"foo/bar.txt", "foo/baz.txt", "foo/bar/baz.txt"},
			},
			want: "foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SharedDirectory(tt.args.paths); got != tt.want {
				t.Errorf("SharedDirectory(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
