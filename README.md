mmslice
=======

[![Build Status](https://travis-ci.org/brentp/mmslice.svg?branch=master)](https://travis-ci.org/brentp/mmslice)

mmslice makes it easy to mmap an slice of []uint16 to a file for
reading and/or writing. It uses [this excellent library for mmap'ing](https://github.com/edsrzf/mmap-go).

```Go

import "github.com/brentp/mmslice/uint16mm"

// create a new file mapping 2000 uint16s
f, _ := uint16mm.Create("/tmp/t.bin", 2000)

// The mapped data is in .A
len(f.A) // 2000

for i := 0; i < 1000; i++ {
	f.A[i] = 22
}
// all entries in the slice are initialized to 0.
for i := 1000; i < 2000; i++ {
	if f.A[i] != 0 {
		panic("WTF!")
	}
}

f.Close() // unmaps the memory and closes the filehandle.
```

To open for reading:

```Go

f, _ := uint16mm.Open("/tmp/t.bin", mmap.RDONLY)
len(f.A) // 2000

f.A[200] // 22
f.A[1009] // 0

f.Close()
```


ToDo
----

there is a code-generator in scripts/ (a bash script with sed commands)
to make this work with other types. Use go generate or just commit those to the repo.
