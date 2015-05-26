// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package tiff

import (
	"image"
	"image/color"
	"reflect"
)

type imageRGB96i struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageRGB96i returns a new imageRGB96i with the given bounds.
func newImageRGB96i(r image.Rectangle) *imageRGB96i {
	return new(imageRGB96i).Init(make([]uint8, 12*r.Dx()*r.Dy()), 12*r.Dx(), r)
}

func (p *imageRGB96i) Init(pix []uint8, stride int, rect image.Rectangle) *imageRGB96i {
	*p = imageRGB96i{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *imageRGB96i) Pix() []byte           { return p.M.Pix }
func (p *imageRGB96i) Stride() int           { return p.M.Stride }
func (p *imageRGB96i) Rect() image.Rectangle { return p.M.Rect }
func (p *imageRGB96i) Channels() int         { return 3 }
func (p *imageRGB96i) Depth() reflect.Kind   { return reflect.Int32 }

func (p *imageRGB96i) ColorModel() color.Model { return colorRGB96iModel }

func (p *imageRGB96i) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageRGB96i) At(x, y int) color.Color {
	return p.RGB96iAt(x, y)
}

func (p *imageRGB96i) RGB96iAt(x, y int) colorRGB96i {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorRGB96i{}
	}
	i := p.PixOffset(x, y)
	return pRGB96iAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageRGB96i) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*12
}

func (p *imageRGB96i) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorRGB96iModel.Convert(c).(colorRGB96i)
	pSetRGB96i(p.M.Pix[i:], c1)
	return
}

func (p *imageRGB96i) SetRGB96i(x, y int, c colorRGB96i) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGB96i(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageRGB96i) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageRGB96i{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageRGB96i).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageRGB96i) Opaque() bool {
	return true
}