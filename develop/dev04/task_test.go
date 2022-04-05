package main

import (
	"reflect"
	"testing"
)

func TestSearchAnagram(t *testing.T) {
	type args struct {
		arr []string
	}
	t1 := []string{"пятаК", "пятка", "тяпка", "каПот"}
	t1Res := []string{"пятка", "тяпка"}
	m1 := make(map[string][]string)
	m1["пятак"] = t1Res

	t2 := []string{"листок", "Слиток", "столИк", "Сколит", ""}
	t2Res := []string{"сколит", "слиток", "столик"}
	m2 := make(map[string][]string)
	m2["листок"] = t2Res

	t3 := []string{"колпа", "КОПАЛ", "покал", "полак", "полка"}
	t3Res := []string{"копал", "покал", "полак", "полка"}
	m3 := make(map[string][]string)
	m3["колпа"] = t3Res

	t4 := []string{" Карп", "крап", "парк", "прак"}
	t4Res := []string{"крап", "парк", "прак"}
	m4 := make(map[string][]string)
	m4["карп"] = t4Res

	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "1",
			args: args{
				arr: t1,
			},
			want: m1,
		},
		{
			name: "2",
			args: args{
				arr: t2,
			},
			want: m2,
		},
		{
			name: "3",
			args: args{
				arr: t3,
			},
			want: m3,
		},
		{
			name: "4",
			args: args{
				arr: t4,
			},
			want: m4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchAnagram(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
