package mongo

import(
	"github.com/mongodb/mongo-go-driver/bson"
)

func MakeDoc(doc []map[string]interface{}) *bson.Document{
	inter := new(bson.ElementConstructor)
	var el_array []*bson.Element
	for _, m := range doc{

		for k, v := range m {
			el := inter.Interface(k, v)
			el_array = append(el_array, el)
		}
	}
	return bson.NewDocument(el_array...)
}