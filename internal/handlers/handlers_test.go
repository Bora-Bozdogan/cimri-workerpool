package handlers

import (
	"testing"
)

func TestHandleUpdate(t *testing.T) {
	MTest.TestHandleUpdate(t)
}

func (m *MainTest) TestHandleUpdate(t *testing.T) {
	m.handler.HandleWorkers()

	ok := m.client.CheckSuccess() //have checksuccess that checks if faultyitem-validitem counters 0 etc

	if !ok {
		t.Fail()
	}
}
