package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody provides operations to call the assign method.
type ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignments property
    assignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable
}
// NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody instantiates a new ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody and sets the default values.
func NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody()(*ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) {
    m := &ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignments gets the assignments property value. The assignments property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) GetAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable) {
    return m.assignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignments", cast)
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
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignments sets the assignments property value. The assignments property
func (m *ReusablePolicySettingsItemReferencingConfigurationPoliciesItemAssignPostRequestBody) SetAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationPolicyAssignmentable)() {
    m.assignments = value
}
