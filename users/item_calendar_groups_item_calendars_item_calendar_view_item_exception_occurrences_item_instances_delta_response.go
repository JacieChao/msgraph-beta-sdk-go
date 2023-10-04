package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse struct {
    ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponse
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse{
        ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponse: *NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponse(), nil
}
// ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaResponseable interface {
    ItemCalendarGroupsItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
