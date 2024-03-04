package minimax_test

import (
	"context"
	"fmt"
	"github.com/luckpunk/go-minimax"
	"os"
	"testing"
)

// go test -v -test.run=TestTextToSpeech ./
func TestTextToSpeech(t *testing.T) {
	client := minimax.NewClient(os.Getenv("MINIMAX_API_TOKEN"), os.Getenv("MINIMAX_GROUP_ID"))
	resp, err := client.CreateTextToSpeech(context.Background(), &minimax.CreateT2ARequest{
		VoiceID: minimax.VoiceFemaleYuJie,
		Text:    "hello",
		Path:    "./",
		Name:    "hello.mp3",
		TimberWeights: []minimax.TimberWeight{
			{
				VoiceID: minimax.VoiceFemaleYuJie,
				Weight:  1,
			},
		},
	})
	if err != nil {
		t.Log(err.Error())
		return
	}
	fmt.Printf("%#v\n", resp)
}

func TestTextToSpeechPro(t *testing.T) {
	client := minimax.NewClient(os.Getenv("MINIMAX_API_TOKEN"), os.Getenv("MINIMAX_GROUP_ID"))
	resp, err := client.CreateTextToSpeechPro(context.Background(), &minimax.CreateT2ARequest{
		Text:    "hello",
		VoiceID: minimax.VoiceFemaleYuJie,
	})
	if err != nil {
		t.Log(err.Error())
		return
	}

	fmt.Printf("%#v\n", resp)
}
