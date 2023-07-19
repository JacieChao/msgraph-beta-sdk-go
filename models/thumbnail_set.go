package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ThumbnailSet 
type ThumbnailSet struct {
    Entity
}
// NewThumbnailSet instantiates a new thumbnailSet and sets the default values.
func NewThumbnailSet()(*ThumbnailSet) {
    m := &ThumbnailSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateThumbnailSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateThumbnailSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewThumbnailSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThumbnailSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["large"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLarge(val.(Thumbnailable))
        }
        return nil
    }
    res["medium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedium(val.(Thumbnailable))
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
    res["small"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmall(val.(Thumbnailable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateThumbnailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(Thumbnailable))
        }
        return nil
    }
    return res
}
// GetLarge gets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) GetLarge()(Thumbnailable) {
    val, err := m.GetBackingStore().Get("large")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Thumbnailable)
    }
    return nil
}
// GetMedium gets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) GetMedium()(Thumbnailable) {
    val, err := m.GetBackingStore().Get("medium")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Thumbnailable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ThumbnailSet) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSmall gets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) GetSmall()(Thumbnailable) {
    val, err := m.GetBackingStore().Get("small")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Thumbnailable)
    }
    return nil
}
// GetSource gets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) GetSource()(Thumbnailable) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Thumbnailable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ThumbnailSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("large", m.GetLarge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("medium", m.GetMedium())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("small", m.GetSmall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLarge sets the large property value. A 1920x1920 scaled thumbnail.
func (m *ThumbnailSet) SetLarge(value Thumbnailable)() {
    err := m.GetBackingStore().Set("large", value)
    if err != nil {
        panic(err)
    }
}
// SetMedium sets the medium property value. A 176x176 scaled thumbnail.
func (m *ThumbnailSet) SetMedium(value Thumbnailable)() {
    err := m.GetBackingStore().Set("medium", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ThumbnailSet) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSmall sets the small property value. A 48x48 cropped thumbnail.
func (m *ThumbnailSet) SetSmall(value Thumbnailable)() {
    err := m.GetBackingStore().Set("small", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. A custom thumbnail image or the original image used to generate other thumbnails.
func (m *ThumbnailSet) SetSource(value Thumbnailable)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
// ThumbnailSetable 
type ThumbnailSetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLarge()(Thumbnailable)
    GetMedium()(Thumbnailable)
    GetOdataType()(*string)
    GetSmall()(Thumbnailable)
    GetSource()(Thumbnailable)
    SetLarge(value Thumbnailable)()
    SetMedium(value Thumbnailable)()
    SetOdataType(value *string)()
    SetSmall(value Thumbnailable)()
    SetSource(value Thumbnailable)()
}
