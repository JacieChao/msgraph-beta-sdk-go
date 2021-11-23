package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// userSet 
type UserSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
    isBackup *bool;
}
// NewUserSet instantiates a new userSet and sets the default values.
func NewUserSet()(*UserSet) {
    m := &UserSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsBackup gets the isBackup property value. For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
func (m *UserSet) GetIsBackup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBackup
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isBackup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBackup(val)
        }
        return nil
    }
    return res
}
func (m *UserSet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isBackup", m.GetIsBackup())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsBackup sets the isBackup property value. For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
func (m *UserSet) SetIsBackup(value *bool)() {
    m.isBackup = value
}
