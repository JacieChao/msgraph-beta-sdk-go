package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Schema 
type Schema struct {
    Entity
}
// NewSchema instantiates a new Schema and sets the default values.
func NewSchema()(*Schema) {
    m := &Schema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchema(), nil
}
// GetBaseType gets the baseType property value. The baseType property
func (m *Schema) GetBaseType()(*string) {
    val, err := m.GetBackingStore().Get("baseType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Schema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseType(val)
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Propertyable, len(val))
            for i, v := range val {
                res[i] = v.(Propertyable)
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
// GetProperties gets the properties property value. The properties property
func (m *Schema) GetProperties()([]Propertyable) {
    val, err := m.GetBackingStore().Get("properties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Propertyable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Schema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("baseType", m.GetBaseType())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseType sets the baseType property value. The baseType property
func (m *Schema) SetBaseType(value *string)() {
    err := m.GetBackingStore().Set("baseType", value)
    if err != nil {
        panic(err)
    }
}
// SetProperties sets the properties property value. The properties property
func (m *Schema) SetProperties(value []Propertyable)() {
    err := m.GetBackingStore().Set("properties", value)
    if err != nil {
        panic(err)
    }
}
// Schemaable 
type Schemaable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseType()(*string)
    GetProperties()([]Propertyable)
    SetBaseType(value *string)()
    SetProperties(value []Propertyable)()
}
