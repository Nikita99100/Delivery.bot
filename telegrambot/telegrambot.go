//Telegram Bot Subscriber
package telegrambot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/handler"
	"gitlab.com/faemproject/backend/delivery/delivery.bot/proto"
	"strconv"
)

type (
	BotClient struct {
		Bot   *tgbotapi.BotAPI
		Token string
	}
	Subscriber struct {
		Bot     *tgbotapi.BotAPI
		Handler *handler.Handler
	}
)

func (b *BotClient) Init() error {
	var err error
	b.Bot, err = tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		return errors.Wrap(err, "SimpleBot init: ")
	}
	return nil
}

func (b *BotClient) DeleteMsg(chatID int64, msgID int) error {
	cfg := tgbotapi.NewDeleteMessage(chatID, msgID)
	_, err := b.Bot.Send(cfg)
	return err
}

func (b *BotClient) UpdateKeyboard(chatID int64, msgID int, newKeyborad proto.ButtonsSet) error {

	var markup tgbotapi.InlineKeyboardMarkup
	if len(newKeyborad.Buttons) == 0 {
		markup = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0)}
	} else {
		markup, _ = getInlineKeyboardMarkup(newKeyborad.Buttons)
	}
	cfg := tgbotapi.NewEditMessageReplyMarkup(chatID, msgID, markup)
	_, err := b.Bot.Send(cfg)
	return err
}

func (b *BotClient) UpdateMessage(chatID int64, msgId int, msg string) (proto.SendedMessage, error) {
	var tlgrMsg proto.SendedMessage
	cfg := tgbotapi.NewEditMessageText(chatID, msgId, msg)
	sendedMsg, err := b.Bot.Send(cfg)
	if err != nil {
		return tlgrMsg, errors.Wrap(err, "failed to send message")
	}
	tlgrMsg.Id = strconv.Itoa(sendedMsg.MessageID)
	return tlgrMsg, err
}

// SendMessage -
func (b *BotClient) SendMessage(chatID int64, msg string, keyboard ...proto.ButtonsSet) (proto.SendedMessage, error) {
	var keys proto.ButtonsSet
	var tlgrMsg proto.SendedMessage
	tlgMsgSender := tgbotapi.NewMessage(chatID, msg)
	if len(keyboard) > 0 && keyboard[0].DisplayLocation.S() != "" {
		keys = keyboard[0]
		if keys.DisplayLocation == "" {
			return tlgrMsg, errors.New("Empty display location for telegram buttons")
		}
		if keys.DisplayLocation == proto.Consts.ButtonsDisplayLocation.Inline || keys.DisplayLocation == proto.Buttons.Display.Inline {
			markup, err := getInlineKeyboardMarkup(keys.Buttons)
			if err != nil {
				return tlgrMsg, errors.Wrap(err, "Reply inline keyboard markup error")
			}
			tlgMsgSender.ReplyMarkup = markup
		}
		if keys.DisplayLocation == proto.Consts.ButtonsDisplayLocation.Reply || keys.DisplayLocation == proto.Buttons.Display.Reply {
			markup, err := getReplyKeyboardMarkup(keys.Buttons)
			if err != nil {
				return tlgrMsg, errors.Wrap(err, "Reply keyboard markup error")
			}
			markup.OneTimeKeyboard = true
			tlgMsgSender.ReplyMarkup = markup
		}
	} else {
		tlgMsgSender.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	}

	tlgMsgSender.ParseMode = "HTML"
	sendedMsg, err := b.Bot.Send(tlgMsgSender)
	if err != nil {
		return tlgrMsg, errors.Wrap(err, "sending message error")
	}
	tlgrMsg.Id = strconv.Itoa(sendedMsg.MessageID)
	return tlgrMsg, nil
}

func getInlineKeyboardMarkup(keys []proto.MsgKeyboardRows) (tgbotapi.InlineKeyboardMarkup, error) {
	var res tgbotapi.InlineKeyboardMarkup
	var kb [][]tgbotapi.InlineKeyboardButton
	for _, i := range keys {
		var button []tgbotapi.InlineKeyboardButton
		for _, btn := range i.MsgButtons {
			if len(btn.Data) > proto.ButtonDataSize {
				return res, errors.Errorf("Слишок большой размер мета данных кнопки (>64)")
			}
			button = append(button, tgbotapi.NewInlineKeyboardButtonData(btn.Text, btn.Data))
		}
		newRow := tgbotapi.NewInlineKeyboardRow(button...)
		kb = append(kb, newRow)
	}
	res = tgbotapi.NewInlineKeyboardMarkup(kb...)
	return res, nil
}

func getReplyKeyboardMarkup(keys []proto.MsgKeyboardRows) (tgbotapi.ReplyKeyboardMarkup, error) {
	var res tgbotapi.ReplyKeyboardMarkup
	var kb [][]tgbotapi.KeyboardButton
	for _, i := range keys {
		var button []tgbotapi.KeyboardButton
		for _, btn := range i.MsgButtons {
			switch btn.Type {
			case proto.Consts.ButtonsTypes.Regular, proto.Buttons.Type.Regular:
				button = append(button, tgbotapi.NewKeyboardButton(btn.Text))
			case proto.Consts.ButtonsTypes.Contact, proto.Buttons.Type.Contact:
				button = append(button, tgbotapi.NewKeyboardButtonContact(btn.Text))
			case proto.Consts.ButtonsTypes.Location, proto.Buttons.Type.Location:
				button = append(button, tgbotapi.NewKeyboardButtonLocation(btn.Text))
			default:
				button = append(button, tgbotapi.NewKeyboardButton(btn.Text))
			}
		}
		newRow := tgbotapi.NewKeyboardButtonRow(button...)
		kb = append(kb, newRow)
	}
	res = tgbotapi.NewReplyKeyboard(kb...)
	return res, nil
}

//Init Telegram bot subscriber
func (s *Subscriber) Init() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := s.Bot.GetUpdatesChan(u)
	if err != nil {
		return errors.Wrap(err, "Failed to get updates chan: ")
	}
	//go s.handleNewMsg(&updates)
	go s.handleIncomeMsg(&updates)
	return nil
}

func (s *Subscriber) handleIncomeMsg(updChan *tgbotapi.UpdatesChannel) {
	for update := range *updChan {
		msgJob := s.Handler.Jobs.GetJobQueue(handler.JobQueueNameMsgs, handler.JobQueueLimitMsgs)
		//создаем очередь обработки сообщений, что бы все обрабатывалось последовательно
		var id int
		if update.CallbackQuery != nil {
			id = update.CallbackQuery.Message.MessageID
		} else if update.Message != nil {
			id = update.Message.MessageID
		}

		err := msgJob.Execute(id, func() error {
			ctx := context.Background()
			s.Handler.TelegramMsgHandler(ctx, &update)
			return nil
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"reason": "msgJob execute error",
			}).Error(err)
		}

	}
}
