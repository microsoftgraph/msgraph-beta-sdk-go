package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceConfiguration struct {
    Entity
    assignments []DeviceConfigurationAssignment;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *string;
    deviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode;
    deviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition;
    deviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion;
    deviceSettingStateSummaries []SettingStateDeviceSummary;
    deviceStatuses []DeviceConfigurationDeviceStatus;
    deviceStatusOverview *DeviceConfigurationDeviceOverview;
    displayName *string;
    groupAssignments []DeviceConfigurationGroupAssignment;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    roleScopeTagIds []string;
    supportsScopeTags *bool;
    userStatuses []DeviceConfigurationUserStatus;
    userStatusOverview *DeviceConfigurationUserOverview;
    version *int32;
}
func NewDeviceConfiguration()(*DeviceConfiguration) {
    m := &DeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceConfiguration) GetAssignments()([]DeviceConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *DeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *DeviceConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleDeviceMode()(*DeviceManagementApplicabilityRuleDeviceMode) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleDeviceMode
    }
}
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsEdition()(*DeviceManagementApplicabilityRuleOsEdition) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleOsEdition
    }
}
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsVersion()(*DeviceManagementApplicabilityRuleOsVersion) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleOsVersion
    }
}
func (m *DeviceConfiguration) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
func (m *DeviceConfiguration) GetDeviceStatuses()([]DeviceConfigurationDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
func (m *DeviceConfiguration) GetDeviceStatusOverview()(*DeviceConfigurationDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusOverview
    }
}
func (m *DeviceConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceConfiguration) GetGroupAssignments()([]DeviceConfigurationGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
func (m *DeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *DeviceConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *DeviceConfiguration) GetSupportsScopeTags()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsScopeTags
    }
}
func (m *DeviceConfiguration) GetUserStatuses()([]DeviceConfigurationUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
func (m *DeviceConfiguration) GetUserStatusOverview()(*DeviceConfigurationUserOverview) {
    if m == nil {
        return nil
    } else {
        return m.userStatusOverview
    }
}
func (m *DeviceConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
func (m *DeviceConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfigurationAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfigurationAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
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
    res["deviceManagementApplicabilityRuleDeviceMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementApplicabilityRuleDeviceMode() })
        if err != nil {
            return err
        }
        m.SetDeviceManagementApplicabilityRuleDeviceMode(val.(*DeviceManagementApplicabilityRuleDeviceMode))
        return nil
    }
    res["deviceManagementApplicabilityRuleOsEdition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementApplicabilityRuleOsEdition() })
        if err != nil {
            return err
        }
        m.SetDeviceManagementApplicabilityRuleOsEdition(val.(*DeviceManagementApplicabilityRuleOsEdition))
        return nil
    }
    res["deviceManagementApplicabilityRuleOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementApplicabilityRuleOsVersion() })
        if err != nil {
            return err
        }
        m.SetDeviceManagementApplicabilityRuleOsVersion(val.(*DeviceManagementApplicabilityRuleOsVersion))
        return nil
    }
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingStateDeviceSummary() })
        if err != nil {
            return err
        }
        res := make([]SettingStateDeviceSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*SettingStateDeviceSummary))
        }
        m.SetDeviceSettingStateSummaries(res)
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationDeviceStatus() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfigurationDeviceStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfigurationDeviceStatus))
        }
        m.SetDeviceStatuses(res)
        return nil
    }
    res["deviceStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationDeviceOverview() })
        if err != nil {
            return err
        }
        m.SetDeviceStatusOverview(val.(*DeviceConfigurationDeviceOverview))
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
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationGroupAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfigurationGroupAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfigurationGroupAssignment))
        }
        m.SetGroupAssignments(res)
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
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["supportsScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSupportsScopeTags(val)
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationUserStatus() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfigurationUserStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfigurationUserStatus))
        }
        m.SetUserStatuses(res)
        return nil
    }
    res["userStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationUserOverview() })
        if err != nil {
            return err
        }
        m.SetUserStatusOverview(val.(*DeviceConfigurationUserOverview))
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
func (m *DeviceConfiguration) IsNil()(bool) {
    return m == nil
}
func (m *DeviceConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
func (m *DeviceConfiguration) SetAssignments(value []DeviceConfigurationAssignment)() {
    m.assignments = value
}
func (m *DeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *DeviceConfiguration) SetDescription(value *string)() {
    m.description = value
}
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleDeviceMode(value *DeviceManagementApplicabilityRuleDeviceMode)() {
    m.deviceManagementApplicabilityRuleDeviceMode = value
}
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsEdition(value *DeviceManagementApplicabilityRuleOsEdition)() {
    m.deviceManagementApplicabilityRuleOsEdition = value
}
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsVersion(value *DeviceManagementApplicabilityRuleOsVersion)() {
    m.deviceManagementApplicabilityRuleOsVersion = value
}
func (m *DeviceConfiguration) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummary)() {
    m.deviceSettingStateSummaries = value
}
func (m *DeviceConfiguration) SetDeviceStatuses(value []DeviceConfigurationDeviceStatus)() {
    m.deviceStatuses = value
}
func (m *DeviceConfiguration) SetDeviceStatusOverview(value *DeviceConfigurationDeviceOverview)() {
    m.deviceStatusOverview = value
}
func (m *DeviceConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceConfiguration) SetGroupAssignments(value []DeviceConfigurationGroupAssignment)() {
    m.groupAssignments = value
}
func (m *DeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *DeviceConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *DeviceConfiguration) SetSupportsScopeTags(value *bool)() {
    m.supportsScopeTags = value
}
func (m *DeviceConfiguration) SetUserStatuses(value []DeviceConfigurationUserStatus)() {
    m.userStatuses = value
}
func (m *DeviceConfiguration) SetUserStatusOverview(value *DeviceConfigurationUserOverview)() {
    m.userStatusOverview = value
}
func (m *DeviceConfiguration) SetVersion(value *int32)() {
    m.version = value
}
