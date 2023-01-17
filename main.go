package main

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/rivo/tview"
	"google.golang.org/api/iterator"
	"os"
	"time"
)

func main() {

	BUCKET := os.Getenv("BUCKET")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	it := client.Bucket(BUCKET).Objects(ctx, nil)

	dir := tview.NewTreeView()
	root := tview.NewTreeNode(".")
	dir.SetRoot(root)

	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			panic(err)
		}

		node := tview.NewTreeNode(battrs.Name)
		root.AddChild(node)
	}

	app := tview.NewApplication()

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
