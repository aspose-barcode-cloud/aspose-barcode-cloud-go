/*
 * MIT License

 * Copyright (c) 2024 Aspose Pty Ltd

 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:

 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.

 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package barcode

// DataBarParams - Databar parameters.
type DataBarParams struct {
	// Height/Width ratio of 2D BarCode module. Used for DataBar stacked.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Columns count.
	Columns int32 `json:"Columns,omitempty"`
	// Rows count.
	Rows int32 `json:"Rows,omitempty"`
	// Enables flag of 2D composite component with DataBar barcode
	Is2DCompositeComponent bool `json:"Is2DCompositeComponent,omitempty"`
	// If this flag is set, it allows only GS1 encoding standard for Databar barcode types
	IsAllowOnlyGS1Encoding bool `json:"IsAllowOnlyGS1Encoding,omitempty"`
}
