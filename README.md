marray
======

marray makes it easy to mmap an array of []uint16 to a file for
reading and/or writing. It uses [this excellent library for mmap'ing](https://github.com/edsrzf/mmap-go).

```Go

// create a new file mapping 2000 uint16s
f, _ := marray.Create("/tmp/t.bin", 2000)

// The mapped data is in .A
len(f.A) // 2000

for i := 0; i < 1000; i++ {
	f.A[i] = 22
}
// all entries in the array are initialized to 0.
for i := 1000; i < 2000; i++ {
	if f.A[i] != 0 {
		panic("WTF!")
	}
}

f.Close() // unmaps the memory and closes the filehandle.
```

To open for reading:

```Go

f, _ := marray.Open("/tmp/t.bin", mmap.RDONLY)
len(f.A) // 2000

f.A[200] // 22
f.A[1009] // 0

f.Close()
```


ToDo
----

Write a code-generator to make this work for any size data.
For now, everything is uint16.
