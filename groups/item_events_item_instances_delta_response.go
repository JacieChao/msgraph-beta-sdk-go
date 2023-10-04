package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEventsItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemInstancesDeltaResponse struct {
    ItemEventsItemInstancesDeltaGetResponse
}
// NewItemEventsItemInstancesDeltaResponse instantiates a new ItemEventsItemInstancesDeltaResponse and sets the default values.
func NewItemEventsItemInstancesDeltaResponse()(*ItemEventsItemInstancesDeltaResponse) {
    m := &ItemEventsItemInstancesDeltaResponse{
        ItemEventsItemInstancesDeltaGetResponse: *NewItemEventsItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemEventsItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEventsItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemInstancesDeltaResponse(), nil
}
// ItemEventsItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemInstancesDeltaResponseable interface {
    ItemEventsItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
