// Package marray makes it simple to map an array of integers to file.
package marray

import (
	"encoding/binary"
	"fmt"
	"os"
	"unsafe"

	mmap "github.com/edsrzf/mmap-go"
)

// An A is the type mapped by this lib
type A []uint16

const maxSize = 1<<31 - 1

// Marray is acts like a slice of uint16.
type Marray struct {
	A
	f   *os.File
	Map mmap.MMap
}

// Close the underlying file-handle
func (m *Marray) Close() error {
	m.Map.Flush()
	m.Map.Unmap()
	m.A = nil
	return m.f.Close()
}

// Flush data to the map
func (m *Marray) Flush() error {
	return m.Map.Flush()
}

// Open return s an Marray object given a file to map.
func Open(f *os.File, mode int) (*Marray, error) {
	var anon int
	if f == nil {
		anon = 1
	}
	b, err := mmap.Map(f, mode, anon)
	if err != nil {
		return nil, err
	}
	len, err := f.Seek(0, 2)
	if err != nil {
		return nil, err
	}
	sz := int64(unsafe.Sizeof(uint16(0)))
	arr := (*[maxSize]uint16)(unsafe.Pointer(&b[0]))[:len/sz]
	return &Marray{arr, f, b}, nil
}

// Create opens a new Marray for reading and writing.
func Create(path string, length int64) (*Marray, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	sz := int64(unsafe.Sizeof(uint16(0)))
	sizeInBytes := sz * length
	if sizeInBytes > maxSize {
		f.Close()
		return nil, fmt.Errorf("length %d too big to map to %s", length, path)
	}
	f.Seek(sizeInBytes-sz, 0)
	err = binary.Write(f, binary.LittleEndian, uint16(0))
	if err != nil {
		f.Close()
		return nil, err
	}
	return Open(f, mmap.RDWR)
}
