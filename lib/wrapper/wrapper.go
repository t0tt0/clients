package wrapper

func Wrap(code int, err error) error {
	return wrapToStackError(2, code, err)
}

func WrapString(code int, err string) error {
	return wrapStringToStackError(2, code, err)
}

func WrapCode(code int) error {
	return wrapCodeToStackError(2, code)
}

func WrapN(skip, code int, err error) error {
	return wrapToStackError(skip, code, err)
}

func WrapStringN(skip, code int, err string) error {
	return wrapStringToStackError(skip, code, err)
}

func WrapCodeN(skip, code int) error {
	return wrapCodeToStackError(skip, code)
}

func FromBytes(b []byte) (f Frame, ok bool) {
	f, ok = stackErrorFromBytes(b)
	return
}

func FromString(s string) (f Frame, ok bool) {
	f, ok = stackErrorFromString(s)
	return
}

func FromError(err error) (Frame, bool) {
	return FromString(err.Error())
}

func StackFromString(s string) (fs Frames, ok bool) {
	ok = true
	var f Frame
	for ok {
		if f, ok = FromString(s); ok {
			fs = append(fs, f)
		}
		s = f.GetErr()
		if len(fs) >= 2 {
			fs[len(fs)-2].ReleaseError()
		}
	}
	return fs, len(fs) > 0
}

func StackFromError(s error) (Frames, bool) {
	return StackFromString(s.Error())
}
