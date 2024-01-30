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

// AztecParams - Aztec parameters.
type AztecParams struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// Level of error correction of Aztec types of barcode. Value should between 10 to 95.
	ErrorLevel int32 `json:"ErrorLevel,omitempty"`
	// Aztec Symbol mode. Default value: AztecSymbolMode.Auto.
	SymbolMode AztecSymbolMode `json:"SymbolMode,omitempty"`
	// DEPRECATED: This property is obsolete and will be removed in future releases. Unicode symbols detection and encoding will be processed in Auto mode with Extended Channel Interpretation charset designator. Using of own encodings requires manual CodeText encoding into byte[] array.  Sets the encoding of codetext.
	TextEncoding string `json:"TextEncoding,omitempty"`
	// Encoding mode for Aztec barcodes. Default value: Auto
	EncodeMode AztecEncodeMode `json:"EncodeMode,omitempty"`
	// Identifies ECI encoding. Used when AztecEncodeMode is Auto. Default value: ISO-8859-1.
	ECIEncoding EciEncodings `json:"ECIEncoding,omitempty"`
	// Used to instruct the reader to interpret the data contained within the symbol as programming for reader initialization.
	IsReaderInitialization bool `json:"IsReaderInitialization,omitempty"`
	// Gets or sets layers count of Aztec symbol. Layers count should be in range from 1 to 3 for Compact mode and in range from 1 to 32 for Full Range mode. Default value: 0 (auto).
	LayersCount int32 `json:"LayersCount,omitempty"`
}
