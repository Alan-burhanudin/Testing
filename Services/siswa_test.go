package siswa

import "testing"

func TestGrade(t *testing.T) {
	type args struct {
		Nilai int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Nilai C",
			args: args{Nilai: 40},
			want: "C",
		},
		{
			name: "Test Nilai B",
			args: args{Nilai: 65},
			want: "B",
		}, {
			name: "Test Nilai A",
			args: args{Nilai: 100},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Grade(tt.args.Nilai); got != tt.want {
				t.Errorf("Grade() = %v, want %v", got, tt.want)
			}
		})
	}
}
