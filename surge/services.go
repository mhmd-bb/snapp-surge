package surge

import (
    "errors"
    "github.com/mhmd-bb/snapp-surge/osm"
    "gorm.io/gorm"
    "time"
)

type SurgeService struct {
    DB      *gorm.DB

    // OsmService is used in surge service to use functionalities it provides like getting districtID
    // from latitude and longitude
    OsmService      *osm.OpenStreetMapService

    // Bucket Length is a constant that you can change in .env file
    // Default value is 600 seconds or 10 minutes
    bucketLength    uint64

    // Window Length is configurable as well default value is 2 hours
    // It is a moving window on buckets
    windowLength    uint64

}

// 1.Get last bucket that hasn't expired yet and increment it's counter by one
// 2.If no active bucket is found then create a new one
func (s *SurgeService) IncrementLastActiveBucket(bucket *Bucket, district uint8) (err error) {

    err = s.GetLastActiveBucketByDistrict(bucket, district)

    // If no active bucket found then create a new bucket
    if errors.Is(err, gorm.ErrRecordNotFound) {
        err = s.CreateBucket(bucket, district)
        return err
    }

    // If there is an active bucket in database already, increment its counter by one
    err = s.IncrementBucketCounter(bucket)

    return err
}

// Get bucket by reference and increment its counter by one
func (s *SurgeService) IncrementBucketCounter(bucket *Bucket) (err error){
    err = s.DB.Model(bucket).Update("counter", bucket.Counter + 1).Error
    return err
}

// Get District's last active bucket (has not expired) returns error when there is no active bucket
func (s *SurgeService) GetLastActiveBucketByDistrict(bucket *Bucket, district uint8) (err error){

    err = s.DB.First(&bucket, "exp_date > ? AND district_id = ?", time.Now(), district).Error
    return err
}

// Create a new bucket
func (s *SurgeService) CreateBucket(bucket *Bucket, district uint8) (err error){

    *bucket = Bucket{DistrictID: district, BucketLength: s.bucketLength}

    err = s.DB.Create(bucket).Error

    return err
}

// Sum All buckets in moving window
func (s *SurgeService) SumAllBucketsInCurrentWindow(n *uint64, district uint8) (err error){
    err = s.DB.Table("buckets").Where("district_id = ? AND created_at > ?", district, time.Now().Add(-time.Second * time.Duration(s.windowLength))).Select("sum(counter) as n").Scan(n).Error
    return err
}

// Get Coefficient by total count of requests in moving window
func (s *SurgeService) CalculateCoefficient(coefficient *float32, counter uint64) (err error){
    var rule Rule

    err = s.DB.Select("coefficient").Where("threshold <= ?", counter).Last(&rule).Error

    *coefficient = rule.Coefficient

    // if we haven't reached to first threshold or we have no threshold return 1
    if errors.Is(err, gorm.ErrRecordNotFound) {
        *coefficient = 1
        return nil
    }

    return err

}

// Uses Osm service's GetDistrictIDFromLocation
func (s *SurgeService) GetDistrictIDFromLocation(districtID *uint8, lat float32, lon float32) (err error){
    *districtID, err = s.OsmService.GetDistrictIDFromLocation(lat, lon)
    return err
}

func NewSurgeService(db *gorm.DB, osmService *osm.OpenStreetMapService, bucketLength uint64, windowLength uint64) *SurgeService {

    return &SurgeService{DB: db, OsmService: osmService, bucketLength: bucketLength, windowLength: windowLength}
}