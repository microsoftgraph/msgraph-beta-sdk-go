package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCompliancePolicy this is the base class for Compliance policy. Compliance policies are platform specific and individual per-platform compliance policies inherit from here. 
type DeviceCompliancePolicy struct {
    Entity
    // The collection of assignments for this compliance policy.
    assignments []DeviceCompliancePolicyAssignmentable
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the Device Configuration.
    description *string
    // Compliance Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummaryable
    // List of DeviceComplianceDeviceStatus.
    deviceStatuses []DeviceComplianceDeviceStatusable
    // Device compliance devices status overview
    deviceStatusOverview DeviceComplianceDeviceOverviewable
    // Admin provided name of the device configuration.
    displayName *string
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // The list of scheduled action for this rule
    scheduledActionsForRule []DeviceComplianceScheduledActionForRuleable
    // List of DeviceComplianceUserStatus.
    userStatuses []DeviceComplianceUserStatusable
    // Device compliance users status overview
    userStatusOverview DeviceComplianceUserOverviewable
    // Version of the device configuration.
    version *int32
}
// NewDeviceCompliancePolicy instantiates a new deviceCompliancePolicy and sets the default values.
func NewDeviceCompliancePolicy()(*DeviceCompliancePolicy) {
    m := &DeviceCompliancePolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidCompliancePolicy":
                        return NewAndroidCompliancePolicy(), nil
                    case "#microsoft.graph.androidDeviceOwnerCompliancePolicy":
                        return NewAndroidDeviceOwnerCompliancePolicy(), nil
                    case "#microsoft.graph.androidForWorkCompliancePolicy":
                        return NewAndroidForWorkCompliancePolicy(), nil
                    case "#microsoft.graph.androidWorkProfileCompliancePolicy":
                        return NewAndroidWorkProfileCompliancePolicy(), nil
                    case "#microsoft.graph.aospDeviceOwnerCompliancePolicy":
                        return NewAospDeviceOwnerCompliancePolicy(), nil
                    case "#microsoft.graph.defaultDeviceCompliancePolicy":
                        return NewDefaultDeviceCompliancePolicy(), nil
                    case "#microsoft.graph.iosCompliancePolicy":
                        return NewIosCompliancePolicy(), nil
                    case "#microsoft.graph.macOSCompliancePolicy":
                        return NewMacOSCompliancePolicy(), nil
                    case "#microsoft.graph.windows10CompliancePolicy":
                        return NewWindows10CompliancePolicy(), nil
                    case "#microsoft.graph.windows10MobileCompliancePolicy":
                        return NewWindows10MobileCompliancePolicy(), nil
                    case "#microsoft.graph.windows81CompliancePolicy":
                        return NewWindows81CompliancePolicy(), nil
                    case "#microsoft.graph.windowsPhone81CompliancePolicy":
                        return NewWindowsPhone81CompliancePolicy(), nil
                }
            }
        }
    }
    return NewDeviceCompliancePolicy(), nil
}
// GetAssignments gets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) GetAssignments()([]DeviceCompliancePolicyAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) GetDescription()(*string) {
    return m.description
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable) {
    return m.deviceSettingStateSummaries
}
// GetDeviceStatuses gets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) GetDeviceStatuses()([]DeviceComplianceDeviceStatusable) {
    return m.deviceStatuses
}
// GetDeviceStatusOverview gets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) GetDeviceStatusOverview()(DeviceComplianceDeviceOverviewable) {
    return m.deviceStatusOverview
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceCompliancePolicyAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["deviceSettingStateSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSettingStateDeviceSummaryFromDiscriminatorValue , m.SetDeviceSettingStateSummaries)
    res["deviceStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceComplianceDeviceStatusFromDiscriminatorValue , m.SetDeviceStatuses)
    res["deviceStatusOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceComplianceDeviceOverviewFromDiscriminatorValue , m.SetDeviceStatusOverview)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["scheduledActionsForRule"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue , m.SetScheduledActionsForRule)
    res["userStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceComplianceUserStatusFromDiscriminatorValue , m.SetUserStatuses)
    res["userStatusOverview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDeviceComplianceUserOverviewFromDiscriminatorValue , m.SetUserStatusOverview)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceCompliancePolicy) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetScheduledActionsForRule gets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceCompliancePolicy) GetScheduledActionsForRule()([]DeviceComplianceScheduledActionForRuleable) {
    return m.scheduledActionsForRule
}
// GetUserStatuses gets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) GetUserStatuses()([]DeviceComplianceUserStatusable) {
    return m.userStatuses
}
// GetUserStatusOverview gets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) GetUserStatusOverview()(DeviceComplianceUserOverviewable) {
    return m.userStatusOverview
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceSettingStateSummaries())
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceStatuses())
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusOverview", m.GetDeviceStatusOverview())
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
    if m.GetUserStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserStatuses())
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusOverview", m.GetUserStatusOverview())
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
// SetAssignments sets the assignments property value. The collection of assignments for this compliance policy.
func (m *DeviceCompliancePolicy) SetAssignments(value []DeviceCompliancePolicyAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *DeviceCompliancePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceCompliancePolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Compliance Setting State Device Summary
func (m *DeviceCompliancePolicy) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)() {
    m.deviceSettingStateSummaries = value
}
// SetDeviceStatuses sets the deviceStatuses property value. List of DeviceComplianceDeviceStatus.
func (m *DeviceCompliancePolicy) SetDeviceStatuses(value []DeviceComplianceDeviceStatusable)() {
    m.deviceStatuses = value
}
// SetDeviceStatusOverview sets the deviceStatusOverview property value. Device compliance devices status overview
func (m *DeviceCompliancePolicy) SetDeviceStatusOverview(value DeviceComplianceDeviceOverviewable)() {
    m.deviceStatusOverview = value
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceCompliancePolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceCompliancePolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceCompliancePolicy) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetScheduledActionsForRule sets the scheduledActionsForRule property value. The list of scheduled action for this rule
func (m *DeviceCompliancePolicy) SetScheduledActionsForRule(value []DeviceComplianceScheduledActionForRuleable)() {
    m.scheduledActionsForRule = value
}
// SetUserStatuses sets the userStatuses property value. List of DeviceComplianceUserStatus.
func (m *DeviceCompliancePolicy) SetUserStatuses(value []DeviceComplianceUserStatusable)() {
    m.userStatuses = value
}
// SetUserStatusOverview sets the userStatusOverview property value. Device compliance users status overview
func (m *DeviceCompliancePolicy) SetUserStatusOverview(value DeviceComplianceUserOverviewable)() {
    m.userStatusOverview = value
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *DeviceCompliancePolicy) SetVersion(value *int32)() {
    m.version = value
}
