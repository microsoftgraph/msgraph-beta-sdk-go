package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplate entity that represents a defined collection of device settings
type DeviceManagementTemplate struct {
    Entity
}
// NewDeviceManagementTemplate instantiates a new DeviceManagementTemplate and sets the default values.
func NewDeviceManagementTemplate()(*DeviceManagementTemplate) {
    m := &DeviceManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.securityBaselineTemplate":
                        return NewSecurityBaselineTemplate(), nil
                }
            }
        }
    }
    return NewDeviceManagementTemplate(), nil
}
// GetCategories gets the categories property value. Collection of setting categories within the template
// returns a []DeviceManagementTemplateSettingCategoryable when successful
func (m *DeviceManagementTemplate) GetCategories()([]DeviceManagementTemplateSettingCategoryable) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementTemplateSettingCategoryable)
    }
    return nil
}
// GetDescription gets the description property value. The template's description
// returns a *string when successful
func (m *DeviceManagementTemplate) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The template's display name
// returns a *string when successful
func (m *DeviceManagementTemplate) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateSettingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateSettingCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTemplateSettingCategoryable)
                }
            }
            m.SetCategories(res)
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
    res["intentCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentCount(val)
        }
        return nil
    }
    res["isDeprecated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeprecated(val)
        }
        return nil
    }
    res["migratableTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementTemplateable)
                }
            }
            m.SetMigratableTo(res)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["publishedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementSettingInstanceable)
                }
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateSubtype"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateSubtype)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateSubtype(val.(*DeviceManagementTemplateSubtype))
        }
        return nil
    }
    res["templateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(*DeviceManagementTemplateType))
        }
        return nil
    }
    res["versionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInfo(val)
        }
        return nil
    }
    return res
}
// GetIntentCount gets the intentCount property value. Number of Intents created from this template.
// returns a *int32 when successful
func (m *DeviceManagementTemplate) GetIntentCount()(*int32) {
    val, err := m.GetBackingStore().Get("intentCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetIsDeprecated gets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
// returns a *bool when successful
func (m *DeviceManagementTemplate) GetIsDeprecated()(*bool) {
    val, err := m.GetBackingStore().Get("isDeprecated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMigratableTo gets the migratableTo property value. Collection of templates this template can migrate to
// returns a []DeviceManagementTemplateable when successful
func (m *DeviceManagementTemplate) GetMigratableTo()([]DeviceManagementTemplateable) {
    val, err := m.GetBackingStore().Get("migratableTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementTemplateable)
    }
    return nil
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
// returns a *PolicyPlatformType when successful
func (m *DeviceManagementTemplate) GetPlatformType()(*PolicyPlatformType) {
    val, err := m.GetBackingStore().Get("platformType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyPlatformType)
    }
    return nil
}
// GetPublishedDateTime gets the publishedDateTime property value. When the template was published
// returns a *Time when successful
func (m *DeviceManagementTemplate) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("publishedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSettings gets the settings property value. Collection of all settings this template has
// returns a []DeviceManagementSettingInstanceable when successful
func (m *DeviceManagementTemplate) GetSettings()([]DeviceManagementSettingInstanceable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementSettingInstanceable)
    }
    return nil
}
// GetTemplateSubtype gets the templateSubtype property value. Template subtype
// returns a *DeviceManagementTemplateSubtype when successful
func (m *DeviceManagementTemplate) GetTemplateSubtype()(*DeviceManagementTemplateSubtype) {
    val, err := m.GetBackingStore().Get("templateSubtype")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementTemplateSubtype)
    }
    return nil
}
// GetTemplateType gets the templateType property value. Template type
// returns a *DeviceManagementTemplateType when successful
func (m *DeviceManagementTemplate) GetTemplateType()(*DeviceManagementTemplateType) {
    val, err := m.GetBackingStore().Get("templateType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementTemplateType)
    }
    return nil
}
// GetVersionInfo gets the versionInfo property value. The template's version information
// returns a *string when successful
func (m *DeviceManagementTemplate) GetVersionInfo()(*string) {
    val, err := m.GetBackingStore().Get("versionInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
        err = writer.WriteInt32Value("intentCount", m.GetIntentCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeprecated", m.GetIsDeprecated())
        if err != nil {
            return err
        }
    }
    if m.GetMigratableTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMigratableTo()))
        for i, v := range m.GetMigratableTo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("migratableTo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateSubtype() != nil {
        cast := (*m.GetTemplateSubtype()).String()
        err = writer.WriteStringValue("templateSubtype", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := (*m.GetTemplateType()).String()
        err = writer.WriteStringValue("templateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionInfo", m.GetVersionInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategories sets the categories property value. Collection of setting categories within the template
func (m *DeviceManagementTemplate) SetCategories(value []DeviceManagementTemplateSettingCategoryable)() {
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The template's description
func (m *DeviceManagementTemplate) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The template's display name
func (m *DeviceManagementTemplate) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIntentCount sets the intentCount property value. Number of Intents created from this template.
func (m *DeviceManagementTemplate) SetIntentCount(value *int32)() {
    err := m.GetBackingStore().Set("intentCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDeprecated sets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
func (m *DeviceManagementTemplate) SetIsDeprecated(value *bool)() {
    err := m.GetBackingStore().Set("isDeprecated", value)
    if err != nil {
        panic(err)
    }
}
// SetMigratableTo sets the migratableTo property value. Collection of templates this template can migrate to
func (m *DeviceManagementTemplate) SetMigratableTo(value []DeviceManagementTemplateable)() {
    err := m.GetBackingStore().Set("migratableTo", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *DeviceManagementTemplate) SetPlatformType(value *PolicyPlatformType)() {
    err := m.GetBackingStore().Set("platformType", value)
    if err != nil {
        panic(err)
    }
}
// SetPublishedDateTime sets the publishedDateTime property value. When the template was published
func (m *DeviceManagementTemplate) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("publishedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Collection of all settings this template has
func (m *DeviceManagementTemplate) SetSettings(value []DeviceManagementSettingInstanceable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateSubtype sets the templateSubtype property value. Template subtype
func (m *DeviceManagementTemplate) SetTemplateSubtype(value *DeviceManagementTemplateSubtype)() {
    err := m.GetBackingStore().Set("templateSubtype", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateType sets the templateType property value. Template type
func (m *DeviceManagementTemplate) SetTemplateType(value *DeviceManagementTemplateType)() {
    err := m.GetBackingStore().Set("templateType", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionInfo sets the versionInfo property value. The template's version information
func (m *DeviceManagementTemplate) SetVersionInfo(value *string)() {
    err := m.GetBackingStore().Set("versionInfo", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]DeviceManagementTemplateSettingCategoryable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIntentCount()(*int32)
    GetIsDeprecated()(*bool)
    GetMigratableTo()([]DeviceManagementTemplateable)
    GetPlatformType()(*PolicyPlatformType)
    GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSettings()([]DeviceManagementSettingInstanceable)
    GetTemplateSubtype()(*DeviceManagementTemplateSubtype)
    GetTemplateType()(*DeviceManagementTemplateType)
    GetVersionInfo()(*string)
    SetCategories(value []DeviceManagementTemplateSettingCategoryable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIntentCount(value *int32)()
    SetIsDeprecated(value *bool)()
    SetMigratableTo(value []DeviceManagementTemplateable)()
    SetPlatformType(value *PolicyPlatformType)()
    SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSettings(value []DeviceManagementSettingInstanceable)()
    SetTemplateSubtype(value *DeviceManagementTemplateSubtype)()
    SetTemplateType(value *DeviceManagementTemplateType)()
    SetVersionInfo(value *string)()
}
