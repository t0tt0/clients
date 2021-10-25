package wrapper

func Wrap(code int, err error) error {
	return wrapToStackError(3, code, err)
}

func WrapString(code int, err string) error {
	return wrapStringToStackError(3, code, err)
}

func WrapCode(code int) error {
	return wrapCodeToStackError(3, code)
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
	if err == nil {
		return nil, false
	}
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
	if s == nil {
		return nil, false
	}
	return StackFromString(s.Error())
}

type Describer struct {
	Pack, Rel string
}

func (d Describer) Describe(e error) string {
	if e == nil {

	}
	if frames, ok := StackFromError(e); ok {
		s, err := frames.Rel(d.Pack, d.Rel)
		if err != nil {
			return e.Error()
		}
		return s
	}
	return e.Error()
}

func Describe(e error) string {
	if frames, ok := StackFromError(e); ok {
		return frames.String()
	}
	return e.Error()
}
