// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package ffmpeg

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

func newRational(r C.struct_AVRational) Rational {
	return NewRational(int(r.num), int(r.den))
}
