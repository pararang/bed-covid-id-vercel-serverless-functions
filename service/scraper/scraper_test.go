package scraper

import (
	"api-bed-covid/model"
	"testing"
)

func TestHospitalDetail_IsEmpty(t *testing.T) {
	type fields struct {
		Name    string
		Address string
		Hotline string
		Room    []model.Room
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "is empty after init",
			fields: fields{},
			want:   true,
		},
		{
			name: "is empty without init",
			want: true,
		},
		{
			name: "no empty",
			fields: fields{
				Name: "Name",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := &model.HospitalDetail{
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
				Hotline: tt.fields.Hotline,
				Room:    tt.fields.Room,
			}
			if got := hd.IsEmpty(); got != tt.want {
				t.Errorf("HospitalDetail.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
