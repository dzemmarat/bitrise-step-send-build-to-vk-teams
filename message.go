package main

import (
	botgolang "github.com/mail-ru-im/bot-golang"
	"log"
)

func sendSuccessMessage(conf Config, string CommitMessage, string UrlToInstallPage) error {
	bot, err := botgolang.NewBot(string(conf.BotToken))
	if err != nil {
		log.Println("Wrong token!")
	}

	message := bot.NewMessage(string(conf.ChatID))
	message.ParseMode = botgolang.ParseModeHTML
	message.Text = "<b>😀 Новый билд!</b>\n" +
		"Последний коммит: <b><i>\"" + CommitMessage + "\"</i></b>\n" +
		"Пользователь: <b><i>" + conf.CommitAuthor + "</i></b>\n" +
		"Версия: <b<i>" + conf.Version + "</b></i>" +
		"Файлы билда доступны по ссылке: <b>" + UrlToInstallPage + "</b>"

	seeNewBuildBtn := botgolang.NewURLButton("Посмотреть билд!", UrlToInstallPage)
	keyboard := botgolang.NewKeyboard()
	keyboard.AddRow(seeNewBuildBtn)
	message.AttachInlineKeyboard(keyboard)

	err = message.Send()
	if err != nil {
		return err
	}
	return err
}

func sendErrorMessage(conf Config) error {
	bot, err := botgolang.NewBot(string(conf.BotToken))
	if err != nil {
		log.Println("Wrong token!")
	}

	message := bot.NewMessage(string(conf.ChatID))
	message.ParseMode = botgolang.ParseModeHTML
	message.Text = "<b>😭 Кажется что-то пошло не так...</b>\n" +
		"Билд №" + conf.Version + " провалился\n"
	err = message.Send()
	if err != nil {
		return err
	}
	return err
}
