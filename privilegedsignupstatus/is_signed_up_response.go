package privilegedsignupstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IsSignedUpResponse 
// Deprecated: This class is obsolete. Use isSignedUpGetResponse instead.
type IsSignedUpResponse struct {
    IsSignedUpGetResponse
}
// NewIsSignedUpResponse instantiates a new IsSignedUpResponse and sets the default values.
func NewIsSignedUpResponse()(*IsSignedUpResponse) {
    m := &IsSignedUpResponse{
        IsSignedUpGetResponse: *NewIsSignedUpGetResponse(),
    }
    return m
}
// CreateIsSignedUpResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIsSignedUpResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIsSignedUpResponse(), nil
}
// IsSignedUpResponseable 
// Deprecated: This class is obsolete. Use isSignedUpGetResponse instead.
type IsSignedUpResponseable interface {
    IsSignedUpGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
