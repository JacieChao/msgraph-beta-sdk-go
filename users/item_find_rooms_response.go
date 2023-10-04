package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemFindRoomsResponse 
// Deprecated: This class is obsolete. Use findRoomsGetResponse instead.
type ItemFindRoomsResponse struct {
    ItemFindRoomsGetResponse
}
// NewItemFindRoomsResponse instantiates a new ItemFindRoomsResponse and sets the default values.
func NewItemFindRoomsResponse()(*ItemFindRoomsResponse) {
    m := &ItemFindRoomsResponse{
        ItemFindRoomsGetResponse: *NewItemFindRoomsGetResponse(),
    }
    return m
}
// CreateItemFindRoomsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemFindRoomsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFindRoomsResponse(), nil
}
// ItemFindRoomsResponseable 
// Deprecated: This class is obsolete. Use findRoomsGetResponse instead.
type ItemFindRoomsResponseable interface {
    ItemFindRoomsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
