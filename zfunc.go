package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

type ZFunc struct {
	desc string
	f    func(z, c complex128) complex128
}

func (zf ZFunc) String() string {
	return fmt.Sprintf("ZFunc: %s", zf.desc)
}

var ZFuncLib = map[string]ZFunc{
	"burning ship": {
		desc: `Burning Ship :: (|x| + i|y|)^2 + c`,
		f: func(z, c complex128) complex128 {
			return cmplx.Pow(-complex(math.Abs(real(z)), math.Abs(imag(z))), 2.0) - c
		},
	},

	"klein": {
		desc: `Klein :: z^2 - y + i|x| + c`,
		f: func(z, c complex128) complex128 {
			return cmplx.Pow(z, 2.0) + complex(-imag(z), math.Abs(real(z))) + c
		},
	},

	"mandelbrot": {
		desc: `Mandelbrot :: z^2 + c`,
		f: func(z, c complex128) complex128 {
			return cmplx.Pow(z, 2.0) + c
		},
	},
}
