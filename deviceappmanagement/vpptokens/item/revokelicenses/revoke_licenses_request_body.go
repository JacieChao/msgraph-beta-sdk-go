package revokelicenses

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RevokeLicensesRequestBody struct {
    additionalData map[string]interface{};
    notifyManagedDevices *bool;
    revokeUntrackedLicenses *bool;
}
func NewRevokeLicensesRequestBody()(*RevokeLicensesRequestBody) {
    m := &RevokeLicensesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RevokeLicensesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RevokeLicensesRequestBody) GetNotifyManagedDevices()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notifyManagedDevices
    }
}
func (m *RevokeLicensesRequestBody) GetRevokeUntrackedLicenses()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.revokeUntrackedLicenses
    }
}
func (m *RevokeLicensesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notifyManagedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNotifyManagedDevices(val)
        return nil
    }
    res["revokeUntrackedLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRevokeUntrackedLicenses(val)
        return nil
    }
    return res
}
func (m *RevokeLicensesRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *RevokeLicensesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("notifyManagedDevices", m.GetNotifyManagedDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("revokeUntrackedLicenses", m.GetRevokeUntrackedLicenses())
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
func (m *RevokeLicensesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RevokeLicensesRequestBody) SetNotifyManagedDevices(value *bool)() {
    m.notifyManagedDevices = value
}
func (m *RevokeLicensesRequestBody) SetRevokeUntrackedLicenses(value *bool)() {
    m.revokeUntrackedLicenses = value
}