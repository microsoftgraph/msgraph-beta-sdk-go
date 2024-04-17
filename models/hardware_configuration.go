package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HardwareConfiguration bIOS configuration and other settings provides customers the ability to configure hardware/bios settings on the enrolled Windows 10/11 Entra ID joined devices by uploading a configuration file generated with their OEM tool (e.g. Dell Command tool). A BIOS configuration policy can be assigned to multiple devices, allowing admins to remotely control a device's hardware properties (e.g. enable Secure Boot) from the Intune Portal.
type HardwareConfiguration struct {
    Entity
}
// NewHardwareConfiguration instantiates a new HardwareConfiguration and sets the default values.
func NewHardwareConfiguration()(*HardwareConfiguration) {
    m := &HardwareConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHardwareConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareConfiguration(), nil
}
// GetAssignments gets the assignments property value. List of the Azure AD user group ids that hardware configuration will be applied to. Only security groups and Office 365 Groups are supported.
// returns a []HardwareConfigurationAssignmentable when successful
func (m *HardwareConfiguration) GetAssignments()([]HardwareConfigurationAssignmentable) {
    val, err := m.GetBackingStore().Get("assignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HardwareConfigurationAssignmentable)
    }
    return nil
}
// GetConfigurationFileContent gets the configurationFileContent property value. File content of the hardware configuration
// returns a []byte when successful
func (m *HardwareConfiguration) GetConfigurationFileContent()([]byte) {
    val, err := m.GetBackingStore().Get("configurationFileContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the hardware configuration was created. This property is read-only.
// returns a *Time when successful
func (m *HardwareConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. Description of the hardware configuration
// returns a *string when successful
func (m *HardwareConfiguration) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceRunStates gets the deviceRunStates property value. List of run states for the hardware configuration across all devices
// returns a []HardwareConfigurationDeviceStateable when successful
func (m *HardwareConfiguration) GetDeviceRunStates()([]HardwareConfigurationDeviceStateable) {
    val, err := m.GetBackingStore().Get("deviceRunStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HardwareConfigurationDeviceStateable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Name of the hardware configuration
// returns a *string when successful
func (m *HardwareConfiguration) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HardwareConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHardwareConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareConfigurationAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HardwareConfigurationAssignmentable)
                }
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["configurationFileContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationFileContent(val)
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
    res["deviceRunStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHardwareConfigurationDeviceStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareConfigurationDeviceStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HardwareConfigurationDeviceStateable)
                }
            }
            m.SetDeviceRunStates(res)
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
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["hardwareConfigurationFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHardwareConfigurationFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareConfigurationFormat(val.(*HardwareConfigurationFormat))
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
    res["perDevicePasswordDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPerDevicePasswordDisabled(val)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["runSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHardwareConfigurationRunSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunSummary(val.(HardwareConfigurationRunSummaryable))
        }
        return nil
    }
    res["userRunStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHardwareConfigurationUserStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HardwareConfigurationUserStateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(HardwareConfigurationUserStateable)
                }
            }
            m.SetUserRunStates(res)
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
// GetFileName gets the fileName property value. File name of the hardware configuration
// returns a *string when successful
func (m *HardwareConfiguration) GetFileName()(*string) {
    val, err := m.GetBackingStore().Get("fileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetHardwareConfigurationFormat gets the hardwareConfigurationFormat property value. Indicates the supported oems of hardware configuration
// returns a *HardwareConfigurationFormat when successful
func (m *HardwareConfiguration) GetHardwareConfigurationFormat()(*HardwareConfigurationFormat) {
    val, err := m.GetBackingStore().Get("hardwareConfigurationFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*HardwareConfigurationFormat)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp of when the hardware configuration was modified. This property is read-only.
// returns a *Time when successful
func (m *HardwareConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPerDevicePasswordDisabled gets the perDevicePasswordDisabled property value. A value indicating whether per devcive pasword disabled
// returns a *bool when successful
func (m *HardwareConfiguration) GetPerDevicePasswordDisabled()(*bool) {
    val, err := m.GetBackingStore().Get("perDevicePasswordDisabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for the hardware configuration
// returns a []string when successful
func (m *HardwareConfiguration) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRunSummary gets the runSummary property value. A summary of the results from an attempt to configure hardware settings
// returns a HardwareConfigurationRunSummaryable when successful
func (m *HardwareConfiguration) GetRunSummary()(HardwareConfigurationRunSummaryable) {
    val, err := m.GetBackingStore().Get("runSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(HardwareConfigurationRunSummaryable)
    }
    return nil
}
// GetUserRunStates gets the userRunStates property value. List of run states for the hardware configuration across all users
// returns a []HardwareConfigurationUserStateable when successful
func (m *HardwareConfiguration) GetUserRunStates()([]HardwareConfigurationUserStateable) {
    val, err := m.GetBackingStore().Get("userRunStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]HardwareConfigurationUserStateable)
    }
    return nil
}
// GetVersion gets the version property value. Version of the hardware configuration (E.g. 1, 2, 3 ...)
// returns a *int32 when successful
func (m *HardwareConfiguration) GetVersion()(*int32) {
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
func (m *HardwareConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRunStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceRunStates()))
        for i, v := range m.GetDeviceRunStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
        cast := (*m.GetHardwareConfigurationFormat()).String()
        err = writer.WriteStringValue("hardwareConfigurationFormat", &cast)
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
    if m.GetRoleScopeTagIds() != nil {
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
    if m.GetUserRunStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserRunStates()))
        for i, v := range m.GetUserRunStates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAssignments sets the assignments property value. List of the Azure AD user group ids that hardware configuration will be applied to. Only security groups and Office 365 Groups are supported.
func (m *HardwareConfiguration) SetAssignments(value []HardwareConfigurationAssignmentable)() {
    err := m.GetBackingStore().Set("assignments", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationFileContent sets the configurationFileContent property value. File content of the hardware configuration
func (m *HardwareConfiguration) SetConfigurationFileContent(value []byte)() {
    err := m.GetBackingStore().Set("configurationFileContent", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the hardware configuration was created. This property is read-only.
func (m *HardwareConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the hardware configuration
func (m *HardwareConfiguration) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceRunStates sets the deviceRunStates property value. List of run states for the hardware configuration across all devices
func (m *HardwareConfiguration) SetDeviceRunStates(value []HardwareConfigurationDeviceStateable)() {
    err := m.GetBackingStore().Set("deviceRunStates", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Name of the hardware configuration
func (m *HardwareConfiguration) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetFileName sets the fileName property value. File name of the hardware configuration
func (m *HardwareConfiguration) SetFileName(value *string)() {
    err := m.GetBackingStore().Set("fileName", value)
    if err != nil {
        panic(err)
    }
}
// SetHardwareConfigurationFormat sets the hardwareConfigurationFormat property value. Indicates the supported oems of hardware configuration
func (m *HardwareConfiguration) SetHardwareConfigurationFormat(value *HardwareConfigurationFormat)() {
    err := m.GetBackingStore().Set("hardwareConfigurationFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp of when the hardware configuration was modified. This property is read-only.
func (m *HardwareConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPerDevicePasswordDisabled sets the perDevicePasswordDisabled property value. A value indicating whether per devcive pasword disabled
func (m *HardwareConfiguration) SetPerDevicePasswordDisabled(value *bool)() {
    err := m.GetBackingStore().Set("perDevicePasswordDisabled", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for the hardware configuration
func (m *HardwareConfiguration) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRunSummary sets the runSummary property value. A summary of the results from an attempt to configure hardware settings
func (m *HardwareConfiguration) SetRunSummary(value HardwareConfigurationRunSummaryable)() {
    err := m.GetBackingStore().Set("runSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetUserRunStates sets the userRunStates property value. List of run states for the hardware configuration across all users
func (m *HardwareConfiguration) SetUserRunStates(value []HardwareConfigurationUserStateable)() {
    err := m.GetBackingStore().Set("userRunStates", value)
    if err != nil {
        panic(err)
    }
}
// SetVersion sets the version property value. Version of the hardware configuration (E.g. 1, 2, 3 ...)
func (m *HardwareConfiguration) SetVersion(value *int32)() {
    err := m.GetBackingStore().Set("version", value)
    if err != nil {
        panic(err)
    }
}
type HardwareConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]HardwareConfigurationAssignmentable)
    GetConfigurationFileContent()([]byte)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceRunStates()([]HardwareConfigurationDeviceStateable)
    GetDisplayName()(*string)
    GetFileName()(*string)
    GetHardwareConfigurationFormat()(*HardwareConfigurationFormat)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPerDevicePasswordDisabled()(*bool)
    GetRoleScopeTagIds()([]string)
    GetRunSummary()(HardwareConfigurationRunSummaryable)
    GetUserRunStates()([]HardwareConfigurationUserStateable)
    GetVersion()(*int32)
    SetAssignments(value []HardwareConfigurationAssignmentable)()
    SetConfigurationFileContent(value []byte)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceRunStates(value []HardwareConfigurationDeviceStateable)()
    SetDisplayName(value *string)()
    SetFileName(value *string)()
    SetHardwareConfigurationFormat(value *HardwareConfigurationFormat)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPerDevicePasswordDisabled(value *bool)()
    SetRoleScopeTagIds(value []string)()
    SetRunSummary(value HardwareConfigurationRunSummaryable)()
    SetUserRunStates(value []HardwareConfigurationUserStateable)()
    SetVersion(value *int32)()
}
