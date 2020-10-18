package surge

type LatLonDto struct {
	Lat float32 `json:"lat" binding:"required,lte=90,gte=-90" example:"51.13199462890625"`
	Lon float32 `json:"lon" binding:"required,lte=180,gte=-180" example:"35.73425097869431"`
}

type RuleDto struct {
	Threshold   uint64  `json:"threshold" binding:"required" example:"12"`
	Coefficient float32 `json:"coefficient" binding:"required" example:"1.12"`
}
