package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPlannerTasksDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPlannerTasksDeltaResponse struct {
    ItemPlannerTasksDeltaGetResponse
}
// NewItemPlannerTasksDeltaResponse instantiates a new ItemPlannerTasksDeltaResponse and sets the default values.
func NewItemPlannerTasksDeltaResponse()(*ItemPlannerTasksDeltaResponse) {
    m := &ItemPlannerTasksDeltaResponse{
        ItemPlannerTasksDeltaGetResponse: *NewItemPlannerTasksDeltaGetResponse(),
    }
    return m
}
// CreateItemPlannerTasksDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPlannerTasksDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPlannerTasksDeltaResponse(), nil
}
// ItemPlannerTasksDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ItemPlannerTasksDeltaResponseable interface {
    ItemPlannerTasksDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
