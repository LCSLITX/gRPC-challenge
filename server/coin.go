package main

type Coin struct {
	Id 					uint32			`bson:"_id,omitempty"`
	Name 				string 			`bson:"name"`
	Short 			string 			`bson:"short"`
	Votes 			int64 			`bson:"votes"`
}
