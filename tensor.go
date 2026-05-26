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
	"github.com/galeone/tensorflow/tensorflow/go/op"
)

// Tensor is an high level abstraction for the
// tf.Output structure, associating a scope to the Tensor
type Tensor struct {
	// Root: Each tensor maintains a pointer to the graph root
	Root *op.Scope
	// Path is the current Tensor full path
	Path *op.Scope
	// Output is the Tensor content
	Output tf.Output
}

// NewTensor creates a *Tensor from a tf.Output
// Place the cloned tensor within the specified scope
func NewTensor(scope *op.Scope, tfout tf.Output) (tensor *Tensor) {
	_ = "STUB: not implemented"
	return nil
}

// Copy the tensor to a new node in the graph

// Check checks if the previous operation caused an error
// and thus tensor.Path.Err is not nil.
// If it's not, panics because we're defining the graph in a wrong way
func (tensor *Tensor) Check() { _ = "STUB: not implemented"; return }

// Scope returns the scope associated to the tensor
func (tensor *Tensor) Scope() *op.Scope {
	_ = "STUB: not implemented"

	// Shape64 returns the shape of the tensor as []int64.
	// If firstDimension is true a 4 elements slice is returned.
	// Otherwise a 3 elements slice is returned.
	return nil
}

func (tensor *Tensor) Shape64(firstDimension bool) []int64 { _ = "STUB: not implemented"; return nil }

// Shape32 returns the shape of the tensor as []int32.
// If firstDimension is true a 4 elements slice is returned.
// Otherwise a 3 elements slice is returned.
func (tensor *Tensor) Shape32(firstDimension bool) []int32 { _ = "STUB: not implemented"; return nil }

// Dtype returns the tensor dtype
func (tensor *Tensor) Dtype() tf.DataType { _ = "STUB: not implemented"; return *new(tf.DataType) }

// --------
// Methods returning *Tensor
// Structs that embed *Tensor should ""override"" them
// Returning their own type
// --------

// Clone returns a copy of the current tensor in a new scope
// Clone is used to create a different tensor
// from the output of an operation.
// The new node is placed at the same level of the current tensor
// it can be seen as a twin tensor
func (tensor *Tensor) Clone() *Tensor { _ = "STUB: not implemented"; return nil }

// Cast casts the current tensor to the requested dtype
func (tensor *Tensor) Cast(dtype tf.DataType) *Tensor { _ = "STUB: not implemented"; return nil }

// Add defines the add operation between the tensor and tfout
// `tfout` dtype is converted to tensor.Dtype() before adding
func (tensor *Tensor) Add(tfout tf.Output) *Tensor { _ = "STUB: not implemented"; return nil }

// Mul defines the multiplication operation between the tensor
// and `tfout`. It's the multiplication element-wise with broadcasting support.
// `tfout` dtype is converted to tensor.Dtype() before multiplying
func (tensor *Tensor) Mul(tfout tf.Output) *Tensor { _ = "STUB: not implemented"; return nil }

// MatMul defines the matrix multiplication operation between the tensor
// and `tfout`.
// `tfout` dtype is converted to tensor.Dtype() before multiplying
func (tensor *Tensor) MatMul(tfout tf.Output) *Tensor { _ = "STUB: not implemented"; return nil }

// Pow defines the pow operation x^y, where x are the tensor values
// y dtype is converted to tensor.Dtype() before executing Pow
func (tensor *Tensor) Pow(y tf.Output) *Tensor { _ = "STUB: not implemented"; return nil }

// Square defines the square operation for the tensor values
func (tensor *Tensor) Square() *Tensor { _ = "STUB: not implemented"; return nil }

// Sqrt defines the square root operation for the tensor values
func (tensor *Tensor) Sqrt() *Tensor { _ = "STUB: not implemented"; return nil }
