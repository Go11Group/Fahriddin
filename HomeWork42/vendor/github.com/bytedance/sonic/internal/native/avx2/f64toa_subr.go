// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__f64toa = 32
)

const (
    _stack__f64toa = 56
)

const (
    _size__f64toa = 4704
)

var (
    _pcsp__f64toa = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {4614, 56},
        {4618, 48},
        {4619, 40},
        {4621, 32},
        {4623, 24},
        {4625, 16},
        {4627, 8},
        {4631, 0},
        {4694, 56},
    }
)

var _cfunc_f64toa = []loader.CFunc{
    {"_f64toa_entry", 0,  _entry__f64toa, 0, nil},
    {"_f64toa", _entry__f64toa, _size__f64toa, _stack__f64toa, _pcsp__f64toa},
}
