package main

import (
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestISEECalc(t *testing.T) {
	tests := []struct {
		salary        int
		hasHouse      bool
		houseArea     int
		familyMembers int
		euroPrice     int
		expected      float64
	}{
		{
			100,
			true,
			100,
			3,
			32000,
			0.0,
		},
	}

	for i, test := range tests {
		usr := &User{&primitive.NilObjectID, 1, &test.salary, &test.hasHouse, &test.houseArea, &test.familyMembers}
		srv := ISEEService{usr}
		Assert(t, test.expected, srv.Calc(test.euroPrice), "Test case %d is not successful\n", i)
	}
}

func Assert(t *testing.T, exp interface{}, act interface{}, msg string, args ...interface{}) {
	t.Helper()
	if !reflect.DeepEqual(act, exp) {
		msg = fmt.Sprintf(msg, args...)
		t.Errorf("%s\nactual: %v\nexpected: %v", msg, act, exp)
	}
}
