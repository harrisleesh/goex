package domain

import (
	"reflect"
	"testing"
)

func TestNewDeck1(t *testing.T) {
	tests := []struct {
		name string
		want deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDeck(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDeck() = %v, want %v", got, tt.want)
			}
		})
	}
}
