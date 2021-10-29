package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
    // The setting type (security or admx) of the Group Policy. Possible values are: unknown, policy, account, securityOptions, userRightsAssignment, auditSetting, windowsFirewallSettings.
    settingType *GroupPolicySettingType;
    // The value of this group policy setting.
    settingValue *string;
    // The display units of this group policy setting value
    settingValueDisplayUnits *string;
    // The value type of this group policy setting.
    settingValueType *string;
}
// Instantiates a new groupPolicySettingMapping and sets the default values.
func NewGroupPolicySettingMapping()(*GroupPolicySettingMapping) {
    m := &GroupPolicySettingMapping{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the admxSettingDefinitionId property value. Admx Group Policy Id
func (m *GroupPolicySettingMapping) GetAdmxSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.admxSettingDefinitionId
    }
}
// Gets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) GetChildIdList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.childIdList
    }
}
// Gets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) GetIntuneSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingDefinitionId
    }
}
// Gets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) GetIntuneSettingUriList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingUriList
    }
}
// Gets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) GetIsMdmSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMdmSupported
    }
}
// Gets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmCspName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmCspName
    }
}
// Gets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) GetMdmMinimumOSVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mdmMinimumOSVersion
    }
}
// Gets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmSettingUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmSettingUri
    }
}
// Gets the mdmSupportedState property value. Indicates if the setting is supported in Mdm or not. Possible values are: unknown, supported, unsupported, deprecated.
func (m *GroupPolicySettingMapping) GetMdmSupportedState()(*MdmSupportedState) {
    if m == nil {
        return nil
    } else {
        return m.mdmSupportedState
    }
}
// Gets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) GetParentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentId
    }
}
// Gets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) GetSettingCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategory
    }
}
// Gets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayName
    }
}
// Gets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayValue
    }
}
// Gets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayValueType
    }
}
// Gets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// Gets the settingScope property value. The scope of the setting. Possible values are: unknown, device, user.
func (m *GroupPolicySettingMapping) GetSettingScope()(*GroupPolicySettingScope) {
    if m == nil {
        return nil
    } else {
        return m.settingScope
    }
}
// Gets the settingType property value. The setting type (security or admx) of the Group Policy. Possible values are: unknown, policy, account, securityOptions, userRightsAssignment, auditSetting, windowsFirewallSettings.
func (m *GroupPolicySettingMapping) GetSettingType()(*GroupPolicySettingType) {
    if m == nil {
        return nil
    } else {
        return m.settingType
    }
}
// Gets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValue
    }
}
// Gets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) GetSettingValueDisplayUnits()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueDisplayUnits
    }
}
// Gets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueType
    }
}
// The deserialization information for the current model
func (m *GroupPolicySettingMapping) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["admxSettingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdmxSettingDefinitionId(val)
        return nil
    }
    res["childIdList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetChildIdList(res)
        return nil
    }
    res["intuneSettingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIntuneSettingDefinitionId(val)
        return nil
    }
    res["intuneSettingUriList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIntuneSettingUriList(res)
        return nil
    }
    res["isMdmSupported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMdmSupported(val)
        return nil
    }
    res["mdmCspName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMdmCspName(val)
        return nil
    }
    res["mdmMinimumOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMdmMinimumOSVersion(val)
        return nil
    }
    res["mdmSettingUri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMdmSettingUri(val)
        return nil
    }
    res["mdmSupportedState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmSupportedState)
        if err != nil {
            return err
        }
        cast := val.(MdmSupportedState)
        m.SetMdmSupportedState(&cast)
        return nil
    }
    res["parentId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentId(val)
        return nil
    }
    res["settingCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingCategory(val)
        return nil
    }
    res["settingDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingDisplayName(val)
        return nil
    }
    res["settingDisplayValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingDisplayValue(val)
        return nil
    }
    res["settingDisplayValueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingDisplayValueType(val)
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
        return nil
    }
    res["settingScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicySettingScope)
        m.SetSettingScope(&cast)
        return nil
    }
    res["settingType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingType)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicySettingType)
        m.SetSettingType(&cast)
        return nil
    }
    res["settingValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingValue(val)
        return nil
    }
    res["settingValueDisplayUnits"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingValueDisplayUnits(val)
        return nil
    }
    res["settingValueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingValueType(val)
        return nil
    }
    return res
}
func (m *GroupPolicySettingMapping) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
    {
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
        cast := m.GetMdmSupportedState().String()
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
        cast := m.GetSettingScope().String()
        err = writer.WriteStringValue("settingScope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingType() != nil {
        cast := m.GetSettingType().String()
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
// Sets the admxSettingDefinitionId property value. Admx Group Policy Id
// Parameters:
//  - value : Value to set for the admxSettingDefinitionId property.
func (m *GroupPolicySettingMapping) SetAdmxSettingDefinitionId(value *string)() {
    m.admxSettingDefinitionId = value
}
// Sets the childIdList property value. List of Child Ids of the group policy setting.
// Parameters:
//  - value : Value to set for the childIdList property.
func (m *GroupPolicySettingMapping) SetChildIdList(value []string)() {
    m.childIdList = value
}
// Sets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
// Parameters:
//  - value : Value to set for the intuneSettingDefinitionId property.
func (m *GroupPolicySettingMapping) SetIntuneSettingDefinitionId(value *string)() {
    m.intuneSettingDefinitionId = value
}
// Sets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
// Parameters:
//  - value : Value to set for the intuneSettingUriList property.
func (m *GroupPolicySettingMapping) SetIntuneSettingUriList(value []string)() {
    m.intuneSettingUriList = value
}
// Sets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
// Parameters:
//  - value : Value to set for the isMdmSupported property.
func (m *GroupPolicySettingMapping) SetIsMdmSupported(value *bool)() {
    m.isMdmSupported = value
}
// Sets the mdmCspName property value. The CSP name this group policy setting maps to.
// Parameters:
//  - value : Value to set for the mdmCspName property.
func (m *GroupPolicySettingMapping) SetMdmCspName(value *string)() {
    m.mdmCspName = value
}
// Sets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
// Parameters:
//  - value : Value to set for the mdmMinimumOSVersion property.
func (m *GroupPolicySettingMapping) SetMdmMinimumOSVersion(value *int32)() {
    m.mdmMinimumOSVersion = value
}
// Sets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
// Parameters:
//  - value : Value to set for the mdmSettingUri property.
func (m *GroupPolicySettingMapping) SetMdmSettingUri(value *string)() {
    m.mdmSettingUri = value
}
// Sets the mdmSupportedState property value. Indicates if the setting is supported in Mdm or not. Possible values are: unknown, supported, unsupported, deprecated.
// Parameters:
//  - value : Value to set for the mdmSupportedState property.
func (m *GroupPolicySettingMapping) SetMdmSupportedState(value *MdmSupportedState)() {
    m.mdmSupportedState = value
}
// Sets the parentId property value. Parent Id of the group policy setting.
// Parameters:
//  - value : Value to set for the parentId property.
func (m *GroupPolicySettingMapping) SetParentId(value *string)() {
    m.parentId = value
}
// Sets the settingCategory property value. The category the group policy setting is in.
// Parameters:
//  - value : Value to set for the settingCategory property.
func (m *GroupPolicySettingMapping) SetSettingCategory(value *string)() {
    m.settingCategory = value
}
// Sets the settingDisplayName property value. The display name of this group policy setting.
// Parameters:
//  - value : Value to set for the settingDisplayName property.
func (m *GroupPolicySettingMapping) SetSettingDisplayName(value *string)() {
    m.settingDisplayName = value
}
// Sets the settingDisplayValue property value. The display value of this group policy setting.
// Parameters:
//  - value : Value to set for the settingDisplayValue property.
func (m *GroupPolicySettingMapping) SetSettingDisplayValue(value *string)() {
    m.settingDisplayValue = value
}
// Sets the settingDisplayValueType property value. The display value type of this group policy setting.
// Parameters:
//  - value : Value to set for the settingDisplayValueType property.
func (m *GroupPolicySettingMapping) SetSettingDisplayValueType(value *string)() {
    m.settingDisplayValueType = value
}
// Sets the settingName property value. The name of this group policy setting.
// Parameters:
//  - value : Value to set for the settingName property.
func (m *GroupPolicySettingMapping) SetSettingName(value *string)() {
    m.settingName = value
}
// Sets the settingScope property value. The scope of the setting. Possible values are: unknown, device, user.
// Parameters:
//  - value : Value to set for the settingScope property.
func (m *GroupPolicySettingMapping) SetSettingScope(value *GroupPolicySettingScope)() {
    m.settingScope = value
}
// Sets the settingType property value. The setting type (security or admx) of the Group Policy. Possible values are: unknown, policy, account, securityOptions, userRightsAssignment, auditSetting, windowsFirewallSettings.
// Parameters:
//  - value : Value to set for the settingType property.
func (m *GroupPolicySettingMapping) SetSettingType(value *GroupPolicySettingType)() {
    m.settingType = value
}
// Sets the settingValue property value. The value of this group policy setting.
// Parameters:
//  - value : Value to set for the settingValue property.
func (m *GroupPolicySettingMapping) SetSettingValue(value *string)() {
    m.settingValue = value
}
// Sets the settingValueDisplayUnits property value. The display units of this group policy setting value
// Parameters:
//  - value : Value to set for the settingValueDisplayUnits property.
func (m *GroupPolicySettingMapping) SetSettingValueDisplayUnits(value *string)() {
    m.settingValueDisplayUnits = value
}
// Sets the settingValueType property value. The value type of this group policy setting.
// Parameters:
//  - value : Value to set for the settingValueType property.
func (m *GroupPolicySettingMapping) SetSettingValueType(value *string)() {
    m.settingValueType = value
}
