package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody provides operations to call the createUploadSession method.
type GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AttachmentItem property
    attachmentItem AttachmentItemable
}
// NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody instantiates a new GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody and sets the default values.
func NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody()(*GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) {
    m := &GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttachmentItem gets the attachmentItem property value. The AttachmentItem property
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) GetAttachmentItem()(AttachmentItemable) {
    return m.attachmentItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttachmentItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentItem(val.(AttachmentItemable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttachmentItem sets the attachmentItem property value. The AttachmentItem property
func (m *GroupsItemCalendarCalendarViewItemExceptionOccurrencesItemAttachmentsCreateUploadSessionPostRequestBody) SetAttachmentItem(value AttachmentItemable)() {
    m.attachmentItem = value
}
