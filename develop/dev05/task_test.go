package main

import (
	"testing"
)

func Test_executeNoFlag(t *testing.T) {
	type args struct {
		grep *Grep
	}
	t1Rune := []rune{}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				&Grep{
					flags:    t1Rune,
					template: "ad",
					fileName: "test3.txt",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executeNoFlag(tt.args.grep)
		})
	}
}

func Test_createContent(t *testing.T) {
	type args struct {
		grep *Grep
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				&Grep{
					fileName: "test3.txt",
					content:  map[int]string{},
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				&Grep{
					fileName: "NO.txt",
					content:  map[int]string{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createContent(tt.args.grep); (err != nil) != tt.wantErr {
				t.Errorf("createContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseFlags(t *testing.T) {
	type args struct {
		flags []rune
	}
	t1 := []rune{'-', 'A', '2'}
	t2 := []rune{'-', 'R'}
	t3 := []rune{'C'}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				flags: t1,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				flags: t2,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				flags: t3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseFlags(tt.args.flags); got != tt.want {
				t.Errorf("parseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
