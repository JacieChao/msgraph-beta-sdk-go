package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse struct {
    ItemCalendarCalendarViewItemExceptionOccurrencesDeltaGetResponse
}
// NewItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse instantiates a new ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse and sets the default values.
func NewItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse()(*ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse) {
    m := &ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse{
        ItemCalendarCalendarViewItemExceptionOccurrencesDeltaGetResponse: *NewItemCalendarCalendarViewItemExceptionOccurrencesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponse(), nil
}
// ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarCalendarViewItemExceptionOccurrencesDeltaResponseable interface {
    ItemCalendarCalendarViewItemExceptionOccurrencesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
