package main

import "testing"

func TestUnpackString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				"a4bc2d5e",
			},
			want: "aaaabccddddde",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				"a4b0c2d5e",
			},
			want: "aaaaccddddde",
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				"abcd",
			},
			want: "abcd",
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				"45",
			},
			want: "",
			wantErr: true,
		},
		{
			name: "5",
			args: args{
				"",
			},
			want: "",
			wantErr: false,
		},
		{
			name: "6",
			args: args{
				`qwe\4\5`,
			},
			want: "qwe45",
			wantErr: false,
		},
		{
			name: "7",
			args: args{
				`qwe\45`,
			},
			want: "qwe44444",
			wantErr: false,
		},
		{
			name: "8",
			args: args{
				`qwe\\5`,
			},
			want: `qwe\\\\\`,
			wantErr: false,
		},
		{
			name: "8",
			args: args{
				`qwe\r5`,
			},
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnpackString() = %v, want %v", got, tt.want)
			}
		})
	}
}
