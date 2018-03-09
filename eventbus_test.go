package eventbus

import (
	"testing"
)

func TestEventBus(t *testing.T) {
	bus := New()
	var caught string
	var data EventData
	bus.Subscribe("test", func(e Event) {
		caught = e.Name
		data = e.Data
	})
	bus.Publish("test")
	if caught != "test" {
		t.Error("expected to caught event")
	}
	caught = ""
	bus.Publish("test", EventData{"pi": 3.14159})
	if _, ok := data["pi"]; !ok {
		t.Error("expected to fetch data")
	}
	caught = ""
	bus.Publish("test", map[string]interface{}{"pi": 3.14159})
	if _, ok := data["pi"]; !ok {
		t.Error("expected to fetch data")
	}
}
