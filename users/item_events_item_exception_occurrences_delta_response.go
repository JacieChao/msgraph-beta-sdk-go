package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEventsItemExceptionOccurrencesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemExceptionOccurrencesDeltaResponse struct {
    ItemEventsItemExceptionOccurrencesDeltaGetResponse
}
// NewItemEventsItemExceptionOccurrencesDeltaResponse instantiates a new ItemEventsItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewItemEventsItemExceptionOccurrencesDeltaResponse()(*ItemEventsItemExceptionOccurrencesDeltaResponse) {
    m := &ItemEventsItemExceptionOccurrencesDeltaResponse{
        ItemEventsItemExceptionOccurrencesDeltaGetResponse: *NewItemEventsItemExceptionOccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemEventsItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEventsItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEventsItemExceptionOccurrencesDeltaResponse(), nil
}
// ItemEventsItemExceptionOccurrencesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemEventsItemExceptionOccurrencesDeltaResponseable interface {
    ItemEventsItemExceptionOccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
