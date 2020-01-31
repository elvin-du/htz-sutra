package model

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"htz/sutra/common/types"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"htz/sutra/common/database"
	"time"
)

var (
	DB_COLLECTION_FILE = "file"
)

func NewFileModel(dbName string) *FileModel {
	am := &FileModel{database.DefaultDB.Client.Database(dbName).Collection(DB_COLLECTION_FILE)}
	return am
}

type FileModel struct {
	collection *mongo.Collection
}

func (fm *FileModel) InsertBlock(f *types.File) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := fm.collection.InsertOne(ctx, f)
	if nil != err {
		log.Errorln(err)
		return "", err
	}

	log.Debugln("insertid:", res.InsertedID)
	return fmt.Sprintf("%s", res.InsertedID), nil
}

//func (bm *AddressesModel) Find(address string) (*Address, error) {
//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//	result := bm.collection.FindOne(ctx, bson.M{"address": address})
//	if result.Err() != nil {
//		log.Errorln(result.Err())
//		return nil, result.Err()
//	}
//	addr := &Address{}
//	err := result.Decode(addr)
//	if err != nil {
//		log.Errorln(result.Err())
//		return nil, err
//	}
//	return addr, nil
//}
//
//func (bm *AddressesModel) Upsert(addr *Address) error {
//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//	filter := bson.M{"address": addr.Address}
//	update := bson.M{"$set": addr}
//	_, err := bm.collection.UpdateOne(ctx, filter, update)
//	if nil != err {
//		log.Errorln(err)
//		return err
//	}
//
//	return nil
//}
