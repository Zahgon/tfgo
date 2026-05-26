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
)

func boxes2batch(scope *op.Scope, boxes []Box) tf.Output {
	_ = "STUB: not implemented"
	return *new(tf.Output)
}

func sizes2batch(scope *op.Scope, sizes []Size) tf.Output {
	_ = "STUB: not implemented"
	return *new(tf.Output)
}

func points2batch(scope *op.Scope, points []Point) tf.Output {
	_ = "STUB: not implemented"
	return *new(tf.Output)
}
