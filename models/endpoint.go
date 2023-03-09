package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Endpoint 
type Endpoint struct {
    DirectoryObject
}
// NewEndpoint instantiates a new endpoint and sets the default values.
func NewEndpoint()(*Endpoint) {
    m := &Endpoint{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.endpoint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndpoint(), nil
}
// GetCapability gets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.) Not nullable. Read-only.
func (m *Endpoint) GetCapability()(*string) {
    val, err := m.GetBackingStore().Get("capability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Endpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["capability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapability(val)
        }
        return nil
    }
    res["providerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderId(val)
        }
        return nil
    }
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["providerResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderResourceId(val)
        }
        return nil
    }
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetProviderId gets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
func (m *Endpoint) GetProviderId()(*string) {
    val, err := m.GetBackingStore().Get("providerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProviderName gets the providerName property value. Name of the publishing underlying service. Read-only.
func (m *Endpoint) GetProviderName()(*string) {
    val, err := m.GetBackingStore().Get("providerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProviderResourceId gets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
func (m *Endpoint) GetProviderResourceId()(*string) {
    val, err := m.GetBackingStore().Get("providerResourceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUri gets the uri property value. URL of the published resource. Not nullable. Read-only.
func (m *Endpoint) GetUri()(*string) {
    val, err := m.GetBackingStore().Get("uri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Endpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("capability", m.GetCapability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerId", m.GetProviderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerName", m.GetProviderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerResourceId", m.GetProviderResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCapability sets the capability property value. Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.) Not nullable. Read-only.
func (m *Endpoint) SetCapability(value *string)() {
    err := m.GetBackingStore().Set("capability", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderId sets the providerId property value. Application id of the publishing underlying service. Not nullable. Read-only.
func (m *Endpoint) SetProviderId(value *string)() {
    err := m.GetBackingStore().Set("providerId", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderName sets the providerName property value. Name of the publishing underlying service. Read-only.
func (m *Endpoint) SetProviderName(value *string)() {
    err := m.GetBackingStore().Set("providerName", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderResourceId sets the providerResourceId property value. For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only.
func (m *Endpoint) SetProviderResourceId(value *string)() {
    err := m.GetBackingStore().Set("providerResourceId", value)
    if err != nil {
        panic(err)
    }
}
// SetUri sets the uri property value. URL of the published resource. Not nullable. Read-only.
func (m *Endpoint) SetUri(value *string)() {
    err := m.GetBackingStore().Set("uri", value)
    if err != nil {
        panic(err)
    }
}
// Endpointable 
type Endpointable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapability()(*string)
    GetProviderId()(*string)
    GetProviderName()(*string)
    GetProviderResourceId()(*string)
    GetUri()(*string)
    SetCapability(value *string)()
    SetProviderId(value *string)()
    SetProviderName(value *string)()
    SetProviderResourceId(value *string)()
    SetUri(value *string)()
}
