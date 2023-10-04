package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInvalidateAllRefreshTokensResponse 
// Deprecated: This class is obsolete. Use invalidateAllRefreshTokensPostResponse instead.
type ItemInvalidateAllRefreshTokensResponse struct {
    ItemInvalidateAllRefreshTokensPostResponse
}
// NewItemInvalidateAllRefreshTokensResponse instantiates a new ItemInvalidateAllRefreshTokensResponse and sets the default values.
func NewItemInvalidateAllRefreshTokensResponse()(*ItemInvalidateAllRefreshTokensResponse) {
    m := &ItemInvalidateAllRefreshTokensResponse{
        ItemInvalidateAllRefreshTokensPostResponse: *NewItemInvalidateAllRefreshTokensPostResponse(),
    }
    return m
}
// CreateItemInvalidateAllRefreshTokensResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInvalidateAllRefreshTokensResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInvalidateAllRefreshTokensResponse(), nil
}
// ItemInvalidateAllRefreshTokensResponseable 
// Deprecated: This class is obsolete. Use invalidateAllRefreshTokensPostResponse instead.
type ItemInvalidateAllRefreshTokensResponseable interface {
    ItemInvalidateAllRefreshTokensPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
