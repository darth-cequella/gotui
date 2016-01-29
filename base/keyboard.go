package base

import "bytes"

type Keyboard struct {
}

//SPECIAL
func (this *Keyboard) Backspace(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{127,0,0}, buffer)
	return
}
func (this *Keyboard) Enter(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{10,91,67}, buffer)
	return
}
func (this *Keyboard) Space(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{32,91,67}, buffer)
	return
}
func (this *Keyboard) F1(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,80}, buffer)
	return
}
func (this *Keyboard) F2(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,81}, buffer)
	return
}
func (this *Keyboard) F3(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,82}, buffer)
	return
}
func (this *Keyboard) F4(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,83}, buffer)
	return
}
func (this *Keyboard) F5(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,91,49}, buffer)
	return
}
func (this *Keyboard) F6(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{53,126,49}, buffer)
	return
}
func (this *Keyboard) F7(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,91,49}, buffer)
	return
}
func (this *Keyboard) F8(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{55,126,49}, buffer)
	return
}
func (this *Keyboard) F9(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,88}, buffer)
	return
}
func (this *Keyboard) F10(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,89}, buffer)
	return
}
func (this *Keyboard) F11(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,90}, buffer)
	return
}
func (this *Keyboard) F12(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,79,91}, buffer)
	return
}

//ARROWS
func (this *Keyboard) UpArrow(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,91,65}, buffer)
	return
}
func (this *Keyboard) BottomArrow(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,91,66}, buffer)
	return
}
func (this *Keyboard) RightArrow(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,91,67}, buffer)
	return
}
func (this *Keyboard) LeftArrow(buffer []byte) (out bool) {
	out = bytes.Equal([]byte{27,91,68}, buffer)
	return
}

func (this *Keyboard) KeyA(buffer []byte) bool {
	possible := [][]byte{ []byte{byte('A'),0,0}, []byte{byte('a'),0,0} }
	out := bytes.Equal(possible[0], buffer)||bytes.Equal(possible[1], buffer)
	return out
}