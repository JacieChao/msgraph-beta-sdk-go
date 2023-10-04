package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse struct {
    ItemEventsItemInstancesItemExceptionOccurrencesDeltaGetResponse
}
// NewItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse instantiates a new ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse()(*ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse) {
    m := &ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse{
        ItemEventsItemInstancesItemExceptionOccurrencesDeltaGetResponse: *NewItemEventsItemInstancesItemExceptionOccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemInstancesItemExceptionOccurrencesDeltaResponse(), nil
}
// ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemInstancesItemExceptionOccurrencesDeltaResponseable interface {
    ItemEventsItemInstancesItemExceptionOccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
