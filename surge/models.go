package surge

import (
    "gorm.io/gorm"
    "time"
)

type Bucket struct {

    BucketLength    uint64    `gorm:"-"`

	gorm.Model

	ExpDate     time.Time

	Counter     uint64    `gorm:"default:1"`

	DistrictID  uint

}

// set the expiration time on save of Bucket
func (b *Bucket) BeforeSave(tx *gorm.DB) (err error) {
    b.ExpDate = time.Now().Add(time.Second * time.Duration(b.BucketLength))
    return
}


type District struct {
	gorm.Model

	Buckets []Bucket

	Code    uint

}