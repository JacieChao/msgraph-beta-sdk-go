package print

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// PrintersCreatePostRequestBody 
type PrintersCreatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The certificateSigningRequest property
    certificateSigningRequest ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintCertificateSigningRequestable
    // The connectorId property
    connectorId *string
    // The displayName property
    displayName *string
    // The hasPhysicalDevice property
    hasPhysicalDevice *bool
    // The manufacturer property
    manufacturer *string
    // The model property
    model *string
    // The physicalDeviceId property
    physicalDeviceId *string
}
// NewPrintersCreatePostRequestBody instantiates a new PrintersCreatePostRequestBody and sets the default values.
func NewPrintersCreatePostRequestBody()(*PrintersCreatePostRequestBody) {
    m := &PrintersCreatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrintersCreatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintersCreatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintersCreatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintersCreatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCertificateSigningRequest gets the certificateSigningRequest property value. The certificateSigningRequest property
func (m *PrintersCreatePostRequestBody) GetCertificateSigningRequest()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintCertificateSigningRequestable) {
    return m.certificateSigningRequest
}
// GetConnectorId gets the connectorId property value. The connectorId property
func (m *PrintersCreatePostRequestBody) GetConnectorId()(*string) {
    return m.connectorId
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PrintersCreatePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintersCreatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificateSigningRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrintCertificateSigningRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSigningRequest(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintCertificateSigningRequestable))
        }
        return nil
    }
    res["connectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
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
    res["hasPhysicalDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPhysicalDevice(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["physicalDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalDeviceId(val)
        }
        return nil
    }
    return res
}
// GetHasPhysicalDevice gets the hasPhysicalDevice property value. The hasPhysicalDevice property
func (m *PrintersCreatePostRequestBody) GetHasPhysicalDevice()(*bool) {
    return m.hasPhysicalDevice
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
func (m *PrintersCreatePostRequestBody) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model property
func (m *PrintersCreatePostRequestBody) GetModel()(*string) {
    return m.model
}
// GetPhysicalDeviceId gets the physicalDeviceId property value. The physicalDeviceId property
func (m *PrintersCreatePostRequestBody) GetPhysicalDeviceId()(*string) {
    return m.physicalDeviceId
}
// Serialize serializes information the current object
func (m *PrintersCreatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("certificateSigningRequest", m.GetCertificateSigningRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorId", m.GetConnectorId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasPhysicalDevice", m.GetHasPhysicalDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("physicalDeviceId", m.GetPhysicalDeviceId())
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
func (m *PrintersCreatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCertificateSigningRequest sets the certificateSigningRequest property value. The certificateSigningRequest property
func (m *PrintersCreatePostRequestBody) SetCertificateSigningRequest(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrintCertificateSigningRequestable)() {
    m.certificateSigningRequest = value
}
// SetConnectorId sets the connectorId property value. The connectorId property
func (m *PrintersCreatePostRequestBody) SetConnectorId(value *string)() {
    m.connectorId = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PrintersCreatePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHasPhysicalDevice sets the hasPhysicalDevice property value. The hasPhysicalDevice property
func (m *PrintersCreatePostRequestBody) SetHasPhysicalDevice(value *bool)() {
    m.hasPhysicalDevice = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *PrintersCreatePostRequestBody) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model property
func (m *PrintersCreatePostRequestBody) SetModel(value *string)() {
    m.model = value
}
// SetPhysicalDeviceId sets the physicalDeviceId property value. The physicalDeviceId property
func (m *PrintersCreatePostRequestBody) SetPhysicalDeviceId(value *string)() {
    m.physicalDeviceId = value
}
