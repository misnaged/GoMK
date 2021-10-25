package spriteplus

import (
	"github.com/faiface/pixel"
	"reflect"
	"testing"
)

func TestMakeAnimation(t *testing.T) {
	type args struct {
		sprites     []*pixel.Sprite
		frameLength int
	}
	tests := []struct {
		name    string
		args    args
		want    *Animation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeAnimation(tt.args.sprites, tt.args.frameLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeAnimation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeAnimation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

