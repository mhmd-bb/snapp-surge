package osm

import (
    "gorm.io/gorm"
    "strconv"
)

type OpenStreetMapService struct {
    DB      *gorm.DB
    QB      *OpenStreetMapQueryBuilder
}


func (osmService *OpenStreetMapService) GetDistrictIDFromLocation(lat float32, lon float32) uint8{

    var districtID string

    osmService.DB.Raw(osmService.QB.GetDistrictIDFromLocation(lat, lon)).Scan(&districtID)

    u, _ := strconv.ParseUint(districtID,10, 64)

    return  uint8(u)

}

func NewOpenStreetMapService(db *gorm.DB) *OpenStreetMapService{

    return &OpenStreetMapService{DB: db, QB: &OpenStreetMapQueryBuilder{}}

}