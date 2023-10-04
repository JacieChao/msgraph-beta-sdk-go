package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemplatesItemCompareWithTemplateIdResponse 
// Deprecated: This class is obsolete. Use compareWithTemplateIdGetResponse instead.
type TemplatesItemCompareWithTemplateIdResponse struct {
    TemplatesItemCompareWithTemplateIdGetResponse
}
// NewTemplatesItemCompareWithTemplateIdResponse instantiates a new TemplatesItemCompareWithTemplateIdResponse and sets the default values.
func NewTemplatesItemCompareWithTemplateIdResponse()(*TemplatesItemCompareWithTemplateIdResponse) {
    m := &TemplatesItemCompareWithTemplateIdResponse{
        TemplatesItemCompareWithTemplateIdGetResponse: *NewTemplatesItemCompareWithTemplateIdGetResponse(),
    }
    return m
}
// CreateTemplatesItemCompareWithTemplateIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemplatesItemCompareWithTemplateIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemplatesItemCompareWithTemplateIdResponse(), nil
}
// TemplatesItemCompareWithTemplateIdResponseable 
// Deprecated: This class is obsolete. Use compareWithTemplateIdGetResponse instead.
type TemplatesItemCompareWithTemplateIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TemplatesItemCompareWithTemplateIdGetResponseable
}
