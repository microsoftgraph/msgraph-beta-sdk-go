package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetVersionItem 
type DocumentSetVersionItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The itemId property
    itemId *string
    // The title property
    title *string
    // The versionId property
    versionId *string
}
// NewDocumentSetVersionItem instantiates a new documentSetVersionItem and sets the default values.
func NewDocumentSetVersionItem()(*DocumentSetVersionItem) {
    m := &DocumentSetVersionItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDocumentSetVersionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentSetVersionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentSetVersionItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DocumentSetVersionItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentSetVersionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["versionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionId(val)
        }
        return nil
    }
    return res
}
// GetItemId gets the itemId property value. The itemId property
func (m *DocumentSetVersionItem) GetItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.itemId
    }
}
// GetTitle gets the title property value. The title property
func (m *DocumentSetVersionItem) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetVersionId gets the versionId property value. The versionId property
func (m *DocumentSetVersionItem) GetVersionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionId
    }
}
// Serialize serializes information the current object
func (m *DocumentSetVersionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionId", m.GetVersionId())
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
func (m *DocumentSetVersionItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetItemId sets the itemId property value. The itemId property
func (m *DocumentSetVersionItem) SetItemId(value *string)() {
    if m != nil {
        m.itemId = value
    }
}
// SetTitle sets the title property value. The title property
func (m *DocumentSetVersionItem) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetVersionId sets the versionId property value. The versionId property
func (m *DocumentSetVersionItem) SetVersionId(value *string)() {
    if m != nil {
        m.versionId = value
    }
}
