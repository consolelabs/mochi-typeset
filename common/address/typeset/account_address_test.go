package typeset

import "testing"

func Test_evmCheck(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid evm",
			args: args{
				hash: "0x8D565bd9596AC3d1e8662Bd5661085b95a5D4F99",
			},
			want: true,
		},
		{
			name: "invalid evm, too long",
			args: args{
				hash: "0x8D565bd9596AC3d1e8662Bd5661085b95a5D4F99123",
			},
			want: false,
		},
		{
			name: "invalid evm, missing 0x",
			args: args{
				hash: "8D565bd9596AC3d1e8662Bd5661085b95a5D4F99123",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evmCheck(tt.args.hash); got != tt.want {
				t.Errorf("evmCheck(): %s, = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func Test_roninCheck(t *testing.T) {
	type args struct {
		hash string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid ronin",
			args: args{
				hash: "ronin:8D565bd9596AC3d1e8662Bd5661085b95a5D4F99",
			},
			want: true,
		},
		{
			name: "invalid ronin, without prefix",
			args: args{
				hash: "8D565bd9596AC3d1e8662Bd5661085b95a5D4F99123",
			},
			want: false,
		},
		{
			name: "invalid ronin, too short",
			args: args{
				hash: "ronin:8D565bd9596AC3d1e8662Bd5661085b95a5D4F9912",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roninCheck(tt.args.hash); got != tt.want {
				t.Errorf("roninCheck(): %s, = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
