package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentFilterState struct {
    additionalData map[string]interface{};
    enabled *bool;
}
func NewAssignmentFilterState()(*AssignmentFilterState) {
    m := &AssignmentFilterState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignmentFilterState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignmentFilterState) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
func (m *AssignmentFilterState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEnabled(val)
        return nil
    }
    return res
}
func (m *AssignmentFilterState) IsNil()(bool) {
    return m == nil
}
func (m *AssignmentFilterState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *AssignmentFilterState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignmentFilterState) SetEnabled(value *bool)() {
    m.enabled = value
}