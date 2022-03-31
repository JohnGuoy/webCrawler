package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v6"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			save(item)
			itemCount++
		}

	}()

	return out
}

func save(item interface{}) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Index().Index("dating_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
