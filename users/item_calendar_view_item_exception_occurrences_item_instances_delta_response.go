package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse struct {
    ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponse
}
// NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse instantiates a new ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse()(*ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse) {
    m := &ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse{
        ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponse: *NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse(), nil
}
// ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable interface {
    ItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
