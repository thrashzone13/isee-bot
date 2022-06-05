package main

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	UserID        int64               `bson:"user_id,omitempty" json:"user_id"`
	Salary        *int                `bson:"salary,omitempty" json:"salary"`
	HasHouse      *bool               `bson:"has_house,omitempty" json:"has_house"`
	HouseArea     *int                `bson:"house_area,omitempty" json:"house_area"`
	FamilyMembers *int                `bson:"family_members,omitempty" json:"family_members"`
}

func (u *User) GetStatus() int {
	if u.Salary == nil {
		return 0
	}
	if u.HasHouse == nil {
		return 1
	}
	if u.HouseArea == nil {
		return 2
	}
	if u.FamilyMembers == nil {
		return 3
	}
	return 4
}

func (u *User) SetSalary(val string) bool {
	salary, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	u.Salary = &salary
	return true
}

func (u *User) SetHasHouse(val string) bool {
	hasHouse, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	u.HasHouse = &hasHouse
	return true
}

func (u *User) SetHouseArea(val string) bool {
	houseArea, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	u.HouseArea = &houseArea
	return true
}

func (u *User) SetFamilyMembers(val string) bool {
	familyMembers, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	u.FamilyMembers = &familyMembers
	return true
}
