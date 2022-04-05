package updatepriorities

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatePrioritiesRequestBody provides operations to call the updatePriorities method.
type UpdatePrioritiesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The officeConfigurationPolicyIds property
    officeConfigurationPolicyIds []string;
    // The officeConfigurationPriorities property
    officeConfigurationPriorities []int32;
}
// NewUpdatePrioritiesRequestBody instantiates a new updatePrioritiesRequestBody and sets the default values.
func NewUpdatePrioritiesRequestBody()(*UpdatePrioritiesRequestBody) {
    m := &UpdatePrioritiesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUpdatePrioritiesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePrioritiesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatePrioritiesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePrioritiesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePrioritiesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["officeConfigurationPolicyIds"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOfficeConfigurationPolicyIds(res)
        }
        return nil
    }
    res["officeConfigurationPriorities"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]int32, len(val))
            for i, v := range val {
                res[i] = *(v.(*int32))
            }
            m.SetOfficeConfigurationPriorities(res)
        }
        return nil
    }
    return res
}
// GetOfficeConfigurationPolicyIds gets the officeConfigurationPolicyIds property value. The officeConfigurationPolicyIds property
func (m *UpdatePrioritiesRequestBody) GetOfficeConfigurationPolicyIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationPolicyIds
    }
}
// GetOfficeConfigurationPriorities gets the officeConfigurationPriorities property value. The officeConfigurationPriorities property
func (m *UpdatePrioritiesRequestBody) GetOfficeConfigurationPriorities()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationPriorities
    }
}
// Serialize serializes information the current object
func (m *UpdatePrioritiesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOfficeConfigurationPolicyIds() != nil {
        err := writer.WriteCollectionOfStringValues("officeConfigurationPolicyIds", m.GetOfficeConfigurationPolicyIds())
        if err != nil {
            return err
        }
    }
    if m.GetOfficeConfigurationPriorities() != nil {
        err := writer.WriteCollectionOfInt32Values("officeConfigurationPriorities", m.GetOfficeConfigurationPriorities())
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
func (m *UpdatePrioritiesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOfficeConfigurationPolicyIds sets the officeConfigurationPolicyIds property value. The officeConfigurationPolicyIds property
func (m *UpdatePrioritiesRequestBody) SetOfficeConfigurationPolicyIds(value []string)() {
    if m != nil {
        m.officeConfigurationPolicyIds = value
    }
}
// SetOfficeConfigurationPriorities sets the officeConfigurationPriorities property value. The officeConfigurationPriorities property
func (m *UpdatePrioritiesRequestBody) SetOfficeConfigurationPriorities(value []int32)() {
    if m != nil {
        m.officeConfigurationPriorities = value
    }
}
