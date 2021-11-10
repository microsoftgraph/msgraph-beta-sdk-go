package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceManagementTemplate and sets the default values.
func NewDeviceManagementTemplate()(*DeviceManagementTemplate) {
    m := &DeviceManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the categories property value. Collection of setting categories within the template
func (m *DeviceManagementTemplate) GetCategories()([]DeviceManagementTemplateSettingCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the description property value. The template's description
func (m *DeviceManagementTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The template's display name
func (m *DeviceManagementTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the intentCount property value. Number of Intents created from this template.
func (m *DeviceManagementTemplate) GetIntentCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.intentCount
    }
}
// Gets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
func (m *DeviceManagementTemplate) GetIsDeprecated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeprecated
    }
}
// Gets the migratableTo property value. Collection of templates this template can migrate to
func (m *DeviceManagementTemplate) GetMigratableTo()([]DeviceManagementTemplate) {
    if m == nil {
        return nil
    } else {
        return m.migratableTo
    }
}
// Gets the platformType property value. The template's platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, all.
func (m *DeviceManagementTemplate) GetPlatformType()(*PolicyPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// Gets the publishedDateTime property value. When the template was published
func (m *DeviceManagementTemplate) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.publishedDateTime
    }
}
// Gets the settings property value. Collection of all settings this template has
func (m *DeviceManagementTemplate) GetSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the templateSubtype property value. The template's subtype. Possible values are: none, firewall, diskEncryption, attackSurfaceReduction, endpointDetectionReponse, accountProtection, antivirus, firewallSharedAppList, firewallSharedIpList, firewallSharedPortlist.
func (m *DeviceManagementTemplate) GetTemplateSubtype()(*DeviceManagementTemplateSubtype) {
    if m == nil {
        return nil
    } else {
        return m.templateSubtype
    }
}
// Gets the templateType property value. The template's type. Possible values are: securityBaseline, specializedDevices, advancedThreatProtectionSecurityBaseline, deviceConfiguration, custom, securityTemplate, microsoftEdgeSecurityBaseline, microsoftOffice365ProPlusSecurityBaseline, deviceCompliance, deviceConfigurationForOffice365, cloudPC, firewallSharedSettings.
func (m *DeviceManagementTemplate) GetTemplateType()(*DeviceManagementTemplateType) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
}
// Gets the versionInfo property value. The template's version information
func (m *DeviceManagementTemplate) GetVersionInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionInfo
    }
}
// The deserialization information for the current model
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
            cast := val.(PolicyPlatformType)
            m.SetPlatformType(&cast)
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
            cast := val.(DeviceManagementTemplateSubtype)
            m.SetTemplateSubtype(&cast)
        }
        return nil
    }
    res["templateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementTemplateType)
            m.SetTemplateType(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
        cast := m.GetPlatformType().String()
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
    {
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
        cast := m.GetTemplateSubtype().String()
        err = writer.WriteStringValue("templateSubtype", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := m.GetTemplateType().String()
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
// Sets the categories property value. Collection of setting categories within the template
// Parameters:
//  - value : Value to set for the categories property.
func (m *DeviceManagementTemplate) SetCategories(value []DeviceManagementTemplateSettingCategory)() {
    m.categories = value
}
// Sets the description property value. The template's description
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementTemplate) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The template's display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the intentCount property value. Number of Intents created from this template.
// Parameters:
//  - value : Value to set for the intentCount property.
func (m *DeviceManagementTemplate) SetIntentCount(value *int32)() {
    m.intentCount = value
}
// Sets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
// Parameters:
//  - value : Value to set for the isDeprecated property.
func (m *DeviceManagementTemplate) SetIsDeprecated(value *bool)() {
    m.isDeprecated = value
}
// Sets the migratableTo property value. Collection of templates this template can migrate to
// Parameters:
//  - value : Value to set for the migratableTo property.
func (m *DeviceManagementTemplate) SetMigratableTo(value []DeviceManagementTemplate)() {
    m.migratableTo = value
}
// Sets the platformType property value. The template's platform. Possible values are: android, androidForWork, iOS, macOS, windowsPhone81, windows81AndLater, windows10AndLater, androidWorkProfile, windows10XProfile, all.
// Parameters:
//  - value : Value to set for the platformType property.
func (m *DeviceManagementTemplate) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// Sets the publishedDateTime property value. When the template was published
// Parameters:
//  - value : Value to set for the publishedDateTime property.
func (m *DeviceManagementTemplate) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// Sets the settings property value. Collection of all settings this template has
// Parameters:
//  - value : Value to set for the settings property.
func (m *DeviceManagementTemplate) SetSettings(value []DeviceManagementSettingInstance)() {
    m.settings = value
}
// Sets the templateSubtype property value. The template's subtype. Possible values are: none, firewall, diskEncryption, attackSurfaceReduction, endpointDetectionReponse, accountProtection, antivirus, firewallSharedAppList, firewallSharedIpList, firewallSharedPortlist.
// Parameters:
//  - value : Value to set for the templateSubtype property.
func (m *DeviceManagementTemplate) SetTemplateSubtype(value *DeviceManagementTemplateSubtype)() {
    m.templateSubtype = value
}
// Sets the templateType property value. The template's type. Possible values are: securityBaseline, specializedDevices, advancedThreatProtectionSecurityBaseline, deviceConfiguration, custom, securityTemplate, microsoftEdgeSecurityBaseline, microsoftOffice365ProPlusSecurityBaseline, deviceCompliance, deviceConfigurationForOffice365, cloudPC, firewallSharedSettings.
// Parameters:
//  - value : Value to set for the templateType property.
func (m *DeviceManagementTemplate) SetTemplateType(value *DeviceManagementTemplateType)() {
    m.templateType = value
}
// Sets the versionInfo property value. The template's version information
// Parameters:
//  - value : Value to set for the versionInfo property.
func (m *DeviceManagementTemplate) SetVersionInfo(value *string)() {
    m.versionInfo = value
}
