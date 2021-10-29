package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecurityBaselineState struct {
    Entity
    displayName *string;
    securityBaselineTemplateId *string;
    settingStates []SecurityBaselineSettingState;
    state *SecurityBaselineComplianceState;
    userPrincipalName *string;
}
func NewSecurityBaselineState()(*SecurityBaselineState) {
    m := &SecurityBaselineState{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SecurityBaselineState) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *SecurityBaselineState) GetSecurityBaselineTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.securityBaselineTemplateId
    }
}
func (m *SecurityBaselineState) GetSettingStates()([]SecurityBaselineSettingState) {
    if m == nil {
        return nil
    } else {
        return m.settingStates
    }
}
func (m *SecurityBaselineState) GetState()(*SecurityBaselineComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *SecurityBaselineState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *SecurityBaselineState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["securityBaselineTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSecurityBaselineTemplateId(val)
        return nil
    }
    res["settingStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityBaselineSettingState() })
        if err != nil {
            return err
        }
        res := make([]SecurityBaselineSettingState, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecurityBaselineSettingState))
        }
        m.SetSettingStates(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityBaselineComplianceState)
        if err != nil {
            return err
        }
        cast := val.(SecurityBaselineComplianceState)
        m.SetState(&cast)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *SecurityBaselineState) IsNil()(bool) {
    return m == nil
}
func (m *SecurityBaselineState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityBaselineTemplateId", m.GetSecurityBaselineTemplateId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SecurityBaselineState) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *SecurityBaselineState) SetSecurityBaselineTemplateId(value *string)() {
    m.securityBaselineTemplateId = value
}
func (m *SecurityBaselineState) SetSettingStates(value []SecurityBaselineSettingState)() {
    m.settingStates = value
}
func (m *SecurityBaselineState) SetState(value *SecurityBaselineComplianceState)() {
    m.state = value
}
func (m *SecurityBaselineState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}