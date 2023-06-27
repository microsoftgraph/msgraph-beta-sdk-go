package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementReusablePolicySetting 
type DeviceManagementReusablePolicySetting struct {
    Entity
}
// NewDeviceManagementReusablePolicySetting instantiates a new DeviceManagementReusablePolicySetting and sets the default values.
func NewDeviceManagementReusablePolicySetting()(*DeviceManagementReusablePolicySetting) {
    m := &DeviceManagementReusablePolicySetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementReusablePolicySetting(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) GetDisplayName()(*string) {
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
func (m *DeviceManagementReusablePolicySetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["referencingConfigurationPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DeviceManagementConfigurationPolicyable)
                }
            }
            m.SetReferencingConfigurationPolicies(res)
        }
        return nil
    }
    res["referencingConfigurationPolicyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencingConfigurationPolicyCount(val)
        }
        return nil
    }
    res["settingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    res["settingInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstance(val.(DeviceManagementConfigurationSettingInstanceable))
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReferencingConfigurationPolicies gets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicyable) {
    val, err := m.GetBackingStore().Get("referencingConfigurationPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationPolicyable)
    }
    return nil
}
// GetReferencingConfigurationPolicyCount gets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetReferencingConfigurationPolicyCount()(*int32) {
    val, err := m.GetBackingStore().Get("referencingConfigurationPolicyCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSettingDefinitionId gets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) GetSettingDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("settingDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingInstance gets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable) {
    val, err := m.GetBackingStore().Get("settingInstance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationSettingInstanceable)
    }
    return nil
}
// GetVersion gets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) GetVersion()(*int32) {
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
func (m *DeviceManagementReusablePolicySetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetReferencingConfigurationPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReferencingConfigurationPolicies()))
        for i, v := range m.GetReferencingConfigurationPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("referencingConfigurationPolicies", cast)
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
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. reusable setting creation date and time. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. reusable setting description supplied by user.
func (m *DeviceManagementReusablePolicySetting) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. reusable setting display name supplied by user.
func (m *DeviceManagementReusablePolicySetting) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. date and time when reusable setting was last modified. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReferencingConfigurationPolicies sets the referencingConfigurationPolicies property value. configuration policies referencing the current reusable setting. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)() {
    err := m.GetBackingStore().Set("referencingConfigurationPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetReferencingConfigurationPolicyCount sets the referencingConfigurationPolicyCount property value. count of configuration policies referencing the current reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetReferencingConfigurationPolicyCount(value *int32)() {
    err := m.GetBackingStore().Set("referencingConfigurationPolicyCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDefinitionId sets the settingDefinitionId property value. setting definition id associated with this reusable setting.
func (m *DeviceManagementReusablePolicySetting) SetSettingDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("settingDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingInstance sets the settingInstance property value. reusable setting configuration instance
func (m *DeviceManagementReusablePolicySetting) SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)() {
    err := m.GetBackingStore().Set("settingInstance", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. version number for reusable setting. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementReusablePolicySetting) SetVersion(value *int32)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementReusablePolicySettingable 
type DeviceManagementReusablePolicySettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReferencingConfigurationPolicies()([]DeviceManagementConfigurationPolicyable)
    GetReferencingConfigurationPolicyCount()(*int32)
    GetSettingDefinitionId()(*string)
    GetSettingInstance()(DeviceManagementConfigurationSettingInstanceable)
    GetVersion()(*int32)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReferencingConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)()
    SetReferencingConfigurationPolicyCount(value *int32)()
    SetSettingDefinitionId(value *string)()
    SetSettingInstance(value DeviceManagementConfigurationSettingInstanceable)()
    SetVersion(value *int32)()
}
