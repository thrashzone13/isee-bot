package main

import (
	"errors"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	bot := createBot()
	db := ConnectMongoDB()
	userRepo := NewUserRepo(db)

	bot.Handle("/start", func(c tele.Context) error {
		user := userRepo.FindOrInsert(c.Sender().ID)

		replyProperMessage(c, user.GetStatus())

		return nil
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		var (
			user = userRepo.Find(c.Sender().ID)
			txt  = c.Text()
		)

		status := user.GetStatus()

		switch status {
		case 0:
			if res := user.SetSalary(txt); !res {
				c.Send("Wrong salary format")
			}
		case 1:
			if res := user.SetHasHouse(txt); !res {
				c.Send("Wrong has house format")
			}
		case 2:
			if res := user.SetHouseArea(txt); !res {
				c.Send("Wrong house area format")
			}
		case 3:
			if res := user.SetFamilyMembers(txt); !res {
				c.Send("Wrong family member count format")
			}
		default:
			log.Panic("Invalid user status")
		}

		userRepo.Update(user)

		replyProperMessage(c, status)
		return nil
	})

	bot.Start()
}

func replyProperMessage(c tele.Context, status int) error {
	switch status {
	case 0:
		c.Send("Enter salary")
	case 1:
		selector := &tele.ReplyMarkup{}
		btnYes := selector.Data("Yes", "true")
		btnNo := selector.Data("No", "false")
		selector.Inline(
			selector.Row(btnYes, btnNo),
		)
		c.Send("Do you own a house?", selector)
	case 2:
		c.Send("Enter house area")
	case 3:
		c.Send("Enter number of family members")
	case 4:
		c.Send("Woop!")
	default:
		return errors.New("invalid user status")
	}

	return nil
}

func createBot() *tele.Bot {
	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	CheckIfError(err)

	return b
}
