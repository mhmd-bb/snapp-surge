package surge


type LatLonDto struct {
    Lat float32  `json:"lat" binding:"required,lte=90,gte=-90"`
    Lon float32  `json:"lon" binding:"required,lte=180,gte=-180"`
}

type RuleDto struct {
    Threshold uint64   `json:"threshold" binding:"required"`
    Coefficient float32   `json:"coefficient" binding:"required"`
}

type DeleteRuleDto struct {
    Id  uint64  `json:"id"  binding:"required"`
}