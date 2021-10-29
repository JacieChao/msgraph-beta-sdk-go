package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TenantCustomizedInformation struct {
    Entity
    contacts []TenantContactInformation;
    displayName *string;
    tenantId *string;
    website *string;
}
func NewTenantCustomizedInformation()(*TenantCustomizedInformation) {
    m := &TenantCustomizedInformation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TenantCustomizedInformation) GetContacts()([]TenantContactInformation) {
    if m == nil {
        return nil
    } else {
        return m.contacts
    }
}
func (m *TenantCustomizedInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *TenantCustomizedInformation) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *TenantCustomizedInformation) GetWebsite()(*string) {
    if m == nil {
        return nil
    } else {
        return m.website
    }
}
func (m *TenantCustomizedInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contacts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantContactInformation() })
        if err != nil {
            return err
        }
        res := make([]TenantContactInformation, len(val))
        for i, v := range val {
            res[i] = *(v.(*TenantContactInformation))
        }
        m.SetContacts(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["website"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebsite(val)
        return nil
    }
    return res
}
func (m *TenantCustomizedInformation) IsNil()(bool) {
    return m == nil
}
func (m *TenantCustomizedInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetContacts()))
        for i, v := range m.GetContacts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("contacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TenantCustomizedInformation) SetContacts(value []TenantContactInformation)() {
    m.contacts = value
}
func (m *TenantCustomizedInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *TenantCustomizedInformation) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *TenantCustomizedInformation) SetWebsite(value *string)() {
    m.website = value
}