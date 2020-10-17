package surge

import (
    "gorm.io/gorm"
    "time"
)

type Bucket struct {
    BucketLength uint64 `gorm:"-"`

    ID uint64 `gorm:"primaryKey"`

    CreatedAt time.Time `gorm:"autoCreateTime;index"`

    ExpDate time.Time `gorm:"index:idx_member"`

    Counter uint64 `gorm:"default:1"`

    DistrictID uint8 `gorm:"index:idx_member"`
}

// set the expiration time on create of Bucket
func (b *Bucket) BeforeCreate(tx *gorm.DB) (err error) {
    b.ExpDate = time.Now().Add(time.Second * time.Duration(b.BucketLength))
    return
}

type Rule struct {
    ID          uint64  `gorm:"primaryKey" json:"id"`
    Threshold   uint64  `gorm:"unique" json:"threshold"`
    Coefficient float32 `json:"coefficient"`
}
