package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	crawler "github.com/kevin51034/Crawler591"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var c *crawler.Crawler

func init() {
	c = crawler.Newcrawler()
}
func main() {
	var err error
	//c := crawler.Newcrawler()

	// set tour option
	setOptions()

	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "items" {
					items, _ := c.ItemandPageNum()
					c.Start(1)
					// flex messages
					jsonData := NewJSONData()
					container, err := linebot.UnmarshalFlexMessageJSON(jsonData)
					// err is returned if invalid JSON is given that cannot be unmarshalled
					if err != nil {
						log.Fatal(err)
					}
					//fmt.Printf("%+v", container)
					var messages []linebot.SendingMessage

					tmp1 := linebot.NewTextMessage("it has " + strconv.Itoa(items) + " items within your conditions!")
					messages = append(messages, tmp1)
					tmp2 := linebot.NewFlexMessage("alt text", container)
					messages = append(messages, tmp2)
					if _, err = bot.ReplyMessage(event.ReplyToken, messages...).Do(); err != nil {
						log.Print(err)
					}
				}
				/*
					// get quota
					quota, err := bot.GetMessageQuota().Do()
					if err != nil {
						log.Println("Quota err:", err)
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
						log.Print(err)
					}*/
			}
		}
	}
}

func setOptions() {
	c.Options.RentPrice = "8000,15000"
	c.Options.Kind = 2
	c.Options.HasImg = "1"
	c.Options.NotCover = "1"
	c.Options.Role = "1"
}

/*

package main

import (
	crawler "github.com/kevin51034/Crawler591"
)

func main() {
	c := crawler.Newcrawler()
	//c.options.RentPrice = "2"
	//fmt.Println(c.ItemandPageNum())
	c.Start(3)
	//c.Scrape(10)
	//doc := NewDoc()
	//findItemandPage(doc)
	//fmt.Println(c.options)

	c.ExportJSON()

	// TODO extract within minuts/hours/days
}

*/
