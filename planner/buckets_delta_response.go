package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BucketsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type BucketsDeltaResponse struct {
    BucketsDeltaGetResponse
}
// NewBucketsDeltaResponse instantiates a new BucketsDeltaResponse and sets the default values.
func NewBucketsDeltaResponse()(*BucketsDeltaResponse) {
    m := &BucketsDeltaResponse{
        BucketsDeltaGetResponse: *NewBucketsDeltaGetResponse(),
    }
    return m
}
// CreateBucketsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBucketsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBucketsDeltaResponse(), nil
}
// BucketsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type BucketsDeltaResponseable interface {
    BucketsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
