package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EducationSynchronizationCustomization 
type EducationSynchronizationCustomization struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEducationSynchronizationCustomization instantiates a new educationSynchronizationCustomization and sets the default values.
func NewEducationSynchronizationCustomization()(*EducationSynchronizationCustomization) {
    m := &EducationSynchronizationCustomization{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationSynchronizationCustomizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationCustomizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationCustomization(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationCustomization) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAllowDisplayNameUpdate gets the allowDisplayNameUpdate property value. Indicates whether the display name of the resource can be overwritten by the sync.
func (m *EducationSynchronizationCustomization) GetAllowDisplayNameUpdate()(*bool) {
    val, err := m.GetBackingStore().Get("allowDisplayNameUpdate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *EducationSynchronizationCustomization) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationCustomization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowDisplayNameUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowDisplayNameUpdate(val)
        }
        return nil
    }
    res["isSyncDeferred"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSyncDeferred(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["optionalPropertiesToSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetOptionalPropertiesToSync(res)
        }
        return nil
    }
    res["synchronizationStartDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizationStartDate(val)
        }
        return nil
    }
    return res
}
// GetIsSyncDeferred gets the isSyncDeferred property value. Indicates whether synchronization of the parent entity is deferred to a later date.
func (m *EducationSynchronizationCustomization) GetIsSyncDeferred()(*bool) {
    val, err := m.GetBackingStore().Get("isSyncDeferred")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationCustomization) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOptionalPropertiesToSync gets the optionalPropertiesToSync property value. The collection of property names to sync. If set to null, all properties will be synchronized. Does not apply to Student Enrollments or Teacher Rosters
func (m *EducationSynchronizationCustomization) GetOptionalPropertiesToSync()([]string) {
    val, err := m.GetBackingStore().Get("optionalPropertiesToSync")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSynchronizationStartDate gets the synchronizationStartDate property value. The date that the synchronization should start. This value should be set to a future date. If set to null, the resource will be synchronized when the profile setup completes. Only applies to Student Enrollments
func (m *EducationSynchronizationCustomization) GetSynchronizationStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("synchronizationStartDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationSynchronizationCustomization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowDisplayNameUpdate", m.GetAllowDisplayNameUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSyncDeferred", m.GetIsSyncDeferred())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOptionalPropertiesToSync() != nil {
        err := writer.WriteCollectionOfStringValues("optionalPropertiesToSync", m.GetOptionalPropertiesToSync())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("synchronizationStartDate", m.GetSynchronizationStartDate())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationSynchronizationCustomization) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowDisplayNameUpdate sets the allowDisplayNameUpdate property value. Indicates whether the display name of the resource can be overwritten by the sync.
func (m *EducationSynchronizationCustomization) SetAllowDisplayNameUpdate(value *bool)() {
    err := m.GetBackingStore().Set("allowDisplayNameUpdate", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *EducationSynchronizationCustomization) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsSyncDeferred sets the isSyncDeferred property value. Indicates whether synchronization of the parent entity is deferred to a later date.
func (m *EducationSynchronizationCustomization) SetIsSyncDeferred(value *bool)() {
    err := m.GetBackingStore().Set("isSyncDeferred", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationSynchronizationCustomization) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOptionalPropertiesToSync sets the optionalPropertiesToSync property value. The collection of property names to sync. If set to null, all properties will be synchronized. Does not apply to Student Enrollments or Teacher Rosters
func (m *EducationSynchronizationCustomization) SetOptionalPropertiesToSync(value []string)() {
    err := m.GetBackingStore().Set("optionalPropertiesToSync", value)
    if err != nil {
        panic(err)
    }
}
// SetSynchronizationStartDate sets the synchronizationStartDate property value. The date that the synchronization should start. This value should be set to a future date. If set to null, the resource will be synchronized when the profile setup completes. Only applies to Student Enrollments
func (m *EducationSynchronizationCustomization) SetSynchronizationStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("synchronizationStartDate", value)
    if err != nil {
        panic(err)
    }
}
// EducationSynchronizationCustomizationable 
type EducationSynchronizationCustomizationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDisplayNameUpdate()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsSyncDeferred()(*bool)
    GetOdataType()(*string)
    GetOptionalPropertiesToSync()([]string)
    GetSynchronizationStartDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAllowDisplayNameUpdate(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsSyncDeferred(value *bool)()
    SetOdataType(value *string)()
    SetOptionalPropertiesToSync(value []string)()
    SetSynchronizationStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
