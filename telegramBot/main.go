package main

import (
	"fmt"
	"goTddProject/telegramBot/service"
	"log"
	"time"
)

func main() {
	//log.Println("Telegram alarm bot program is start")
	//
	//// bot token value 1209633330:AAEArq32PNoeMex1IhlH1mKENmLbWyrwE-Y
	//bot, err := tgbotapi.NewBotAPI("토큰값")
	//if err != nil {
	//	log.Fatalln("Telegram Bot API connect token connect fail")
	//	panic(err)
	//}
	//
	//bot.Debug = true
	//
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 30
	//
	////updates, err := bot.GetUpdatesChan(u)
	//
	//msg := tgbotapi.NewMessage(-499899417, "맥주병으로 헤드샷 날리게???")
	////msg.ReplyToMessageID = update.Message.MessageID
	//
	//bot.Send(msg)


	//for update := range updates {
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//	}
	//
	//	log.Println("updateChatId : ",  update.Message.Chat.ID)
	//
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "test message send")
	//	msg.ReplyToMessageID = update.Message.MessageID
	//
	//	bot.Send(msg)
	//}


	//********************************** 단순 메세지 보내기 **********************************//
	//t, err := telegram.NewTelegram("1209633330:AAEArq32PNoeMex1IhlH1mKENmLbWyrwE-Y")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//msg := "야근하십니꽈???!!"
	//
	//if err := t.Send(-499899417, msg); err != nil {
	//	log.Fatalln(err)
	//}
	//********************************** 단순 메세지 보내기 **********************************//




	//********************************** 특정 시간까지 남은 시간 **********************************//
	t, err := telegram.NewTelegram("1209633330:AAEArq32PNoeMex1IhlH1mKENmLbWyrwE-Y")
	if err != nil {
		log.Fatalln(err)
	}

	if err := t.Send(-499899417, "참치까지 남은시간 D-Timer 동작 시작합니다."); err != nil {
		log.Fatalln(err)
	}

	endTime := time.Date(2020,10,15,19,0,0,0, time.Local)
	fmt.Println(endTime.Format(time.Kitchen))

	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		timeDif := endTime.Sub(time.Now())

		msg := "참치까지 남은시간은 정확하게 " + timeDif.String() + " 남았습니다"

		if err := t.Send(-499899417, msg); err != nil {
			log.Fatalln(err)
		}
	}
	//********************************** 특정 시간까지 남은 시간 **********************************//

}