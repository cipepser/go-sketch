package util

import "testing"

func TestCalcMD5Hash(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcMD5Hash(tt.args.str); got != tt.want {
				t.Errorf("CalcMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleHashing(t *testing.T) {
	type args struct {
		hashA uint64
		hashB uint64
		n     uint64
		r     uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoubleHashing(tt.args.hashA, tt.args.hashB, tt.args.n, tt.args.r); got != tt.want {
				t.Errorf("DoubleHashing() = %v, want %v", got, tt.want)
			}
		})
	}
}