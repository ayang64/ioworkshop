

* Copy Copies from a reader to a Writer

- Copy(dst, src) copies from src to dst until EOF is reached or an error
	occurrs.

- CopyN(dst, src, n) copies n bytes from src to dst.  Under the hood, it simply
	creates an io.LimtedReader from src.

- CopyBuffer(dst, src, buf) - Copy from src to dest, but allows the user to
	supply an underlying buffer

Lets pick one to talk about in more detail:

Under the hood, io.Copy() is wrapper for the private io.copyBuffer() which, if
left on its own will allocate (using make()) a 32k intermediate buffer that is
used to transfer bytes from the source to destination.

This is how io.Copy() works -- every call will to io.Copy() will allocate 32k.

io.CopyBuffer allows you to supply a byte slice buffer and re-use it across
copy calls.  This will reduce the number of total allocation.  You can also
specify a specific buffer size which could help you optimize your buffer for
whatever your use case (disk io?  network io?)

- show benchmarks with -benchmem
- show pprof mem profile
	- go test -bench <benchname> -memprofile=buffered.protobuf.gz
	- go test -bench <benchname> -memprofile=unbuffered.protobuf.gz
	- go tool pprof *
	- top10 -cum
	- go test -bench <benchname> -trace ./trace.out ; go tool trace ./trace.out
