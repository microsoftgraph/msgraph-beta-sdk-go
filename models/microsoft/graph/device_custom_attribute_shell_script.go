package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCustomAttributeShellScript 
type DeviceCustomAttributeShellScript struct {
    Entity
    // The list of group assignments for the device management script.
    assignments []DeviceManagementScriptAssignment;
    // The date and time the device management script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of the custom attribute.
    customAttributeName *string;
    // The expected type of the custom attribute's value. Possible values are: integer, string, dateTime.
    customAttributeType *DeviceCustomAttributeValueType;
    // Optional description for the device management script.
    description *string;
    // List of run states for this script across all devices.
    deviceRunStates []DeviceManagementScriptDeviceState;
    // Name of the device management script.
    displayName *string;
    // Script file name.
    fileName *string;
    // The list of group assignments for the device management script.
    groupAssignments []DeviceManagementScriptGroupAssignment;
    // The date and time the device management script was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tag IDs for this PowerShellScript instance.
    roleScopeTagIds []string;
    // Indicates the type of execution context. Possible values are: system, user.
    runAsAccount *RunAsAccountType;
    // Run summary for device management script.
    runSummary *DeviceManagementScriptRunSummary;
    // The script content.
    scriptContent []byte;
    // List of run states for this script across all users.
    userRunStates []DeviceManagementScriptUserState;
}
// NewDeviceCustomAttributeShellScript instantiates a new deviceCustomAttributeShellScript and sets the default values.
func NewDeviceCustomAttributeShellScript()(*DeviceCustomAttributeShellScript) {
    m := &DeviceCustomAttributeShellScript{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) GetAssignments()([]DeviceManagementScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceCustomAttributeShellScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustomAttributeName gets the customAttributeName property value. The name of the custom attribute.
func (m *DeviceCustomAttributeShellScript) GetCustomAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customAttributeName
    }
}
// GetCustomAttributeType gets the customAttributeType property value. The expected type of the custom attribute's value. Possible values are: integer, string, dateTime.
func (m *DeviceCustomAttributeShellScript) GetCustomAttributeType()(*DeviceCustomAttributeValueType) {
    if m == nil {
        return nil
    } else {
        return m.customAttributeType
    }
}
// GetDescription gets the description property value. Optional description for the device management script.
func (m *DeviceCustomAttributeShellScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceCustomAttributeShellScript) GetDeviceRunStates()([]DeviceManagementScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// GetDisplayName gets the displayName property value. Name of the device management script.
func (m *DeviceCustomAttributeShellScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFileName gets the fileName property value. Script file name.
func (m *DeviceCustomAttributeShellScript) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetGroupAssignments gets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) GetGroupAssignments()([]DeviceManagementScriptGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceCustomAttributeShellScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceCustomAttributeShellScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceCustomAttributeShellScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// GetRunSummary gets the runSummary property value. Run summary for device management script.
func (m *DeviceCustomAttributeShellScript) GetRunSummary()(*DeviceManagementScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// GetScriptContent gets the scriptContent property value. The script content.
func (m *DeviceCustomAttributeShellScript) GetScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.scriptContent
    }
}
// GetUserRunStates gets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceCustomAttributeShellScript) GetUserRunStates()([]DeviceManagementScriptUserState) {
    if m == nil {
        return nil
    } else {
        return m.userRunStates
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCustomAttributeShellScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementScriptAssignment))
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
    res["customAttributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomAttributeName(val)
        }
        return nil
    }
    res["customAttributeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceCustomAttributeValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomAttributeType(val.(*DeviceCustomAttributeValueType))
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
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementScriptDeviceState))
            }
            m.SetDeviceRunStates(res)
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
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptGroupAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptGroupAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementScriptGroupAssignment))
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
    res["runAsAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptRunSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(*DeviceManagementScriptRunSummary))
        }
        return nil
    }
    res["scriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptContent(val)
        }
        return nil
    }
    res["userRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptUserState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementScriptUserState, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementScriptUserState))
            }
            m.SetUserRunStates(res)
        }
        return nil
    }
    return res
}
func (m *DeviceCustomAttributeShellScript) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceCustomAttributeShellScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
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
        err = writer.WriteStringValue("customAttributeName", m.GetCustomAttributeName())
        if err != nil {
            return err
        }
    }
    if m.GetCustomAttributeType() != nil {
        cast := (*m.GetCustomAttributeType()).String()
        err = writer.WriteStringValue("customAttributeType", &cast)
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
    if m.GetDeviceRunStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceRunStates", cast)
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
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupAssignments() != nil {
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
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("runSummary", m.GetRunSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("scriptContent", m.GetScriptContent())
        if err != nil {
            return err
        }
    }
    if m.GetUserRunStates() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRunStates()))
        for i, v := range m.GetUserRunStates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userRunStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) SetAssignments(value []DeviceManagementScriptAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceCustomAttributeShellScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustomAttributeName sets the customAttributeName property value. The name of the custom attribute.
func (m *DeviceCustomAttributeShellScript) SetCustomAttributeName(value *string)() {
    if m != nil {
        m.customAttributeName = value
    }
}
// SetCustomAttributeType sets the customAttributeType property value. The expected type of the custom attribute's value. Possible values are: integer, string, dateTime.
func (m *DeviceCustomAttributeShellScript) SetCustomAttributeType(value *DeviceCustomAttributeValueType)() {
    if m != nil {
        m.customAttributeType = value
    }
}
// SetDescription sets the description property value. Optional description for the device management script.
func (m *DeviceCustomAttributeShellScript) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceCustomAttributeShellScript) SetDeviceRunStates(value []DeviceManagementScriptDeviceState)() {
    if m != nil {
        m.deviceRunStates = value
    }
}
// SetDisplayName sets the displayName property value. Name of the device management script.
func (m *DeviceCustomAttributeShellScript) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFileName sets the fileName property value. Script file name.
func (m *DeviceCustomAttributeShellScript) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetGroupAssignments sets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceCustomAttributeShellScript) SetGroupAssignments(value []DeviceManagementScriptGroupAssignment)() {
    if m != nil {
        m.groupAssignments = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceCustomAttributeShellScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceCustomAttributeShellScript) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceCustomAttributeShellScript) SetRunAsAccount(value *RunAsAccountType)() {
    if m != nil {
        m.runAsAccount = value
    }
}
// SetRunSummary sets the runSummary property value. Run summary for device management script.
func (m *DeviceCustomAttributeShellScript) SetRunSummary(value *DeviceManagementScriptRunSummary)() {
    if m != nil {
        m.runSummary = value
    }
}
// SetScriptContent sets the scriptContent property value. The script content.
func (m *DeviceCustomAttributeShellScript) SetScriptContent(value []byte)() {
    if m != nil {
        m.scriptContent = value
    }
}
// SetUserRunStates sets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceCustomAttributeShellScript) SetUserRunStates(value []DeviceManagementScriptUserState)() {
    if m != nil {
        m.userRunStates = value
    }
}
