package base

import "testing"

func TestEquals(t *testing.T) {

	t.Run("test the same type of value ", func(t *testing.T) {
		a := &Entity{Id: 10}
		b := &Entity{Id: 10}
		got := a.equals(b)
		if !got {
			t.Errorf("want true ,got %t", got)
		}
	})
}
