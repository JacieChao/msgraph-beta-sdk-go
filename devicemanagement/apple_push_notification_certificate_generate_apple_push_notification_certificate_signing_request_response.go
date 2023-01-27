package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse 
type ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value *string
}
// NewApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse instantiates a new ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse and sets the default values.
func NewApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse()(*ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) {
    m := &ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestResponse) SetValue(value *string)() {
    m.value = value
}