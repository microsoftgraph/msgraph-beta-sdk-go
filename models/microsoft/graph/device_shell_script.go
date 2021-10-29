package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceShellScript struct {
    Entity
    // The list of group assignments for the device management script.
    assignments []DeviceManagementScriptAssignment;
    // Does not notify the user a script is being executed
    blockExecutionNotifications *bool;
    // The date and time the device management script was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional description for the device management script.
    description *string;
    // List of run states for this script across all devices.
    deviceRunStates []DeviceManagementScriptDeviceState;
    // Name of the device management script.
    displayName *string;
    // The interval for script to run. If not defined the script will run once
    executionFrequency *string;
    // Script file name.
    fileName *string;
    // The list of group assignments for the device management script.
    groupAssignments []DeviceManagementScriptGroupAssignment;
    // The date and time the device management script was last modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of times for the script to be retried if it fails
    retryCount *int32;
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
// Instantiates a new deviceShellScript and sets the default values.
func NewDeviceShellScript()(*DeviceShellScript) {
    m := &DeviceShellScript{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the device management script.
func (m *DeviceShellScript) GetAssignments()([]DeviceManagementScriptAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the blockExecutionNotifications property value. Does not notify the user a script is being executed
func (m *DeviceShellScript) GetBlockExecutionNotifications()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.blockExecutionNotifications
    }
}
// Gets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
func (m *DeviceShellScript) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Optional description for the device management script.
func (m *DeviceShellScript) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceRunStates property value. List of run states for this script across all devices.
func (m *DeviceShellScript) GetDeviceRunStates()([]DeviceManagementScriptDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// Gets the displayName property value. Name of the device management script.
func (m *DeviceShellScript) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the executionFrequency property value. The interval for script to run. If not defined the script will run once
func (m *DeviceShellScript) GetExecutionFrequency()(*string) {
    if m == nil {
        return nil
    } else {
        return m.executionFrequency
    }
}
// Gets the fileName property value. Script file name.
func (m *DeviceShellScript) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Gets the groupAssignments property value. The list of group assignments for the device management script.
func (m *DeviceShellScript) GetGroupAssignments()([]DeviceManagementScriptGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// Gets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
func (m *DeviceShellScript) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the retryCount property value. Number of times for the script to be retried if it fails
func (m *DeviceShellScript) GetRetryCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.retryCount
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
func (m *DeviceShellScript) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
func (m *DeviceShellScript) GetRunAsAccount()(*RunAsAccountType) {
    if m == nil {
        return nil
    } else {
        return m.runAsAccount
    }
}
// Gets the runSummary property value. Run summary for device management script.
func (m *DeviceShellScript) GetRunSummary()(*DeviceManagementScriptRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// Gets the scriptContent property value. The script content.
func (m *DeviceShellScript) GetScriptContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.scriptContent
    }
}
// Gets the userRunStates property value. List of run states for this script across all users.
func (m *DeviceShellScript) GetUserRunStates()([]DeviceManagementScriptUserState) {
    if m == nil {
        return nil
    } else {
        return m.userRunStates
    }
}
// The deserialization information for the current model
func (m *DeviceShellScript) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["blockExecutionNotifications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetBlockExecutionNotifications(val)
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
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptDeviceState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptDeviceState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptDeviceState))
        }
        m.SetDeviceRunStates(res)
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
    res["executionFrequency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExecutionFrequency(val)
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptGroupAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptGroupAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptGroupAssignment))
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
    res["retryCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRetryCount(val)
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
    res["runAsAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        cast := val.(RunAsAccountType)
        m.SetRunAsAccount(&cast)
        return nil
    }
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptRunSummary() })
        if err != nil {
            return err
        }
        m.SetRunSummary(val.(*DeviceManagementScriptRunSummary))
        return nil
    }
    res["scriptContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetScriptContent(val)
        return nil
    }
    res["userRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementScriptUserState() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementScriptUserState, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementScriptUserState))
        }
        m.SetUserRunStates(res)
        return nil
    }
    return res
}
func (m *DeviceShellScript) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceShellScript) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("blockExecutionNotifications", m.GetBlockExecutionNotifications())
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
        err = writer.WriteStringValue("executionFrequency", m.GetExecutionFrequency())
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
        err = writer.WriteInt32Value("retryCount", m.GetRetryCount())
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
    if m.GetRunAsAccount() != nil {
        cast := m.GetRunAsAccount().String()
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
    {
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
// Sets the assignments property value. The list of group assignments for the device management script.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceShellScript) SetAssignments(value []DeviceManagementScriptAssignment)() {
    m.assignments = value
}
// Sets the blockExecutionNotifications property value. Does not notify the user a script is being executed
// Parameters:
//  - value : Value to set for the blockExecutionNotifications property.
func (m *DeviceShellScript) SetBlockExecutionNotifications(value *bool)() {
    m.blockExecutionNotifications = value
}
// Sets the createdDateTime property value. The date and time the device management script was created. This property is read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceShellScript) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Optional description for the device management script.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceShellScript) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceRunStates property value. List of run states for this script across all devices.
// Parameters:
//  - value : Value to set for the deviceRunStates property.
func (m *DeviceShellScript) SetDeviceRunStates(value []DeviceManagementScriptDeviceState)() {
    m.deviceRunStates = value
}
// Sets the displayName property value. Name of the device management script.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceShellScript) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the executionFrequency property value. The interval for script to run. If not defined the script will run once
// Parameters:
//  - value : Value to set for the executionFrequency property.
func (m *DeviceShellScript) SetExecutionFrequency(value *string)() {
    m.executionFrequency = value
}
// Sets the fileName property value. Script file name.
// Parameters:
//  - value : Value to set for the fileName property.
func (m *DeviceShellScript) SetFileName(value *string)() {
    m.fileName = value
}
// Sets the groupAssignments property value. The list of group assignments for the device management script.
// Parameters:
//  - value : Value to set for the groupAssignments property.
func (m *DeviceShellScript) SetGroupAssignments(value []DeviceManagementScriptGroupAssignment)() {
    m.groupAssignments = value
}
// Sets the lastModifiedDateTime property value. The date and time the device management script was last modified. This property is read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceShellScript) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the retryCount property value. Number of times for the script to be retried if it fails
// Parameters:
//  - value : Value to set for the retryCount property.
func (m *DeviceShellScript) SetRetryCount(value *int32)() {
    m.retryCount = value
}
// Sets the roleScopeTagIds property value. List of Scope Tag IDs for this PowerShellScript instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceShellScript) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the runAsAccount property value. Indicates the type of execution context. Possible values are: system, user.
// Parameters:
//  - value : Value to set for the runAsAccount property.
func (m *DeviceShellScript) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// Sets the runSummary property value. Run summary for device management script.
// Parameters:
//  - value : Value to set for the runSummary property.
func (m *DeviceShellScript) SetRunSummary(value *DeviceManagementScriptRunSummary)() {
    m.runSummary = value
}
// Sets the scriptContent property value. The script content.
// Parameters:
//  - value : Value to set for the scriptContent property.
func (m *DeviceShellScript) SetScriptContent(value []byte)() {
    m.scriptContent = value
}
// Sets the userRunStates property value. List of run states for this script across all users.
// Parameters:
//  - value : Value to set for the userRunStates property.
func (m *DeviceShellScript) SetUserRunStates(value []DeviceManagementScriptUserState)() {
    m.userRunStates = value
}
