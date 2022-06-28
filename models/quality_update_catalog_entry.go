package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// QualityUpdateCatalogEntry 
type QualityUpdateCatalogEntry struct {
    SoftwareUpdateCatalogEntry
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the content can be deployed as an expedited quality update. Read-only.
    isExpeditable *bool
}
// NewQualityUpdateCatalogEntry instantiates a new QualityUpdateCatalogEntry and sets the default values.
func NewQualityUpdateCatalogEntry()(*QualityUpdateCatalogEntry) {
    m := &QualityUpdateCatalogEntry{
        SoftwareUpdateCatalogEntry: *NewSoftwareUpdateCatalogEntry(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateQualityUpdateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQualityUpdateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQualityUpdateCatalogEntry(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *QualityUpdateCatalogEntry) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *QualityUpdateCatalogEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SoftwareUpdateCatalogEntry.GetFieldDeserializers()
    res["isExpeditable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExpeditable(val)
        }
        return nil
    }
    return res
}
// GetIsExpeditable gets the isExpeditable property value. Indicates whether the content can be deployed as an expedited quality update. Read-only.
func (m *QualityUpdateCatalogEntry) GetIsExpeditable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isExpeditable
    }
}
// Serialize serializes information the current object
func (m *QualityUpdateCatalogEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SoftwareUpdateCatalogEntry.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isExpeditable", m.GetIsExpeditable())
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
func (m *QualityUpdateCatalogEntry) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsExpeditable sets the isExpeditable property value. Indicates whether the content can be deployed as an expedited quality update. Read-only.
func (m *QualityUpdateCatalogEntry) SetIsExpeditable(value *bool)() {
    if m != nil {
        m.isExpeditable = value
    }
}
