package timing

import (
	"testing"
	"time"
)

func TestStopTime(t *testing.T) {
	fn := func() {
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}

	got, _ := StopTime(func() error {
		fn()
		return nil
	})

	want := int64(100) // Expect around 100 milliseconds

	if got < want-10 || got > want+10 {
		t.Errorf("StopTime() = %d, want around %d", got, want)
	}
}
