package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type UsersDeltaResponse struct {
    UsersDeltaGetResponse
}
// NewUsersDeltaResponse instantiates a new UsersDeltaResponse and sets the default values.
func NewUsersDeltaResponse()(*UsersDeltaResponse) {
    m := &UsersDeltaResponse{
        UsersDeltaGetResponse: *NewUsersDeltaGetResponse(),
    }
    return m
}
// CreateUsersDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersDeltaResponse(), nil
}
// UsersDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type UsersDeltaResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UsersDeltaGetResponseable
}
