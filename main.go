package main

import (
	"log"

	"github.com/slack-go/slack"
)

func main() {
	token := "token"
	c := slack.New(token)

	_, _, err := c.PostMessage("@hono3bono3", slack.MsgOptionBlocks(
		slack.NewSectionBlock(
			// テキスト
			&slack.TextBlockObject{Type: "mrkdwn", Text: "*稼働状況の通知*"},
			// フィールド
			[]*slack.TextBlockObject{
				{Type: "mrkdwn", Text: "*勤務日数*:\n`12`日"},
				{Type: "mrkdwn", Text: "*平均勤務時間*:\n`8.15`時間"},
				{Type: "mrkdwn", Text: "*稼働時間*:\n`97.80`時間"},
				{Type: "mrkdwn", Text: "*残日数*:\n`8`日"},
			},
			nil,
		),
	))
	if err != nil {
		log.Fatalln(err)
	}
}
