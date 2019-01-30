package main

type MultiWriter struct {
	w []io.Writer
}

type multiWriterError []error

func (mwe *multiWriterError) Error() string {
	s = ""
	for _, err := range mwe {
		s += err.Error() + "\n"
	}
	return s
}

func (mw *MultiWriter) Write(b []byte) (int, error) {
	if len(mw) == 0 {
		return len(b), nil
	}

	work := func() <-chan error {
		errchan := make(chan error, len(mw.w))
		go func() {
			wg := sync.WaitGroup{}
			wg.Add(len(mw.w))
			for i := 0; i < len(mw.w); i++ {
				go func(i int) {
					nb, err := mw.w[i].Write(b)
					if err != nil {
						errchan <- err
					}
					wg.Done()
					return nil
				}()
			}
		}()
		return errchan
	}

	rc := []error{}

	for err := range work() {
		rc = append(rc, err)
	}

	return len(mw), multiWriterError(rc)
}

func MultiWriter(writers ...io.Writer) *MultiWriter {
	rc := MultiWriter{}
	for _, w := range writers {
		rc.w = append(rc.w, w)
	}
	return &rc
}
