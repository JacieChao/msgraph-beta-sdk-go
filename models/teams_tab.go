package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsTab 
type TeamsTab struct {
    Entity
}
// NewTeamsTab instantiates a new teamsTab and sets the default values.
func NewTeamsTab()(*TeamsTab) {
    m := &TeamsTab{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsTabFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsTabFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsTab(), nil
}
// GetConfiguration gets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) GetConfiguration()(TeamsTabConfigurationable) {
    val, err := m.GetBackingStore().Get("configuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamsTabConfigurationable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the tab.
func (m *TeamsTab) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsTab) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsTabConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(TeamsTabConfigurationable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["messageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["sortOrderIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortOrderIndex(val)
        }
        return nil
    }
    res["teamsApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsApp(val.(TeamsAppable))
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the messageId property value. The messageId property
func (m *TeamsTab) GetMessageId()(*string) {
    val, err := m.GetBackingStore().Get("messageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSortOrderIndex gets the sortOrderIndex property value. Index of the order used for sorting tabs.
func (m *TeamsTab) GetSortOrderIndex()(*string) {
    val, err := m.GetBackingStore().Get("sortOrderIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamsApp gets the teamsApp property value. The application that is linked to the tab.
func (m *TeamsTab) GetTeamsApp()(TeamsAppable) {
    val, err := m.GetBackingStore().Get("teamsApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamsAppable)
    }
    return nil
}
// GetTeamsAppId gets the teamsAppId property value. The teamsAppId property
func (m *TeamsTab) GetTeamsAppId()(*string) {
    val, err := m.GetBackingStore().Get("teamsAppId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebUrl gets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) GetWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("webUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamsTab) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
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
        err = writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sortOrderIndex", m.GetSortOrderIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) SetConfiguration(value TeamsTabConfigurationable)() {
    err := m.GetBackingStore().Set("configuration", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the tab.
func (m *TeamsTab) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetMessageId sets the messageId property value. The messageId property
func (m *TeamsTab) SetMessageId(value *string)() {
    err := m.GetBackingStore().Set("messageId", value)
    if err != nil {
        panic(err)
    }
}
// SetSortOrderIndex sets the sortOrderIndex property value. Index of the order used for sorting tabs.
func (m *TeamsTab) SetSortOrderIndex(value *string)() {
    err := m.GetBackingStore().Set("sortOrderIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsApp sets the teamsApp property value. The application that is linked to the tab.
func (m *TeamsTab) SetTeamsApp(value TeamsAppable)() {
    err := m.GetBackingStore().Set("teamsApp", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsAppId sets the teamsAppId property value. The teamsAppId property
func (m *TeamsTab) SetTeamsAppId(value *string)() {
    err := m.GetBackingStore().Set("teamsAppId", value)
    if err != nil {
        panic(err)
    }
}
// SetWebUrl sets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) SetWebUrl(value *string)() {
    err := m.GetBackingStore().Set("webUrl", value)
    if err != nil {
        panic(err)
    }
}
// TeamsTabable 
type TeamsTabable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(TeamsTabConfigurationable)
    GetDisplayName()(*string)
    GetMessageId()(*string)
    GetSortOrderIndex()(*string)
    GetTeamsApp()(TeamsAppable)
    GetTeamsAppId()(*string)
    GetWebUrl()(*string)
    SetConfiguration(value TeamsTabConfigurationable)()
    SetDisplayName(value *string)()
    SetMessageId(value *string)()
    SetSortOrderIndex(value *string)()
    SetTeamsApp(value TeamsAppable)()
    SetTeamsAppId(value *string)()
    SetWebUrl(value *string)()
}
