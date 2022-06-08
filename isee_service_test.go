package main

import (
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestISEECalc(t *testing.T) {
	tests := []struct {
		Salary        int
		HasHouse      bool
		HouseArea     int
		FamilyMembers int
		EuroPrice     int
		Expected      int
	}{
		{
			10000000,
			false,
			100,
			4,
			5000,
			9756,
		},
		{
			10000000,
			true,
			70,
			4,
			35000,
			4239,
		},
	}

	for i, test := range tests {
		usr := &User{&primitive.NilObjectID, 1, &test.Salary, &test.HasHouse, &test.HouseArea, &test.FamilyMembers}
		srv := ISEEService{usr}
		Assert(t, test.Expected, srv.Calc(test.EuroPrice), "Test case %d is not successful\n", i)
	}
}

func Assert(t *testing.T, exp interface{}, act interface{}, msg string, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(act, exp) {
		msg = fmt.Sprintf(msg, args...)
		t.Errorf("%s\nactual: %v\nexpected: %v", msg, act, exp)
	}
}
