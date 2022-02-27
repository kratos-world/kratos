package base

import (
	"testing"
)

func TestName(t *testing.T) {
	name := "user"
	if SmallCamel(name) != name {
		t.Fatal("SmallCamel error", SmallCamel(name))
	}
	if BigCamel(name) != "User" {
		t.Fatal("BigCamel error", BigCamel(name))
	}

	name = "user_service"
	if SmallCamel(name) != "userService" {
		t.Fatal("SmallCamel error", SmallCamel(name))
	}
	if BigCamel(name) != "UserService" {
		t.Fatal("BigCamel error", BigCamel(name))
	}

	name = "UserService"
	if camelToUnderline(name) != "user_service" {
		t.Fatal("camelToUnderline error", camelToUnderline(name))
	}
	if SmallCamel(name) != "userService" {
		t.Fatal("SmallCamel error", SmallCamel(name))
	}
	if BigCamel(name) != "UserService" {
		t.Fatal("BigCamel error", BigCamel(name))
	}

	name = "userService"
	if camelToUnderline(name) != "user_service" {
		t.Fatal("camelToUnderline error", camelToUnderline(name))
	}
	if SmallCamel(name) != "userService" {
		t.Fatal("SmallCamel error", SmallCamel(name))
	}
	if BigCamel(name) != "UserService" {
		t.Fatal("BigCamel error", BigCamel(name))
	}
}
