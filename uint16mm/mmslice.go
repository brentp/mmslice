// Package uint16mm makes it simple to map a slice of integers to file.
package uint16mm

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	mmap "github.com/edsrzf/mmap-go"
)

// An A is the type mapped by this lib
type A []uint16

const maxSize int64 = 1<<37 - 1

// Slice holds uint16's backed by a mmap.
type Slice struct {
	A
	f   *os.File
	Map mmap.MMap
}

// Close the underlying file-handle
func (m *Slice) Close() error {
	m.Map.Flush()
	m.Map.Unmap()
	m.A = nil
	if m.f == nil {
		return nil
	}
	return m.f.Close()
}

// Flush data to the map
func (m *Slice) Flush() error {
	return m.Map.Flush()
}

// Open return s an Marray object given a file to map.
// if f is nil then the 2nd argument is the size of an
// anonymously mapped array.
func Open(f *os.File, mode int) (*Slice, error) {
	var b mmap.MMap
	var err error
	var length int64
	sz := int64(unsafe.Sizeof(uint16(0)))
	if f == nil {
		length = int64(mode) * int64(sz)
		mode = mmap.RDWR
		b, err = mmap.MapRegion(nil, int(length), mode, 1, 0)
	} else {
		b, err = mmap.Map(f, mode, 0)
	}
	if err != nil {
		return nil, err
	}
	syscall.Madvise(b, syscall.MADV_SEQUENTIAL|syscall.MADV_WILLNEED)
	if f != nil {
		length, err = f.Seek(0, 2)
		if err != nil {
			return nil, err
		}
	}
	arr := (*[maxSize]uint16)(unsafe.Pointer(&b[0]))[:length/sz]
	return &Slice{arr, f, b}, nil
}

// Create opens a new Marray for reading and writing.
func Create(path string, length int64) (*Slice, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	sz := int64(unsafe.Sizeof(uint16(0)))
	sizeInBytes := sz * length
	if sizeInBytes > maxSize {
		f.Close()
		return nil, fmt.Errorf("length %d too big to map to %s (max: %d)", length, path, maxSize/sz)
	}
	f.Seek(sizeInBytes-sz, 0)
	err = binary.Write(f, binary.LittleEndian, uint16(0))
	if err != nil {
		f.Close()
		return nil, err
	}
	return Open(f, mmap.RDWR)
}
