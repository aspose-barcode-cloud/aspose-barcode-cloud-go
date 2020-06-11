/*
 * MIT License

 * Copyright (c) 2020 Aspose Pty Ltd

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

// PDF417 parameters.
type Pdf417Params struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Encoding of codetext.
	TextEncoding string `json:"TextEncoding,omitempty"`
	// Columns count.
	Columns int32 `json:"Columns,omitempty"`
	// Pdf417 symbology type of BarCode's compaction mode. Default value: Pdf417CompactionMode.Auto.
	CompactionMode Pdf417CompactionMode `json:"CompactionMode,omitempty"`
	// Pdf417 symbology type of BarCode's error correction level ranging from level0 to level8, level0 means no error correction info, level8 means best error correction which means a larger picture.
	ErrorLevel Pdf417ErrorLevel `json:"ErrorLevel,omitempty"`
	// Macro Pdf417 barcode's file ID. Used for MacroPdf417.
	MacroFileID int32 `json:"MacroFileID,omitempty"`
	// Macro Pdf417 barcode's segment ID, which starts from 0, to MacroSegmentsCount - 1.
	MacroSegmentID int32 `json:"MacroSegmentID,omitempty"`
	// Macro Pdf417 barcode segments count.
	MacroSegmentsCount int32 `json:"MacroSegmentsCount,omitempty"`
	// Rows count.
	Rows int32 `json:"Rows,omitempty"`
	// Whether Pdf417 symbology type of BarCode is truncated (to reduce space).
	Truncate bool `json:"Truncate,omitempty"`
}
