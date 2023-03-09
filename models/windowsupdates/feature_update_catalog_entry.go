package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureUpdateCatalogEntry 
type FeatureUpdateCatalogEntry struct {
    SoftwareUpdateCatalogEntry
}
// NewFeatureUpdateCatalogEntry instantiates a new FeatureUpdateCatalogEntry and sets the default values.
func NewFeatureUpdateCatalogEntry()(*FeatureUpdateCatalogEntry) {
    m := &FeatureUpdateCatalogEntry{
        SoftwareUpdateCatalogEntry: *NewSoftwareUpdateCatalogEntry(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.featureUpdateCatalogEntry"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFeatureUpdateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureUpdateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureUpdateCatalogEntry(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureUpdateCatalogEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SoftwareUpdateCatalogEntry.GetFieldDeserializers()
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetVersion gets the version property value. The version of the feature update. Read-only.
func (m *FeatureUpdateCatalogEntry) GetVersion()(*string) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FeatureUpdateCatalogEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SoftwareUpdateCatalogEntry.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVersion sets the version property value. The version of the feature update. Read-only.
func (m *FeatureUpdateCatalogEntry) SetVersion(value *string)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// FeatureUpdateCatalogEntryable 
type FeatureUpdateCatalogEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SoftwareUpdateCatalogEntryable
    GetVersion()(*string)
    SetVersion(value *string)()
}
