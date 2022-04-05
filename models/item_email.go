package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemEmail 
type ItemEmail struct {
    ItemFacet
    // The email address itself.
    address *string;
    // The name or label a user has associated with a particular email address.
    displayName *string;
    // The type of email address. Possible values are: unknown, work, personal, main, other.
    type_escaped *EmailType;
}
// NewItemEmail instantiates a new itemEmail and sets the default values.
func NewItemEmail()(*ItemEmail) {
    m := &ItemEmail{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// CreateItemEmailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemEmailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemEmail(), nil
}
// GetAddress gets the address property value. The email address itself.
func (m *ItemEmail) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetDisplayName gets the displayName property value. The name or label a user has associated with a particular email address.
func (m *ItemEmail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemEmail) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*EmailType))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type of email address. Possible values are: unknown, work, personal, main, other.
func (m *ItemEmail) GetType()(*EmailType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *ItemEmail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("address", m.GetAddress())
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The email address itself.
func (m *ItemEmail) SetAddress(value *string)() {
    if m != nil {
        m.address = value
    }
}
// SetDisplayName sets the displayName property value. The name or label a user has associated with a particular email address.
func (m *ItemEmail) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetType sets the type property value. The type of email address. Possible values are: unknown, work, personal, main, other.
func (m *ItemEmail) SetType(value *EmailType)() {
    if m != nil {
        m.type_escaped = value
    }
}