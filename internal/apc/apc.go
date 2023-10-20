package apc

type AP struct {
	N1 float64 `json:"n1"` // first elem
	D  float64 `json:"d"`  // delta
}

func (ap *AP) Count() {
	ap.N1 = ap.N1 + ap.D
}
