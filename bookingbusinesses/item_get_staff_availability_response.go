package bookingbusinesses

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemGetStaffAvailabilityResponse 
// Deprecated: This class is obsolete. Use getStaffAvailabilityPostResponse instead.
type ItemGetStaffAvailabilityResponse struct {
    ItemGetStaffAvailabilityPostResponse
}
// NewItemGetStaffAvailabilityResponse instantiates a new ItemGetStaffAvailabilityResponse and sets the default values.
func NewItemGetStaffAvailabilityResponse()(*ItemGetStaffAvailabilityResponse) {
    m := &ItemGetStaffAvailabilityResponse{
        ItemGetStaffAvailabilityPostResponse: *NewItemGetStaffAvailabilityPostResponse(),
    }
    return m
}
// CreateItemGetStaffAvailabilityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemGetStaffAvailabilityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemGetStaffAvailabilityResponse(), nil
}
// ItemGetStaffAvailabilityResponseable 
// Deprecated: This class is obsolete. Use getStaffAvailabilityPostResponse instead.
type ItemGetStaffAvailabilityResponseable interface {
    ItemGetStaffAvailabilityPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
