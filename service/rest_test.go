package service

import (
	"reflect"
	"testing"
)

func TestMapStringInt_GetKeys(t *testing.T) {
	tests := []struct {
		name string
		m    MapStringInt
		want []string
	}{
		{
			name: "empty",
			m:    MapStringInt{},
			want: []string{},
		},
		{
			name: "one",
			m:    MapStringInt{"ONE": 1},
			want: []string{"ONE"},
		},
		{
			name: "two",
			m:    MapStringInt{"ONE": 1, "TWO": 2},
			want: []string{"ONE", "TWO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.GetKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapStringInt.GetKeys() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestMapStringInt_GetListForOptions(t *testing.T) {
	tests := []struct {
		name string
		m    MapStringInt
		want []Option
	}{
		{
			name: "empty",
			m:    MapStringInt{},
			want: []Option{},
		},
		{
			name: "one",
			m:    MapStringInt{"ONE": 1},
			want: []Option{
				{
					ID:    1,
					Label: "ONE",
				},
			},
		},
		{
			name: "sorted asc on label",
			m:    MapStringInt{"XONE": 1, "ATWO": 2},
			want: []Option{
				{
					ID:    2,
					Label: "ATWO",
				},
				{
					ID:    1,
					Label: "XONE",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.GetListForOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapStringInt.GetListForOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
