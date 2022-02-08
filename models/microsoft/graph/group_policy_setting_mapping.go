package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicySettingMapping 
type GroupPolicySettingMapping struct {
    Entity
    // Admx Group Policy Id
    admxSettingDefinitionId *string;
    // List of Child Ids of the group policy setting.
    childIdList []string;
    // The Intune Setting Definition Id
    intuneSettingDefinitionId *string;
    // The list of Intune Setting URIs this group policy setting maps to
    intuneSettingUriList []string;
    // Indicates if the setting is supported by Intune or not
    isMdmSupported *bool;
    // The CSP name this group policy setting maps to.
    mdmCspName *string;
    // The minimum OS version this mdm setting supports.
    mdmMinimumOSVersion *int32;
    // The MDM CSP URI this group policy setting maps to.
    mdmSettingUri *string;
    // Indicates if the setting is supported in Mdm or not. Possible values are: unknown, supported, unsupported, deprecated.
    mdmSupportedState *MdmSupportedState;
    // Parent Id of the group policy setting.
    parentId *string;
    // The category the group policy setting is in.
    settingCategory *string;
    // The display name of this group policy setting.
    settingDisplayName *string;
    // The display value of this group policy setting.
    settingDisplayValue *string;
    // The display value type of this group policy setting.
    settingDisplayValueType *string;
    // The name of this group policy setting.
    settingName *string;
    // The scope of the setting. Possible values are: unknown, device, user.
    settingScope *GroupPolicySettingScope;
    // The setting type (security or admx) of the Group Policy. Possible values are: unknown, policy, account, securityOptions, userRightsAssignment, auditSetting, windowsFirewallSettings, appLockerRuleCollection, dataSourcesSettings, devicesSettings, driveMapSettings, environmentVariables, filesSettings, folderOptions, folders, iniFiles, internetOptions, localUsersAndGroups, networkOptions, networkShares, ntServices, powerOptions, printers, regionalOptionsSettings, registrySettings, scheduledTasks, shortcutSettings, startMenuSettings.
    settingType *GroupPolicySettingType;
    // The value of this group policy setting.
    settingValue *string;
    // The display units of this group policy setting value
    settingValueDisplayUnits *string;
    // The value type of this group policy setting.
    settingValueType *string;
}
// NewGroupPolicySettingMapping instantiates a new groupPolicySettingMapping and sets the default values.
func NewGroupPolicySettingMapping()(*GroupPolicySettingMapping) {
    m := &GroupPolicySettingMapping{
        Entity: *NewEntity(),
    }
    return m
}
// GetAdmxSettingDefinitionId gets the admxSettingDefinitionId property value. Admx Group Policy Id
func (m *GroupPolicySettingMapping) GetAdmxSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.admxSettingDefinitionId
    }
}
// GetChildIdList gets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) GetChildIdList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.childIdList
    }
}
// GetIntuneSettingDefinitionId gets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) GetIntuneSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingDefinitionId
    }
}
// GetIntuneSettingUriList gets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) GetIntuneSettingUriList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingUriList
    }
}
// GetIsMdmSupported gets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) GetIsMdmSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMdmSupported
    }
}
// GetMdmCspName gets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmCspName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmCspName
    }
}
// GetMdmMinimumOSVersion gets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) GetMdmMinimumOSVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mdmMinimumOSVersion
    }
}
// GetMdmSettingUri gets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmSettingUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmSettingUri
    }
}
// GetMdmSupportedState gets the mdmSupportedState property value. Indicates if the setting is supported in Mdm or not. Possible values are: unknown, supported, unsupported, deprecated.
func (m *GroupPolicySettingMapping) GetMdmSupportedState()(*MdmSupportedState) {
    if m == nil {
        return nil
    } else {
        return m.mdmSupportedState
    }
}
// GetParentId gets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) GetParentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentId
    }
}
// GetSettingCategory gets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) GetSettingCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategory
    }
}
// GetSettingDisplayName gets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayName
    }
}
// GetSettingDisplayValue gets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayValue
    }
}
// GetSettingDisplayValueType gets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayValueType
    }
}
// GetSettingName gets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetSettingScope gets the settingScope property value. The scope of the setting. Possible values are: unknown, device, user.
func (m *GroupPolicySettingMapping) GetSettingScope()(*GroupPolicySettingScope) {
    if m == nil {
        return nil
    } else {
        return m.settingScope
    }
}
// GetSettingType gets the settingType property value. The setting type (security or admx) of the Group Policy. Possible values are: unknown, policy, account, securityOptions, userRightsAssignment, auditSetting, windowsFirewallSettings, appLockerRuleCollection, dataSourcesSettings, devicesSettings, driveMapSettings, environmentVariables, filesSettings, folderOptions, folders, iniFiles, internetOptions, localUsersAndGroups, networkOptions, networkShares, ntServices, powerOptions, printers, regionalOptionsSettings, registrySettings, scheduledTasks, shortcutSettings, startMenuSettings.
func (m *GroupPolicySettingMapping) GetSettingType()(*GroupPolicySettingType) {
    if m == nil {
        return nil
    } else {
        return m.settingType
    }
}
// GetSettingValue gets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValue
    }
}
// GetSettingValueDisplayUnits gets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) GetSettingValueDisplayUnits()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueDisplayUnits
    }
}
// GetSettingValueType gets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicySettingMapping) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["admxSettingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmxSettingDefinitionId(val)
        }
        return nil
    }
    res["childIdList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetChildIdList(res)
        }
        return nil
    }
    res["intuneSettingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneSettingDefinitionId(val)
        }
        return nil
    }
    res["intuneSettingUriList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIntuneSettingUriList(res)
        }
        return nil
    }
    res["isMdmSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMdmSupported(val)
        }
        return nil
    }
    res["mdmCspName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmCspName(val)
        }
        return nil
    }
    res["mdmMinimumOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmMinimumOSVersion(val)
        }
        return nil
    }
    res["mdmSettingUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmSettingUri(val)
        }
        return nil
    }
    res["mdmSupportedState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmSupportedState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmSupportedState(val.(*MdmSupportedState))
        }
        return nil
    }
    res["parentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentId(val)
        }
        return nil
    }
    res["settingCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategory(val)
        }
        return nil
    }
    res["settingDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayName(val)
        }
        return nil
    }
    res["settingDisplayValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayValue(val)
        }
        return nil
    }
    res["settingDisplayValueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayValueType(val)
        }
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["settingScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingScope(val.(*GroupPolicySettingScope))
        }
        return nil
    }
    res["settingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*GroupPolicySettingType))
        }
        return nil
    }
    res["settingValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValue(val)
        }
        return nil
    }
    res["settingValueDisplayUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueDisplayUnits(val)
        }
        return nil
    }
    res["settingValueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueType(val)
        }
        return nil
    }
    return res
}
func (m *GroupPolicySettingMapping) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GroupPolicySettingMapping) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("admxSettingDefinitionId", m.GetAdmxSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetChildIdList() != nil {
        err = writer.WriteCollectionOfStringValues("childIdList", m.GetChildIdList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("intuneSettingDefinitionId", m.GetIntuneSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetIntuneSettingUriList() != nil {
        err = writer.WriteCollectionOfStringValues("intuneSettingUriList", m.GetIntuneSettingUriList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMdmSupported", m.GetIsMdmSupported())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmCspName", m.GetMdmCspName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mdmMinimumOSVersion", m.GetMdmMinimumOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmSettingUri", m.GetMdmSettingUri())
        if err != nil {
            return err
        }
    }
    if m.GetMdmSupportedState() != nil {
        cast := (*m.GetMdmSupportedState()).String()
        err = writer.WriteStringValue("mdmSupportedState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentId", m.GetParentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingCategory", m.GetSettingCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDisplayName", m.GetSettingDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDisplayValue", m.GetSettingDisplayValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDisplayValueType", m.GetSettingDisplayValueType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetSettingScope() != nil {
        cast := (*m.GetSettingScope()).String()
        err = writer.WriteStringValue("settingScope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingType() != nil {
        cast := (*m.GetSettingType()).String()
        err = writer.WriteStringValue("settingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingValue", m.GetSettingValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingValueDisplayUnits", m.GetSettingValueDisplayUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingValueType", m.GetSettingValueType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdmxSettingDefinitionId sets the admxSettingDefinitionId property value. Admx Group Policy Id
func (m *GroupPolicySettingMapping) SetAdmxSettingDefinitionId(value *string)() {
    if m != nil {
        m.admxSettingDefinitionId = value
    }
}
// SetChildIdList sets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) SetChildIdList(value []string)() {
    if m != nil {
        m.childIdList = value
    }
}
// SetIntuneSettingDefinitionId sets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) SetIntuneSettingDefinitionId(value *string)() {
    if m != nil {
        m.intuneSettingDefinitionId = value
    }
}
// SetIntuneSettingUriList sets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) SetIntuneSettingUriList(value []string)() {
    if m != nil {
        m.intuneSettingUriList = value
    }
}
// SetIsMdmSupported sets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) SetIsMdmSupported(value *bool)() {
    if m != nil {
        m.isMdmSupported = value
    }
}
// SetMdmCspName sets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) SetMdmCspName(value *string)() {
    if m != nil {
        m.mdmCspName = value
    }
}
// SetMdmMinimumOSVersion sets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) SetMdmMinimumOSVersion(value *int32)() {
    if m != nil {
        m.mdmMinimumOSVersion = value
    }
}
// SetMdmSettingUri sets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) SetMdmSettingUri(value *string)() {
    if m != nil {
        m.mdmSettingUri = value
    }
}
// SetMdmSupportedState sets the mdmSupportedState property value. Indicates if the setting is supported in Mdm or not. Possible values are: unknown, supported, unsupported, deprecated.
func (m *GroupPolicySettingMapping) SetMdmSupportedState(value *MdmSupportedState)() {
    if m != nil {
        m.mdmSupportedState = value
    }
}
// SetParentId sets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) SetParentId(value *string)() {
    if m != nil {
        m.parentId = value
    }
}
// SetSettingCategory sets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) SetSettingCategory(value *string)() {
    if m != nil {
        m.settingCategory = value
    }
}
// SetSettingDisplayName sets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayName(value *string)() {
    if m != nil {
        m.settingDisplayName = value
    }
}
// SetSettingDisplayValue sets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayValue(value *string)() {
    if m != nil {
        m.settingDisplayValue = value
    }
}
// SetSettingDisplayValueType sets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayValueType(value *string)() {
    if m != nil {
        m.settingDisplayValueType = value
    }
}
// SetSettingName sets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetSettingScope sets the settingScope property value. The scope of the setting. Possible values are: unknown, device, user.
func (m *GroupPolicySettingMapping) SetSettingScope(value *GroupPolicySettingScope)() {
    if m != nil {
        m.settingScope = value
    }
}
// SetSettingType sets the settingType property value. The setting type (security or admx) of the Group Policy. Possible values are: unknown, policy, account, securityOptions, userRightsAssignment, auditSetting, windowsFirewallSettings, appLockerRuleCollection, dataSourcesSettings, devicesSettings, driveMapSettings, environmentVariables, filesSettings, folderOptions, folders, iniFiles, internetOptions, localUsersAndGroups, networkOptions, networkShares, ntServices, powerOptions, printers, regionalOptionsSettings, registrySettings, scheduledTasks, shortcutSettings, startMenuSettings.
func (m *GroupPolicySettingMapping) SetSettingType(value *GroupPolicySettingType)() {
    if m != nil {
        m.settingType = value
    }
}
// SetSettingValue sets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingValue(value *string)() {
    if m != nil {
        m.settingValue = value
    }
}
// SetSettingValueDisplayUnits sets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) SetSettingValueDisplayUnits(value *string)() {
    if m != nil {
        m.settingValueDisplayUnits = value
    }
}
// SetSettingValueType sets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingValueType(value *string)() {
    if m != nil {
        m.settingValueType = value
    }
}
