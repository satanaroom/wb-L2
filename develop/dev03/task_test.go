package main

import (
	"reflect"
	"testing"
)

func Test_parseFlag(t *testing.T) {
	type args struct {
		flags string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				"-rk1M",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				"k2u",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				"-t2u",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				"-k2uy",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFlag(tt.args.flags); got != tt.want {
				t.Errorf("parseFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortK(t *testing.T) {
	type args struct {
		lines []string
	}
	t1 := []string{"Vak", "mak", "dak", "tak"}
	t1Res := []string{"Vak", "dak", "mak", "tak"}
	t2 := []string{"xak", "wak", "aak", "pak"}
	t2Res := []string{"aak", "pak", "wak", "xak"}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				lines: t1,
			},
			want: t1Res,
		},
		{
			name: "2",
			args: args{
				lines: t2,
			},
			want: t2Res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortK(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortN(t *testing.T) {
	type args struct {
		lines []string
	}
	t1 := []string{"56", "1", "12", "89"}
	t1Res := []string{"1", "12", "56", "89"}
	t2 := []string{"-2", "-1", "0", "5"}
	t2Res := []string{"-2", "-1", "0", "5"}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				lines: t1,
			},
			want: t1Res,
		},
		{
			name: "2",
			args: args{
				lines: t2,
			},
			want: t2Res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortN(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortR(t *testing.T) {
	type args struct {
		lines []string
	}
	t1 := []string{"Vak", "mak", "dak", "tak"}
	t1Res := []string{"tak", "mak", "dak", "Vak"}
	t2 := []string{"xak", "wak", "aak", "pak"}
	t2Res := []string{"xak", "wak", "pak", "aak"}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				lines: t1,
			},
			want: t1Res,
		},
		{
			name: "2",
			args: args{
				lines: t2,
			},
			want: t2Res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortR(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		str []string
	}
	t1 := []string{"Vak", "mak", "dak", "tak", "mak"}
	t1Res := []string{"Vak", "mak", "dak", "tak"}
	t2 := []string{"xak", "wak", "aak", "pak", "xak", "aak"}
	t2Res := []string{"xak", "wak", "aak", "pak"}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				str: t1,
			},
			want: t1Res,
		},
		{
			name: "2",
			args: args{
				str: t2,
			},
			want: t2Res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNoFlag(t *testing.T) {
	type args struct {
		fileName string
	}
	t1 := []byte("hello\r\n")
	t2 := []byte("Vak\r\nmak\r\ndak\r\ntak\r\n")
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 bool
	}{
		{
			name: "1",
			args: args{
				"test4.txt",
			},
			want: t1,
			want1: true,
		},
		{
			name: "2",
			args: args{
				"test3.txt",
			},
			want: t2,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := checkNoFlag(tt.args.fileName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("checkNoFlag() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("checkNoFlag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
