package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PasswordSingleSignOnSettings struct {
    additionalData map[string]interface{};
    fields []PasswordSingleSignOnField;
}
func NewPasswordSingleSignOnSettings()(*PasswordSingleSignOnSettings) {
    m := &PasswordSingleSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PasswordSingleSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PasswordSingleSignOnSettings) GetFields()([]PasswordSingleSignOnField) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
func (m *PasswordSingleSignOnSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordSingleSignOnField() })
        if err != nil {
            return err
        }
        res := make([]PasswordSingleSignOnField, len(val))
        for i, v := range val {
            res[i] = *(v.(*PasswordSingleSignOnField))
        }
        m.SetFields(res)
        return nil
    }
    return res
}
func (m *PasswordSingleSignOnSettings) IsNil()(bool) {
    return m == nil
}
func (m *PasswordSingleSignOnSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
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
func (m *PasswordSingleSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PasswordSingleSignOnSettings) SetFields(value []PasswordSingleSignOnField)() {
    m.fields = value
}