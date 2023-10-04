package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse struct {
    ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponse
}
// NewItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse instantiates a new ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse()(*ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse) {
    m := &ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse{
        ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponse: *NewItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponse(), nil
}
// ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaResponseable interface {
    ItemCalendarsItemEventsItemExceptionOccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
