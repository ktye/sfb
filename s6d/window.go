package s6d

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
)

type window struct {
}

func (w window) Release()                    {}
func (w window) Send(event interface{})      {}
func (w window) SendFirst(event interface{}) {}
func (w window) NextEvent() interface{}
func (w window) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {}
func (w window) Fill(dr image.Rectangle, src color.Color, op draw.Op) {

}

func (w window) Draw(src2dst f64.Aff3, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
}
func (w window) DrawUniform(src2dst f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
}
func (w window) Copy(dp image.Point, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
}
func (w window) Scale(dr image.Rectangle, src screen.Texture, sr image.Rectangle, op draw.Op, opts *screen.DrawOptions) {
}
func (w window) Publish() screen.PublishResult { return screen.PublishResult{true} }

/*
type Window interface {
    // Release closes the window.
    //
    // The behavior of the Window after Release, whether calling its methods or
    // passing it as an argument, is undefined.
    Release()

    EventDeque

    Uploader

    Drawer

    // Publish flushes any pending Upload and Draw calls to the window, and
    // swaps the back buffer to the front.
    Publish() PublishResult
}
type EventDeque interface {
    // Send adds an event to the end of the deque. They are returned by
    // NextEvent in FIFO order.
    Send(event interface{})

    // SendFirst adds an event to the start of the deque. They are returned by
    // NextEvent in LIFO order, and have priority over events sent via Send.
    SendFirst(event interface{})

    // NextEvent returns the next event in the deque. It blocks until such an
    // event has been sent.
    //
    // Typical event types include:
    //	- lifecycle.Event
    //	- size.Event
    //	- paint.Event
    //	- key.Event
    //	- mouse.Event
    //	- touch.Event
    // from the golang.org/x/mobile/event/... packages. Other packages may send
    // events, of those types above or of other types, via Send or SendFirst.
    NextEvent() interface{}
}
type Drawer interface {
    // Draw draws the sub-Texture defined by src and sr to the destination (the
    // method receiver). src2dst defines how to transform src coordinates to
    // dst coordinates. For example, if src2dst is the matrix
    //
    // m00 m01 m02
    // m10 m11 m12
    //
    // then the src-space point (sx, sy) maps to the dst-space point
    // (m00*sx + m01*sy + m02, m10*sx + m11*sy + m12).
    Draw(src2dst f64.Aff3, src Texture, sr image.Rectangle, op draw.Op, opts *DrawOptions)

    // DrawUniform is like Draw except that the src is a uniform color instead
    // of a Texture.
    DrawUniform(src2dst f64.Aff3, src color.Color, sr image.Rectangle, op draw.Op, opts *DrawOptions)

    // Copy copies the sub-Texture defined by src and sr to the destination
    // (the method receiver), such that sr.Min in src-space aligns with dp in
    // dst-space.
    Copy(dp image.Point, src Texture, sr image.Rectangle, op draw.Op, opts *DrawOptions)

    // Scale scales the sub-Texture defined by src and sr to the destination
    // (the method receiver), such that sr in src-space is mapped to dr in
    // dst-space.
    Scale(dr image.Rectangle, src Texture, sr image.Rectangle, op draw.Op, opts *DrawOptions)
}
type PublishResult struct {
    // BackBufferPreserved is whether the contents of the back buffer was
    // preserved. If false, the contents are undefined.
    BackBufferPreserved bool
}
*/
