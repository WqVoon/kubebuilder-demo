package v1

import "testing"

func TestCommandEqualTo(t *testing.T) {
	var obj1, obj2 *Cmd
	if !obj1.CommandIsEqualTo(obj2) || !obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should equal", obj2.Spec)
	}

	obj1 = &Cmd{
		Spec: CmdSpec{
			Command: "c1",
		},
	}
	if obj1.CommandIsEqualTo(obj2) || obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should not equal", obj2.Spec)
	}

	obj2 = &Cmd{
		Spec: CmdSpec{
			Command: "c1",
		},
	}
	if !obj1.CommandIsEqualTo(obj2) || !obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should equal", obj2.Spec)
	}

	obj1.Spec.Args = []string{"a1"}
	obj2.Spec.Args = []string{"a1"}
	if !obj1.CommandIsEqualTo(obj2) || !obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should equal", obj2.Spec)
	}

	obj1.Spec.Args = append(obj1.Spec.Args, "a2")
	if obj1.CommandIsEqualTo(obj2) || obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should not equal", obj2.Spec)
	}

	obj2.Spec.Args = nil
	if obj1.CommandIsEqualTo(obj2) || obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should not equal", obj2.Spec)
	}

	obj1.Spec.Args = nil
	obj1.Spec.Command = "c2"
	if obj1.CommandIsEqualTo(obj2) || obj2.CommandIsEqualTo(obj1) {
		t.Fatal(obj1.Spec, "should not equal", obj2.Spec)
	}
}
