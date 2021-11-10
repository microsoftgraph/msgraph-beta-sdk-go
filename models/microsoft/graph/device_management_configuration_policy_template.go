package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationPolicyTemplate struct {
    Entity
    // Allow unmanaged setting templates
    allowUnmanagedSettings *bool;
    // Template base identifier
    baseId *string;
    // Template description
    description *string;
    // Template display name
    displayName *string;
    // Description of template version
    displayVersion *string;
    // Indicate current lifecycle state of template. Possible values are: invalid, draft, active, superseded, deprecated, retired.
    lifecycleState *DeviceManagementTemplateLifecycleState;
    // Platforms for this template. Possible values are: none, android, iOS, macOS, windows10X, windows10.
    platforms *DeviceManagementConfigurationPlatforms;
    // Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
    settingTemplateCount *int32;
    // Setting templates
    settingTemplates []DeviceManagementConfigurationSettingTemplate;
    // Technologies for this template. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
    technologies *DeviceManagementConfigurationTechnologies;
    // TemplateFamily for this template. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
    templateFamily *DeviceManagementConfigurationTemplateFamily;
    // Template version. Valid values 1 to 2147483647. This property is read-only.
    version *int32;
}
// Instantiates a new deviceManagementConfigurationPolicyTemplate and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplate()(*DeviceManagementConfigurationPolicyTemplate) {
    m := &DeviceManagementConfigurationPolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetAllowUnmanagedSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUnmanagedSettings
    }
}
// Gets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) GetBaseId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseId
    }
}
// Gets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// Gets the lifecycleState property value. Indicate current lifecycle state of template. Possible values are: invalid, draft, active, superseded, deprecated, retired.
func (m *DeviceManagementConfigurationPolicyTemplate) GetLifecycleState()(*DeviceManagementTemplateLifecycleState) {
    if m == nil {
        return nil
    } else {
        return m.lifecycleState
    }
}
// Gets the platforms property value. Platforms for this template. Possible values are: none, android, iOS, macOS, windows10X, windows10.
func (m *DeviceManagementConfigurationPolicyTemplate) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// Gets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.settingTemplateCount
    }
}
// Gets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplate) {
    if m == nil {
        return nil
    } else {
        return m.settingTemplates
    }
}
// Gets the technologies property value. Technologies for this template. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
func (m *DeviceManagementConfigurationPolicyTemplate) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    if m == nil {
        return nil
    } else {
        return m.technologies
    }
}
// Gets the templateFamily property value. TemplateFamily for this template. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
func (m *DeviceManagementConfigurationPolicyTemplate) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    if m == nil {
        return nil
    } else {
        return m.templateFamily
    }
}
// Gets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowUnmanagedSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUnmanagedSettings(val)
        }
        return nil
    }
    res["baseId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseId(val)
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
    res["displayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["lifecycleState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateLifecycleState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementTemplateLifecycleState)
            m.SetLifecycleState(&cast)
        }
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementConfigurationPlatforms)
            m.SetPlatforms(&cast)
        }
        return nil
    }
    res["settingTemplateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingTemplateCount(val)
        }
        return nil
    }
    res["settingTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingTemplate() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingTemplate, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationSettingTemplate))
            }
            m.SetSettingTemplates(res)
        }
        return nil
    }
    res["technologies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementConfigurationTechnologies)
            m.SetTechnologies(&cast)
        }
        return nil
    }
    res["templateFamily"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementConfigurationTemplateFamily)
            m.SetTemplateFamily(&cast)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DeviceManagementConfigurationPolicyTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationPolicyTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetLifecycleState().String()
        err = writer.WriteStringValue("lifecycleState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatforms() != nil {
        cast := m.GetPlatforms().String()
        err = writer.WriteStringValue("platforms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingTemplateCount", m.GetSettingTemplateCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingTemplates()))
        for i, v := range m.GetSettingTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settingTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnologies() != nil {
        cast := m.GetTechnologies().String()
        err = writer.WriteStringValue("technologies", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateFamily() != nil {
        cast := m.GetTemplateFamily().String()
        err = writer.WriteStringValue("templateFamily", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowUnmanagedSettings property value. Allow unmanaged setting templates
// Parameters:
//  - value : Value to set for the allowUnmanagedSettings property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetAllowUnmanagedSettings(value *bool)() {
    m.allowUnmanagedSettings = value
}
// Sets the baseId property value. Template base identifier
// Parameters:
//  - value : Value to set for the baseId property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetBaseId(value *string)() {
    m.baseId = value
}
// Sets the description property value. Template description
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Template display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the displayVersion property value. Description of template version
// Parameters:
//  - value : Value to set for the displayVersion property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
// Sets the lifecycleState property value. Indicate current lifecycle state of template. Possible values are: invalid, draft, active, superseded, deprecated, retired.
// Parameters:
//  - value : Value to set for the lifecycleState property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetLifecycleState(value *DeviceManagementTemplateLifecycleState)() {
    m.lifecycleState = value
}
// Sets the platforms property value. Platforms for this template. Possible values are: none, android, iOS, macOS, windows10X, windows10.
// Parameters:
//  - value : Value to set for the platforms property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// Sets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
// Parameters:
//  - value : Value to set for the settingTemplateCount property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplateCount(value *int32)() {
    m.settingTemplateCount = value
}
// Sets the settingTemplates property value. Setting templates
// Parameters:
//  - value : Value to set for the settingTemplates property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplate)() {
    m.settingTemplates = value
}
// Sets the technologies property value. Technologies for this template. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
// Parameters:
//  - value : Value to set for the technologies property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
// Sets the templateFamily property value. TemplateFamily for this template. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
// Parameters:
//  - value : Value to set for the templateFamily property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    m.templateFamily = value
}
// Sets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceManagementConfigurationPolicyTemplate) SetVersion(value *int32)() {
    m.version = value
}
