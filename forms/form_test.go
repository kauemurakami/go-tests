package form

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle area", func(t *testing.T) {
		r := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := r.Area()

		if expectedArea != receivedArea {
			//fatal vai parar os testes
			t.Fatalf("Received area %f, expected is %f", receivedArea, expectedArea)
		}
	})
	t.Run("Circle area", func(t *testing.T) {
		c := Circle{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := c.Area()
		if expectedArea != receivedArea {
			//fatal vai parar os testes
			t.Fatalf("Received area %f, expected is %f", receivedArea, expectedArea)
		}

	})
}
