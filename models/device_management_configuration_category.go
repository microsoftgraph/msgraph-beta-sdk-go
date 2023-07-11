package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationCategory device Management Configuration Policy
type DeviceManagementConfigurationCategory struct {
    Entity
}
// NewDeviceManagementConfigurationCategory instantiates a new deviceManagementConfigurationCategory and sets the default values.
func NewDeviceManagementConfigurationCategory()(*DeviceManagementConfigurationCategory) {
    m := &DeviceManagementConfigurationCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationCategory(), nil
}
// GetCategoryDescription gets the categoryDescription property value. Description of the category header in policy summary.
func (m *DeviceManagementConfigurationCategory) GetCategoryDescription()(*string) {
    val, err := m.GetBackingStore().Get("categoryDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetChildCategoryIds gets the childCategoryIds property value. List of child ids of the category.
func (m *DeviceManagementConfigurationCategory) GetChildCategoryIds()([]string) {
    val, err := m.GetBackingStore().Get("childCategoryIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the category. For example: Display
func (m *DeviceManagementConfigurationCategory) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the category. For example: Device Lock
func (m *DeviceManagementConfigurationCategory) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categoryDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategoryDescription(val)
        }
        return nil
    }
    res["childCategoryIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetChildCategoryIds(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["helpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpText(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["parentCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentCategoryId(val)
        }
        return nil
    }
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(*DeviceManagementConfigurationPlatforms))
        }
        return nil
    }
    res["rootCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCategoryId(val)
        }
        return nil
    }
    res["settingUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSettingUsage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingUsage(val.(*DeviceManagementConfigurationSettingUsage))
        }
        return nil
    }
    res["technologies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnologies(val.(*DeviceManagementConfigurationTechnologies))
        }
        return nil
    }
    return res
}
// GetHelpText gets the helpText property value. Help text of the category. Give more details of the category.
func (m *DeviceManagementConfigurationCategory) GetHelpText()(*string) {
    val, err := m.GetBackingStore().Get("helpText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetName gets the name property value. Name of the item
func (m *DeviceManagementConfigurationCategory) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationCategory) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParentCategoryId gets the parentCategoryId property value. Direct parent id of the category. If the category is the root, the parent id is same as its id.
func (m *DeviceManagementConfigurationCategory) GetParentCategoryId()(*string) {
    val, err := m.GetBackingStore().Get("parentCategoryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationCategory) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    val, err := m.GetBackingStore().Get("platforms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationPlatforms)
    }
    return nil
}
// GetRootCategoryId gets the rootCategoryId property value. Root id of the category.
func (m *DeviceManagementConfigurationCategory) GetRootCategoryId()(*string) {
    val, err := m.GetBackingStore().Get("rootCategoryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingUsage gets the settingUsage property value. Supported setting types
func (m *DeviceManagementConfigurationCategory) GetSettingUsage()(*DeviceManagementConfigurationSettingUsage) {
    val, err := m.GetBackingStore().Get("settingUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationSettingUsage)
    }
    return nil
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationCategory) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    val, err := m.GetBackingStore().Get("technologies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationTechnologies)
    }
    return nil
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetCategoryDescription sets the categoryDescription property value. Description of the category header in policy summary.
func (m *DeviceManagementConfigurationCategory) SetCategoryDescription(value *string)() {
    err := m.GetBackingStore().Set("categoryDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetChildCategoryIds sets the childCategoryIds property value. List of child ids of the category.
func (m *DeviceManagementConfigurationCategory) SetChildCategoryIds(value []string)() {
    err := m.GetBackingStore().Set("childCategoryIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the category. For example: Display
func (m *DeviceManagementConfigurationCategory) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the category. For example: Device Lock
func (m *DeviceManagementConfigurationCategory) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHelpText sets the helpText property value. Help text of the category. Give more details of the category.
func (m *DeviceManagementConfigurationCategory) SetHelpText(value *string)() {
    err := m.GetBackingStore().Set("helpText", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name of the item
func (m *DeviceManagementConfigurationCategory) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationCategory) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetParentCategoryId sets the parentCategoryId property value. Direct parent id of the category. If the category is the root, the parent id is same as its id.
func (m *DeviceManagementConfigurationCategory) SetParentCategoryId(value *string)() {
    err := m.GetBackingStore().Set("parentCategoryId", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationCategory) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    err := m.GetBackingStore().Set("platforms", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCategoryId sets the rootCategoryId property value. Root id of the category.
func (m *DeviceManagementConfigurationCategory) SetRootCategoryId(value *string)() {
    err := m.GetBackingStore().Set("rootCategoryId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingUsage sets the settingUsage property value. Supported setting types
func (m *DeviceManagementConfigurationCategory) SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)() {
    err := m.GetBackingStore().Set("settingUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationCategory) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    err := m.GetBackingStore().Set("technologies", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationCategoryable 
type DeviceManagementConfigurationCategoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategoryDescription()(*string)
    GetChildCategoryIds()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHelpText()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetParentCategoryId()(*string)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetRootCategoryId()(*string)
    GetSettingUsage()(*DeviceManagementConfigurationSettingUsage)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    SetCategoryDescription(value *string)()
    SetChildCategoryIds(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHelpText(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetParentCategoryId(value *string)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetRootCategoryId(value *string)()
    SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
}
