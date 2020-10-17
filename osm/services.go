package osm

import (
    "gorm.io/gorm"
    "strconv"
)

type OpenStreetMapService struct {
    DB      *gorm.DB
    QB      *OpenStreetMapQueryBuilder
}


func (osmService *OpenStreetMapService) GetDistrictIDFromLocation(lat float32, lon float32) (district uint8, err error){

    var districtID string

    err = osmService.DB.Raw(osmService.QB.GetDistrictIDFromLocation(lat, lon)).Scan(&districtID).Error

    u, _ := strconv.ParseUint(districtID,10, 64)

    district = uint8(u)

    return  district, err

}

func NewOpenStreetMapService(db *gorm.DB) *OpenStreetMapService{

    return &OpenStreetMapService{DB: db, QB: &OpenStreetMapQueryBuilder{}}

}