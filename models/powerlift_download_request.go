package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PowerliftDownloadRequest request used to download app diagnostic files.
type PowerliftDownloadRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The list of files to download
    files []string
    // The OdataType property
    odataType *string
    // The unique id for the request
    powerliftId *string
}
// NewPowerliftDownloadRequest instantiates a new powerliftDownloadRequest and sets the default values.
func NewPowerliftDownloadRequest()(*PowerliftDownloadRequest) {
    m := &PowerliftDownloadRequest{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.powerliftDownloadRequest";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePowerliftDownloadRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePowerliftDownloadRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPowerliftDownloadRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PowerliftDownloadRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PowerliftDownloadRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["files"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFiles(res)
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
    res["powerliftId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPowerliftId(val)
        }
        return nil
    }
    return res
}
// GetFiles gets the files property value. The list of files to download
func (m *PowerliftDownloadRequest) GetFiles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.files
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PowerliftDownloadRequest) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetPowerliftId gets the powerliftId property value. The unique id for the request
func (m *PowerliftDownloadRequest) GetPowerliftId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.powerliftId
    }
}
// Serialize serializes information the current object
func (m *PowerliftDownloadRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFiles() != nil {
        err := writer.WriteCollectionOfStringValues("files", m.GetFiles())
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
        err := writer.WriteStringValue("powerliftId", m.GetPowerliftId())
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
func (m *PowerliftDownloadRequest) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFiles sets the files property value. The list of files to download
func (m *PowerliftDownloadRequest) SetFiles(value []string)() {
    if m != nil {
        m.files = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PowerliftDownloadRequest) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetPowerliftId sets the powerliftId property value. The unique id for the request
func (m *PowerliftDownloadRequest) SetPowerliftId(value *string)() {
    if m != nil {
        m.powerliftId = value
    }
}
