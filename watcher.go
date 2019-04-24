package gomocket

import "testing"

var Watcher *MockWatcher

type MockWatcher struct {
    queries []QueryMock
    called []int64
}

func (self MockWatcher) GetQueryMocks() []QueryMock {
    return self.queries
}

func (self *MockWatcher) Reset() {
    self.queries = make([]QueryMock, 0)
    self.called = make([]int64, 0)
}

func (self *MockWatcher) WatchQuery(query QueryMock) {
    self.queries = append(self.queries, query)
}

func (self MockWatcher) AssertWasCalled(query QueryMock, t *testing.T) {
    if self.WasCalled(query) == false {
        t.Fail()
    }
}

func (self MockWatcher) WasCalled(query QueryMock) bool {
    called := false

    for _, id := range self.called {
        if id == query.GetId() {
            called = true
            break
        }
    }

    return called
}

func (self *MockWatcher) ProcessQuery(query string, args []driver.NamedValue) {
    // @todo Implement here every check
}

func init() {
    Watcher = &MockWatcher{}
}
