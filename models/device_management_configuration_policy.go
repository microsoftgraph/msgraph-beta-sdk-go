package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicy device Management Configuration Policy
type DeviceManagementConfigurationPolicy struct {
    Entity
}
// NewDeviceManagementConfigurationPolicy instantiates a new deviceManagementConfigurationPolicy and sets the default values.
func NewDeviceManagementConfigurationPolicy()(*DeviceManagementConfigurationPolicy) {
    m := &DeviceManagementConfigurationPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationPolicy(), nil
}
// GetAssignments gets the assignments property value. Policy assignments
func (m *DeviceManagementConfigurationPolicy) GetAssignments()([]DeviceManagementConfigurationPolicyAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationPolicyAssignmentable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Policy creation date and time
func (m *DeviceManagementConfigurationPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreationSource gets the creationSource property value. Policy creation source
func (m *DeviceManagementConfigurationPolicy) GetCreationSource()(*string) {
    val, err := m.GetBackingStore().Get("creationSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Policy description
func (m *DeviceManagementConfigurationPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
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
    res["creationSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSource(val)
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
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(*DeviceManagementConfigurationPlatforms))
        }
        return nil
    }
    res["priorityMetaData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementPriorityMetaDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriorityMetaData(val.(DeviceManagementPriorityMetaDataable))
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["settingCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCount(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingable)
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["technologies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnologies(val.(*DeviceManagementConfigurationTechnologies))
        }
        return nil
    }
    res["templateReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationPolicyTemplateReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateReference(val.(DeviceManagementConfigurationPolicyTemplateReferenceable))
        }
        return nil
    }
    return res
}
// GetIsAssigned gets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementConfigurationPolicy) GetIsAssigned()(*bool) {
    val, err := m.GetBackingStore().Get("isAssigned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Policy last modification date and time
func (m *DeviceManagementConfigurationPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetName gets the name property value. Policy name
func (m *DeviceManagementConfigurationPolicy) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicy) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    val, err := m.GetBackingStore().Get("platforms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationPlatforms)
    }
    return nil
}
// GetPriorityMetaData gets the priorityMetaData property value. Indicates the priority of each policies that are selected by the admin during enrollment process
func (m *DeviceManagementConfigurationPolicy) GetPriorityMetaData()(DeviceManagementPriorityMetaDataable) {
    val, err := m.GetBackingStore().Get("priorityMetaData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementPriorityMetaDataable)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementConfigurationPolicy) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSettingCount gets the settingCount property value. Number of settings
func (m *DeviceManagementConfigurationPolicy) GetSettingCount()(*int32) {
    val, err := m.GetBackingStore().Get("settingCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSettings gets the settings property value. Policy settings
func (m *DeviceManagementConfigurationPolicy) GetSettings()([]DeviceManagementConfigurationSettingable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceManagementConfigurationSettingable)
    }
    return nil
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicy) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    val, err := m.GetBackingStore().Get("technologies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementConfigurationTechnologies)
    }
    return nil
}
// GetTemplateReference gets the templateReference property value. Template reference information
func (m *DeviceManagementConfigurationPolicy) GetTemplateReference()(DeviceManagementConfigurationPolicyTemplateReferenceable) {
    val, err := m.GetBackingStore().Get("templateReference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceManagementConfigurationPolicyTemplateReferenceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := (*m.GetPlatforms()).String()
        err = writer.WriteStringValue("platforms", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("priorityMetaData", m.GetPriorityMetaData())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
    if m.GetTechnologies() != nil {
        cast := (*m.GetTechnologies()).String()
        err = writer.WriteStringValue("technologies", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateReference", m.GetTemplateReference())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. Policy assignments
func (m *DeviceManagementConfigurationPolicy) SetAssignments(value []DeviceManagementConfigurationPolicyAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Policy creation date and time
func (m *DeviceManagementConfigurationPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCreationSource sets the creationSource property value. Policy creation source
func (m *DeviceManagementConfigurationPolicy) SetCreationSource(value *string)() {
    err := m.GetBackingStore().Set("creationSource", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Policy description
func (m *DeviceManagementConfigurationPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAssigned sets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementConfigurationPolicy) SetIsAssigned(value *bool)() {
    err := m.GetBackingStore().Set("isAssigned", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Policy last modification date and time
func (m *DeviceManagementConfigurationPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Policy name
func (m *DeviceManagementConfigurationPolicy) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicy) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    err := m.GetBackingStore().Set("platforms", value)
    if err != nil {
        panic(err)
    }
}
// SetPriorityMetaData sets the priorityMetaData property value. Indicates the priority of each policies that are selected by the admin during enrollment process
func (m *DeviceManagementConfigurationPolicy) SetPriorityMetaData(value DeviceManagementPriorityMetaDataable)() {
    err := m.GetBackingStore().Set("priorityMetaData", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementConfigurationPolicy) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingCount sets the settingCount property value. Number of settings
func (m *DeviceManagementConfigurationPolicy) SetSettingCount(value *int32)() {
    err := m.GetBackingStore().Set("settingCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Policy settings
func (m *DeviceManagementConfigurationPolicy) SetSettings(value []DeviceManagementConfigurationSettingable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicy) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    err := m.GetBackingStore().Set("technologies", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateReference sets the templateReference property value. Template reference information
func (m *DeviceManagementConfigurationPolicy) SetTemplateReference(value DeviceManagementConfigurationPolicyTemplateReferenceable)() {
    err := m.GetBackingStore().Set("templateReference", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementConfigurationPolicyable 
type DeviceManagementConfigurationPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceManagementConfigurationPolicyAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreationSource()(*string)
    GetDescription()(*string)
    GetIsAssigned()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetPriorityMetaData()(DeviceManagementPriorityMetaDataable)
    GetRoleScopeTagIds()([]string)
    GetSettingCount()(*int32)
    GetSettings()([]DeviceManagementConfigurationSettingable)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    GetTemplateReference()(DeviceManagementConfigurationPolicyTemplateReferenceable)
    SetAssignments(value []DeviceManagementConfigurationPolicyAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreationSource(value *string)()
    SetDescription(value *string)()
    SetIsAssigned(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetPriorityMetaData(value DeviceManagementPriorityMetaDataable)()
    SetRoleScopeTagIds(value []string)()
    SetSettingCount(value *int32)()
    SetSettings(value []DeviceManagementConfigurationSettingable)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
    SetTemplateReference(value DeviceManagementConfigurationPolicyTemplateReferenceable)()
}
