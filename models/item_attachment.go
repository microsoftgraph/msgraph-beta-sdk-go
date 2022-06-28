package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAttachment 
type ItemAttachment struct {
    Attachment
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attached contact, message or event. Navigation property.
    item OutlookItemable
}
// NewItemAttachment instantiates a new ItemAttachment and sets the default values.
func NewItemAttachment()(*ItemAttachment) {
    m := &ItemAttachment{
        Attachment: *NewAttachment(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAttachment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAttachment) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Attachment.GetFieldDeserializers()
    res["item"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(OutlookItemable))
        }
        return nil
    }
    return res
}
// GetItem gets the item property value. The attached contact, message or event. Navigation property.
func (m *ItemAttachment) GetItem()(OutlookItemable) {
    if m == nil {
        return nil
    } else {
        return m.item
    }
}
// Serialize serializes information the current object
func (m *ItemAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Attachment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("item", m.GetItem())
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
func (m *ItemAttachment) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetItem sets the item property value. The attached contact, message or event. Navigation property.
func (m *ItemAttachment) SetItem(value OutlookItemable)() {
    if m != nil {
        m.item = value
    }
}
