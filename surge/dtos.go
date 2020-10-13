package surge


type LatLonDto struct {
    Lat float32  `json:"lat" binding:"required,lte=90,gte=-90"`
    Lon float32  `json:"lon" binding:"required,lte=180,gte=-180"`
}