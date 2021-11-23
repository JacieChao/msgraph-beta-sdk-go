package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// windowsInformationProtectionApp 
type WindowsInformationProtectionApp struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If true, app is denied protection or exemption.
    denied *bool;
    // The app's description.
    description *string;
    // App display name.
    displayName *string;
    // The product name.
    productName *string;
    // The publisher name
    publisherName *string;
}
// NewWindowsInformationProtectionApp instantiates a new windowsInformationProtectionApp and sets the default values.
func NewWindowsInformationProtectionApp()(*WindowsInformationProtectionApp) {
    m := &WindowsInformationProtectionApp{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsInformationProtectionApp) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDenied gets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) GetDenied()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.denied
    }
}
// GetDescription gets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetProductName gets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) GetProductName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productName
    }
}
// GetPublisherName gets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) GetPublisherName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisherName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["denied"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDenied(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["productName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisherName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    return res
}
func (m *WindowsInformationProtectionApp) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("denied", m.GetDenied())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("publisherName", m.GetPublisherName())
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
func (m *WindowsInformationProtectionApp) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDenied sets the denied property value. If true, app is denied protection or exemption.
func (m *WindowsInformationProtectionApp) SetDenied(value *bool)() {
    m.denied = value
}
// SetDescription sets the description property value. The app's description.
func (m *WindowsInformationProtectionApp) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. App display name.
func (m *WindowsInformationProtectionApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetProductName sets the productName property value. The product name.
func (m *WindowsInformationProtectionApp) SetProductName(value *string)() {
    m.productName = value
}
// SetPublisherName sets the publisherName property value. The publisher name
func (m *WindowsInformationProtectionApp) SetPublisherName(value *string)() {
    m.publisherName = value
}
