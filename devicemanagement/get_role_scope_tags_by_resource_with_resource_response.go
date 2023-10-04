package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetRoleScopeTagsByResourceWithResourceResponse 
// Deprecated: This class is obsolete. Use getRoleScopeTagsByResourceWithResourceGetResponse instead.
type GetRoleScopeTagsByResourceWithResourceResponse struct {
    GetRoleScopeTagsByResourceWithResourceGetResponse
}
// NewGetRoleScopeTagsByResourceWithResourceResponse instantiates a new GetRoleScopeTagsByResourceWithResourceResponse and sets the default values.
func NewGetRoleScopeTagsByResourceWithResourceResponse()(*GetRoleScopeTagsByResourceWithResourceResponse) {
    m := &GetRoleScopeTagsByResourceWithResourceResponse{
        GetRoleScopeTagsByResourceWithResourceGetResponse: *NewGetRoleScopeTagsByResourceWithResourceGetResponse(),
    }
    return m
}
// CreateGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetRoleScopeTagsByResourceWithResourceResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetRoleScopeTagsByResourceWithResourceResponse(), nil
}
// GetRoleScopeTagsByResourceWithResourceResponseable 
// Deprecated: This class is obsolete. Use getRoleScopeTagsByResourceWithResourceGetResponse instead.
type GetRoleScopeTagsByResourceWithResourceResponseable interface {
    GetRoleScopeTagsByResourceWithResourceGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
