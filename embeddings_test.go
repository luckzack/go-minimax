package minimax_test

import (
	"context"
	"fmt"
	"github.com/luckpunk/go-minimax"
	"os"
	"testing"
)

func TestCreateEmbeddingsByDbType(t *testing.T) {
	client := minimax.NewClient(os.Getenv("MINIMAX_API_TOKEN"), os.Getenv("MINIMAX_GROUP_ID"))
	resp, err := client.CreateEmbeddings(context.Background(), &minimax.CreateEmbeddingsRequest{
		Texts: []string{"hello"},
		Type:  minimax.EmbeddingsDbType,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("%#v\n", resp)
}

func TestCreateEmbeddingsByQueryType(t *testing.T) {
	client := minimax.NewClient(os.Getenv("MINIMAX_API_TOKEN"), os.Getenv("MINIMAX_GROUP_ID"))
	resp, err := client.CreateEmbeddings(context.Background(), &minimax.CreateEmbeddingsRequest{
		Texts: []string{"hello"},
		Type:  minimax.EmbeddingsQueryType,
	})
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("%#v\n", resp)
}
