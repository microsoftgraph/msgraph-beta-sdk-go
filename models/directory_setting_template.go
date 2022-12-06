package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectorySettingTemplate 
type DirectorySettingTemplate struct {
    DirectoryObject
    // Description of the template. Read-only.
    description *string
    // Display name of the template. Read-only.
    displayName *string
    // Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
    values []SettingTemplateValueable
}
// NewDirectorySettingTemplate instantiates a new DirectorySettingTemplate and sets the default values.
func NewDirectorySettingTemplate()(*DirectorySettingTemplate) {
    m := &DirectorySettingTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.directorySettingTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDirectorySettingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectorySettingTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectorySettingTemplate(), nil
}
// GetDescription gets the description property value. Description of the template. Read-only.
func (m *DirectorySettingTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the template. Read-only.
func (m *DirectorySettingTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySettingTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["values"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSettingTemplateValueFromDiscriminatorValue , m.SetValues)
    return res
}
// GetValues gets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
func (m *DirectorySettingTemplate) GetValues()([]SettingTemplateValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *DirectorySettingTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValues())
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of the template. Read-only.
func (m *DirectorySettingTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the template. Read-only.
func (m *DirectorySettingTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetValues sets the values property value. Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template.  Read-only.
func (m *DirectorySettingTemplate) SetValues(value []SettingTemplateValueable)() {
    m.values = value
}
