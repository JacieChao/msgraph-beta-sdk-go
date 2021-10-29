package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnPremisesAgentGroup struct {
    Entity
    agents []OnPremisesAgent;
    displayName *string;
    isDefault *bool;
    publishedResources []PublishedResource;
    publishingType *OnPremisesPublishingType;
}
func NewOnPremisesAgentGroup()(*OnPremisesAgentGroup) {
    m := &OnPremisesAgentGroup{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OnPremisesAgentGroup) GetAgents()([]OnPremisesAgent) {
    if m == nil {
        return nil
    } else {
        return m.agents
    }
}
func (m *OnPremisesAgentGroup) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *OnPremisesAgentGroup) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
func (m *OnPremisesAgentGroup) GetPublishedResources()([]PublishedResource) {
    if m == nil {
        return nil
    } else {
        return m.publishedResources
    }
}
func (m *OnPremisesAgentGroup) GetPublishingType()(*OnPremisesPublishingType) {
    if m == nil {
        return nil
    } else {
        return m.publishingType
    }
}
func (m *OnPremisesAgentGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesAgent() })
        if err != nil {
            return err
        }
        res := make([]OnPremisesAgent, len(val))
        for i, v := range val {
            res[i] = *(v.(*OnPremisesAgent))
        }
        m.SetAgents(res)
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
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["publishedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublishedResource() })
        if err != nil {
            return err
        }
        res := make([]PublishedResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*PublishedResource))
        }
        m.SetPublishedResources(res)
        return nil
    }
    res["publishingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnPremisesPublishingType)
        if err != nil {
            return err
        }
        cast := val.(OnPremisesPublishingType)
        m.SetPublishingType(&cast)
        return nil
    }
    return res
}
func (m *OnPremisesAgentGroup) IsNil()(bool) {
    return m == nil
}
func (m *OnPremisesAgentGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAgents()))
        for i, v := range m.GetAgents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("agents", cast)
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
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPublishedResources()))
        for i, v := range m.GetPublishedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("publishedResources", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublishingType() != nil {
        cast := m.GetPublishingType().String()
        err = writer.WriteStringValue("publishingType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OnPremisesAgentGroup) SetAgents(value []OnPremisesAgent)() {
    m.agents = value
}
func (m *OnPremisesAgentGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *OnPremisesAgentGroup) SetIsDefault(value *bool)() {
    m.isDefault = value
}
func (m *OnPremisesAgentGroup) SetPublishedResources(value []PublishedResource)() {
    m.publishedResources = value
}
func (m *OnPremisesAgentGroup) SetPublishingType(value *OnPremisesPublishingType)() {
    m.publishingType = value
}