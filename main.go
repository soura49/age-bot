package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func botCalculactor(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) string {
	year := request.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		fmt.Println(err)
		return "Invalid Command"
	}
	age := 2023 - yob

	r := fmt.Sprintf("age is %d", age)

	return r
}

func replyToslackBot(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	res := botCalculactor(botCtx, request, response)
	response.Reply(res)
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
	runtime.Gosched()
	time.Sleep(10 * time.Millisecond)
	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler:     replyToslackBot,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("***************************************************")
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("***************************************************")
	}
}
