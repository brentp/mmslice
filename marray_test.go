package marray_test

import (
	"os"
	"testing"

	"github.com/brentp/marray"
	mmap "github.com/edsrzf/mmap-go"
)

func TestCreate(t *testing.T) {

	f, err := marray.Create("/tmp/t.bin", 2000)
	if err != nil {
		t.Errorf("error creating file")
	}
	_ = f

	if len(f.A) != 2000 {
		t.Errorf("expected length 2000. found %d\n", len(f.A))
	}

	for i := 0; i < 1000; i++ {
		f.A[i] = 22
	}
	for i := 0; i < 1000; i++ {
		if f.A[i] != 22 {
			t.Errorf("expected to find 22. found %d\n", f.A[i])
		}
	}
	for i := 1000; i < 2000; i++ {
		if f.A[i] != 0 {
			t.Errorf("expected to find 0 at index %d. found %d\n", i, f.A[i])
		}
	}

	if err := f.Flush(); err != nil {
		t.Errorf("error on flushing: %s", err)
	}
	if err := f.Close(); err != nil {
		t.Errorf("error on flushing: %s", err)
	}

	fh, err := os.Open("/tmp/t.bin")
	if err != nil {
		t.Errorf("error on opening: %s", err)
	}

	r, err := marray.Open(fh, mmap.RDONLY)
	if err != nil {
		t.Errorf("error on opening: %s", err)
	}
	if len(r.A) != 2000 {
		t.Errorf("expected length 2000. found %d\n", len(r.A))
	}

	for i := 0; i < 1000; i++ {
		if r.A[i] != 22 {
			t.Errorf("expected to find 22. found %d\n", r.A[i])
		}
	}
	for i := 1000; i < 2000; i++ {
		if r.A[i] != 0 {
			t.Errorf("expected to find 0 at index %d. found %d\n", i, r.A[i])
		}
	}
}
