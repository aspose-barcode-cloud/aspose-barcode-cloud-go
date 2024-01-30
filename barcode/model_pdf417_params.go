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

import (
	"time"
)

// Pdf417Params - PDF417 parameters.
type Pdf417Params struct {
	// Height/Width ratio of 2D BarCode module.
	AspectRatio float64 `json:"AspectRatio,omitempty"`
	// DEPRECATED: This property is obsolete and will be removed in future releases. Unicode symbols detection and encoding will be processed in Auto mode with Extended Channel Interpretation charset designator. Using of own encodings requires manual CodeText encoding into byte[] array.  Sets the encoding of codetext.
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
	// Extended Channel Interpretation Identifiers. It is used to tell the barcode reader details about the used references for encoding the data in the symbol. Current implementation consists all well known charset encodings.
	Pdf417ECIEncoding EciEncodings `json:"Pdf417ECIEncoding,omitempty"`
	// Used to instruct the reader to interpret the data contained within the symbol as programming for reader initialization
	IsReaderInitialization bool `json:"IsReaderInitialization,omitempty"`
	// Macro Pdf417 barcode time stamp
	MacroTimeStamp time.Time `json:"MacroTimeStamp,omitempty"`
	// Macro Pdf417 barcode sender name
	MacroSender string `json:"MacroSender,omitempty"`
	// Macro Pdf417 file size. The file size field contains the size in bytes of the entire source file
	MacroFileSize int32 `json:"MacroFileSize,omitempty"`
	// Macro Pdf417 barcode checksum. The checksum field contains the value of the 16-bit (2 bytes) CRC checksum using the CCITT-16 polynomial
	MacroChecksum int32 `json:"MacroChecksum,omitempty"`
	// Macro Pdf417 barcode file name
	MacroFileName string `json:"MacroFileName,omitempty"`
	// Macro Pdf417 barcode addressee name
	MacroAddressee string `json:"MacroAddressee,omitempty"`
	// Extended Channel Interpretation Identifiers. Applies for Macro PDF417 text fields.
	MacroECIEncoding EciEncodings `json:"MacroECIEncoding,omitempty"`
	// DEPRECATED: This property is obsolete and will be removed in future releases. See samples of using new parameters on https://releases.aspose.com/barcode/net/release-notes/2023/aspose-barcode-for-net-23-10-release-notes/ Function codeword for Code 128 emulation. Applied for MicroPDF417 only. Ignored for PDF417 and MacroPDF417 barcodes.
	Code128Emulation Code128Emulation `json:"Code128Emulation,omitempty"`
	// Can be used only with MicroPdf417 and encodes Code 128 emulation modes. Can encode FNC1 in second position modes 908 and 909, also can encode 910 and 911 which just indicate that recognized MicroPdf417 can be interpret as Code 128.
	IsCode128Emulation bool `json:"IsCode128Emulation,omitempty"`
	// Used to tell the encoder whether to add Macro PDF417 Terminator (codeword 922) to the segment. Applied only for Macro PDF417.
	Pdf417MacroTerminator Pdf417MacroTerminator `json:"Pdf417MacroTerminator,omitempty"`
	// Defines linked modes with GS1MicroPdf417, MicroPdf417 and Pdf417 barcodes. With GS1MicroPdf417 symbology encodes 906, 907, 912, 913, 914, 915 “Linked” UCC/EAN-128 modes. With MicroPdf417 and Pdf417 symbologies encodes 918 linkage flag to associated linear component other than an EAN.UCC.
	IsLinked bool `json:"IsLinked,omitempty"`
	// Macro Characters 05 and 06 values are used to obtain more compact encoding in special modes. Can be used only with MicroPdf417 and encodes 916 and 917 MicroPdf417 modes. Default value: MacroCharacters.None.
	MacroCharacters MacroCharacter `json:"MacroCharacters,omitempty"`
}
