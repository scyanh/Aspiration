package mapper

import (
	"errors"
	"unicode"
)

type ObjMapper struct {
	frequency int
	Text      string
	offset    int
}

type Interface interface {
	TransformRune(pos int, runes []rune)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	r := i.GetValueAsRuneSlice()
	for pos, _ := range r {
		i.TransformRune(pos, r)
	}
}

func (o *ObjMapper) TransformRune(pos int, runes []rune) {
	runeValid := false
	if runes[pos] >= 97 && runes[pos] <= 122 ||
		runes[pos] >= 65 && runes[pos] <= 90 ||
		runes[pos] >= 48 && runes[pos] <= 57 {
		runeValid = true
	}
	if !runeValid {
		o.offset++
		return
	}

	if o.frequency == 1 {
		runes[pos] = unicode.ToUpper(runes[pos])
	} else if pos > 0 && (pos+1-o.offset)%o.frequency == 0 {
		runes[pos] = unicode.ToUpper(runes[pos])
	} else {
		runes[pos] = unicode.ToLower(runes[pos])
	}

	o.Text = string(runes)
}

func (o *ObjMapper) GetValueAsRuneSlice() []rune {
	return []rune(o.Text)
}
func (o *ObjMapper) String() string {
	return o.Text
}

func NewSkipString(frequency int, text string) (ObjMapper, error) {
	o := ObjMapper{
		frequency: frequency,
		Text:      text,
		offset:    0,
	}
	if frequency == 0 {
		return o, errors.New("frequency must be greater than 0")
	}
	return o, nil
}
