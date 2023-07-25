package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// QualityUpdateCatalogEntry 
type QualityUpdateCatalogEntry struct {
    SoftwareUpdateCatalogEntry
}
// NewQualityUpdateCatalogEntry instantiates a new qualityUpdateCatalogEntry and sets the default values.
func NewQualityUpdateCatalogEntry()(*QualityUpdateCatalogEntry) {
    m := &QualityUpdateCatalogEntry{
        SoftwareUpdateCatalogEntry: *NewSoftwareUpdateCatalogEntry(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntry"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateQualityUpdateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQualityUpdateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQualityUpdateCatalogEntry(), nil
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
    res["qualityUpdateClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseQualityUpdateClassification)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateClassification(val.(*QualityUpdateClassification))
        }
        return nil
    }
    return res
}
// GetIsExpeditable gets the isExpeditable property value. Indicates whether the content can be deployed as an expedited quality update. Read-only.
func (m *QualityUpdateCatalogEntry) GetIsExpeditable()(*bool) {
    val, err := m.GetBackingStore().Get("isExpeditable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetQualityUpdateClassification gets the qualityUpdateClassification property value. The qualityUpdateClassification property
func (m *QualityUpdateCatalogEntry) GetQualityUpdateClassification()(*QualityUpdateClassification) {
    val, err := m.GetBackingStore().Get("qualityUpdateClassification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*QualityUpdateClassification)
    }
    return nil
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
    if m.GetQualityUpdateClassification() != nil {
        cast := (*m.GetQualityUpdateClassification()).String()
        err = writer.WriteStringValue("qualityUpdateClassification", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsExpeditable sets the isExpeditable property value. Indicates whether the content can be deployed as an expedited quality update. Read-only.
func (m *QualityUpdateCatalogEntry) SetIsExpeditable(value *bool)() {
    err := m.GetBackingStore().Set("isExpeditable", value)
    if err != nil {
        panic(err)
    }
}
// SetQualityUpdateClassification sets the qualityUpdateClassification property value. The qualityUpdateClassification property
func (m *QualityUpdateCatalogEntry) SetQualityUpdateClassification(value *QualityUpdateClassification)() {
    err := m.GetBackingStore().Set("qualityUpdateClassification", value)
    if err != nil {
        panic(err)
    }
}
// QualityUpdateCatalogEntryable 
type QualityUpdateCatalogEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SoftwareUpdateCatalogEntryable
    GetIsExpeditable()(*bool)
    GetQualityUpdateClassification()(*QualityUpdateClassification)
    SetIsExpeditable(value *bool)()
    SetQualityUpdateClassification(value *QualityUpdateClassification)()
}
