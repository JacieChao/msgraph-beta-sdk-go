package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody 
type ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The AttachmentItem property
    attachmentItem ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentItemable
}
// NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody()(*ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttachmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentItem(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentItemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentItem", m.GetAttachmentItem())
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
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *ItemCalendarEventsItemInstancesItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttachmentItemable)() {
    m.attachmentItem = value
}