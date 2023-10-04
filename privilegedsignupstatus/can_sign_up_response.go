package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CanSignUpResponse 
// Deprecated: This class is obsolete. Use canSignUpGetResponse instead.
type CanSignUpResponse struct {
    CanSignUpGetResponse
}
// NewCanSignUpResponse instantiates a new CanSignUpResponse and sets the default values.
func NewCanSignUpResponse()(*CanSignUpResponse) {
    m := &CanSignUpResponse{
        CanSignUpGetResponse: *NewCanSignUpGetResponse(),
    }
    return m
}
// CreateCanSignUpResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCanSignUpResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCanSignUpResponse(), nil
}
// CanSignUpResponseable 
// Deprecated: This class is obsolete. Use canSignUpGetResponse instead.
type CanSignUpResponseable interface {
    CanSignUpGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
