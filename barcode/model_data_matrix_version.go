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

// DataMatrixVersion :
type DataMatrixVersion string

// List of DataMatrixVersion
const (
	DataMatrixVersionAuto             DataMatrixVersion = "Auto"
	DataMatrixVersionRowsColumns      DataMatrixVersion = "RowsColumns"
	DataMatrixVersionECC000_9x9       DataMatrixVersion = "ECC000_9x9"
	DataMatrixVersionECC000_050_11x11 DataMatrixVersion = "ECC000_050_11x11"
	DataMatrixVersionECC000_100_13x13 DataMatrixVersion = "ECC000_100_13x13"
	DataMatrixVersionECC000_100_15x15 DataMatrixVersion = "ECC000_100_15x15"
	DataMatrixVersionECC000_140_17x17 DataMatrixVersion = "ECC000_140_17x17"
	DataMatrixVersionECC000_140_19x19 DataMatrixVersion = "ECC000_140_19x19"
	DataMatrixVersionECC000_140_21x21 DataMatrixVersion = "ECC000_140_21x21"
	DataMatrixVersionECC000_140_23x23 DataMatrixVersion = "ECC000_140_23x23"
	DataMatrixVersionECC000_140_25x25 DataMatrixVersion = "ECC000_140_25x25"
	DataMatrixVersionECC000_140_27x27 DataMatrixVersion = "ECC000_140_27x27"
	DataMatrixVersionECC000_140_29x29 DataMatrixVersion = "ECC000_140_29x29"
	DataMatrixVersionECC000_140_31x31 DataMatrixVersion = "ECC000_140_31x31"
	DataMatrixVersionECC000_140_33x33 DataMatrixVersion = "ECC000_140_33x33"
	DataMatrixVersionECC000_140_35x35 DataMatrixVersion = "ECC000_140_35x35"
	DataMatrixVersionECC000_140_37x37 DataMatrixVersion = "ECC000_140_37x37"
	DataMatrixVersionECC000_140_39x39 DataMatrixVersion = "ECC000_140_39x39"
	DataMatrixVersionECC000_140_41x41 DataMatrixVersion = "ECC000_140_41x41"
	DataMatrixVersionECC000_140_43x43 DataMatrixVersion = "ECC000_140_43x43"
	DataMatrixVersionECC000_140_45x45 DataMatrixVersion = "ECC000_140_45x45"
	DataMatrixVersionECC000_140_47x47 DataMatrixVersion = "ECC000_140_47x47"
	DataMatrixVersionECC000_140_49x49 DataMatrixVersion = "ECC000_140_49x49"
	DataMatrixVersionECC200_10x10     DataMatrixVersion = "ECC200_10x10"
	DataMatrixVersionECC200_12x12     DataMatrixVersion = "ECC200_12x12"
	DataMatrixVersionECC200_14x14     DataMatrixVersion = "ECC200_14x14"
	DataMatrixVersionECC200_16x16     DataMatrixVersion = "ECC200_16x16"
	DataMatrixVersionECC200_18x18     DataMatrixVersion = "ECC200_18x18"
	DataMatrixVersionECC200_20x20     DataMatrixVersion = "ECC200_20x20"
	DataMatrixVersionECC200_22x22     DataMatrixVersion = "ECC200_22x22"
	DataMatrixVersionECC200_24x24     DataMatrixVersion = "ECC200_24x24"
	DataMatrixVersionECC200_26x26     DataMatrixVersion = "ECC200_26x26"
	DataMatrixVersionECC200_32x32     DataMatrixVersion = "ECC200_32x32"
	DataMatrixVersionECC200_36x36     DataMatrixVersion = "ECC200_36x36"
	DataMatrixVersionECC200_40x40     DataMatrixVersion = "ECC200_40x40"
	DataMatrixVersionECC200_44x44     DataMatrixVersion = "ECC200_44x44"
	DataMatrixVersionECC200_48x48     DataMatrixVersion = "ECC200_48x48"
	DataMatrixVersionECC200_52x52     DataMatrixVersion = "ECC200_52x52"
	DataMatrixVersionECC200_64x64     DataMatrixVersion = "ECC200_64x64"
	DataMatrixVersionECC200_72x72     DataMatrixVersion = "ECC200_72x72"
	DataMatrixVersionECC200_80x80     DataMatrixVersion = "ECC200_80x80"
	DataMatrixVersionECC200_88x88     DataMatrixVersion = "ECC200_88x88"
	DataMatrixVersionECC200_96x96     DataMatrixVersion = "ECC200_96x96"
	DataMatrixVersionECC200_104x104   DataMatrixVersion = "ECC200_104x104"
	DataMatrixVersionECC200_120x120   DataMatrixVersion = "ECC200_120x120"
	DataMatrixVersionECC200_132x132   DataMatrixVersion = "ECC200_132x132"
	DataMatrixVersionECC200_144x144   DataMatrixVersion = "ECC200_144x144"
	DataMatrixVersionECC200_8x18      DataMatrixVersion = "ECC200_8x18"
	DataMatrixVersionECC200_8x32      DataMatrixVersion = "ECC200_8x32"
	DataMatrixVersionECC200_12x26     DataMatrixVersion = "ECC200_12x26"
	DataMatrixVersionECC200_12x36     DataMatrixVersion = "ECC200_12x36"
	DataMatrixVersionECC200_16x36     DataMatrixVersion = "ECC200_16x36"
	DataMatrixVersionECC200_16x48     DataMatrixVersion = "ECC200_16x48"
	DataMatrixVersionDMRE_8x48        DataMatrixVersion = "DMRE_8x48"
	DataMatrixVersionDMRE_8x64        DataMatrixVersion = "DMRE_8x64"
	DataMatrixVersionDMRE_8x80        DataMatrixVersion = "DMRE_8x80"
	DataMatrixVersionDMRE_8x96        DataMatrixVersion = "DMRE_8x96"
	DataMatrixVersionDMRE_8x120       DataMatrixVersion = "DMRE_8x120"
	DataMatrixVersionDMRE_8x144       DataMatrixVersion = "DMRE_8x144"
	DataMatrixVersionDMRE_12x64       DataMatrixVersion = "DMRE_12x64"
	DataMatrixVersionDMRE_12x88       DataMatrixVersion = "DMRE_12x88"
	DataMatrixVersionDMRE_16x64       DataMatrixVersion = "DMRE_16x64"
	DataMatrixVersionDMRE_20x36       DataMatrixVersion = "DMRE_20x36"
	DataMatrixVersionDMRE_20x44       DataMatrixVersion = "DMRE_20x44"
	DataMatrixVersionDMRE_20x64       DataMatrixVersion = "DMRE_20x64"
	DataMatrixVersionDMRE_22x48       DataMatrixVersion = "DMRE_22x48"
	DataMatrixVersionDMRE_24x48       DataMatrixVersion = "DMRE_24x48"
	DataMatrixVersionDMRE_24x64       DataMatrixVersion = "DMRE_24x64"
	DataMatrixVersionDMRE_26x40       DataMatrixVersion = "DMRE_26x40"
	DataMatrixVersionDMRE_26x48       DataMatrixVersion = "DMRE_26x48"
	DataMatrixVersionDMRE_26x64       DataMatrixVersion = "DMRE_26x64"
)
