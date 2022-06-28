package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileAttachment 
type FileAttachment struct {
    Attachment
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The base64-encoded contents of the file.
    contentBytes []byte
    // The ID of the attachment in the Exchange store.
    contentId *string
    // Do not use this property as it is not supported.
    contentLocation *string
}
// NewFileAttachment instantiates a new FileAttachment and sets the default values.
func NewFileAttachment()(*FileAttachment) {
    m := &FileAttachment{
        Attachment: *NewAttachment(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFileAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileAttachment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileAttachment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentBytes gets the contentBytes property value. The base64-encoded contents of the file.
func (m *FileAttachment) GetContentBytes()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.contentBytes
    }
}
// GetContentId gets the contentId property value. The ID of the attachment in the Exchange store.
func (m *FileAttachment) GetContentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentId
    }
}
// GetContentLocation gets the contentLocation property value. Do not use this property as it is not supported.
func (m *FileAttachment) GetContentLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentLocation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Attachment.GetFieldDeserializers()
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
    res["contentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentId(val)
        }
        return nil
    }
    res["contentLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentLocation(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *FileAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Attachment.Serialize(writer)
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
        err = writer.WriteStringValue("contentId", m.GetContentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentLocation", m.GetContentLocation())
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
func (m *FileAttachment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentBytes sets the contentBytes property value. The base64-encoded contents of the file.
func (m *FileAttachment) SetContentBytes(value []byte)() {
    if m != nil {
        m.contentBytes = value
    }
}
// SetContentId sets the contentId property value. The ID of the attachment in the Exchange store.
func (m *FileAttachment) SetContentId(value *string)() {
    if m != nil {
        m.contentId = value
    }
}
// SetContentLocation sets the contentLocation property value. Do not use this property as it is not supported.
func (m *FileAttachment) SetContentLocation(value *string)() {
    if m != nil {
        m.contentLocation = value
    }
}
