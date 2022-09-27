package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCompliancePolicy device Management Compliance Policy
type DeviceManagementCompliancePolicy struct {
    Entity
    // Policy assignments
    assignments []DeviceManagementConfigurationPolicyAssignmentable
    // Policy creation date and time. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Policy creation source
    creationSource *string
    // Policy description
    description *string
    // Policy assignment status. This property is read-only.
    isAssigned *bool
    // Policy last modification date and time. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Policy name
    name *string
    // Supported platform types.
    platforms *DeviceManagementConfigurationPlatforms
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // The list of scheduled action for this rule
    scheduledActionsForRule []DeviceManagementComplianceScheduledActionForRuleable
    // Number of settings. This property is read-only.
    settingCount *int32
    // Policy settings
    settings []DeviceManagementConfigurationSettingable
    // Describes which technology this setting can be deployed with
    technologies *DeviceManagementConfigurationTechnologies
}
// NewDeviceManagementCompliancePolicy instantiates a new deviceManagementCompliancePolicy and sets the default values.
func NewDeviceManagementCompliancePolicy()(*DeviceManagementCompliancePolicy) {
    m := &DeviceManagementCompliancePolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementCompliancePolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCompliancePolicy(), nil
}
// GetAssignments gets the assignments property value. Policy assignments
func (m *DeviceManagementCompliancePolicy) GetAssignments()([]DeviceManagementConfigurationPolicyAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. Policy creation date and time. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreationSource gets the creationSource property value. Policy creation source
func (m *DeviceManagementCompliancePolicy) GetCreationSource()(*string) {
    return m.creationSource
}
// GetDescription gets the description property value. Policy description
func (m *DeviceManagementCompliancePolicy) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["creationSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreationSource)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["isAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAssigned)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["platforms"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationPlatforms , m.SetPlatforms)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["scheduledActionsForRule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementComplianceScheduledActionForRuleFromDiscriminatorValue , m.SetScheduledActionsForRule)
    res["settingCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSettingCount)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingFromDiscriminatorValue , m.SetSettings)
    res["technologies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationTechnologies , m.SetTechnologies)
    return res
}
// GetIsAssigned gets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetIsAssigned()(*bool) {
    return m.isAssigned
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Policy last modification date and time. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. Policy name
func (m *DeviceManagementCompliancePolicy) GetName()(*string) {
    return m.name
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementCompliancePolicy) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    return m.platforms
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementCompliancePolicy) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetScheduledActionsForRule gets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceManagementCompliancePolicy) GetScheduledActionsForRule()([]DeviceManagementComplianceScheduledActionForRuleable) {
    return m.scheduledActionsForRule
}
// GetSettingCount gets the settingCount property value. Number of settings. This property is read-only.
func (m *DeviceManagementCompliancePolicy) GetSettingCount()(*int32) {
    return m.settingCount
}
// GetSettings gets the settings property value. Policy settings
func (m *DeviceManagementCompliancePolicy) GetSettings()([]DeviceManagementConfigurationSettingable) {
    return m.settings
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementCompliancePolicy) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    return m.technologies
}
// Serialize serializes information the current object
func (m *DeviceManagementCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetScheduledActionsForRule() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetScheduledActionsForRule())
        err = writer.WriteCollectionOfObjectValues("scheduledActionsForRule", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSettings())
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
    return nil
}
// SetAssignments sets the assignments property value. Policy assignments
func (m *DeviceManagementCompliancePolicy) SetAssignments(value []DeviceManagementConfigurationPolicyAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. Policy creation date and time. This property is read-only.
func (m *DeviceManagementCompliancePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreationSource sets the creationSource property value. Policy creation source
func (m *DeviceManagementCompliancePolicy) SetCreationSource(value *string)() {
    m.creationSource = value
}
// SetDescription sets the description property value. Policy description
func (m *DeviceManagementCompliancePolicy) SetDescription(value *string)() {
    m.description = value
}
// SetIsAssigned sets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementCompliancePolicy) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Policy last modification date and time. This property is read-only.
func (m *DeviceManagementCompliancePolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. Policy name
func (m *DeviceManagementCompliancePolicy) SetName(value *string)() {
    m.name = value
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementCompliancePolicy) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementCompliancePolicy) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetScheduledActionsForRule sets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceManagementCompliancePolicy) SetScheduledActionsForRule(value []DeviceManagementComplianceScheduledActionForRuleable)() {
    m.scheduledActionsForRule = value
}
// SetSettingCount sets the settingCount property value. Number of settings. This property is read-only.
func (m *DeviceManagementCompliancePolicy) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// SetSettings sets the settings property value. Policy settings
func (m *DeviceManagementCompliancePolicy) SetSettings(value []DeviceManagementConfigurationSettingable)() {
    m.settings = value
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementCompliancePolicy) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
