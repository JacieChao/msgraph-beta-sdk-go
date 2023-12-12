package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DayNote 
type DayNote struct {
    ChangeTrackedEntity
}
// NewDayNote instantiates a new dayNote and sets the default values.
func NewDayNote()(*DayNote) {
    m := &DayNote{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.dayNote"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDayNoteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDayNoteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDayNote(), nil
}
// GetDayNoteDate gets the dayNoteDate property value. The date of the day note.
func (m *DayNote) GetDayNoteDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("dayNoteDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetDraftDayNote gets the draftDayNote property value. The draft version of this day note that is viewable by managers. Only contentType text is supported.
func (m *DayNote) GetDraftDayNote()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("draftDayNote")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DayNote) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["dayNoteDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayNoteDate(val)
        }
        return nil
    }
    res["draftDayNote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftDayNote(val.(ItemBodyable))
        }
        return nil
    }
    res["sharedDayNote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedDayNote(val.(ItemBodyable))
        }
        return nil
    }
    return res
}
// GetSharedDayNote gets the sharedDayNote property value. The shared version of this day note that is viewable by both employees and managers. Only contentType text is supported.
func (m *DayNote) GetSharedDayNote()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("sharedDayNote")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DayNote) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("dayNoteDate", m.GetDayNoteDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("draftDayNote", m.GetDraftDayNote())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedDayNote", m.GetSharedDayNote())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDayNoteDate sets the dayNoteDate property value. The date of the day note.
func (m *DayNote) SetDayNoteDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("dayNoteDate", value)
    if err != nil {
        panic(err)
    }
}
// SetDraftDayNote sets the draftDayNote property value. The draft version of this day note that is viewable by managers. Only contentType text is supported.
func (m *DayNote) SetDraftDayNote(value ItemBodyable)() {
    err := m.GetBackingStore().Set("draftDayNote", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedDayNote sets the sharedDayNote property value. The shared version of this day note that is viewable by both employees and managers. Only contentType text is supported.
func (m *DayNote) SetSharedDayNote(value ItemBodyable)() {
    err := m.GetBackingStore().Set("sharedDayNote", value)
    if err != nil {
        panic(err)
    }
}
// DayNoteable 
type DayNoteable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDayNoteDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDraftDayNote()(ItemBodyable)
    GetSharedDayNote()(ItemBodyable)
    SetDayNoteDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDraftDayNote(value ItemBodyable)()
    SetSharedDayNote(value ItemBodyable)()
}
