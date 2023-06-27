/*
 * MIT License

 * Copyright (c) 2023 Aspose Pty Ltd

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

// HanXinParams - HanXin params.
type HanXinParams struct {
	// Encoding mode for XanXin barcodes. Default value: HanXinEncodeMode.Auto.
	EncodeMode HanXinEncodeMode `json:"EncodeMode,omitempty"`
	// Allowed Han Xin error correction levels from L1 to L4. Default value: HanXinErrorLevel.L1.
	ErrorLevel HanXinErrorLevel `json:"ErrorLevel,omitempty"`
	// Allowed Han Xin versions, Auto and Version01 - Version84. Default value: HanXinVersion.Auto.
	Version HanXinVersion `json:"Version,omitempty"`
	// Extended Channel Interpretation Identifiers. It is used to tell the barcode reader details about the used references for encoding the data in the symbol. Current implementation consists all well known charset encodings. Default value: ECIEncodings.ISO_8859_1
	ECIEncoding EciEncodings `json:"ECIEncoding,omitempty"`
}
