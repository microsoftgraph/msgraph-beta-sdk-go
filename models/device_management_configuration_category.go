package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationCategory device Management Configuration Policy
type DeviceManagementConfigurationCategory struct {
    Entity
    // Description of the category header
    categoryDescription *string
    // List of child ids of the category.
    childCategoryIds []string
    // Description of the item
    description *string
    // Display name of the item
    displayName *string
    // Help text of the item
    helpText *string
    // Name of the item
    name *string
    // Parent id of the category.
    parentCategoryId *string
    // Supported platform types.
    platforms *DeviceManagementConfigurationPlatforms
    // Root id of the category.
    rootCategoryId *string
    // Supported setting types
    settingUsage *DeviceManagementConfigurationSettingUsage
    // Describes which technology this setting can be deployed with
    technologies *DeviceManagementConfigurationTechnologies
}
// NewDeviceManagementConfigurationCategory instantiates a new deviceManagementConfigurationCategory and sets the default values.
func NewDeviceManagementConfigurationCategory()(*DeviceManagementConfigurationCategory) {
    m := &DeviceManagementConfigurationCategory{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationCategory";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationCategory(), nil
}
// GetCategoryDescription gets the categoryDescription property value. Description of the category header
func (m *DeviceManagementConfigurationCategory) GetCategoryDescription()(*string) {
    return m.categoryDescription
}
// GetChildCategoryIds gets the childCategoryIds property value. List of child ids of the category.
func (m *DeviceManagementConfigurationCategory) GetChildCategoryIds()([]string) {
    return m.childCategoryIds
}
// GetDescription gets the description property value. Description of the item
func (m *DeviceManagementConfigurationCategory) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the item
func (m *DeviceManagementConfigurationCategory) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryDescription"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCategoryDescription)
    res["childCategoryIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetChildCategoryIds)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["helpText"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHelpText)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["parentCategoryId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetParentCategoryId)
    res["platforms"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationPlatforms , m.SetPlatforms)
    res["rootCategoryId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRootCategoryId)
    res["settingUsage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationSettingUsage , m.SetSettingUsage)
    res["technologies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationTechnologies , m.SetTechnologies)
    return res
}
// GetHelpText gets the helpText property value. Help text of the item
func (m *DeviceManagementConfigurationCategory) GetHelpText()(*string) {
    return m.helpText
}
// GetName gets the name property value. Name of the item
func (m *DeviceManagementConfigurationCategory) GetName()(*string) {
    return m.name
}
// GetParentCategoryId gets the parentCategoryId property value. Parent id of the category.
func (m *DeviceManagementConfigurationCategory) GetParentCategoryId()(*string) {
    return m.parentCategoryId
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationCategory) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    return m.platforms
}
// GetRootCategoryId gets the rootCategoryId property value. Root id of the category.
func (m *DeviceManagementConfigurationCategory) GetRootCategoryId()(*string) {
    return m.rootCategoryId
}
// GetSettingUsage gets the settingUsage property value. Supported setting types
func (m *DeviceManagementConfigurationCategory) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    return m.settingUsage
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationCategory) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    return m.technologies
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("categoryDescription", m.GetCategoryDescription())
        if err != nil {
            return err
        }
    }
    if m.GetChildCategoryIds() != nil {
        err = writer.WriteCollectionOfStringValues("childCategoryIds", m.GetChildCategoryIds())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteStringValue("helpText", m.GetHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentCategoryId", m.GetParentCategoryId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatforms() != nil {
        cast := (*m.GetPlatforms()).String()
        err = writer.WriteStringValue("platforms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rootCategoryId", m.GetRootCategoryId())
        if err != nil {
            return err
        }
    }
    if m.GetSettingUsage() != nil {
        cast := (*m.GetSettingUsage()).String()
        err = writer.WriteStringValue("settingUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnologies() != nil {
        cast := (*m.GetTechnologies()).String()
        err = writer.WriteStringValue("technologies", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategoryDescription sets the categoryDescription property value. Description of the category header
func (m *DeviceManagementConfigurationCategory) SetCategoryDescription(value *string)() {
    m.categoryDescription = value
}
// SetChildCategoryIds sets the childCategoryIds property value. List of child ids of the category.
func (m *DeviceManagementConfigurationCategory) SetChildCategoryIds(value []string)() {
    m.childCategoryIds = value
}
// SetDescription sets the description property value. Description of the item
func (m *DeviceManagementConfigurationCategory) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the item
func (m *DeviceManagementConfigurationCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHelpText sets the helpText property value. Help text of the item
func (m *DeviceManagementConfigurationCategory) SetHelpText(value *string)() {
    m.helpText = value
}
// SetName sets the name property value. Name of the item
func (m *DeviceManagementConfigurationCategory) SetName(value *string)() {
    m.name = value
}
// SetParentCategoryId sets the parentCategoryId property value. Parent id of the category.
func (m *DeviceManagementConfigurationCategory) SetParentCategoryId(value *string)() {
    m.parentCategoryId = value
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationCategory) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// SetRootCategoryId sets the rootCategoryId property value. Root id of the category.
func (m *DeviceManagementConfigurationCategory) SetRootCategoryId(value *string)() {
    m.rootCategoryId = value
}
// SetSettingUsage sets the settingUsage property value. Supported setting types
func (m *DeviceManagementConfigurationCategory) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    m.settingUsage = value
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationCategory) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
