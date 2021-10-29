package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ServiceAnnouncement struct {
    Entity
    healthOverviews []ServiceHealth;
    issues []ServiceHealthIssue;
    messages []ServiceUpdateMessage;
}
func NewServiceAnnouncement()(*ServiceAnnouncement) {
    m := &ServiceAnnouncement{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ServiceAnnouncement) GetHealthOverviews()([]ServiceHealth) {
    if m == nil {
        return nil
    } else {
        return m.healthOverviews
    }
}
func (m *ServiceAnnouncement) GetIssues()([]ServiceHealthIssue) {
    if m == nil {
        return nil
    } else {
        return m.issues
    }
}
func (m *ServiceAnnouncement) GetMessages()([]ServiceUpdateMessage) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
func (m *ServiceAnnouncement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthOverviews"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealth() })
        if err != nil {
            return err
        }
        res := make([]ServiceHealth, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceHealth))
        }
        m.SetHealthOverviews(res)
        return nil
    }
    res["issues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssue() })
        if err != nil {
            return err
        }
        res := make([]ServiceHealthIssue, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceHealthIssue))
        }
        m.SetIssues(res)
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceUpdateMessage() })
        if err != nil {
            return err
        }
        res := make([]ServiceUpdateMessage, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceUpdateMessage))
        }
        m.SetMessages(res)
        return nil
    }
    return res
}
func (m *ServiceAnnouncement) IsNil()(bool) {
    return m == nil
}
func (m *ServiceAnnouncement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHealthOverviews()))
        for i, v := range m.GetHealthOverviews() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("healthOverviews", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIssues()))
        for i, v := range m.GetIssues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("issues", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ServiceAnnouncement) SetHealthOverviews(value []ServiceHealth)() {
    m.healthOverviews = value
}
func (m *ServiceAnnouncement) SetIssues(value []ServiceHealthIssue)() {
    m.issues = value
}
func (m *ServiceAnnouncement) SetMessages(value []ServiceUpdateMessage)() {
    m.messages = value
}