package main

type ISEEService struct {
	User *User
}

func (s *ISEEService) Calc(euroPrice int) float64 {
	totalSalary := (*s.User.Salary * 12) / euroPrice
	houseThreshold := (s.calcHouseWorthiness() * 20) / 100
	return float64(totalSalary+houseThreshold) / s.calcFamilyMembersThreshold()
}

func (s *ISEEService) calcHouseWorthiness() int {
	if !*s.User.HasHouse {
		return 0
	}
	return *s.User.HouseArea * int(500)
}

func (s *ISEEService) calcFamilyMembersThreshold() float64 {
	switch *s.User.FamilyMembers {
	case 1:
		return 1.0
	case 2:
		return 1.75
	case 3:
		return 2.04
	case 4:
		return 2.46
	case 5:
		return 2.85
	default:
		diff := *s.User.FamilyMembers - 5
		return 2.85 + (float64(diff) * 0.35)
	}
}
