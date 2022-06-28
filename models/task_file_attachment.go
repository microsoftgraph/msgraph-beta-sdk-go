package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TaskFileAttachment 
type TaskFileAttachment struct {
    AttachmentBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentBytes property
    contentBytes []byte
}
// NewTaskFileAttachment instantiates a new TaskFileAttachment and sets the default values.
func NewTaskFileAttachment()(*TaskFileAttachment) {
    m := &TaskFileAttachment{
        AttachmentBase: *NewAttachmentBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTaskFileAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskFileAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskFileAttachment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TaskFileAttachment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentBytes gets the contentBytes property value. The contentBytes property
func (m *TaskFileAttachment) GetContentBytes()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.contentBytes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskFileAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AttachmentBase.GetFieldDeserializers()
    res["contentBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentBytes(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TaskFileAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AttachmentBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("contentBytes", m.GetContentBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TaskFileAttachment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentBytes sets the contentBytes property value. The contentBytes property
func (m *TaskFileAttachment) SetContentBytes(value []byte)() {
    if m != nil {
        m.contentBytes = value
    }
}
