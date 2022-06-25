package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	for idx := 0; idx < b.N; idx++ {
		AlienOrder(words)
	}
}
func TestAlienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "[wrt,wrf,er,ett,rftt]",
			args: args{words: []string{"wrt", "wrf", "er", "ett", "rftt"}},
			want: "wertf",
		},
		{
			name: "[z,x]",
			args: args{words: []string{"z", "x"}},
			want: "zx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlienOrder(tt.args.words); got != tt.want {
				t.Errorf("AlienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
