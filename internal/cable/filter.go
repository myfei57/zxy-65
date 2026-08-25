package cable

type Filter struct {
	size  int
	buf   []float64
	count int
}

func NewFilter(size int) *Filter {
	return &Filter{size: size, buf: make([]float64, size)}
}

func (f *Filter) Add(sample float64) float64 {
	if f.count < f.size {
		f.buf[f.count] = sample
		f.count++
	} else {
		copy(f.buf, f.buf[1:])
		f.buf[f.size-1] = sample
	}
	return f.Average()
}

func (f *Filter) Average() float64 {
	if f.count == 0 {
		return 0
	}
	sum := 0.0
	for i := 0; i < f.count; i++ {
		sum += f.buf[i]
	}
	return sum / float64(f.count)
}

func (f *Filter) Reset() {
	f.count = 0
	for i := range f.buf {
		f.buf[i] = 0
	}
}
