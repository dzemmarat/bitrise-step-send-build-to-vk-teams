package main

import (
	botgolang "github.com/mail-ru-im/bot-golang"
	"log"
)

func sendMessage(conf Config) error {
	bot, err := botgolang.NewBot(string(conf.BotToken))
	if err != nil {
		log.Println("Wrong token!")
	}

	message := bot.NewMessage(string(conf.ChatID))
	message.ParseMode = botgolang.ParseModeHTML
	message.Text = "<b>😀 Новый билд!</b>\n" +
		"Последний коммит: <b><i>\"" + conf.CommitMessage + "\"</i></b>\n" +
		"Пользователь: <b><i>" + conf.CommitAuthor + "</i></b>\n" +
		"Файлы билда доступны по ссылке: <b>" + conf.FileURL + "</b>"
	err = message.Send()
	if err != nil {
		return err
	}
	return err
}
