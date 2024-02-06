package coremodule

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/mongodbinteractions"
)

func NewVerifiedTheHiveFormatAlert(
	input <-chan datamodels.ChanOutputDecodeJSON,
	done <-chan bool,
	esm *elasticsearchinteractions.ElasticSearchModule,
	mongodbm *mongodbinteractions.MongoDBModule,
	logging chan<- datamodels.MessageLogging,
) {

}
