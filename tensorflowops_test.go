package tensorflowops

import "testing"

func TestRegisteredOps(t *testing.T) {
	ops, err := RegisteredOps()
	if err != nil {
		t.Fatal(err)
	}
	if len(ops) == 0 {
		t.Fatal("expected non empty ops list")
	}
	t.Log(ops)
}
