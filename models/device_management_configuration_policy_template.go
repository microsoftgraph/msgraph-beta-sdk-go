package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicyTemplate device Management Configuration Policy Template
type DeviceManagementConfigurationPolicyTemplate struct {
    Entity
}
// NewDeviceManagementConfigurationPolicyTemplate instantiates a new deviceManagementConfigurationPolicyTemplate and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplate()(*DeviceManagementConfigurationPolicyTemplate) {
    m := &DeviceManagementConfigurationPolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationPolicyTemplate(), nil
}
// GetAllowUnmanagedSettings gets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetAllowUnmanagedSettings()(*bool) {
    val, err := m.GetBackingStore().Get("allowUnmanagedSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBaseId gets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) GetBaseId()(*string) {
    val, err := m.GetBackingStore().Get("baseId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayVersion gets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayVersion()(*string) {
    val, err := m.GetBackingStore().Get("displayVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowUnmanagedSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUnmanagedSettings(val)
        }
        return nil
    }
    res["baseId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseId(val)
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
    res["displayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["lifecycleState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateLifecycleState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifecycleState(val.(*DeviceManagementTemplateLifecycleState))
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
    res["settingTemplateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingTemplateCount(val)
        }
        return nil
    }
    res["settingTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingTemplateable)
            }
            m.SetSettingTemplates(res)
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
    res["templateFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateFamily(val.(*DeviceManagementConfigurationTemplateFamily))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLifecycleState gets the lifecycleState property value. Describes current lifecycle state of a template
func (m *DeviceManagementConfigurationPolicyTemplate) GetLifecycleState()(*DeviceManagementTemplateLifecycleState) {
    val, err := m.GetBackingStore().Get("lifecycleState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementTemplateLifecycleState)
    }
    return nil
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicyTemplate) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    val, err := m.GetBackingStore().Get("platforms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationPlatforms)
    }
    return nil
}
// GetSettingTemplateCount gets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplateCount()(*int32) {
    val, err := m.GetBackingStore().Get("settingTemplateCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSettingTemplates gets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplateable) {
    val, err := m.GetBackingStore().Get("settingTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingTemplateable)
    }
    return nil
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicyTemplate) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    val, err := m.GetBackingStore().Get("technologies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationTechnologies)
    }
    return nil
}
// GetTemplateFamily gets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplate) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    val, err := m.GetBackingStore().Get("templateFamily")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationTemplateFamily)
    }
    return nil
}
// GetVersion gets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetVersion()(*int32) {
    val, err := m.GetBackingStore().Get("version")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicyTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowUnmanagedSettings", m.GetAllowUnmanagedSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("baseId", m.GetBaseId())
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
        err = writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetLifecycleState() != nil {
        cast := (*m.GetLifecycleState()).String()
        err = writer.WriteStringValue("lifecycleState", &cast)
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
    if m.GetSettingTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingTemplates()))
        for i, v := range m.GetSettingTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingTemplates", cast)
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
    if m.GetTemplateFamily() != nil {
        cast := (*m.GetTemplateFamily()).String()
        err = writer.WriteStringValue("templateFamily", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedSettings sets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) SetAllowUnmanagedSettings(value *bool)() {
    err := m.GetBackingStore().Set("allowUnmanagedSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetBaseId sets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) SetBaseId(value *string)() {
    err := m.GetBackingStore().Set("baseId", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayVersion sets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayVersion(value *string)() {
    err := m.GetBackingStore().Set("displayVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetLifecycleState sets the lifecycleState property value. Describes current lifecycle state of a template
func (m *DeviceManagementConfigurationPolicyTemplate) SetLifecycleState(value *DeviceManagementTemplateLifecycleState)() {
    err := m.GetBackingStore().Set("lifecycleState", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicyTemplate) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    err := m.GetBackingStore().Set("platforms", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingTemplateCount sets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplateCount(value *int32)() {
    err := m.GetBackingStore().Set("settingTemplateCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingTemplates sets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplateable)() {
    err := m.GetBackingStore().Set("settingTemplates", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicyTemplate) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    err := m.GetBackingStore().Set("technologies", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateFamily sets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplate) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    err := m.GetBackingStore().Set("templateFamily", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) SetVersion(value *int32)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationPolicyTemplateable 
type DeviceManagementConfigurationPolicyTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowUnmanagedSettings()(*bool)
    GetBaseId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDisplayVersion()(*string)
    GetLifecycleState()(*DeviceManagementTemplateLifecycleState)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetSettingTemplateCount()(*int32)
    GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplateable)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily)
    GetVersion()(*int32)
    SetAllowUnmanagedSettings(value *bool)()
    SetBaseId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDisplayVersion(value *string)()
    SetLifecycleState(value *DeviceManagementTemplateLifecycleState)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetSettingTemplateCount(value *int32)()
    SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplateable)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
    SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)()
    SetVersion(value *int32)()
}
