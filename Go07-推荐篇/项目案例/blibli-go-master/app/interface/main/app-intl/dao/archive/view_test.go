package archive

import (
	"context"
	"testing"
)

func TestView3(t *testing.T) {
	Convey(t.Name(), t, func() {
		d.View3(context.Background(), 1)
	})
}

func TestViewCache(t *testing.T) {
	Convey(t.Name(), t, func() {
		d.ViewCache(context.Background(), 1)
	})
}

func TestDescription(t *testing.T) {
	Convey(t.Name(), t, func() {
		d.Description(context.Background(), 2)
	})
}
