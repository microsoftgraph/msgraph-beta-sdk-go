package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HardwareConfiguration 
type HardwareConfiguration struct {
    Entity
    // List of group assignments for the hardware configuration
    assignments []HardwareConfigurationAssignment;
    // File content of the hardware configuration
    configurationFileContent []byte;
    // Timestamp of when the hardware configuration was created. This property is read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description of the hardware configuration
    description *string;
    // List of run states for the hardware configuration across all devices
    deviceRunStates []HardwareConfigurationDeviceState;
    // Name of the hardware configuration
    displayName *string;
    // File name of the hardware configuration
    fileName *string;
    // Oem type of the hardware configuration. Possible values are: dell, surface, surfaceDock.
    hardwareConfigurationFormat *HardwareConfigurationFormat;
    // Timestamp of when the hardware configuration was modified. This property is read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A value indicating whether per devcive pasword disabled
    perDevicePasswordDisabled *bool;
    // List of Scope Tag IDs for the hardware configuration
    roleScopeTagIds []string;
    // Run summary for hardware configuration
    runSummary *HardwareConfigurationRunSummary;
    // List of run states for the hardware configuration across all users
    userRunStates []HardwareConfigurationUserState;
    // Version of the hardware configuration
    version *int32;
}
// NewHardwareConfiguration instantiates a new hardwareConfiguration and sets the default values.
func NewHardwareConfiguration()(*HardwareConfiguration) {
    m := &HardwareConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. List of group assignments for the hardware configuration
func (m *HardwareConfiguration) GetAssignments()([]HardwareConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetConfigurationFileContent gets the configurationFileContent property value. File content of the hardware configuration
func (m *HardwareConfiguration) GetConfigurationFileContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.configurationFileContent
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the hardware configuration was created. This property is read-only.
func (m *HardwareConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description of the hardware configuration
func (m *HardwareConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the hardware configuration across all devices
func (m *HardwareConfiguration) GetDeviceRunStates()([]HardwareConfigurationDeviceState) {
    if m == nil {
        return nil
    } else {
        return m.deviceRunStates
    }
}
// GetDisplayName gets the displayName property value. Name of the hardware configuration
func (m *HardwareConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFileName gets the fileName property value. File name of the hardware configuration
func (m *HardwareConfiguration) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// GetHardwareConfigurationFormat gets the hardwareConfigurationFormat property value. Oem type of the hardware configuration. Possible values are: dell, surface, surfaceDock.
func (m *HardwareConfiguration) GetHardwareConfigurationFormat()(*HardwareConfigurationFormat) {
    if m == nil {
        return nil
    } else {
        return m.hardwareConfigurationFormat
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp of when the hardware configuration was modified. This property is read-only.
func (m *HardwareConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPerDevicePasswordDisabled gets the perDevicePasswordDisabled property value. A value indicating whether per devcive pasword disabled
func (m *HardwareConfiguration) GetPerDevicePasswordDisabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.perDevicePasswordDisabled
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the hardware configuration
func (m *HardwareConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetRunSummary gets the runSummary property value. Run summary for hardware configuration
func (m *HardwareConfiguration) GetRunSummary()(*HardwareConfigurationRunSummary) {
    if m == nil {
        return nil
    } else {
        return m.runSummary
    }
}
// GetUserRunStates gets the userRunStates property value. List of run states for the hardware configuration across all users
func (m *HardwareConfiguration) GetUserRunStates()([]HardwareConfigurationUserState) {
    if m == nil {
        return nil
    } else {
        return m.userRunStates
    }
}
// GetVersion gets the version property value. Version of the hardware configuration
func (m *HardwareConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwareConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHardwareConfigurationAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareConfigurationAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*HardwareConfigurationAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["configurationFileContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationFileContent(val)
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
    res["deviceRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHardwareConfigurationDeviceState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareConfigurationDeviceState, len(val))
            for i, v := range val {
                res[i] = *(v.(*HardwareConfigurationDeviceState))
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
    res["hardwareConfigurationFormat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseHardwareConfigurationFormat)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(HardwareConfigurationFormat)
            m.SetHardwareConfigurationFormat(&cast)
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
    res["perDevicePasswordDisabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerDevicePasswordDisabled(val)
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
    res["runSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHardwareConfigurationRunSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(*HardwareConfigurationRunSummary))
        }
        return nil
    }
    res["userRunStates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHardwareConfigurationUserState() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareConfigurationUserState, len(val))
            for i, v := range val {
                res[i] = *(v.(*HardwareConfigurationUserState))
            }
            m.SetUserRunStates(res)
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
func (m *HardwareConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HardwareConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteByteArrayValue("configurationFileContent", m.GetConfigurationFileContent())
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
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    if m.GetHardwareConfigurationFormat() != nil {
        cast := m.GetHardwareConfigurationFormat().String()
        err = writer.WriteStringValue("hardwareConfigurationFormat", &cast)
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
        err = writer.WriteBoolValue("perDevicePasswordDisabled", m.GetPerDevicePasswordDisabled())
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
        err = writer.WriteObjectValue("runSummary", m.GetRunSummary())
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
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. List of group assignments for the hardware configuration
func (m *HardwareConfiguration) SetAssignments(value []HardwareConfigurationAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetConfigurationFileContent sets the configurationFileContent property value. File content of the hardware configuration
func (m *HardwareConfiguration) SetConfigurationFileContent(value []byte)() {
    if m != nil {
        m.configurationFileContent = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the hardware configuration was created. This property is read-only.
func (m *HardwareConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description of the hardware configuration
func (m *HardwareConfiguration) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the hardware configuration across all devices
func (m *HardwareConfiguration) SetDeviceRunStates(value []HardwareConfigurationDeviceState)() {
    if m != nil {
        m.deviceRunStates = value
    }
}
// SetDisplayName sets the displayName property value. Name of the hardware configuration
func (m *HardwareConfiguration) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFileName sets the fileName property value. File name of the hardware configuration
func (m *HardwareConfiguration) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
// SetHardwareConfigurationFormat sets the hardwareConfigurationFormat property value. Oem type of the hardware configuration. Possible values are: dell, surface, surfaceDock.
func (m *HardwareConfiguration) SetHardwareConfigurationFormat(value *HardwareConfigurationFormat)() {
    if m != nil {
        m.hardwareConfigurationFormat = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp of when the hardware configuration was modified. This property is read-only.
func (m *HardwareConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPerDevicePasswordDisabled sets the perDevicePasswordDisabled property value. A value indicating whether per devcive pasword disabled
func (m *HardwareConfiguration) SetPerDevicePasswordDisabled(value *bool)() {
    if m != nil {
        m.perDevicePasswordDisabled = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the hardware configuration
func (m *HardwareConfiguration) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetRunSummary sets the runSummary property value. Run summary for hardware configuration
func (m *HardwareConfiguration) SetRunSummary(value *HardwareConfigurationRunSummary)() {
    if m != nil {
        m.runSummary = value
    }
}
// SetUserRunStates sets the userRunStates property value. List of run states for the hardware configuration across all users
func (m *HardwareConfiguration) SetUserRunStates(value []HardwareConfigurationUserState)() {
    if m != nil {
        m.userRunStates = value
    }
}
// SetVersion sets the version property value. Version of the hardware configuration
func (m *HardwareConfiguration) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
