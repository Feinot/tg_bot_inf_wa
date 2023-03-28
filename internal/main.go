package main

import (
	"bot/comands/info"
	"bot/comands/stat"

	"fmt"
	"log"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var (
	telegramBotToken string = "6193690429:AAGR5rY6GwnAfanZx4fbjSD49vh3v_4FJZc"
)

// func init() {

// 	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
// 	flag.Parse()

// 	if telegramBotToken == "" {
// 		log.Print("-telegrambottoken is required")
// 		os.Exit(1)
// 	}
// }

func main() {

	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println(err)
	}

	for update := range updates {

		reply := "Не знаю что сказать"

		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":

			reply = `I can tell you the info and statistics, /statistics and /city (example Томск = /tomsk, Омск = /omsk )`
			stat.DbUpdate(update.Message.From.UserName)

		case "statistics":
			a, b := stat.Stat(update.Message.From.UserName)
			reply = fmt.Sprintf("quantity requests = %d date first requests = %s", a, b)

		default:

			reply = info.Inf(update.Message.Command())
			stat.DbUpdate(update.Message.From.UserName)

		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		bot.Send(msg)

	}
}

// Привет!

// Задача - написать телеграм бота на Go. Использовать Postgres для хранения данных.

// Бот должен поддерживать две команды:
// Информация - по запросу юзера искать какую-то простую информацию в открытых источниках. Какую информацию - на твой выбор. Может быть погода в городе из запроса, текущее время в городе из запроса, курс валюты из запроса, и т.д.
// Статистика - показать юзеру когда был его первый запрос, сколько всего запросов было. И можно добавить ещё показателей, которые мы можем получить из хранимых данных.

// Бот должен быть упакован в Docker контейнер. Для удобства разработки и тестирования можно приложить docker-compose файл.

// Обрати внимание на обработку ошибок и логирование.
