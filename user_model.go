package main

import (
	"errors"
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
	if *u.HasHouse && u.HouseArea == nil {
		return 2
	}
	if u.FamilyMembers == nil {
		return 3
	}

	return 4
}

func (u *User) SetSalary(txt string) error {
	val, err := strconv.Atoi(txt)
	if err != nil {
		return errors.New("فرمت دریافتی ماهانه اشتباه است. دوباره وارد کنید")
	}

	u.Salary = &val
	return nil
}

func (u *User) SetHasHouse(val bool) {
	u.HasHouse = &val
}

func (u *User) SetHouseArea(txt string) error {
	val, err := strconv.Atoi(txt)
	if err != nil {
		return errors.New("فرمت متراژ ملک اشتباه است. دوباره وارد کنید")
	}

	u.HouseArea = &val
	return nil
}

func (u *User) SetFamilyMembers(txt string) error {
	val, err := strconv.Atoi(txt)
	if err != nil {
		return errors.New("فرمت تعداد افراد تحت تکلف سرپرست اشتباه است. دوباره وارد کنید")
	}

	u.FamilyMembers = &val
	return nil
}

func (u *User) CalculateISEE() int {
	return 0
}
