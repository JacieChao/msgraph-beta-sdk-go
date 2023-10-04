package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WebPartData 
type WebPartData struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWebPartData instantiates a new webPartData and sets the default values.
func NewWebPartData()(*WebPartData) {
    m := &WebPartData{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWebPartDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebPartDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebPartData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WebPartData) GetAdditionalData()(map[string]any) {
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
// GetAudiences gets the audiences property value. Audience information of the web part. By using this property, specific content is prioritized to specific audiences.
func (m *WebPartData) GetAudiences()([]string) {
    val, err := m.GetBackingStore().Get("audiences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *WebPartData) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDataVersion gets the dataVersion property value. Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
func (m *WebPartData) GetDataVersion()(*string) {
    val, err := m.GetBackingStore().Get("dataVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the web part.
func (m *WebPartData) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebPartData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["audiences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAudiences(res)
        }
        return nil
    }
    res["dataVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataVersion(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(Jsonable))
        }
        return nil
    }
    res["serverProcessedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServerProcessedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerProcessedContent(val.(ServerProcessedContentable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WebPartData) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProperties gets the properties property value. Properties bag of the web part.
func (m *WebPartData) GetProperties()(Jsonable) {
    val, err := m.GetBackingStore().Get("properties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Jsonable)
    }
    return nil
}
// GetServerProcessedContent gets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *WebPartData) GetServerProcessedContent()(ServerProcessedContentable) {
    val, err := m.GetBackingStore().Get("serverProcessedContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ServerProcessedContentable)
    }
    return nil
}
// GetTitle gets the title property value. Title of the web part.
func (m *WebPartData) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WebPartData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAudiences() != nil {
        err := writer.WriteCollectionOfStringValues("audiences", m.GetAudiences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataVersion", m.GetDataVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    {
        err := writer.WriteObjectValue("properties", m.GetProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serverProcessedContent", m.GetServerProcessedContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *WebPartData) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAudiences sets the audiences property value. Audience information of the web part. By using this property, specific content is prioritized to specific audiences.
func (m *WebPartData) SetAudiences(value []string)() {
    err := m.GetBackingStore().Set("audiences", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WebPartData) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDataVersion sets the dataVersion property value. Data version of the web part. The value is defined by the web part developer. Different dataVersions usually refers to a different property structure.
func (m *WebPartData) SetDataVersion(value *string)() {
    err := m.GetBackingStore().Set("dataVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the web part.
func (m *WebPartData) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WebPartData) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProperties sets the properties property value. Properties bag of the web part.
func (m *WebPartData) SetProperties(value Jsonable)() {
    err := m.GetBackingStore().Set("properties", value)
    if err != nil {
        panic(err)
    }
}
// SetServerProcessedContent sets the serverProcessedContent property value. Contains collections of data that can be processed by server side services like search index and link fixup.
func (m *WebPartData) SetServerProcessedContent(value ServerProcessedContentable)() {
    err := m.GetBackingStore().Set("serverProcessedContent", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Title of the web part.
func (m *WebPartData) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// WebPartDataable 
type WebPartDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudiences()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDataVersion()(*string)
    GetDescription()(*string)
    GetOdataType()(*string)
    GetProperties()(Jsonable)
    GetServerProcessedContent()(ServerProcessedContentable)
    GetTitle()(*string)
    SetAudiences(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDataVersion(value *string)()
    SetDescription(value *string)()
    SetOdataType(value *string)()
    SetProperties(value Jsonable)()
    SetServerProcessedContent(value ServerProcessedContentable)()
    SetTitle(value *string)()
}
