/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-29 14:00
**/

package lemongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndReplace struct {
	collection               *mongo.Collection
	findOneAndReplaceOptions options.FindOneAndReplaceOptions
	filter                   interface{}
	replacement              interface{}
}

func (f *FindOneAndReplace) Sort(sort interface{}) *FindOneAndReplace {
	f.findOneAndReplaceOptions.Sort = sort
	return f
}

func (f *FindOneAndReplace) Projection(projection interface{}) *FindOneAndReplace {
	f.findOneAndReplaceOptions.Projection = projection
	return f
}

func (f *FindOneAndReplace) Upsert() *FindOneAndReplace {
	var t = true
	f.findOneAndReplaceOptions.Upsert = &t
	return f
}

func (f *FindOneAndReplace) ReturnDocument() *FindOneAndReplace {
	var t = options.After
	f.findOneAndReplaceOptions.ReturnDocument = &t
	return f
}

func (f *FindOneAndReplace) Get(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOneAndReplace(context.Background(), f.filter, f.replacement, &f.findOneAndReplaceOptions)}
	return res.Get(result)
}
