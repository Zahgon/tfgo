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

package tfgo

import (
	tf "github.com/galeone/tensorflow/tensorflow/go"
)

// IsInteger returns true if dtype is a tensorflow integer type
func IsInteger(dtype tf.DataType) bool { _ = "STUB: not implemented"; return false }

// IsFloat returns true if dtype is a tensorfow float type
func IsFloat(dtype tf.DataType) bool { _ = "STUB: not implemented"; return false }

// MaxValue returns the maximum value accepted for the dtype
func MaxValue(dtype tf.DataType) float64 { _ = "STUB: not implemented"; return 0 }

// No support for Quantized types

// MinValue returns the minimum representable value for the specified dtype
func MinValue(dtype tf.DataType) float64 { _ = "STUB: not implemented"; return 0 }

// According to: https://www.khronos.org/opengl/wiki/Small_Float_Formats

// No support for Quantized types
