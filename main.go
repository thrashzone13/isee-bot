package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/enescakir/emoji"
	tele "gopkg.in/telebot.v3"
)

var (
	calcMenu          = &tele.ReplyMarkup{ResizeKeyboard: true}
	btnOfficialEuro   = calcMenu.Text("محاسبه با یورو دولتی")
	btnUnOfficialEuro = calcMenu.Text("محاسبه با یورو آزاد")
	btnReset          = calcMenu.Text("ورود مجدد اطلاعات")

	selector = &tele.ReplyMarkup{}
	btnYes   = selector.Data("بله", "true")
	btnNo    = selector.Data("خیر", "false")
)

func main() {
	bot := createBot()
	db := ConnectMongoDB()
	userRepo := NewUserRepo(db)

	bot.Use(SanitizePersianDigits, logUserMessage)

	bot.Handle(tele.OnText, func(c tele.Context) error {
		isNew, usr := userRepo.FindOrInsert(c.Sender().ID)

		if isNew {
			calcMenu.RemoveKeyboard = true
			c.Send(fmt.Sprintf(`
			سلام %s! %s %s

به ربات محاسبه عدد ایزه بورس استانی ایتالیا خوش اومدی. برای شروع لطفا اطلاعاتی که ازت خواسته میشه رو وارد کن.

%sنکته خیلی مهم%s

این ربات با استفاده از فرمول های موجود در اینترنت محاسبات رو انجام میده و عدد به دست آمده به هیچ عنوان قابل تضمین نیست !
			`, c.Sender().FirstName, emoji.WavingHand, emoji.Parse(":flag-it:"), emoji.ExclamationMark, emoji.ExclamationMark), calcMenu)
		} else {
			recieveInfo(c, usr)
			userRepo.Update(usr)
		}

		askInfo(c, usr)

		return nil
	})

	bot.Handle(&btnYes, func(c tele.Context) error {
		_, usr := userRepo.FindOrInsert(c.Sender().ID)
		usr.SetHasHouse(false)
		userRepo.Update(usr)

		return askInfo(c, usr)
	})

	bot.Handle(&btnNo, func(c tele.Context) error {
		_, usr := userRepo.FindOrInsert(c.Sender().ID)
		usr.SetHasHouse(true)
		userRepo.Update(usr)

		return askInfo(c, usr)
	})

	bot.Handle(&btnOfficialEuro, func(c tele.Context) error {
		usr := userRepo.Find(c.Sender().ID)

		if usr.GetStatus() != 4 {
			return askInfo(c, usr)
		}

		srv := ISEEService{usr}
		return c.Send(ThousandSeparator(srv.Calc(5000)))
	})

	bot.Handle(&btnUnOfficialEuro, func(c tele.Context) error {
		usr := userRepo.Find(c.Sender().ID)

		if usr.GetStatus() != 4 {
			return askInfo(c, usr)
		}

		srv := ISEEService{usr}
		return c.Send(ThousandSeparator(srv.Calc(35000)))
	})

	bot.Handle(&btnReset, func(c tele.Context) error {
		usr := userRepo.Find(c.Sender().ID)

		usr.Salary = nil
		usr.HasHouse = nil
		usr.HouseArea = nil
		usr.FamilyMembers = nil
		userRepo.Update(usr)

		calcMenu.RemoveKeyboard = true
		c.Send("اطلاعات خود را مجدد وارد کنید", calcMenu)

		return askInfo(c, usr)
	})

	bot.Start()
}

func recieveInfo(c tele.Context, usr *User) error {
	txt := c.Text()

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
		return showCalculationKeyboard(c)
	default:
		return c.Send("Invalid user status")
	}

	return nil
}

func askInfo(c tele.Context, usr *User) error {
	switch usr.GetStatus() {
	case 0:
		c.Send("مقدار دریافتی ماهانه سرپرست را به تومان وارد کنید")
	case 1:
		selector.Inline(
			selector.Row(btnYes, btnNo),
		)
		c.Send("آیا ملک مسکونی استیجاری است؟", selector)
	case 2:
		c.EditOrSend("متراژ ملک مسکونی را وارد کنید")
	case 3:
		c.EditOrSend(`
تعداد افراد تحت تکلف سرپرست را وارد کنید
در صورتی که خودتان سرپرست هستید عدد ۱ را وارد کنید
		`)
	case 4:
		showCalculationKeyboard(c)
	default:
		return errors.New("invalid user status")
	}

	return nil
}

func showCalculationKeyboard(c tele.Context) error {
	calcMenu.Reply(
		calcMenu.Row(btnOfficialEuro, btnUnOfficialEuro),
		calcMenu.Row(btnReset),
	)
	return c.Send("برای محاسبه عدد ایزه یکی از انواع نرخ ارز را انتخاب کنید", calcMenu)
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
