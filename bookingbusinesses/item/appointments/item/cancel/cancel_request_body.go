package cancel

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CancelRequestBody provides operations to call the cancel method.
type CancelRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The cancellationMessage property
    cancellationMessage *string;
}
// NewCancelRequestBody instantiates a new cancelRequestBody and sets the default values.
func NewCancelRequestBody()(*CancelRequestBody) {
    m := &CancelRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCancelRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCancelRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCancelRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CancelRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCancellationMessage gets the cancellationMessage property value. The cancellationMessage property
func (m *CancelRequestBody) GetCancellationMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cancellationMessage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CancelRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cancellationMessage"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancellationMessage(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CancelRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cancellationMessage", m.GetCancellationMessage())
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
func (m *CancelRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCancellationMessage sets the cancellationMessage property value. The cancellationMessage property
func (m *CancelRequestBody) SetCancellationMessage(value *string)() {
    if m != nil {
        m.cancellationMessage = value
    }
}
