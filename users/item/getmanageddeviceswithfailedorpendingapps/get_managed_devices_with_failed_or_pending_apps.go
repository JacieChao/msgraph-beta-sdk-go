package getmanageddeviceswithfailedorpendingapps

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// getManagedDevicesWithFailedOrPendingApps 
type GetManagedDevicesWithFailedOrPendingApps struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewGetManagedDevicesWithFailedOrPendingApps instantiates a new getManagedDevicesWithFailedOrPendingApps and sets the default values.
func NewGetManagedDevicesWithFailedOrPendingApps()(*GetManagedDevicesWithFailedOrPendingApps) {
    m := &GetManagedDevicesWithFailedOrPendingApps{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetManagedDevicesWithFailedOrPendingApps) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetManagedDevicesWithFailedOrPendingApps) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *GetManagedDevicesWithFailedOrPendingApps) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetManagedDevicesWithFailedOrPendingApps) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetManagedDevicesWithFailedOrPendingApps) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
