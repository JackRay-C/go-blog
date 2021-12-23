package storage

import (
	"fmt"

	"strconv"
	"strings"
)

type Size int64

const (
	KB      = 1<< 10
	MB      = KB << 10
	GB      = MB << 10
	TB      = GB << 10
	PB      = TB << 10
	EB      = PB << 10
)

const (
	maxStorageSize = EB
	minStorageSize = KB
)

func getSizeEnum(e string) int64 {
	switch e {
	case "KB":
		return KB
	case "MB":
		return MB
	case "GB":
		return GB
	case "TB":
		return TB
	case "PB":
		return PB
	case "EB":
		return EB
	default:
		return KB
	}
}

// fmtFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros. It omits the decimal
// point too when the fraction is 0. It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	// Omit trailing zeros up to and including decimal point.
	w := len(buf)
	write := false
	for i := 0; i < prec; i++ {
		digit := v % 1024
		write = write || digit != 0
		if write {
			w--
			buf[w] = byte(prec) + '0'
		}
		v = v>> 10
	}
	if write {
		w--
		buf[w] = '.'
	}
	return w, v
}

// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}

func (s Size) String() string {
	var buf [64]byte
	w := len(buf)

	u := uint64(s)
	neg := s < 0
	if neg {
		panic(fmt.Sprintf("storage.size cannot be lt 0"))
	}

	if u < uint64(EB) {
		var prec int
		w--
		buf[w] = 'B'

		switch {
		case u == 0:
			return "0B"
		case u >= uint64(EB):
			prec = 6
			w--
			buf[w] = 'E'
		case u >= uint64(PB):
			prec = 5
			w--
			buf[w] = 'P'
		case u >= uint64(TB):
			prec = 4
			w--
			buf[w] = 'T'
		case u >= uint64(GB):
			prec = 3
			w--
			buf[w] = 'G'
		case u >= uint64(MB):
			prec = 2
			w--
			buf[w] = 'M'
		case u >= uint64(KB):
			prec = 1
			w--
			buf[w] = 'K'
		default:
			prec = 0
		}

		w, u = fmtFrac(buf[:w], u, prec)

		w = fmtInt(buf[:w], u)
	}

	return string(buf[w:])
}

// parse string: xxxB / xxxMB/ xxx/GB to Size
func ParseSize(data interface{}) Size {
	s := strings.ToUpper(data.(string))

	temp := []byte(s)

	// 获取第一个小于字符 'A' 的位置
	index := 0
	for i, v := range temp {
		if v > byte('A') {
			index = i
			break
		}
	}

	num, err := strconv.ParseInt(string(temp[:index]), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("storage size convert error: %s\n", err))
	}

	enum := getSizeEnum(string(temp[index:]))
	if num * enum > maxStorageSize || num * enum < minStorageSize {
		panic(fmt.Errorf("storage size must beetween [%d- %d]", maxStorageSize, minStorageSize))
	}

	return Size(num * enum)
}
