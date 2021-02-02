package iterator_test

import (
	"testing"

	"github.com/y00rb/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
