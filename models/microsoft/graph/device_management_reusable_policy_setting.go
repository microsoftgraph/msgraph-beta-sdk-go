package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementReusablePolicySetting 
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
    // configuration policies referencing the current reusable setting. This property is read-only.
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
// NewDeviceManagementReusablePolicySetting instantiates a new deviceManagementReusablePolicySetting and sets the default values.
func NewDeviceManagementReusablePolicySetting()(*DeviceManagementReusablePolicySetting) {
    m := &DeviceManagementReusablePolicySetting{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetReferencingConfigurationPolicies gets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicy) {
    if m == nil {
        return nil
    } else {
        return m.referencingConfigurationPolicies
    }
}
// GetReferencingConfigurationPolicyCount gets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicyCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.referencingConfigurationPolicyCount
    }
}
// GetSettingDefinitionId gets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
// GetSettingInstance gets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) GetSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settingInstance
    }
}
// GetVersion gets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementReusablePolicySetting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["referencingConfigurationPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationPolicy))
            }
            m.SetReferencingConfigurationPolicies(res)
        }
        return nil
    }
    res["referencingConfigurationPolicyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencingConfigurationPolicyCount(val)
        }
        return nil
    }
    res["settingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    res["settingInstance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstance(val.(*DeviceManagementConfigurationSettingInstance))
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
func (m *DeviceManagementReusablePolicySetting) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetReferencingConfigurationPolicies() != nil {
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
// SetCreatedDateTime sets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetReferencingConfigurationPolicies sets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicy)() {
    if m != nil {
        m.referencingConfigurationPolicies = value
    }
}
// SetReferencingConfigurationPolicyCount sets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicyCount(value *int32)() {
    if m != nil {
        m.referencingConfigurationPolicyCount = value
    }
}
// SetSettingDefinitionId sets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) SetSettingDefinitionId(value *string)() {
    if m != nil {
        m.settingDefinitionId = value
    }
}
// SetSettingInstance sets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) SetSettingInstance(value *DeviceManagementConfigurationSettingInstance)() {
    if m != nil {
        m.settingInstance = value
    }
}
// SetVersion sets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
