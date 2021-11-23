package windowsdefenderscan

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// windowsDefenderScanRequestBody 
type WindowsDefenderScanRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    quickScan *bool;
}
// NewWindowsDefenderScanRequestBody instantiates a new windowsDefenderScanRequestBody and sets the default values.
func NewWindowsDefenderScanRequestBody()(*WindowsDefenderScanRequestBody) {
    m := &WindowsDefenderScanRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDefenderScanRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetQuickScan gets the quickScan property value. 
func (m *WindowsDefenderScanRequestBody) GetQuickScan()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.quickScan
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderScanRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["quickScan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScan(val)
        }
        return nil
    }
    return res
}
func (m *WindowsDefenderScanRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDefenderScanRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("quickScan", m.GetQuickScan())
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
func (m *WindowsDefenderScanRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetQuickScan sets the quickScan property value. 
func (m *WindowsDefenderScanRequestBody) SetQuickScan(value *bool)() {
    m.quickScan = value
}
