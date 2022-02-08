package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyMigrationReport 
type GroupPolicyMigrationReport struct {
    Entity
    // The date and time at which the GroupPolicyMigrationReport was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of Group Policy Object from the GPO Xml Content
    displayName *string;
    // The date and time at which the GroupPolicyMigrationReport was created.
    groupPolicyCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time at which the GroupPolicyMigrationReport was last modified.
    groupPolicyLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Group Policy Object GUID from GPO Xml content
    groupPolicyObjectId *string;
    // A list of group policy settings to MDM/Intune mappings.
    groupPolicySettingMappings []GroupPolicySettingMapping;
    // The date and time at which the GroupPolicyMigrationReport was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Intune coverage for the associated Group Policy Object file. Possible values are: none, partial, complete, error, notApplicable.
    migrationReadiness *GroupPolicyMigrationReadiness;
    // The distinguished name of the OU.
    ouDistinguishedName *string;
    // The number of Group Policy Settings supported by Intune.
    supportedSettingsCount *int32;
    // The Percentage of Group Policy Settings supported by Intune.
    supportedSettingsPercent *int32;
    // The Targeted in AD property from GPO Xml Content
    targetedInActiveDirectory *bool;
    // The total number of Group Policy Settings from GPO file.
    totalSettingsCount *int32;
    // A list of unsupported group policy extensions inside the Group Policy Object.
    unsupportedGroupPolicyExtensions []UnsupportedGroupPolicyExtension;
}
// NewGroupPolicyMigrationReport instantiates a new groupPolicyMigrationReport and sets the default values.
func NewGroupPolicyMigrationReport()(*GroupPolicyMigrationReport) {
    m := &GroupPolicyMigrationReport{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGroupPolicyCreatedDateTime gets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetGroupPolicyCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCreatedDateTime
    }
}
// GetGroupPolicyLastModifiedDateTime gets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetGroupPolicyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyLastModifiedDateTime
    }
}
// GetGroupPolicyObjectId gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) GetGroupPolicyObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectId
    }
}
// GetGroupPolicySettingMappings gets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) GetGroupPolicySettingMappings()([]GroupPolicySettingMapping) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicySettingMappings
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetMigrationReadiness gets the migrationReadiness property value. The Intune coverage for the associated Group Policy Object file. Possible values are: none, partial, complete, error, notApplicable.
func (m *GroupPolicyMigrationReport) GetMigrationReadiness()(*GroupPolicyMigrationReadiness) {
    if m == nil {
        return nil
    } else {
        return m.migrationReadiness
    }
}
// GetOuDistinguishedName gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) GetOuDistinguishedName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ouDistinguishedName
    }
}
// GetSupportedSettingsCount gets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supportedSettingsCount
    }
}
// GetSupportedSettingsPercent gets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsPercent()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supportedSettingsPercent
    }
}
// GetTargetedInActiveDirectory gets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) GetTargetedInActiveDirectory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.targetedInActiveDirectory
    }
}
// GetTotalSettingsCount gets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) GetTotalSettingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalSettingsCount
    }
}
// GetUnsupportedGroupPolicyExtensions gets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) GetUnsupportedGroupPolicyExtensions()([]UnsupportedGroupPolicyExtension) {
    if m == nil {
        return nil
    } else {
        return m.unsupportedGroupPolicyExtensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyMigrationReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["groupPolicyCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyCreatedDateTime(val)
        }
        return nil
    }
    res["groupPolicyLastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLastModifiedDateTime(val)
        }
        return nil
    }
    res["groupPolicyObjectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectId(val)
        }
        return nil
    }
    res["groupPolicySettingMappings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicySettingMapping() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicySettingMapping, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupPolicySettingMapping))
            }
            m.SetGroupPolicySettingMappings(res)
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
    res["migrationReadiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyMigrationReadiness)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMigrationReadiness(val.(*GroupPolicyMigrationReadiness))
        }
        return nil
    }
    res["ouDistinguishedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuDistinguishedName(val)
        }
        return nil
    }
    res["supportedSettingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSettingsCount(val)
        }
        return nil
    }
    res["supportedSettingsPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSettingsPercent(val)
        }
        return nil
    }
    res["targetedInActiveDirectory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedInActiveDirectory(val)
        }
        return nil
    }
    res["totalSettingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSettingsCount(val)
        }
        return nil
    }
    res["unsupportedGroupPolicyExtensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnsupportedGroupPolicyExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnsupportedGroupPolicyExtension, len(val))
            for i, v := range val {
                res[i] = *(v.(*UnsupportedGroupPolicyExtension))
            }
            m.SetUnsupportedGroupPolicyExtensions(res)
        }
        return nil
    }
    return res
}
func (m *GroupPolicyMigrationReport) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicyMigrationReport) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("groupPolicyObjectId", m.GetGroupPolicyObjectId())
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicySettingMappings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupPolicySettingMappings()))
        for i, v := range m.GetGroupPolicySettingMappings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUnsupportedGroupPolicyExtensions()))
        for i, v := range m.GetUnsupportedGroupPolicyExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGroupPolicyCreatedDateTime sets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) SetGroupPolicyCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.groupPolicyCreatedDateTime = value
    }
}
// SetGroupPolicyLastModifiedDateTime sets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) SetGroupPolicyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.groupPolicyLastModifiedDateTime = value
    }
}
// SetGroupPolicyObjectId sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) SetGroupPolicyObjectId(value *string)() {
    if m != nil {
        m.groupPolicyObjectId = value
    }
}
// SetGroupPolicySettingMappings sets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) SetGroupPolicySettingMappings(value []GroupPolicySettingMapping)() {
    if m != nil {
        m.groupPolicySettingMappings = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetMigrationReadiness sets the migrationReadiness property value. The Intune coverage for the associated Group Policy Object file. Possible values are: none, partial, complete, error, notApplicable.
func (m *GroupPolicyMigrationReport) SetMigrationReadiness(value *GroupPolicyMigrationReadiness)() {
    if m != nil {
        m.migrationReadiness = value
    }
}
// SetOuDistinguishedName sets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) SetOuDistinguishedName(value *string)() {
    if m != nil {
        m.ouDistinguishedName = value
    }
}
// SetSupportedSettingsCount sets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsCount(value *int32)() {
    if m != nil {
        m.supportedSettingsCount = value
    }
}
// SetSupportedSettingsPercent sets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsPercent(value *int32)() {
    if m != nil {
        m.supportedSettingsPercent = value
    }
}
// SetTargetedInActiveDirectory sets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) SetTargetedInActiveDirectory(value *bool)() {
    if m != nil {
        m.targetedInActiveDirectory = value
    }
}
// SetTotalSettingsCount sets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) SetTotalSettingsCount(value *int32)() {
    if m != nil {
        m.totalSettingsCount = value
    }
}
// SetUnsupportedGroupPolicyExtensions sets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) SetUnsupportedGroupPolicyExtensions(value []UnsupportedGroupPolicyExtension)() {
    if m != nil {
        m.unsupportedGroupPolicyExtensions = value
    }
}
