package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/rivo/tview"
	"google.golang.org/api/iterator"
	"time"
)

const PROJECT_ID = "core-dev-341718"

func main() {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	it := client.Buckets(ctx, PROJECT_ID)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(battrs.Name)

	}

	app := tview.NewApplication()

	dir := tview.NewBox().SetBorder(true).SetTitle("Dir")

	leftPane := tview.NewBox().SetBorder(true).SetTitle("Left")
	rightPane := tview.NewBox().SetBorder(true).SetTitle("Right")

	flex := tview.NewFlex().
		AddItem(dir, 40, 1, true).
		AddItem(leftPane, 0, 1, false).
		AddItem(rightPane, 0, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
