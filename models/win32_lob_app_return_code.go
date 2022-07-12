package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppReturnCode contains return code properties for a Win32 App
type Win32LobAppReturnCode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Return code.
    returnCode *int32
    // Indicates the type of return code.
    type_escaped *Win32LobAppReturnCodeType
}
// NewWin32LobAppReturnCode instantiates a new win32LobAppReturnCode and sets the default values.
func NewWin32LobAppReturnCode()(*Win32LobAppReturnCode) {
    m := &Win32LobAppReturnCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWin32LobAppReturnCodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppReturnCodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppReturnCode(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Win32LobAppReturnCode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppReturnCode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["returnCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnCode(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppReturnCodeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*Win32LobAppReturnCodeType))
        }
        return nil
    }
    return res
}
// GetReturnCode gets the returnCode property value. Return code.
func (m *Win32LobAppReturnCode) GetReturnCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.returnCode
    }
}
// GetType gets the type property value. Indicates the type of return code.
func (m *Win32LobAppReturnCode) GetType()(*Win32LobAppReturnCodeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *Win32LobAppReturnCode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("returnCode", m.GetReturnCode())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Win32LobAppReturnCode) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReturnCode sets the returnCode property value. Return code.
func (m *Win32LobAppReturnCode) SetReturnCode(value *int32)() {
    if m != nil {
        m.returnCode = value
    }
}
// SetType sets the type property value. Indicates the type of return code.
func (m *Win32LobAppReturnCode) SetType(value *Win32LobAppReturnCodeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
