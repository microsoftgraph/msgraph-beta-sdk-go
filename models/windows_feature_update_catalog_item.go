package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFeatureUpdateCatalogItem 
type WindowsFeatureUpdateCatalogItem struct {
    WindowsUpdateCatalogItem
    // The feature update version
    version *string
}
// NewWindowsFeatureUpdateCatalogItem instantiates a new WindowsFeatureUpdateCatalogItem and sets the default values.
func NewWindowsFeatureUpdateCatalogItem()(*WindowsFeatureUpdateCatalogItem) {
    m := &WindowsFeatureUpdateCatalogItem{
        WindowsUpdateCatalogItem: *NewWindowsUpdateCatalogItem(),
    }
    odataTypeValue := "#microsoft.graph.windowsFeatureUpdateCatalogItem";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsFeatureUpdateCatalogItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFeatureUpdateCatalogItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFeatureUpdateCatalogItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsFeatureUpdateCatalogItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsUpdateCatalogItem.GetFieldDeserializers()
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersion)
    return res
}
// GetVersion gets the version property value. The feature update version
func (m *WindowsFeatureUpdateCatalogItem) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *WindowsFeatureUpdateCatalogItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsUpdateCatalogItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVersion sets the version property value. The feature update version
func (m *WindowsFeatureUpdateCatalogItem) SetVersion(value *string)() {
    m.version = value
}
