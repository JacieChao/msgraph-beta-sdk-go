package planner

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BucketsItemTasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type BucketsItemTasksDeltaResponse struct {
    BucketsItemTasksDeltaGetResponse
}
// NewBucketsItemTasksDeltaResponse instantiates a new BucketsItemTasksDeltaResponse and sets the default values.
func NewBucketsItemTasksDeltaResponse()(*BucketsItemTasksDeltaResponse) {
    m := &BucketsItemTasksDeltaResponse{
        BucketsItemTasksDeltaGetResponse: *NewBucketsItemTasksDeltaGetResponse(),
    }
    return m
}
// CreateBucketsItemTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBucketsItemTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBucketsItemTasksDeltaResponse(), nil
}
// BucketsItemTasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type BucketsItemTasksDeltaResponseable interface {
    BucketsItemTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
