package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserRegistrationCount struct {
    additionalData map[string]interface{};
    registrationCount *int64;
    registrationStatus *RegistrationStatusType;
}
func NewUserRegistrationCount()(*UserRegistrationCount) {
    m := &UserRegistrationCount{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserRegistrationCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserRegistrationCount) GetRegistrationCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.registrationCount
    }
}
func (m *UserRegistrationCount) GetRegistrationStatus()(*RegistrationStatusType) {
    if m == nil {
        return nil
    } else {
        return m.registrationStatus
    }
}
func (m *UserRegistrationCount) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["registrationCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetRegistrationCount(val)
        return nil
    }
    res["registrationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistrationStatusType)
        if err != nil {
            return err
        }
        cast := val.(RegistrationStatusType)
        m.SetRegistrationStatus(&cast)
        return nil
    }
    return res
}
func (m *UserRegistrationCount) IsNil()(bool) {
    return m == nil
}
func (m *UserRegistrationCount) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("registrationCount", m.GetRegistrationCount())
        if err != nil {
            return err
        }
    }
    if m.GetRegistrationStatus() != nil {
        cast := m.GetRegistrationStatus().String()
        err := writer.WriteStringValue("registrationStatus", &cast)
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
func (m *UserRegistrationCount) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserRegistrationCount) SetRegistrationCount(value *int64)() {
    m.registrationCount = value
}
func (m *UserRegistrationCount) SetRegistrationStatus(value *RegistrationStatusType)() {
    m.registrationStatus = value
}