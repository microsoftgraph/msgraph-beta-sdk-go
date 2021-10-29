package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new groupPolicyMigrationReport and sets the default values.
func NewGroupPolicyMigrationReport()(*GroupPolicyMigrationReport) {
    m := &GroupPolicyMigrationReport{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetGroupPolicyCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyCreatedDateTime
    }
}
// Gets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetGroupPolicyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyLastModifiedDateTime
    }
}
// Gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) GetGroupPolicyObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectId
    }
}
// Gets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) GetGroupPolicySettingMappings()([]GroupPolicySettingMapping) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicySettingMappings
    }
}
// Gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the migrationReadiness property value. The Intune coverage for the associated Group Policy Object file. Possible values are: none, partial, complete, error, notApplicable.
func (m *GroupPolicyMigrationReport) GetMigrationReadiness()(*GroupPolicyMigrationReadiness) {
    if m == nil {
        return nil
    } else {
        return m.migrationReadiness
    }
}
// Gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) GetOuDistinguishedName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ouDistinguishedName
    }
}
// Gets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supportedSettingsCount
    }
}
// Gets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsPercent()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supportedSettingsPercent
    }
}
// Gets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) GetTargetedInActiveDirectory()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.targetedInActiveDirectory
    }
}
// Gets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) GetTotalSettingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalSettingsCount
    }
}
// Gets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) GetUnsupportedGroupPolicyExtensions()([]UnsupportedGroupPolicyExtension) {
    if m == nil {
        return nil
    } else {
        return m.unsupportedGroupPolicyExtensions
    }
}
// The deserialization information for the current model
func (m *GroupPolicyMigrationReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["groupPolicyCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetGroupPolicyCreatedDateTime(val)
        return nil
    }
    res["groupPolicyLastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetGroupPolicyLastModifiedDateTime(val)
        return nil
    }
    res["groupPolicyObjectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupPolicyObjectId(val)
        return nil
    }
    res["groupPolicySettingMappings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupPolicySettingMapping() })
        if err != nil {
            return err
        }
        res := make([]GroupPolicySettingMapping, len(val))
        for i, v := range val {
            res[i] = *(v.(*GroupPolicySettingMapping))
        }
        m.SetGroupPolicySettingMappings(res)
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
    res["migrationReadiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyMigrationReadiness)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicyMigrationReadiness)
        m.SetMigrationReadiness(&cast)
        return nil
    }
    res["ouDistinguishedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOuDistinguishedName(val)
        return nil
    }
    res["supportedSettingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSupportedSettingsCount(val)
        return nil
    }
    res["supportedSettingsPercent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSupportedSettingsPercent(val)
        return nil
    }
    res["targetedInActiveDirectory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTargetedInActiveDirectory(val)
        return nil
    }
    res["totalSettingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalSettingsCount(val)
        return nil
    }
    res["unsupportedGroupPolicyExtensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUnsupportedGroupPolicyExtension() })
        if err != nil {
            return err
        }
        res := make([]UnsupportedGroupPolicyExtension, len(val))
        for i, v := range val {
            res[i] = *(v.(*UnsupportedGroupPolicyExtension))
        }
        m.SetUnsupportedGroupPolicyExtensions(res)
        return nil
    }
    return res
}
func (m *GroupPolicyMigrationReport) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
        cast := m.GetMigrationReadiness().String()
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
    {
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
// Sets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *GroupPolicyMigrationReport) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the displayName property value. The name of Group Policy Object from the GPO Xml Content
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GroupPolicyMigrationReport) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
// Parameters:
//  - value : Value to set for the groupPolicyCreatedDateTime property.
func (m *GroupPolicyMigrationReport) SetGroupPolicyCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.groupPolicyCreatedDateTime = value
}
// Sets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
// Parameters:
//  - value : Value to set for the groupPolicyLastModifiedDateTime property.
func (m *GroupPolicyMigrationReport) SetGroupPolicyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.groupPolicyLastModifiedDateTime = value
}
// Sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
// Parameters:
//  - value : Value to set for the groupPolicyObjectId property.
func (m *GroupPolicyMigrationReport) SetGroupPolicyObjectId(value *string)() {
    m.groupPolicyObjectId = value
}
// Sets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
// Parameters:
//  - value : Value to set for the groupPolicySettingMappings property.
func (m *GroupPolicyMigrationReport) SetGroupPolicySettingMappings(value []GroupPolicySettingMapping)() {
    m.groupPolicySettingMappings = value
}
// Sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *GroupPolicyMigrationReport) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the migrationReadiness property value. The Intune coverage for the associated Group Policy Object file. Possible values are: none, partial, complete, error, notApplicable.
// Parameters:
//  - value : Value to set for the migrationReadiness property.
func (m *GroupPolicyMigrationReport) SetMigrationReadiness(value *GroupPolicyMigrationReadiness)() {
    m.migrationReadiness = value
}
// Sets the ouDistinguishedName property value. The distinguished name of the OU.
// Parameters:
//  - value : Value to set for the ouDistinguishedName property.
func (m *GroupPolicyMigrationReport) SetOuDistinguishedName(value *string)() {
    m.ouDistinguishedName = value
}
// Sets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
// Parameters:
//  - value : Value to set for the supportedSettingsCount property.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsCount(value *int32)() {
    m.supportedSettingsCount = value
}
// Sets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
// Parameters:
//  - value : Value to set for the supportedSettingsPercent property.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsPercent(value *int32)() {
    m.supportedSettingsPercent = value
}
// Sets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
// Parameters:
//  - value : Value to set for the targetedInActiveDirectory property.
func (m *GroupPolicyMigrationReport) SetTargetedInActiveDirectory(value *bool)() {
    m.targetedInActiveDirectory = value
}
// Sets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
// Parameters:
//  - value : Value to set for the totalSettingsCount property.
func (m *GroupPolicyMigrationReport) SetTotalSettingsCount(value *int32)() {
    m.totalSettingsCount = value
}
// Sets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
// Parameters:
//  - value : Value to set for the unsupportedGroupPolicyExtensions property.
func (m *GroupPolicyMigrationReport) SetUnsupportedGroupPolicyExtensions(value []UnsupportedGroupPolicyExtension)() {
    m.unsupportedGroupPolicyExtensions = value
}
