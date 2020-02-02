package upstream

import "testing"

func TestMockBNIStorage(t *testing.T) {
	b := &MockBNIStorage{}
	tester := bNIStorageTestSet{b}
	b.insertMockData(tester.MockingData())
	tester.RunTests(t)
}
