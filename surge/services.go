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

func (s *SurgeService) IncrementLastActiveBucket(bucket *Bucket, district uint8) (err error) {

    // 1.get last bucket that hasn't expired yet and increment it's counter by one
    // 2.if no bucket found create a new one

    err = s.GetLastActiveBucketByDistrict(bucket, district)

    if err != nil {
        err = s.CreateBucket(bucket, district)
        return err
    }

    err = s.IncrementBucketCounter(bucket)

    return err
}

func (s *SurgeService) IncrementBucketCounter(bucket *Bucket) (err error){
    bucket.Counter += 1
    err = s.DB.Save(bucket).Error
    return err
}

func (s *SurgeService) GetLastActiveBucketByDistrict(bucket *Bucket, district uint8) (err error){

    err = s.DB.First(&bucket, "exp_date > ? AND district_id = ?", time.Now(), district).Error
    return err
}

func (s *SurgeService) CreateBucket(bucket *Bucket, district uint8) (err error){

    *bucket = Bucket{DistrictID: district, BucketLength: 600}

    err = s.DB.Create(bucket).Error

    return err
}

func (s *SurgeService) SumAllBucketsAfterGivenTime(district uint8, windowLength uint64) {

}

func (s *SurgeService) CalculateCoefficient(requests uint64) {

}

func (s *SurgeService) GetDistrictIDFromLocation(lat float32, lon float32) uint8{
    return s.OsmService.GetDistrictIDFromLocation(lat, lon)
}

func NewSurgeService(db *gorm.DB, osmService *osm.OpenStreetMapService) *SurgeService {

    return &SurgeService{DB: db, OsmService: osmService}
}