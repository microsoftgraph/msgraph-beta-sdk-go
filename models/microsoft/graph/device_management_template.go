package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTemplate 
type DeviceManagementTemplate struct {
    Entity
    // Collection of setting categories within the template
    categories []DeviceManagementTemplateSettingCategory;
    // The template's description
    description *string;
    // The template's display name
    displayName *string;
    // Number of Intents created from this template.
    intentCount *int32;
    // The template is deprecated or not. Intents cannot be created from a deprecated template.
    isDeprecated *bool;
    // Collection of templates this template can migrate to
    migratableTo []DeviceManagementTemplate;
    // The template's platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, all.
    platformType *PolicyPlatformType;
    // When the template was published
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Collection of all settings this template has
    settings []DeviceManagementSettingInstance;
    // The template's subtype. Possible values are: none, firewall, diskEncryption, attackSurfaceReduction, endpointDetectionReponse, accountProtection, antivirus, firewallSharedAppList, firewallSharedIpList, firewallSharedPortlist.
    templateSubtype *DeviceManagementTemplateSubtype;
    // The template's type. Possible values are: securityBaseline, specializedDevices, advancedThreatProtectionSecurityBaseline, deviceConfiguration, custom, securityTemplate, microsoftEdgeSecurityBaseline, microsoftOffice365ProPlusSecurityBaseline, deviceCompliance, deviceConfigurationForOffice365, cloudPC, firewallSharedSettings.
    templateType *DeviceManagementTemplateType;
    // The template's version information
    versionInfo *string;
}
// NewDeviceManagementTemplate instantiates a new deviceManagementTemplate and sets the default values.
func NewDeviceManagementTemplate()(*DeviceManagementTemplate) {
    m := &DeviceManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategories gets the categories property value. Collection of setting categories within the template
func (m *DeviceManagementTemplate) GetCategories()([]DeviceManagementTemplateSettingCategory) {
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
func (m *DeviceManagementTemplate) GetMigratableTo()([]DeviceManagementTemplate) {
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
func (m *DeviceManagementTemplate) GetSettings()([]DeviceManagementSettingInstance) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTemplateSettingCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateSettingCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementTemplateSettingCategory))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["intentCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentCount(val)
        }
        return nil
    }
    res["isDeprecated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeprecated(val)
        }
        return nil
    }
    res["migratableTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplate, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementTemplate))
            }
            m.SetMigratableTo(res)
        }
        return nil
    }
    res["platformType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["publishedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementSettingInstance))
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateSubtype"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateSubtype)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateSubtype(val.(*DeviceManagementTemplateSubtype))
        }
        return nil
    }
    res["templateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(*DeviceManagementTemplateType))
        }
        return nil
    }
    res["versionInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DeviceManagementTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMigratableTo()))
        for i, v := range m.GetMigratableTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *DeviceManagementTemplate) SetCategories(value []DeviceManagementTemplateSettingCategory)() {
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
func (m *DeviceManagementTemplate) SetMigratableTo(value []DeviceManagementTemplate)() {
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
func (m *DeviceManagementTemplate) SetSettings(value []DeviceManagementSettingInstance)() {
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
