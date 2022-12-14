package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CloudPCsItemReprovisionPostRequestBody provides operations to call the reprovision method.
type CloudPCsItemReprovisionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The osVersion property
    osVersion *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem
    // The userAccountType property
    userAccountType *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType
}
// NewCloudPCsItemReprovisionPostRequestBody instantiates a new CloudPCsItemReprovisionPostRequestBody and sets the default values.
func NewCloudPCsItemReprovisionPostRequestBody()(*CloudPCsItemReprovisionPostRequestBody) {
    m := &CloudPCsItemReprovisionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPCsItemReprovisionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCsItemReprovisionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPCsItemReprovisionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPCsItemReprovisionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPCsItemReprovisionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem))
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType))
        }
        return nil
    }
    return res
}
// GetOsVersion gets the osVersion property value. The osVersion property
func (m *CloudPCsItemReprovisionPostRequestBody) GetOsVersion()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem) {
    return m.osVersion
}
// GetUserAccountType gets the userAccountType property value. The userAccountType property
func (m *CloudPCsItemReprovisionPostRequestBody) GetUserAccountType()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType) {
    return m.userAccountType
}
// Serialize serializes information the current object
func (m *CloudPCsItemReprovisionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetOsVersion() != nil {
        cast := (*m.GetOsVersion()).String()
        err := writer.WriteStringValue("osVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err := writer.WriteStringValue("userAccountType", &cast)
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
func (m *CloudPCsItemReprovisionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOsVersion sets the osVersion property value. The osVersion property
func (m *CloudPCsItemReprovisionPostRequestBody) SetOsVersion(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcOperatingSystem)() {
    m.osVersion = value
}
// SetUserAccountType sets the userAccountType property value. The userAccountType property
func (m *CloudPCsItemReprovisionPostRequestBody) SetUserAccountType(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcUserAccountType)() {
    m.userAccountType = value
}
