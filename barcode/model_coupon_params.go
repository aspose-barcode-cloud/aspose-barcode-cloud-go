package barcode

// CouponParams - Coupon parameters. Used for UpcaGs1DatabarCoupon, UpcaGs1Code128Coupon.
type CouponParams struct {
	// Space between main the BarCode and supplement BarCode in Unit value.
	SupplementSpace float64 `json:"SupplementSpace,omitempty"`
}
