package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementReusablePolicySetting struct {
    Entity
    // reusable setting creation date and time. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // reusable setting description supplied by user.
    description *string;
    // reusable setting display name supplied by user.
    displayName *string;
    // date and time when reusable setting was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // configuration policies referencing the current reusable setting
    referencingConfigurationPolicies []DeviceManagementConfigurationPolicy;
    // count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
    referencingConfigurationPolicyCount *int32;
    // setting definition id associated with this reusable setting.
    settingDefinitionId *string;
    // reusable setting configuration instance
    settingInstance *DeviceManagementConfigurationSettingInstance;
    // version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
    version *int32;
}
// Instantiates a new deviceManagementReusablePolicySetting and sets the default values.
func NewDeviceManagementReusablePolicySetting()(*DeviceManagementReusablePolicySetting) {
    m := &DeviceManagementReusablePolicySetting{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.referencingConfigurationPolicies
    }
}
// Gets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicyCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.referencingConfigurationPolicyCount
    }
}
// Gets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
// Gets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) GetSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settingInstance
    }
}
// Gets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceManagementReusablePolicySetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["referencingConfigurationPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationPolicy() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementConfigurationPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementConfigurationPolicy))
        }
        m.SetReferencingConfigurationPolicies(res)
        return nil
    }
    res["referencingConfigurationPolicyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetReferencingConfigurationPolicyCount(val)
        return nil
    }
    res["settingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingDefinitionId(val)
        return nil
    }
    res["settingInstance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstance() })
        if err != nil {
            return err
        }
        m.SetSettingInstance(val.(*DeviceManagementConfigurationSettingInstance))
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *DeviceManagementReusablePolicySetting) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementReusablePolicySetting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReferencingConfigurationPolicies()))
        for i, v := range m.GetReferencingConfigurationPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("referencingConfigurationPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("referencingConfigurationPolicyCount", m.GetReferencingConfigurationPolicyCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settingInstance", m.GetSettingInstance())
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
// Sets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceManagementReusablePolicySetting) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. reusable setting description supplied by user.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementReusablePolicySetting) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. reusable setting display name supplied by user.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceManagementReusablePolicySetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceManagementReusablePolicySetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting
// Parameters:
//  - value : Value to set for the referencingConfigurationPolicies property.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicy)() {
    m.referencingConfigurationPolicies = value
}
// Sets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
// Parameters:
//  - value : Value to set for the referencingConfigurationPolicyCount property.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicyCount(value *int32)() {
    m.referencingConfigurationPolicyCount = value
}
// Sets the settingDefinitionId property value. setting definition id associated with this reusable setting.
// Parameters:
//  - value : Value to set for the settingDefinitionId property.
func (m *DeviceManagementReusablePolicySetting) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
// Sets the settingInstance property value. reusable setting configuration instance
// Parameters:
//  - value : Value to set for the settingInstance property.
func (m *DeviceManagementReusablePolicySetting) SetSettingInstance(value *DeviceManagementConfigurationSettingInstance)() {
    m.settingInstance = value
}
// Sets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceManagementReusablePolicySetting) SetVersion(value *int32)() {
    m.version = value
}
