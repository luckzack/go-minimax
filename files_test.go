package minimax_test

import (
	"context"
	"fmt"
	"github.com/luckpunk/go-minimax"
	"testing"
)

// go test -v -test.run=TestUploadFile ./
func TestUploadFile(t *testing.T) {
	res, err := client.UploadFile(context.Background(), &minimax.FileRequest{
		Purpose:  minimax.Retrieval,
		FilePath: "./testdata/wonderland.txt",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}

func TestListFiles(t *testing.T) {
	res, err := client.ListFiles(context.Background(), minimax.Retrieval)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range res.Files {
		fmt.Printf("%+v\n", v)
	}
}

func TestRetrieveFile(t *testing.T) {
	res, err := client.RetrieveFile(context.Background(), -1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", res.File)
}

func TestDeleteFile(t *testing.T) {
	res, err := client.DeleteFile(context.Background(), &minimax.DeleteFileRequest{
		FileId: -1,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}
