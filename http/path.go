package http

import (
	"strconv"
)

type PathValue string

func (p PathValue) TryInt(i *int) bool {
	v, err := p.parse(32)
	if err != nil {
		return false
	}

	*i = int(v)

	return true
}

func (p PathValue) TryInt64(i *int64) bool {
	v, err := p.parse(64)
	if err != nil {
		return false
	}

	*i = v

	return true
}

func (p PathValue) IntOrDefault(r int) int {
	v, err := p.parse(32)
	if err != nil {
		return r
	}

	return int(v)
}

func (p PathValue) Int64OrDefault(r int64) int64 {
	v, err := p.parse(64)
	if err != nil {
		return r
	}

	return v
}

func (p PathValue) Int() (int, error) {
	v, err := p.parse(64)
	return int(v), err
}

func (p PathValue) Int64() (int64, error) {
	return p.parse(64)
}

func (p PathValue) parse(bitSize int) (int64, error) {
	value := string(p)

	v, err := strconv.ParseInt(value, 10, bitSize)
	if err != nil {
		return -1, err
	}

	return v, nil
}
