package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationPolicyTemplate 
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
    // Platforms for this template. Possible values are: none, android, iOS, macOS, windows10X, windows10, linux, unknownFutureValue.
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
// NewDeviceManagementConfigurationPolicyTemplate instantiates a new deviceManagementConfigurationPolicyTemplate and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplate()(*DeviceManagementConfigurationPolicyTemplate) {
    m := &DeviceManagementConfigurationPolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowUnmanagedSettings gets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetAllowUnmanagedSettings()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowUnmanagedSettings
    }
}
// GetBaseId gets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) GetBaseId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseId
    }
}
// GetDescription gets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDisplayVersion gets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayVersion
    }
}
// GetLifecycleState gets the lifecycleState property value. Indicate current lifecycle state of template. Possible values are: invalid, draft, active, superseded, deprecated, retired.
func (m *DeviceManagementConfigurationPolicyTemplate) GetLifecycleState()(*DeviceManagementTemplateLifecycleState) {
    if m == nil {
        return nil
    } else {
        return m.lifecycleState
    }
}
// GetPlatforms gets the platforms property value. Platforms for this template. Possible values are: none, android, iOS, macOS, windows10X, windows10, linux, unknownFutureValue.
func (m *DeviceManagementConfigurationPolicyTemplate) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// GetSettingTemplateCount gets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.settingTemplateCount
    }
}
// GetSettingTemplates gets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplate) {
    if m == nil {
        return nil
    } else {
        return m.settingTemplates
    }
}
// GetTechnologies gets the technologies property value. Technologies for this template. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
func (m *DeviceManagementConfigurationPolicyTemplate) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    if m == nil {
        return nil
    } else {
        return m.technologies
    }
}
// GetTemplateFamily gets the templateFamily property value. TemplateFamily for this template. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
func (m *DeviceManagementConfigurationPolicyTemplate) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    if m == nil {
        return nil
    } else {
        return m.templateFamily
    }
}
// GetVersion gets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetLifecycleState(val.(*DeviceManagementTemplateLifecycleState))
        }
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(*DeviceManagementConfigurationPlatforms))
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
            m.SetTechnologies(val.(*DeviceManagementConfigurationTechnologies))
        }
        return nil
    }
    res["templateFamily"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateFamily(val.(*DeviceManagementConfigurationTemplateFamily))
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
// Serialize serializes information the current object
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
    {
        err = writer.WriteInt32Value("settingTemplateCount", m.GetSettingTemplateCount())
        if err != nil {
            return err
        }
    }
    if m.GetSettingTemplates() != nil {
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
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedSettings sets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) SetAllowUnmanagedSettings(value *bool)() {
    if m != nil {
        m.allowUnmanagedSettings = value
    }
}
// SetBaseId sets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) SetBaseId(value *string)() {
    if m != nil {
        m.baseId = value
    }
}
// SetDescription sets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDisplayVersion sets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayVersion(value *string)() {
    if m != nil {
        m.displayVersion = value
    }
}
// SetLifecycleState sets the lifecycleState property value. Indicate current lifecycle state of template. Possible values are: invalid, draft, active, superseded, deprecated, retired.
func (m *DeviceManagementConfigurationPolicyTemplate) SetLifecycleState(value *DeviceManagementTemplateLifecycleState)() {
    if m != nil {
        m.lifecycleState = value
    }
}
// SetPlatforms sets the platforms property value. Platforms for this template. Possible values are: none, android, iOS, macOS, windows10X, windows10, linux, unknownFutureValue.
func (m *DeviceManagementConfigurationPolicyTemplate) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    if m != nil {
        m.platforms = value
    }
}
// SetSettingTemplateCount sets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplateCount(value *int32)() {
    if m != nil {
        m.settingTemplateCount = value
    }
}
// SetSettingTemplates sets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplate)() {
    if m != nil {
        m.settingTemplates = value
    }
}
// SetTechnologies sets the technologies property value. Technologies for this template. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
func (m *DeviceManagementConfigurationPolicyTemplate) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    if m != nil {
        m.technologies = value
    }
}
// SetTemplateFamily sets the templateFamily property value. TemplateFamily for this template. Possible values are: none, endpointSecurityAntivirus, endpointSecurityDiskEncryption, endpointSecurityFirewall, endpointSecurityEndpointDetectionAndResponse, endpointSecurityAttackSurfaceReduction, endpointSecurityAccountProtection, endpointSecurityApplicationControl, baseline.
func (m *DeviceManagementConfigurationPolicyTemplate) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    if m != nil {
        m.templateFamily = value
    }
}
// SetVersion sets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
