package telegram

import "github.com/Bossnicks/read_adviser_bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
}

func New(client *telegram.Client) {

}
