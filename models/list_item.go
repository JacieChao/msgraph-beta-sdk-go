package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListItem 
type ListItem struct {
    BaseItem
}
// NewListItem instantiates a new listItem and sets the default values.
func NewListItem()(*ListItem) {
    m := &ListItem{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.listItem"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateListItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewListItem(), nil
}
// GetActivities gets the activities property value. The list of recent activities that took place on this item.
func (m *ListItem) GetActivities()([]ItemActivityOLDable) {
    val, err := m.GetBackingStore().Get("activities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemActivityOLDable)
    }
    return nil
}
// GetAnalytics gets the analytics property value. Analytics about the view activities that took place on this item.
func (m *ListItem) GetAnalytics()(ItemAnalyticsable) {
    val, err := m.GetBackingStore().Get("analytics")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemAnalyticsable)
    }
    return nil
}
// GetContentType gets the contentType property value. The content type of this list item
func (m *ListItem) GetContentType()(ContentTypeInfoable) {
    val, err := m.GetBackingStore().Get("contentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ContentTypeInfoable)
    }
    return nil
}
// GetDeleted gets the deleted property value. The deleted property
func (m *ListItem) GetDeleted()(Deletedable) {
    val, err := m.GetBackingStore().Get("deleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Deletedable)
    }
    return nil
}
// GetDocumentSetVersions gets the documentSetVersions property value. Version information for a document set version created by a user.
func (m *ListItem) GetDocumentSetVersions()([]DocumentSetVersionable) {
    val, err := m.GetBackingStore().Get("documentSetVersions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DocumentSetVersionable)
    }
    return nil
}
// GetDriveItem gets the driveItem property value. For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
func (m *ListItem) GetDriveItem()(DriveItemable) {
    val, err := m.GetBackingStore().Get("driveItem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DriveItemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["activities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemActivityOLDFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemActivityOLDable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ItemActivityOLDable)
                }
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["analytics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalytics(val.(ItemAnalyticsable))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
        }
        return nil
    }
    res["deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeletedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val.(Deletedable))
        }
        return nil
    }
    res["documentSetVersions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDocumentSetVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentSetVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DocumentSetVersionable)
                }
            }
            m.SetDocumentSetVersions(res)
        }
        return nil
    }
    res["driveItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFieldValueSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(FieldValueSetable))
        }
        return nil
    }
    res["sharepointIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListItemVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItemVersionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ListItemVersionable)
                }
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The values of the columns set on this list item.
func (m *ListItem) GetFields()(FieldValueSetable) {
    val, err := m.GetBackingStore().Get("fields")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(FieldValueSetable)
    }
    return nil
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *ListItem) GetSharepointIds()(SharepointIdsable) {
    val, err := m.GetBackingStore().Get("sharepointIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharepointIdsable)
    }
    return nil
}
// GetVersions gets the versions property value. The list of previous versions of the list item.
func (m *ListItem) GetVersions()([]ListItemVersionable) {
    val, err := m.GetBackingStore().Get("versions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ListItemVersionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ListItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("analytics", m.GetAnalytics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deleted", m.GetDeleted())
        if err != nil {
            return err
        }
    }
    if m.GetDocumentSetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDocumentSetVersions()))
        for i, v := range m.GetDocumentSetVersions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("documentSetVersions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivities sets the activities property value. The list of recent activities that took place on this item.
func (m *ListItem) SetActivities(value []ItemActivityOLDable)() {
    err := m.GetBackingStore().Set("activities", value)
    if err != nil {
        panic(err)
    }
}
// SetAnalytics sets the analytics property value. Analytics about the view activities that took place on this item.
func (m *ListItem) SetAnalytics(value ItemAnalyticsable)() {
    err := m.GetBackingStore().Set("analytics", value)
    if err != nil {
        panic(err)
    }
}
// SetContentType sets the contentType property value. The content type of this list item
func (m *ListItem) SetContentType(value ContentTypeInfoable)() {
    err := m.GetBackingStore().Set("contentType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeleted sets the deleted property value. The deleted property
func (m *ListItem) SetDeleted(value Deletedable)() {
    err := m.GetBackingStore().Set("deleted", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentSetVersions sets the documentSetVersions property value. Version information for a document set version created by a user.
func (m *ListItem) SetDocumentSetVersions(value []DocumentSetVersionable)() {
    err := m.GetBackingStore().Set("documentSetVersions", value)
    if err != nil {
        panic(err)
    }
}
// SetDriveItem sets the driveItem property value. For document libraries, the driveItem relationship exposes the listItem as a [driveItem][]
func (m *ListItem) SetDriveItem(value DriveItemable)() {
    err := m.GetBackingStore().Set("driveItem", value)
    if err != nil {
        panic(err)
    }
}
// SetFields sets the fields property value. The values of the columns set on this list item.
func (m *ListItem) SetFields(value FieldValueSetable)() {
    err := m.GetBackingStore().Set("fields", value)
    if err != nil {
        panic(err)
    }
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *ListItem) SetSharepointIds(value SharepointIdsable)() {
    err := m.GetBackingStore().Set("sharepointIds", value)
    if err != nil {
        panic(err)
    }
}
// SetVersions sets the versions property value. The list of previous versions of the list item.
func (m *ListItem) SetVersions(value []ListItemVersionable)() {
    err := m.GetBackingStore().Set("versions", value)
    if err != nil {
        panic(err)
    }
}
// ListItemable 
type ListItemable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivities()([]ItemActivityOLDable)
    GetAnalytics()(ItemAnalyticsable)
    GetContentType()(ContentTypeInfoable)
    GetDeleted()(Deletedable)
    GetDocumentSetVersions()([]DocumentSetVersionable)
    GetDriveItem()(DriveItemable)
    GetFields()(FieldValueSetable)
    GetSharepointIds()(SharepointIdsable)
    GetVersions()([]ListItemVersionable)
    SetActivities(value []ItemActivityOLDable)()
    SetAnalytics(value ItemAnalyticsable)()
    SetContentType(value ContentTypeInfoable)()
    SetDeleted(value Deletedable)()
    SetDocumentSetVersions(value []DocumentSetVersionable)()
    SetDriveItem(value DriveItemable)()
    SetFields(value FieldValueSetable)()
    SetSharepointIds(value SharepointIdsable)()
    SetVersions(value []ListItemVersionable)()
}
