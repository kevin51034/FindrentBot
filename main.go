package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"encoding/json"


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
				if index, err := strconv.Atoi(message.Text); err == nil {
					messages := ConstructFlexJSON(index)
					// flex messages
					if _, err = bot.ReplyMessage(event.ReplyToken, messages...).Do(); err != nil {
						log.Print(err)
					}
				}else if message.Text == "Items" {
					items, _ := c.ItemandPageNum()
					//c.Start(1)
					message := linebot.NewTextMessage("it has " + strconv.Itoa(items) + " items within your conditions!")
					if _, err = bot.ReplyMessage(event.ReplyToken, message).Do(); err != nil {
						log.Print(err)
					}
				} else if message.Text == "Quota" {
					// get quota
					quota, err := bot.GetMessageQuota().Do()
					if err != nil {
						log.Println("Quota err:", err)
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}

func ConstructFlexJSON(index int) []linebot.SendingMessage {
	pages := (index+5)/30 + 1
	fmt.Println(pages)
	c.Start(pages)
	var messages []linebot.SendingMessage
	for i:=index; i<index+5; i++ {
		jsonData := NewJSONData()
		var structdata FlexMessage
		json.Unmarshal(jsonData, &structdata)
		structdata.Header.Contents[0].Contents[0].URL = c.Houselist[i].ImgSrc
		structdata.Body.Contents[0].Contents[0].Contents[0].Text = c.Houselist[i].Title
		structdata.Body.Contents[0].Contents[0].Contents[1].Text = c.Houselist[i].Kind + " $" + c.Houselist[i].Price + "/月"
		structdata.Body.Contents[0].Contents[1].Contents[0].Text = c.Houselist[i].Address + " " + c.Houselist[i].Floor + " \n" + c.Houselist[i].UpdateTime
		structdata.Header.Contents[0].Contents[0].Action.URI = c.Houselist[i].URL

		b, err := json.MarshalIndent(structdata, "", "  ")
		container, err := linebot.UnmarshalFlexMessageJSON(b)
		// err is returned if invalid JSON is given that cannot be unmarshalled
		if err != nil {
			log.Fatal(err)
		}
		tmpmessage := linebot.NewFlexMessage("House Information", container)
		messages = append(messages, tmpmessage)
	}
	return messages
}

func setOptions() {
	c.Options.RentPrice = "10000,15000"
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
