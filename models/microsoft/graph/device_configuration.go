package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfiguration 
type DeviceConfiguration struct {
    Entity
    // The list of assignments for the device configuration profile.
    assignments []DeviceConfigurationAssignmentable;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the Device Configuration.
    description *string;
    // The device mode applicability rule for this Policy.
    deviceManagementApplicabilityRuleDeviceMode DeviceManagementApplicabilityRuleDeviceModeable;
    // The OS edition applicability for this Policy.
    deviceManagementApplicabilityRuleOsEdition DeviceManagementApplicabilityRuleOsEditionable;
    // The OS version applicability rule for this Policy.
    deviceManagementApplicabilityRuleOsVersion DeviceManagementApplicabilityRuleOsVersionable;
    // Device Configuration Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummaryable;
    // Device configuration installation status by device.
    deviceStatuses []DeviceConfigurationDeviceStatusable;
    // Device Configuration devices status overview
    deviceStatusOverview DeviceConfigurationDeviceOverviewable;
    // Admin provided name of the device configuration.
    displayName *string;
    // The list of group assignments for the device configuration profile.
    groupAssignments []DeviceConfigurationGroupAssignmentable;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
    supportsScopeTags *bool;
    // Device configuration installation status by user.
    userStatuses []DeviceConfigurationUserStatusable;
    // Device Configuration users status overview
    userStatusOverview DeviceConfigurationUserOverviewable;
    // Version of the device configuration.
    version *int32;
}
// NewDeviceConfiguration instantiates a new deviceConfiguration and sets the default values.
func NewDeviceConfiguration()(*DeviceConfiguration) {
    m := &DeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) GetAssignments()([]DeviceConfigurationAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceManagementApplicabilityRuleDeviceMode gets the deviceManagementApplicabilityRuleDeviceMode property value. The device mode applicability rule for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleDeviceMode()(DeviceManagementApplicabilityRuleDeviceModeable) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleDeviceMode
    }
}
// GetDeviceManagementApplicabilityRuleOsEdition gets the deviceManagementApplicabilityRuleOsEdition property value. The OS edition applicability for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsEdition()(DeviceManagementApplicabilityRuleOsEditionable) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleOsEdition
    }
}
// GetDeviceManagementApplicabilityRuleOsVersion gets the deviceManagementApplicabilityRuleOsVersion property value. The OS version applicability rule for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsVersion()(DeviceManagementApplicabilityRuleOsVersionable) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleOsVersion
    }
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) GetDeviceStatuses()([]DeviceConfigurationDeviceStatusable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDeviceStatusOverview gets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) GetDeviceStatusOverview()(DeviceConfigurationDeviceOverviewable) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusOverview
    }
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationAssignmentable)
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
    res["deviceManagementApplicabilityRuleDeviceMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementApplicabilityRuleDeviceModeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleDeviceMode(val.(DeviceManagementApplicabilityRuleDeviceModeable))
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleOsEdition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementApplicabilityRuleOsEditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleOsEdition(val.(DeviceManagementApplicabilityRuleOsEditionable))
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementApplicabilityRuleOsVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleOsVersion(val.(DeviceManagementApplicabilityRuleOsVersionable))
        }
        return nil
    }
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingStateDeviceSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingStateDeviceSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(SettingStateDeviceSummaryable)
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationDeviceStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusOverview(val.(DeviceConfigurationDeviceOverviewable))
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
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationGroupAssignmentable)
            }
            m.SetGroupAssignments(res)
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
    res["supportsScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsScopeTags(val)
        }
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationUserStatusable)
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationUserOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusOverview(val.(DeviceConfigurationUserOverviewable))
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
// GetGroupAssignments gets the groupAssignments property value. The list of group assignments for the device configuration profile.
func (m *DeviceConfiguration) GetGroupAssignments()([]DeviceConfigurationGroupAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetSupportsScopeTags gets the supportsScopeTags property value. Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
func (m *DeviceConfiguration) GetSupportsScopeTags()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsScopeTags
    }
}
// GetUserStatuses gets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) GetUserStatuses()([]DeviceConfigurationUserStatusable) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// GetUserStatusOverview gets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) GetUserStatusOverview()(DeviceConfigurationUserOverviewable) {
    if m == nil {
        return nil
    } else {
        return m.userStatusOverview
    }
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *DeviceConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleDeviceMode", m.GetDeviceManagementApplicabilityRuleDeviceMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleOsEdition", m.GetDeviceManagementApplicabilityRuleOsEdition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleOsVersion", m.GetDeviceManagementApplicabilityRuleOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
    if m.GetGroupAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
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
    {
        err = writer.WriteBoolValue("supportsScopeTags", m.GetSupportsScopeTags())
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
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
// SetAssignments sets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) SetAssignments(value []DeviceConfigurationAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceManagementApplicabilityRuleDeviceMode sets the deviceManagementApplicabilityRuleDeviceMode property value. The device mode applicability rule for this Policy.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleDeviceMode(value DeviceManagementApplicabilityRuleDeviceModeable)() {
    if m != nil {
        m.deviceManagementApplicabilityRuleDeviceMode = value
    }
}
// SetDeviceManagementApplicabilityRuleOsEdition sets the deviceManagementApplicabilityRuleOsEdition property value. The OS edition applicability for this Policy.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsEdition(value DeviceManagementApplicabilityRuleOsEditionable)() {
    if m != nil {
        m.deviceManagementApplicabilityRuleOsEdition = value
    }
}
// SetDeviceManagementApplicabilityRuleOsVersion sets the deviceManagementApplicabilityRuleOsVersion property value. The OS version applicability rule for this Policy.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsVersion(value DeviceManagementApplicabilityRuleOsVersionable)() {
    if m != nil {
        m.deviceManagementApplicabilityRuleOsVersion = value
    }
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)() {
    if m != nil {
        m.deviceSettingStateSummaries = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) SetDeviceStatuses(value []DeviceConfigurationDeviceStatusable)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDeviceStatusOverview sets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) SetDeviceStatusOverview(value DeviceConfigurationDeviceOverviewable)() {
    if m != nil {
        m.deviceStatusOverview = value
    }
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGroupAssignments sets the groupAssignments property value. The list of group assignments for the device configuration profile.
func (m *DeviceConfiguration) SetGroupAssignments(value []DeviceConfigurationGroupAssignmentable)() {
    if m != nil {
        m.groupAssignments = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceConfiguration) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetSupportsScopeTags sets the supportsScopeTags property value. Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
func (m *DeviceConfiguration) SetSupportsScopeTags(value *bool)() {
    if m != nil {
        m.supportsScopeTags = value
    }
}
// SetUserStatuses sets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) SetUserStatuses(value []DeviceConfigurationUserStatusable)() {
    if m != nil {
        m.userStatuses = value
    }
}
// SetUserStatusOverview sets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) SetUserStatusOverview(value DeviceConfigurationUserOverviewable)() {
    if m != nil {
        m.userStatusOverview = value
    }
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
