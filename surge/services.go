package surge

import (
    "github.com/mhmd-bb/snapp-surge/osm"
    "gorm.io/gorm"
    "time"
)

type SurgeService struct {
    DB      *gorm.DB
    OsmService      *osm.OpenStreetMapService
}

func (s *SurgeService) IncrementLastActiveBucket(district uint8) {

    // 1.get last bucket that hasn't expired yet and increment it's counter by one
    // 2.if no bucket found create a new one
}

func (s *SurgeService) CreateBucket(district uint8) {

}

func (s *SurgeService) SumAllBucketsAfterGivenTime(district uint8, after time.Time) {

}

func (s *SurgeService) CalculateCoefficient(requests uint64) {

}

func (s *SurgeService) GetDistrictIDFromLocation(lat float32, lon float32) uint8{
    return s.OsmService.GetDistrictIDFromLocation(lat, lon)
}

func NewSurgeService(db *gorm.DB, osmService *osm.OpenStreetMapService) *SurgeService {

    return &SurgeService{DB: db, OsmService: osmService}
}