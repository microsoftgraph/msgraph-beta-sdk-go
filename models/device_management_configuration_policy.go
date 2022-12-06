package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicy device Management Configuration Policy
type DeviceManagementConfigurationPolicy struct {
    Entity
    // Policy assignments
    assignments []DeviceManagementConfigurationPolicyAssignmentable
    // Policy creation date and time
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Policy creation source
    creationSource *string
    // Policy description
    description *string
    // Policy assignment status. This property is read-only.
    isAssigned *bool
    // Policy last modification date and time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Policy name
    name *string
    // Supported platform types.
    platforms *DeviceManagementConfigurationPlatforms
    // Indicates the priority of each policies that are selected by the admin during enrollment process
    priorityMetaData DeviceManagementPriorityMetaDataable
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // Number of settings
    settingCount *int32
    // Policy settings
    settings []DeviceManagementConfigurationSettingable
    // Describes which technology this setting can be deployed with
    technologies *DeviceManagementConfigurationTechnologies
    // Template reference information
    templateReference DeviceManagementConfigurationPolicyTemplateReferenceable
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
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. Policy creation date and time
func (m *DeviceManagementConfigurationPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreationSource gets the creationSource property value. Policy creation source
func (m *DeviceManagementConfigurationPolicy) GetCreationSource()(*string) {
    return m.creationSource
}
// GetDescription gets the description property value. Policy description
func (m *DeviceManagementConfigurationPolicy) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["creationSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreationSource)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["isAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAssigned)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["platforms"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationPlatforms , m.SetPlatforms)
    res["priorityMetaData"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementPriorityMetaDataFromDiscriminatorValue , m.SetPriorityMetaData)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["settingCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSettingCount)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingFromDiscriminatorValue , m.SetSettings)
    res["technologies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementConfigurationTechnologies , m.SetTechnologies)
    res["templateReference"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceManagementConfigurationPolicyTemplateReferenceFromDiscriminatorValue , m.SetTemplateReference)
    return res
}
// GetIsAssigned gets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementConfigurationPolicy) GetIsAssigned()(*bool) {
    return m.isAssigned
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Policy last modification date and time
func (m *DeviceManagementConfigurationPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. Policy name
func (m *DeviceManagementConfigurationPolicy) GetName()(*string) {
    return m.name
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicy) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    return m.platforms
}
// GetPriorityMetaData gets the priorityMetaData property value. Indicates the priority of each policies that are selected by the admin during enrollment process
func (m *DeviceManagementConfigurationPolicy) GetPriorityMetaData()(DeviceManagementPriorityMetaDataable) {
    return m.priorityMetaData
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementConfigurationPolicy) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSettingCount gets the settingCount property value. Number of settings
func (m *DeviceManagementConfigurationPolicy) GetSettingCount()(*int32) {
    return m.settingCount
}
// GetSettings gets the settings property value. Policy settings
func (m *DeviceManagementConfigurationPolicy) GetSettings()([]DeviceManagementConfigurationSettingable) {
    return m.settings
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicy) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    return m.technologies
}
// GetTemplateReference gets the templateReference property value. Template reference information
func (m *DeviceManagementConfigurationPolicy) GetTemplateReference()(DeviceManagementConfigurationPolicyTemplateReferenceable) {
    return m.templateReference
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. Policy creation date and time
func (m *DeviceManagementConfigurationPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreationSource sets the creationSource property value. Policy creation source
func (m *DeviceManagementConfigurationPolicy) SetCreationSource(value *string)() {
    m.creationSource = value
}
// SetDescription sets the description property value. Policy description
func (m *DeviceManagementConfigurationPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetIsAssigned sets the isAssigned property value. Policy assignment status. This property is read-only.
func (m *DeviceManagementConfigurationPolicy) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Policy last modification date and time
func (m *DeviceManagementConfigurationPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. Policy name
func (m *DeviceManagementConfigurationPolicy) SetName(value *string)() {
    m.name = value
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicy) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// SetPriorityMetaData sets the priorityMetaData property value. Indicates the priority of each policies that are selected by the admin during enrollment process
func (m *DeviceManagementConfigurationPolicy) SetPriorityMetaData(value DeviceManagementPriorityMetaDataable)() {
    m.priorityMetaData = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceManagementConfigurationPolicy) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSettingCount sets the settingCount property value. Number of settings
func (m *DeviceManagementConfigurationPolicy) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// SetSettings sets the settings property value. Policy settings
func (m *DeviceManagementConfigurationPolicy) SetSettings(value []DeviceManagementConfigurationSettingable)() {
    m.settings = value
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicy) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
// SetTemplateReference sets the templateReference property value. Template reference information
func (m *DeviceManagementConfigurationPolicy) SetTemplateReference(value DeviceManagementConfigurationPolicyTemplateReferenceable)() {
    m.templateReference = value
}
