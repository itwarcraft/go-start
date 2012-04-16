package mongo

import (
	"launchpad.net/mgo/bson"
	"github.com/ungerik/go-start/model"
	"launchpad.net/mgo"
)

///////////////////////////////////////////////////////////////////////////////
// QueryError

// QueryError implements Query, but all it does is hold and return an
// error generated by another query.
type QueryError struct {
	parent Query
	Err    error
}

func (self *QueryError) subDocumentSelector() string {
	return ""
}

func (self *QueryError) bsonSelector() bson.M {
	return nil
}

func (self *QueryError) mongoQuery() (q *mgo.Query, err error) {
	return nil, self.Err
}

func (self *QueryError) Selector() string {
	return ""
}

func (self *QueryError) ParentQuery() Query {
	return self
}

func (self *QueryError) Collection() *Collection {
	var query Query
	for query = self; query != nil; query = query.ParentQuery() {
		if collection, ok := query.(*Collection); ok {
			return collection
		}
	}
	return nil
}

func (self *QueryError) SubDocument(selector string) Query {
	return self
}

func (self *QueryError) Skip(int) Query {
	return self
}

func (self *QueryError) Limit(int) Query {
	return self
}

func (self *QueryError) Sort(selector string) Query {
	return self
}

func (self *QueryError) SortReverse(selector string) Query {
	return self
}

func (self *QueryError) SortFunc(less func(a, b interface{}) bool) model.Iterator {
	return self.Iterator()
}

func (self *QueryError) IsFilter() bool {
	return false
}

func (self *QueryError) Filter(selector string, value interface{}) Query {
	return self
}

func (self *QueryError) FilterWhere(javascript string) Query {
	return self
}

func (self *QueryError) FilterFunc(passFilter model.FilterFunc) model.Iterator {
	return self.Iterator()
}

func (self *QueryError) FilterRef(selector string, ref ...Ref) Query {
	return self
}

func (self *QueryError) FilterEqualCaseInsensitive(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterNotEqual(selector string, value interface{}) Query {
	return self
}

func (self *QueryError) FilterLess(selector string, value interface{}) Query {
	return self
}

func (self *QueryError) FilterGreater(selector string, value interface{}) Query {
	return self
}

func (self *QueryError) FilterLessEqual(selector string, value interface{}) Query {
	return self
}

func (self *QueryError) FilterGreaterEqual(selector string, value interface{}) Query {
	return self
}

func (self *QueryError) FilterModulo(selector string, divisor, result interface{}) Query {
	return self
}

func (self *QueryError) FilterIn(selector string, values ...interface{}) Query {
	return self
}

func (self *QueryError) FilterNotIn(selector string, values ...interface{}) Query {
	return self
}

func (self *QueryError) FilterAllIn(selector string, values ...interface{}) Query {
	return self
}

func (self *QueryError) FilterArraySize(selector string, size int) Query {
	return self
}

func (self *QueryError) FilterStartsWith(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterStartsWithCaseInsensitive(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterEndsWith(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterEndsWithCaseInsensitive(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterContains(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterContainsCaseInsensitive(selector string, str string) Query {
	return self
}

func (self *QueryError) FilterExists(selector string, exists bool) Query {
	return self
}

func (self *QueryError) Or() Query {
	return self
}

func (self *QueryError) Count() (n int, err error) {
	return 0, self.Err
}

func (self *QueryError) Explain() string {
	return self.Err.Error()
}

func (self *QueryError) One() (document interface{}, err error) {
	return nil, self.Err
}

func (self *QueryError) TryOne() (document interface{}, found bool, err error) {
	return nil, false, self.Err
}

func (self *QueryError) GetOrCreateOne() (document interface{}, found bool, err error) {
	return nil, false, self.Err
}

func (self *QueryError) Iterator() model.Iterator {
	return model.NewErrorIterator(self.Err)
}

func (self *QueryError) OneID() (id bson.ObjectId, err error) {
	return "", self.Err
}

func (self *QueryError) TryOneID() (id bson.ObjectId, found bool, err error) {
	return "", false, self.Err
}

func (self *QueryError) IDs() (ids []bson.ObjectId, err error) {
	return nil, self.Err
}

func (self *QueryError) Refs() (refs []Ref, err error) {
	return nil, self.Err
}

func (self *QueryError) RemoveAll() error {
	return self.Err
}