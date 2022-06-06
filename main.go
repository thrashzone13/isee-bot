package main

import (
	"errors"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	selector = &tele.ReplyMarkup{}
	btnYes   = selector.Data("بله", "true")
	btnNo    = selector.Data("خیر", "false")
)

func main() {
	bot := createBot()
	db := ConnectMongoDB()
	userRepo := NewUserRepo(db)

	bot.Handle("/start", func(c tele.Context) error {
		var (
			sender = c.Sender()
			usr    = userRepo.FindOrInsert(sender.ID)
		)

		replyProperMessage(c, usr)

		return nil
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		var (
			usr = userRepo.FindOrInsert(c.Sender().ID)
			txt = c.Text()
		)

		switch usr.GetStatus() {
		case 0:
			if err := usr.SetSalary(txt); err != nil {
				return c.Send(err.Error())
			}
		case 1:
			return nil
		case 2:
			if err := usr.SetHouseArea(txt); err != nil {
				return c.Send(err.Error())
			}
		case 3:
			if err := usr.SetFamilyMembers(txt); err != nil {
				return c.Send(err.Error())
			}
		case 4:
			return nil
		default:
			return c.Send("Invalid user status")
		}

		userRepo.Update(usr)
		replyProperMessage(c, usr)

		return nil
	})

	bot.Handle(&btnYes, func(c tele.Context) error {
		usr := userRepo.FindOrInsert(c.Sender().ID)
		usr.SetHasHouse(false)
		userRepo.Update(usr)

		return c.Edit("تعداد افراد تحت تکلف سرپرست را وارد کنید")
	})

	bot.Handle(&btnNo, func(c tele.Context) error {
		usr := userRepo.FindOrInsert(c.Sender().ID)
		usr.SetHasHouse(true)
		userRepo.Update(usr)

		return c.Edit("متراژ ملک مسکونی را وارد کنید")
	})

	bot.Start()
}

func replyProperMessage(c tele.Context, usr *User) error {
	switch usr.GetStatus() {
	case 0:
		c.Send("مقدار دریافتی ماهانه سرپرست را به تومان وارد کنید")
	case 1:
		selector.Inline(
			selector.Row(btnYes, btnNo),
		)
		c.Send("آیا ملک مسکونی استیجاری است؟", selector)
	case 2:
		c.Send("متراژ ملک مسکونی را وارد کنید")
	case 3:
		c.Send("تعداد افراد تحت تکلف سرپرست را وارد کنید")
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
