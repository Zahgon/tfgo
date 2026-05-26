/*
Copyright 2017-2022 Paolo Galeone. All right reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package image

import (
	tf "github.com/galeone/tensorflow/tensorflow/go"
	"github.com/galeone/tensorflow/tensorflow/go/op"
	"github.com/galeone/tfgo/image/padding"
)

// ReadJPEG reads the JPEG image whose path is `imagePath` that has `channels` channels
// it returns an Image
func ReadJPEG(scope *op.Scope, imagePath string, channels int64) *Image {
	_ = "STUB: not implemented"
	return nil
}

// ReadPNG reads the PNG image whose path is `imagePath` that has `channels` channels
// it returns an Image
func ReadPNG(scope *op.Scope, imagePath string, channels int64) *Image {
	_ = "STUB: not implemented"
	return nil
}

// ReadGIF reads the GIF image whose path is `imagePath` and returns an Image
func ReadGIF(scope *op.Scope, imagePath string) *Image { _ = "STUB: not implemented"; return nil }

// DecodeGif returns a Tensor of type uint8. 4-D with shape [num_frames, height, width, 3]. RGB order

// Read searches for the `imagePath` extensions and uses the Read<format> function
// to decode and load the right image. Panics if the format is unknown.
func Read(scope *op.Scope, imagePath string, channels int64) *Image {
	_ = "STUB: not implemented"
	return nil
}

// NewImage creates an *Image from a 3 or 4D input tensor
// Place the created image within the specified scope
func NewImage(scope *op.Scope, tensor tf.Output) *Image { _ = "STUB: not implemented"; return nil }

// Copy the tensor to a new node in the graph

// Value returns the 3D tensor that represents a single image
// in the Tensorflow environment.
// If the image is a GIF the returned tensor is 4D
func (image *Image) Value() tf.Output { _ = "STUB: not implemented"; return *new(tf.Output) }

// Scale scales the image range value to be within [min, max]
func (image *Image) Scale(min, max float32) *Image { _ = "STUB: not implemented"; return nil }

// Normalize computes the mean and the stddev of the pixel values
// and normalizes every pixel subtracting the mean (centering) and dividing by
// the stddev (scale)
func (image *Image) Normalize() *Image { _ = "STUB: not implemented"; return nil }

// Avoid division by zero

// Center computes the mean value of the pixel values and subtracts this value
// from every pixel: this operation centers the data
func (image *Image) Center() *Image { _ = "STUB: not implemented"; return nil }

// SaturateCast casts the image to dtype handling overflow and underflow problems, saturate the exceeding values to
// to minimum/maximum accepted value of the dtype
func (image *Image) SaturateCast(dtype tf.DataType) *Image { _ = "STUB: not implemented"; return nil }

// ConvertDtype converts the Image dtype to dtype, uses SaturatesCast if required
func (image *Image) ConvertDtype(dtype tf.DataType, saturate bool) *Image {
	_ = "STUB: not implemented"
	return nil
}

// AdjustBrightness adds delta to the image
func (image *Image) AdjustBrightness(delta float32) *Image { _ = "STUB: not implemented"; return nil }

// AdjustContrast changes the contrast by contrastFactor
func (image *Image) AdjustContrast(contrastFactor float32) *Image {
	_ = "STUB: not implemented"
	return nil
}

// AdjustGamma performs gamma correction on the image
func (image *Image) AdjustGamma(gamma, gain float32) *Image { _ = "STUB: not implemented"; return nil }

// adjusted_img = (img / scale) ** gamma * scale * gain

// AdjustHue changes toe Hue by delta
func (image *Image) AdjustHue(delta float32) *Image { _ = "STUB: not implemented"; return nil }

// AdjustSaturation changes the saturation by saturationFactor
func (image *Image) AdjustSaturation(saturationFactor float32) *Image {
	_ = "STUB: not implemented"
	return nil
}

// CentralCrop extracts from the center of the image a portion of image with an area equals to the centralFraction
func (image *Image) CentralCrop(centralFraction float32) *Image {
	_ = "STUB: not implemented"
	return nil
}

// CropAndResize crops the image to the specified box and resizes the result to size
func (image *Image) CropAndResize(box Box, size Size, optional ...op.CropAndResizeAttr) *Image {
	_ = "STUB: not implemented"
	return nil
}

// DrawBoundingBoxes draws the specified boxes to the image
func (image *Image) DrawBoundingBoxes(boxes []Box) *Image { _ = "STUB: not implemented"; return nil }

// EncodeJPEG encodes the image in the JPEG format and returns an evaluable tensor
func (image *Image) EncodeJPEG(optional ...op.EncodeJpegAttr) tf.Output {
	_ = "STUB: not implemented"
	return *new(tf.Output)
}

// EncodePNG encodes the image in the PNG format and returns an evaluable tensor
func (image *Image) EncodePNG(optional ...op.EncodePngAttr) tf.Output {
	_ = "STUB: not implemented"
	return *new(tf.Output)
}

// ExtractGlimpse extracts a glimpse with the specified size at the specified offset
func (image *Image) ExtractGlimpse(size Size, offset Point, optional ...op.ExtractGlimpseAttr) tf.Output {
	_ = "STUB: not implemented"
	return *new(tf.Output)
}

// RGBToGrayscale converts the image from RGB to Grayscale
func (image *Image) RGBToGrayscale() *Image { _ = "STUB: not implemented"; return nil }

// HSVToRGB performs the colorspace transformation from HSV to RGB
func (image *Image) HSVToRGB() *Image { _ = "STUB: not implemented"; return nil }

// RGBToHSV performs the colorspace transformation from RGB to HSV
func (image *Image) RGBToHSV() *Image { _ = "STUB: not implemented"; return nil }

// ResizeArea resizes the image to the specified size using the Area interpolation
func (image *Image) ResizeArea(size Size, optional ...op.ResizeAreaAttr) *Image {
	_ = "STUB: not implemented"
	return nil
}

// ResizeBicubic resizes the image to the specified size using the Bicubic interpolation
func (image *Image) ResizeBicubic(size Size, optional ...op.ResizeBicubicAttr) *Image {
	_ = "STUB: not implemented"
	return nil
}

// ResizeBilinear resizes the image to the specified size using the Bilinear interpolation
func (image *Image) ResizeBilinear(size Size, optional ...op.ResizeBilinearAttr) *Image {
	_ = "STUB: not implemented"
	return nil
}

// ResizeNearestNeighbor resizes the image to the specified size using the NN interpolation
func (image *Image) ResizeNearestNeighbor(size Size, optional ...op.ResizeNearestNeighborAttr) *Image {
	_ = "STUB: not implemented"
	return nil
}

// Convolve executes the convolution operation between the current image and the passed filter
// The strides parameter rules the stride along each dimension
// Padding is a padding type to specify the type of padding
func (image *Image) Convolve(filter tf.Output, stride Stride, padding padding.Padding) *Image {
	_ = "STUB: not implemented"
	return nil
}

// filp the kernel in order to use the correlation operation (here called convolution)
// like a real convolution operation

// Correlate executes the correlation operation between the current image and the passed filter
// The strides parameter rules the stride along each dimension
// Padding is a padding type to specify the type of padding
func (image *Image) Correlate(filter tf.Output, stride Stride, padding padding.Padding) *Image {
	_ = "STUB: not implemented"
	return nil
}

// Dilate executes the dilatation operation between the current image and the padded filter
// The strides parameter rules the stride along each dimension, in output.
// The rate parameter rules the input stride for atrous morphological dilatation
// Padding is a padding type to specify the type of padding
func (image *Image) Dilate(filter tf.Output, stride, rate Stride, padding padding.Padding) *Image {
	_ = "STUB: not implemented"
	return nil
}

// If the filter is a convolutional filter [height, widht, depth, batch]
// we convert it to a dilatation filter [height, widht, depth]

// Erode executes the erosion operation between the current image and the padded filter
// The strides parameter rules the stride along each dimension
// The rate parameter rules the input stride for atrous morphological dilatation
// Padding is a padding type to specify the type of padding
func (image *Image) Erode(filter tf.Output, stride, rate Stride, padding padding.Padding) *Image {
	_ = "STUB: not implemented"
	return nil
}

// Negate the input

// Flip the kernel
