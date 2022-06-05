package dao

import (
	"context"
	"example/user/industrialExercise/internal/model"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(index string) (chan model.Item, error) {
	//Use elastic search as data storage
	//Must turn off sniff for docker-based server
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan model.Item)
	go func(){
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil{
				log.Printf("Item saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}


func Save(client *elastic.Client, index string, item model.Item) error {
	indexService := client.Index().Index(index).BodyJson(item)
	_, err := indexService.Do(context.Background())
	if err != nil{
		return err
	}
	return nil
}