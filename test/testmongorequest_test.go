package testhandlerobservables

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CollectionName = "case_collection"
	//CollectionName = "stix_object_collection"
)

func NewConnect() (*mongo.Collection, error) {
	var (
		Host = "192.168.9.208"
		Port = 27117
		//Port = 37017
		NameDB = "placeholder_elasticsearch"
		//NameDB = "isems-mrsict"
		User = "module_placeholder_elasticsearch"
		//User = "module-isems-mrsict"
		Passwd = "gDbv5cf7*F2"
		//Passwd = "vkL6Znj$Pmt1e1"
	)

	clientOption := options.Client().SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    NameDB,
		Username:      User,
		Password:      Passwd,
	})

	confPath := fmt.Sprintf("mongodb://%s:%d/%s", Host, Port, NameDB)
	ctx, _ := context.WithTimeout(context.Background(), 7*time.Second)
	conn, err := mongo.Connect(ctx, clientOption.ApplyURI(confPath))
	if err != nil {
		return nil, err
	}

	return conn.Database(NameDB).Collection(CollectionName), nil
}

func TestMongoReguestAggregateDataType(t *testing.T) {
	collection, err := NewConnect()
	assert.NoError(t, err)

	opts := options.Aggregate().SetAllowDiskUse(true)

	cur, err := collection.Aggregate(
		context.Background(),
		bson.A{
			bson.D{
				{"$unwind",
					bson.D{
						{"path", "$observablesmessagethehive.observables"},
						{"includeArrayIndex", "string"},
						{"preserveNullAndEmptyArrays", true},
					},
				},
			},
			bson.D{
				{"$group",
					bson.D{
						{"_id", primitive.Null{}},
						{"distinctGenres", bson.D{{"$addToSet", "$observablesmessagethehive.observables.dataType"}}},
					},
				},
			},
		},
		opts)
	assert.NoError(t, err)

	type Distinct struct {
		DistinctGenres []string `bson:"distinctGenres"`
	}

	list := []string(nil)
	for cur.Next(context.Background()) {
		dist := Distinct{}
		err = cur.Decode(&dist)
		assert.NoError(t, err)

		list = append(list, dist.DistinctGenres...)
	}

	for _, v := range list {
		log.Println(v)
	}

	assert.Equal(t, len(list), 14)
}
