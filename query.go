package gomocket

var Counter int64

type QueryMock struct {
    id int64
    definitions map[string]interface{}
}

func (self *QueryMock) SetId(id int64) {
    self.id = id
}

func (self QueryMock) GetId() int64 {
    return self.id
}

func (self *QueryMock) SetDefinitions(definitions map[string]interface{}) {
    self.definitions = definitions
}

func (self QueryMock) GetDefinitions() map[string]interface{} {
    return self.definitions
}

func NewQueryMock(definitions map[string]interface{}) QueryMock {
    queryMock := QueryMock{}
    queryMock.SetId(Counter)
    queryMock.SetDefinitions(definitions)

    Counter = Counter + 1

    return queryMock;
}

func inti() {
    Counter = 0
}
