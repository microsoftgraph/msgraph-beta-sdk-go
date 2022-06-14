package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplate entity that represents a defined collection of device settings
type DeviceManagementTemplate struct {
    Entity
    // Collection of setting categories within the template
    categories []DeviceManagementTemplateSettingCategoryable
    // The template's description
    description *string
    // The template's display name
    displayName *string
    // Number of Intents created from this template.
    intentCount *int32
    // The template is deprecated or not. Intents cannot be created from a deprecated template.
    isDeprecated *bool
    // Collection of templates this template can migrate to
    migratableTo []DeviceManagementTemplateable
    // The template's platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, all.
    platformType *PolicyPlatformType
    // When the template was published
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Collection of all settings this template has
    settings []DeviceManagementSettingInstanceable
    // The template's subtype. Possible values are: none, firewall, diskEncryption, attackSurfaceReduction, endpointDetectionReponse, accountProtection, antivirus, firewallSharedAppList, firewallSharedIpList, firewallSharedPortlist.
    templateSubtype *DeviceManagementTemplateSubtype
    // The template's type. Possible values are: securityBaseline, specializedDevices, advancedThreatProtectionSecurityBaseline, deviceConfiguration, custom, securityTemplate, microsoftEdgeSecurityBaseline, microsoftOffice365ProPlusSecurityBaseline, deviceCompliance, deviceConfigurationForOffice365, cloudPC, firewallSharedSettings.
    templateType *DeviceManagementTemplateType
    // The template's version information
    versionInfo *string
}
// NewDeviceManagementTemplate instantiates a new deviceManagementTemplate and sets the default values.
func NewDeviceManagementTemplate()(*DeviceManagementTemplate) {
    m := &DeviceManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.securityBaselineTemplate":
                        return NewSecurityBaselineTemplate(), nil
                }
            }
        }
    }
    return NewDeviceManagementTemplate(), nil
}
// GetCategories gets the categories property value. Collection of setting categories within the template
func (m *DeviceManagementTemplate) GetCategories()([]DeviceManagementTemplateSettingCategoryable) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetDescription gets the description property value. The template's description
func (m *DeviceManagementTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The template's display name
func (m *DeviceManagementTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
                res[i] = v.(DeviceManagementTemplateSettingCategoryable)
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
                res[i] = v.(DeviceManagementTemplateable)
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
                res[i] = v.(DeviceManagementSettingInstanceable)
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
func (m *DeviceManagementTemplate) GetIntentCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.intentCount
    }
}
// GetIsDeprecated gets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
func (m *DeviceManagementTemplate) GetIsDeprecated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeprecated
    }
}
// GetMigratableTo gets the migratableTo property value. Collection of templates this template can migrate to
func (m *DeviceManagementTemplate) GetMigratableTo()([]DeviceManagementTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.migratableTo
    }
}
// GetPlatformType gets the platformType property value. The template's platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, all.
func (m *DeviceManagementTemplate) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// GetPublishedDateTime gets the publishedDateTime property value. When the template was published
func (m *DeviceManagementTemplate) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.publishedDateTime
    }
}
// GetSettings gets the settings property value. Collection of all settings this template has
func (m *DeviceManagementTemplate) GetSettings()([]DeviceManagementSettingInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetTemplateSubtype gets the templateSubtype property value. The template's subtype. Possible values are: none, firewall, diskEncryption, attackSurfaceReduction, endpointDetectionReponse, accountProtection, antivirus, firewallSharedAppList, firewallSharedIpList, firewallSharedPortlist.
func (m *DeviceManagementTemplate) GetTemplateSubtype()(*DeviceManagementTemplateSubtype) {
    if m == nil {
        return nil
    } else {
        return m.templateSubtype
    }
}
// GetTemplateType gets the templateType property value. The template's type. Possible values are: securityBaseline, specializedDevices, advancedThreatProtectionSecurityBaseline, deviceConfiguration, custom, securityTemplate, microsoftEdgeSecurityBaseline, microsoftOffice365ProPlusSecurityBaseline, deviceCompliance, deviceConfigurationForOffice365, cloudPC, firewallSharedSettings.
func (m *DeviceManagementTemplate) GetTemplateType()(*DeviceManagementTemplateType) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
}
// GetVersionInfo gets the versionInfo property value. The template's version information
func (m *DeviceManagementTemplate) GetVersionInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionInfo
    }
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m != nil {
        m.categories = value
    }
}
// SetDescription sets the description property value. The template's description
func (m *DeviceManagementTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The template's display name
func (m *DeviceManagementTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIntentCount sets the intentCount property value. Number of Intents created from this template.
func (m *DeviceManagementTemplate) SetIntentCount(value *int32)() {
    if m != nil {
        m.intentCount = value
    }
}
// SetIsDeprecated sets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
func (m *DeviceManagementTemplate) SetIsDeprecated(value *bool)() {
    if m != nil {
        m.isDeprecated = value
    }
}
// SetMigratableTo sets the migratableTo property value. Collection of templates this template can migrate to
func (m *DeviceManagementTemplate) SetMigratableTo(value []DeviceManagementTemplateable)() {
    if m != nil {
        m.migratableTo = value
    }
}
// SetPlatformType sets the platformType property value. The template's platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, all.
func (m *DeviceManagementTemplate) SetPlatformType(value *PolicyPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
// SetPublishedDateTime sets the publishedDateTime property value. When the template was published
func (m *DeviceManagementTemplate) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.publishedDateTime = value
    }
}
// SetSettings sets the settings property value. Collection of all settings this template has
func (m *DeviceManagementTemplate) SetSettings(value []DeviceManagementSettingInstanceable)() {
    if m != nil {
        m.settings = value
    }
}
// SetTemplateSubtype sets the templateSubtype property value. The template's subtype. Possible values are: none, firewall, diskEncryption, attackSurfaceReduction, endpointDetectionReponse, accountProtection, antivirus, firewallSharedAppList, firewallSharedIpList, firewallSharedPortlist.
func (m *DeviceManagementTemplate) SetTemplateSubtype(value *DeviceManagementTemplateSubtype)() {
    if m != nil {
        m.templateSubtype = value
    }
}
// SetTemplateType sets the templateType property value. The template's type. Possible values are: securityBaseline, specializedDevices, advancedThreatProtectionSecurityBaseline, deviceConfiguration, custom, securityTemplate, microsoftEdgeSecurityBaseline, microsoftOffice365ProPlusSecurityBaseline, deviceCompliance, deviceConfigurationForOffice365, cloudPC, firewallSharedSettings.
func (m *DeviceManagementTemplate) SetTemplateType(value *DeviceManagementTemplateType)() {
    if m != nil {
        m.templateType = value
    }
}
// SetVersionInfo sets the versionInfo property value. The template's version information
func (m *DeviceManagementTemplate) SetVersionInfo(value *string)() {
    if m != nil {
        m.versionInfo = value
    }
}
