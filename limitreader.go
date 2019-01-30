func (l *LimitedReader) Read(p []byte) (n int, err error) {
  if l.N <= 0 { 
   return 0, EOF 
  }
  if int64(len(p)) > l.N {
   p = p[0:l.N]
  }
  n, err = l.R.Read(p)
  l.N -= int64(n)
  return
}

