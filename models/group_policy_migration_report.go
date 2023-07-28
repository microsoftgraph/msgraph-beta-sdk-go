package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyMigrationReport the Group Policy migration report.
type GroupPolicyMigrationReport struct {
    Entity
}
// NewGroupPolicyMigrationReport instantiates a new groupPolicyMigrationReport and sets the default values.
func NewGroupPolicyMigrationReport()(*GroupPolicyMigrationReport) {
    m := &GroupPolicyMigrationReport{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyMigrationReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyMigrationReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyMigrationReport(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) GetDisplayName()(*string) {
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
func (m *GroupPolicyMigrationReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["groupPolicyCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyCreatedDateTime(val)
        }
        return nil
    }
    res["groupPolicyLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLastModifiedDateTime(val)
        }
        return nil
    }
    res["groupPolicyObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectId(val)
        }
        return nil
    }
    res["groupPolicySettingMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicySettingMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicySettingMappingable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupPolicySettingMappingable)
                }
            }
            m.SetGroupPolicySettingMappings(res)
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
    res["migrationReadiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyMigrationReadiness)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMigrationReadiness(val.(*GroupPolicyMigrationReadiness))
        }
        return nil
    }
    res["ouDistinguishedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuDistinguishedName(val)
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
    res["supportedSettingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSettingsCount(val)
        }
        return nil
    }
    res["supportedSettingsPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSettingsPercent(val)
        }
        return nil
    }
    res["targetedInActiveDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedInActiveDirectory(val)
        }
        return nil
    }
    res["totalSettingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSettingsCount(val)
        }
        return nil
    }
    res["unsupportedGroupPolicyExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnsupportedGroupPolicyExtensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UnsupportedGroupPolicyExtensionable)
                }
            }
            m.SetUnsupportedGroupPolicyExtensions(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyCreatedDateTime gets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetGroupPolicyCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("groupPolicyCreatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetGroupPolicyLastModifiedDateTime gets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetGroupPolicyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("groupPolicyLastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetGroupPolicyObjectId gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) GetGroupPolicyObjectId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("groupPolicyObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetGroupPolicySettingMappings gets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) GetGroupPolicySettingMappings()([]GroupPolicySettingMappingable) {
    val, err := m.GetBackingStore().Get("groupPolicySettingMappings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]GroupPolicySettingMappingable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMigrationReadiness gets the migrationReadiness property value. Indicates if the Group Policy Object file is covered and ready for Intune migration.
func (m *GroupPolicyMigrationReport) GetMigrationReadiness()(*GroupPolicyMigrationReadiness) {
    val, err := m.GetBackingStore().Get("migrationReadiness")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicyMigrationReadiness)
    }
    return nil
}
// GetOuDistinguishedName gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) GetOuDistinguishedName()(*string) {
    val, err := m.GetBackingStore().Get("ouDistinguishedName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyMigrationReport) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSupportedSettingsCount gets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsCount()(*int32) {
    val, err := m.GetBackingStore().Get("supportedSettingsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSupportedSettingsPercent gets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsPercent()(*int32) {
    val, err := m.GetBackingStore().Get("supportedSettingsPercent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTargetedInActiveDirectory gets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) GetTargetedInActiveDirectory()(*bool) {
    val, err := m.GetBackingStore().Get("targetedInActiveDirectory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTotalSettingsCount gets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) GetTotalSettingsCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalSettingsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnsupportedGroupPolicyExtensions gets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) GetUnsupportedGroupPolicyExtensions()([]UnsupportedGroupPolicyExtensionable) {
    val, err := m.GetBackingStore().Get("unsupportedGroupPolicyExtensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UnsupportedGroupPolicyExtensionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicyMigrationReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteTimeValue("groupPolicyCreatedDateTime", m.GetGroupPolicyCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("groupPolicyLastModifiedDateTime", m.GetGroupPolicyLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("groupPolicyObjectId", m.GetGroupPolicyObjectId())
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicySettingMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicySettingMappings()))
        for i, v := range m.GetGroupPolicySettingMappings() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicySettingMappings", cast)
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
    if m.GetMigrationReadiness() != nil {
        cast := (*m.GetMigrationReadiness()).String()
        err = writer.WriteStringValue("migrationReadiness", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ouDistinguishedName", m.GetOuDistinguishedName())
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
        err = writer.WriteInt32Value("supportedSettingsCount", m.GetSupportedSettingsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supportedSettingsPercent", m.GetSupportedSettingsPercent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("targetedInActiveDirectory", m.GetTargetedInActiveDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalSettingsCount", m.GetTotalSettingsCount())
        if err != nil {
            return err
        }
    }
    if m.GetUnsupportedGroupPolicyExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnsupportedGroupPolicyExtensions()))
        for i, v := range m.GetUnsupportedGroupPolicyExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("unsupportedGroupPolicyExtensions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyCreatedDateTime sets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) SetGroupPolicyCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("groupPolicyCreatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyLastModifiedDateTime sets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) SetGroupPolicyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("groupPolicyLastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicyObjectId sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) SetGroupPolicyObjectId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("groupPolicyObjectId", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupPolicySettingMappings sets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) SetGroupPolicySettingMappings(value []GroupPolicySettingMappingable)() {
    err := m.GetBackingStore().Set("groupPolicySettingMappings", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMigrationReadiness sets the migrationReadiness property value. Indicates if the Group Policy Object file is covered and ready for Intune migration.
func (m *GroupPolicyMigrationReport) SetMigrationReadiness(value *GroupPolicyMigrationReadiness)() {
    err := m.GetBackingStore().Set("migrationReadiness", value)
    if err != nil {
        panic(err)
    }
}
// SetOuDistinguishedName sets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) SetOuDistinguishedName(value *string)() {
    err := m.GetBackingStore().Set("ouDistinguishedName", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyMigrationReport) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedSettingsCount sets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsCount(value *int32)() {
    err := m.GetBackingStore().Set("supportedSettingsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedSettingsPercent sets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsPercent(value *int32)() {
    err := m.GetBackingStore().Set("supportedSettingsPercent", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedInActiveDirectory sets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) SetTargetedInActiveDirectory(value *bool)() {
    err := m.GetBackingStore().Set("targetedInActiveDirectory", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalSettingsCount sets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) SetTotalSettingsCount(value *int32)() {
    err := m.GetBackingStore().Set("totalSettingsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUnsupportedGroupPolicyExtensions sets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) SetUnsupportedGroupPolicyExtensions(value []UnsupportedGroupPolicyExtensionable)() {
    err := m.GetBackingStore().Set("unsupportedGroupPolicyExtensions", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicyMigrationReportable 
type GroupPolicyMigrationReportable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetGroupPolicyCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupPolicyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupPolicyObjectId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetGroupPolicySettingMappings()([]GroupPolicySettingMappingable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMigrationReadiness()(*GroupPolicyMigrationReadiness)
    GetOuDistinguishedName()(*string)
    GetRoleScopeTagIds()([]string)
    GetSupportedSettingsCount()(*int32)
    GetSupportedSettingsPercent()(*int32)
    GetTargetedInActiveDirectory()(*bool)
    GetTotalSettingsCount()(*int32)
    GetUnsupportedGroupPolicyExtensions()([]UnsupportedGroupPolicyExtensionable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetGroupPolicyCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupPolicyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupPolicyObjectId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetGroupPolicySettingMappings(value []GroupPolicySettingMappingable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMigrationReadiness(value *GroupPolicyMigrationReadiness)()
    SetOuDistinguishedName(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetSupportedSettingsCount(value *int32)()
    SetSupportedSettingsPercent(value *int32)()
    SetTargetedInActiveDirectory(value *bool)()
    SetTotalSettingsCount(value *int32)()
    SetUnsupportedGroupPolicyExtensions(value []UnsupportedGroupPolicyExtensionable)()
}
