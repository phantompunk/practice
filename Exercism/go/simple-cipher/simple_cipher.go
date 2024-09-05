package cipher

import (
	"strings"
)

type shift struct {
	distance int
}

type vigenere struct {
	distance []int
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	return shift{distance: distance}
}

func NewVigenere(key string) Cipher {
	return nil
}

func shiftRune(r rune, distance int) rune {
	if distance < 0 {
		distance = 26 + distance
	}
	idx := (int(r-'a') + distance) % 26
	return rune('a' + idx)
}

func shiftEncode(r rune, shift int) rune {
	return rune((int(r-'a')+shift)%26 + 'a')
}

func (s shift) Encode(phrase string) string {
	e := ""
	for _, v := range strings.ToLower(phrase) {
		if v >= 'a' && v <= 'z' {
			e += string(shiftEncode(v, s.distance))
		}
	}
	return e
}

func caesarShiftDecode(r rune) rune {
	return (((r - 'a') + 23) % 26) + 'a'
}

func (c shift) Decode(phrase string) string {
	d := ""
	for _, v := range strings.ToLower(phrase) {
		if v >= 'a' && v <= 'z' {
			d += string(shiftEncode(v))
		}
	}
	return d
}
