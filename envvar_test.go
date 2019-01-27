package envvar

import (
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Setenv("VAR1", "testvalue")
	value := GetString("VAR1")
	if value != "testvalue" {
		t.Errorf("Get string, got: %s, want: %s.", value, "testvalue")
	}
}

func TestDefineString(t *testing.T) {
	DefineString("VAR2", "defvalue1")
	value := GetString("VAR2")
	if value != "defvalue1" {
		t.Errorf("Define string, got: %s, want: %s.", value, "defvalue1")
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("VAR1_INT", "3306")
	value := GetInt("VAR1_INT")
	if value != 3306 {
		t.Errorf("Get string, got: %d, want: %d.", value, 3306)
	}
}

func TestDefineInt(t *testing.T) {
	DefineInt("VAR2_INT", 3306)
	value := GetInt("VAR2_INT")
	if value != 3306 {
		t.Errorf("Define string, got: %d, want: %d.", value, 3306)
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("VAR1_BOOL", "true")
	value := GetBool("VAR1_BOOL")
	if value != true {
		t.Errorf("Get string, got: %v, want: %v.", value, true)
	}
}

func TestDefineBool(t *testing.T) {
	DefineBool("VAR2_BOOL", true)
	value := GetBool("VAR2_BOOL")
	if value != true {
		t.Errorf("Define string, got: %v, want: %v.", value, true)
	}
}
