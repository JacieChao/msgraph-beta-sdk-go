package identity

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody 
type B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The newAssignmentOrder property
    newAssignmentOrder ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable
}
// NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody instantiates a new B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody and sets the default values.
func NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody()(*B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) {
    m := &B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewB2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newAssignmentOrder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignmentOrderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewAssignmentOrder(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable))
        }
        return nil
    }
    return res
}
// GetNewAssignmentOrder gets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) GetNewAssignmentOrder()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable) {
    return m.newAssignmentOrder
}
// Serialize serializes information the current object
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("newAssignmentOrder", m.GetNewAssignmentOrder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewAssignmentOrder sets the newAssignmentOrder property value. The newAssignmentOrder property
func (m *B2cUserFlowsItemUserAttributeAssignmentsSetOrderPostRequestBody) SetNewAssignmentOrder(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignmentOrderable)() {
    m.newAssignmentOrder = value
}