package handler

import (
	"testing"
)

func TestPassword(t *testing.T){
	tests := []struct {
		name  string
		pass  string
		valid bool
	}{
		{
			"CorrectPassword",
			"Aman1?1234",
			true,
		},
		{
			"Empty string",
			"",
			false,
		},
		{
			"WithoutUpperCaseString",
			"aa1?1234",
			false,
		},
		{
			"WithoutLowerCaseString",
			"AA1?1234",
			false,
		},
		{
			"WithoutNumber",
			"Ama?aaaa",
			false,
		},
		{
			"WithoutSymbol",
			"Aa101234",
			false,
		},
		{
			"LessThanRequiredMinimumLength",
			"Aa1?123",
			false,
		},
		{
			//WithSameUsernameandPassword
			"Amanjn82@",
			"Amanjn82@",
			false,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			valid,_:=PasswordChecker(c.pass,c.name)
			if c.valid != valid {
				t.Fatal("invalid password")
			}
		})
	}
}
