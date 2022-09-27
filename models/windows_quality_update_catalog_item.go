package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsQualityUpdateCatalogItem 
type WindowsQualityUpdateCatalogItem struct {
    WindowsUpdateCatalogItem
    // Windows quality update classification
    classification *WindowsQualityUpdateClassification
    // Flag indicating if update qualifies for expedite
    isExpeditable *bool
    // Knowledge base article id
    kbArticleId *string
}
// NewWindowsQualityUpdateCatalogItem instantiates a new WindowsQualityUpdateCatalogItem and sets the default values.
func NewWindowsQualityUpdateCatalogItem()(*WindowsQualityUpdateCatalogItem) {
    m := &WindowsQualityUpdateCatalogItem{
        WindowsUpdateCatalogItem: *NewWindowsUpdateCatalogItem(),
    }
    odataTypeValue := "#microsoft.graph.windowsQualityUpdateCatalogItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsQualityUpdateCatalogItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsQualityUpdateCatalogItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateCatalogItem(), nil
}
// GetClassification gets the classification property value. Windows quality update classification
func (m *WindowsQualityUpdateCatalogItem) GetClassification()(*WindowsQualityUpdateClassification) {
    return m.classification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsQualityUpdateCatalogItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateCatalogItem.GetFieldDeserializers()
    res["classification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWindowsQualityUpdateClassification , m.SetClassification)
    res["isExpeditable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsExpeditable)
    res["kbArticleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetKbArticleId)
    return res
}
// GetIsExpeditable gets the isExpeditable property value. Flag indicating if update qualifies for expedite
func (m *WindowsQualityUpdateCatalogItem) GetIsExpeditable()(*bool) {
    return m.isExpeditable
}
// GetKbArticleId gets the kbArticleId property value. Knowledge base article id
func (m *WindowsQualityUpdateCatalogItem) GetKbArticleId()(*string) {
    return m.kbArticleId
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateCatalogItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateCatalogItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isExpeditable", m.GetIsExpeditable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kbArticleId", m.GetKbArticleId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassification sets the classification property value. Windows quality update classification
func (m *WindowsQualityUpdateCatalogItem) SetClassification(value *WindowsQualityUpdateClassification)() {
    m.classification = value
}
// SetIsExpeditable sets the isExpeditable property value. Flag indicating if update qualifies for expedite
func (m *WindowsQualityUpdateCatalogItem) SetIsExpeditable(value *bool)() {
    m.isExpeditable = value
}
// SetKbArticleId sets the kbArticleId property value. Knowledge base article id
func (m *WindowsQualityUpdateCatalogItem) SetKbArticleId(value *string)() {
    m.kbArticleId = value
}
