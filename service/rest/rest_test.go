package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
			if got := tt.m.GetKeys(); !assert.ElementsMatch(t, tt.m.GetKeys(), tt.want) {
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

func TestResponseSuccessWriter(t *testing.T) {
	var message = "success"

	t.Run("data not null", func(t *testing.T) {
		var w = httptest.NewRecorder()
		var data = map[string]interface{}{"key": "value", "number": float64(1)}
		var obj Response

		ResponseSuccessWriter(w, message, data)

		response := w.Result()

		defer response.Body.Close()

		assert.Equal(t, []string{"application/json"}, response.Header["Content-Type"])

		bytesResponse, _ := ioutil.ReadAll(response.Body)

		json.Unmarshal(bytesResponse, &obj)

		assert.Equal(t, statusOK, obj.Status)
		assert.Equal(t, message, obj.Message)
		assert.Equal(t, data, obj.Data.(map[string]interface{}))
	})

	t.Run("data null", func(t *testing.T) {
		var w = httptest.NewRecorder()
		var obj Response

		ResponseSuccessWriter(w, message, nil)

		response := w.Result()

		defer response.Body.Close()

		assert.Equal(t, []string{"application/json"}, response.Header["Content-Type"])

		bytesResponse, _ := ioutil.ReadAll(response.Body)

		json.Unmarshal(bytesResponse, &obj)

		assert.Equal(t, statusOK, obj.Status)
		assert.Equal(t, message, obj.Message)
		assert.Equal(t, nil, obj.Data)
	})
}

func TestResponseFailWriter(t *testing.T) {
	var w = httptest.NewRecorder()
	var message = "ooops!"
	var obj Response

	ResponseFailWriter(w, message)

	response := w.Result()

	defer response.Body.Close()

	assert.Equal(t, []string{"application/json"}, response.Header["Content-Type"])

	bytesResponse, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(bytesResponse, &obj)

	assert.Equal(t, statusError, obj.Status)
	assert.Equal(t, message, obj.Message)
	assert.Equal(t, nil, obj.Data)
}
