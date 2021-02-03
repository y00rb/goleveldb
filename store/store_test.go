package store

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/raft"
)

func testLevelDBStoreLow(t testing.TB) *LevelDBStore {
	return testLevelDBStore(t, Low)
}

func testLevelDBStoreHigh(t testing.TB) *LevelDBStore {
	return testLevelDBStore(t, High)
}

func testLevelDBStore(t testing.TB, durability Level) *LevelDBStore {
	fh, err := ioutil.TempFile("", "bolt")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	os.Remove(fh.Name())

	// Successfully creates and returns a store
	store, err := NewLevelDBStore(fh.Name(), durability)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	return store
}

func testRaftLog(idx uint64, data string) *raft.Log {
	return &raft.Log{
		Data:  []byte(data),
		Index: idx,
	}
}
