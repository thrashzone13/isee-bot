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

	bot.Use(SanitizePersianDigits)
	bot.Handle("/start", func(c tele.Context) error {
		var (
			sender     = c.Sender()
			isNew, usr = userRepo.FindOrInsert(sender.ID)
		)

		if isNew {
			c.Send(fmt.Sprintf(`
			سلام %s! %s %s

به ربات محاسبه عدد ایزه بورس استانی ایتالیا خوش اومدی. برای شروع لطفا اطلاعاتی که ازت خواسته میشه رو وارد کن.

%sنکته خیلی مهم%s

این ربات با استفاده از فرمول های موجود در اینترنت محاسبات رو انجام میده و عدد به دست آمده به هیچ عنوان قابل تضمین نیست !
			`, c.Sender().FirstName, emoji.WavingHand, emoji.Parse(":flag-it:"), emoji.ExclamationMark, emoji.ExclamationMark))
		}

		replyProperMessage(c, usr)

		return nil
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		var (
			_, usr = userRepo.FindOrInsert(c.Sender().ID)
			txt    = c.Text()
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
			return showCalculationKeyboard(c)
		default:
			return c.Send("Invalid user status")
		}

		userRepo.Update(usr)
		replyProperMessage(c, usr)

		return nil
	})

	bot.Handle(&btnYes, func(c tele.Context) error {
		_, usr := userRepo.FindOrInsert(c.Sender().ID)
		usr.SetHasHouse(false)
		userRepo.Update(usr)

		return c.Edit("تعداد افراد تحت تکلف سرپرست را وارد کنید")
	})

	bot.Handle(&btnNo, func(c tele.Context) error {
		_, usr := userRepo.FindOrInsert(c.Sender().ID)
		usr.SetHasHouse(true)
		userRepo.Update(usr)

		return c.Edit("متراژ ملک مسکونی را وارد کنید")
	})

	bot.Handle(&btnOfficialEuro, func(c tele.Context) error {
		usr := userRepo.Find(c.Sender().ID)
		srv := ISEEService{usr}
		return c.Send(fmt.Sprintf("%f", srv.Calc(5000)))
	})

	bot.Handle(&btnUnOfficialEuro, func(c tele.Context) error {
		usr := userRepo.Find(c.Sender().ID)
		srv := ISEEService{usr}
		return c.Send(fmt.Sprintf("%f", srv.Calc(32000)))
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
