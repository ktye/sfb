package s6d

/*
type texture struct {
	m *image.RGBA
}

func (t texture) Release()                {}
func (t texture) Size() image.Point       { return t.m.Max }
func (t texture) Bounds() image.Rectangle { return t.m.Bounds() }
func (t texture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	r := image.Rectangle{dp, dp.Add(sr.Size())}
	draw.Draw(t.m, r, src.RGBA(), sr.Min, draw.Src)
}
func (t texture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	draw.Draw(t.m, t.m.Bounds(), &uniform{src}, image.ZP, op)
}
*/

/*
type Texture interface {
    // Release releases the Texture's resources, after all pending uploads and
    // draws resolve.
    //
    // The behavior of the Texture after Release, whether calling its methods
    // or passing it as an argument, is undefined.
    Release()

    // Size returns the size of the Texture's image.
    Size() image.Point

    // Bounds returns the bounds of the Texture's image. It is equal to
    // image.Rectangle{Max: t.Size()}.
    Bounds() image.Rectangle

    Uploader
}
type Uploader interface {
    // Upload uploads the sub-Buffer defined by src and sr to the destination
    // (the method receiver), such that sr.Min in src-space aligns with dp in
    // dst-space. The destination's contents are overwritten; the draw operator
    // is implicitly draw.Src.
    //
    // It is valid to upload a Buffer while another upload of the same Buffer
    // is in progress, but a Buffer's image.RGBA pixel contents should not be
    // accessed while it is uploading. A Buffer is re-usable, in that its pixel
    // contents can be further modified, once all outstanding calls to Upload
    // have returned.
    //
    // TODO: make it optional that a Buffer's contents is preserved after
    // Upload? Undoing a swizzle is a non-trivial amount of work, and can be
    // redundant if the next paint cycle starts by clearing the buffer.
    //
    // When uploading to a Window, there will not be any visible effect until
    // Publish is called.
    Upload(dp image.Point, src Buffer, sr image.Rectangle)

    // Fill fills that part of the destination (the method receiver) defined by
    // dr with the given color.
    //
    // When filling a Window, there will not be any visible effect until
    // Publish is called.
    Fill(dr image.Rectangle, src color.Color, op draw.Op)
}
*/
