package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementCompliancePolicy struct {
    Entity
    // Policy assignments
    assignments []DeviceManagementConfigurationPolicyAssignment;
    // Policy creation date and time. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Policy creation source
    creationSource *string;
    // Policy description
    description *string;
    // Policy assignment status. This property is read-only.
    isAssigned *bool;
    // Policy last modification date and time. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Policy name
    name *string;
    // Platforms for this policy. Possible values are: none, android, iOS, macOS, windows10X, windows10.
    platforms *DeviceManagementConfigurationPlatforms;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // The list of scheduled action for this rule
    scheduledActionsForRule []DeviceManagementComplianceScheduledActionForRule;
    // Number of settings. This property is read-only.
    settingCount *int32;
    // Policy settings
    settings []DeviceManagementConfigurationSetting;
    // Technologies for this policy. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
    technologies *DeviceManagementConfigurationTechnologies;
}
// Instantiates a new deviceManagementCompliancePolicy and sets the default values.
func NewDeviceManagementCompliancePolicy()(*DeviceManagementCompliancePolicy) {
    m := &DeviceManagementCompliancePolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. Policy assignments
func (m *DeviceManagementCompliancePolicy) GetAssignments()([]DeviceManagementConfigurationPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. Policy creation date and time. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the creationSource property value. Policy creation source
func (m *DeviceManagementCompliancePolicy) GetCreationSource()(*string) {
    if m == nil {
        return nil
    } else {
        return m.creationSource
    }
}
// Gets the description property value. Policy description
func (m *DeviceManagementCompliancePolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// Gets the lastModifiedDateTime property value. Policy last modification date and time. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the name property value. Policy name
func (m *DeviceManagementCompliancePolicy) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the platforms property value. Platforms for this policy. Possible values are: none, android, iOS, macOS, windows10X, windows10.
func (m *DeviceManagementCompliancePolicy) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementCompliancePolicy) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceManagementCompliancePolicy) GetScheduledActionsForRule()([]DeviceManagementComplianceScheduledActionForRule) {
    if m == nil {
        return nil
    } else {
        return m.scheduledActionsForRule
    }
}
// Gets the settingCount property value. Number of settings. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetSettingCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.settingCount
    }
}
// Gets the settings property value. Policy settings
func (m *DeviceManagementCompliancePolicy) GetSettings()([]DeviceManagementConfigurationSetting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the technologies property value. Technologies for this policy. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
func (m *DeviceManagementCompliancePolicy) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    if m == nil {
        return nil
    } else {
        return m.technologies
    }
}
// The deserialization information for the current model
func (m *DeviceManagementCompliancePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationPolicyAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationPolicyAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["creationSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSource(val)
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
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["scheduledActionsForRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementComplianceScheduledActionForRule() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementComplianceScheduledActionForRule, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementComplianceScheduledActionForRule))
            }
            m.SetScheduledActionsForRule(res)
        }
        return nil
    }
    res["settingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCount(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationSetting))
            }
            m.SetSettings(res)
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
    return res
}
func (m *DeviceManagementCompliancePolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementCompliancePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("creationSource", m.GetCreationSource())
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
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    if m.GetPlatforms() != nil {
        cast := m.GetPlatforms().String()
        err = writer.WriteStringValue("platforms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetScheduledActionsForRule()))
        for i, v := range m.GetScheduledActionsForRule() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("scheduledActionsForRule", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingCount", m.GetSettingCount())
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
    if m.GetTechnologies() != nil {
        cast := m.GetTechnologies().String()
        err = writer.WriteStringValue("technologies", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. Policy assignments
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceManagementCompliancePolicy) SetAssignments(value []DeviceManagementConfigurationPolicyAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. Policy creation date and time. This property is read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceManagementCompliancePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the creationSource property value. Policy creation source
// Parameters:
//  - value : Value to set for the creationSource property.
func (m *DeviceManagementCompliancePolicy) SetCreationSource(value *string)() {
    m.creationSource = value
}
// Sets the description property value. Policy description
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementCompliancePolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the isAssigned property value. Policy assignment status. This property is read-only.
// Parameters:
//  - value : Value to set for the isAssigned property.
func (m *DeviceManagementCompliancePolicy) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// Sets the lastModifiedDateTime property value. Policy last modification date and time. This property is read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceManagementCompliancePolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the name property value. Policy name
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceManagementCompliancePolicy) SetName(value *string)() {
    m.name = value
}
// Sets the platforms property value. Platforms for this policy. Possible values are: none, android, iOS, macOS, windows10X, windows10.
// Parameters:
//  - value : Value to set for the platforms property.
func (m *DeviceManagementCompliancePolicy) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceManagementCompliancePolicy) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the scheduledActionsForRule property value. The list of scheduled action for this rule
// Parameters:
//  - value : Value to set for the scheduledActionsForRule property.
func (m *DeviceManagementCompliancePolicy) SetScheduledActionsForRule(value []DeviceManagementComplianceScheduledActionForRule)() {
    m.scheduledActionsForRule = value
}
// Sets the settingCount property value. Number of settings. This property is read-only.
// Parameters:
//  - value : Value to set for the settingCount property.
func (m *DeviceManagementCompliancePolicy) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// Sets the settings property value. Policy settings
// Parameters:
//  - value : Value to set for the settings property.
func (m *DeviceManagementCompliancePolicy) SetSettings(value []DeviceManagementConfigurationSetting)() {
    m.settings = value
}
// Sets the technologies property value. Technologies for this policy. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
// Parameters:
//  - value : Value to set for the technologies property.
func (m *DeviceManagementCompliancePolicy) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
