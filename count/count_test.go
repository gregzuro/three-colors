package count

import (
	"image"
	"reflect"
	"testing"
)

func TestCountColorsFromImage(t *testing.T) { //comment
	type args struct {
		image image.Image
	}
	tests := []struct {
		name string
		args args
		want map[string]uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountColorsFromImage(tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountColorsFromImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTopThreeFromCounts(t *testing.T) {
	type args struct {
		colorCounts map[string]uint
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTopThreeFromCounts(tt.args.colorCounts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTopThreeFromCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
