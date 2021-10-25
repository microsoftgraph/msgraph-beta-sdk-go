package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GroupPolicySettingMapping struct {
    Entity
    admxSettingDefinitionId *string;
    childIdList []string;
    intuneSettingDefinitionId *string;
    intuneSettingUriList []string;
    isMdmSupported *bool;
    mdmCspName *string;
    mdmMinimumOSVersion *int32;
    mdmSettingUri *string;
    mdmSupportedState *MdmSupportedState;
    parentId *string;
    settingCategory *string;
    settingDisplayName *string;
    settingDisplayValue *string;
    settingDisplayValueType *string;
    settingName *string;
    settingScope *GroupPolicySettingScope;
    settingType *GroupPolicySettingType;
    settingValue *string;
    settingValueDisplayUnits *string;
    settingValueType *string;
}
func NewGroupPolicySettingMapping()(*GroupPolicySettingMapping) {
    m := &GroupPolicySettingMapping{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GroupPolicySettingMapping) GetAdmxSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.admxSettingDefinitionId
    }
}
func (m *GroupPolicySettingMapping) GetChildIdList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.childIdList
    }
}
func (m *GroupPolicySettingMapping) GetIntuneSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingDefinitionId
    }
}
func (m *GroupPolicySettingMapping) GetIntuneSettingUriList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.intuneSettingUriList
    }
}
func (m *GroupPolicySettingMapping) GetIsMdmSupported()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMdmSupported
    }
}
func (m *GroupPolicySettingMapping) GetMdmCspName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmCspName
    }
}
func (m *GroupPolicySettingMapping) GetMdmMinimumOSVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mdmMinimumOSVersion
    }
}
func (m *GroupPolicySettingMapping) GetMdmSettingUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmSettingUri
    }
}
func (m *GroupPolicySettingMapping) GetMdmSupportedState()(*MdmSupportedState) {
    if m == nil {
        return nil
    } else {
        return m.mdmSupportedState
    }
}
func (m *GroupPolicySettingMapping) GetParentId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentId
    }
}
func (m *GroupPolicySettingMapping) GetSettingCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingCategory
    }
}
func (m *GroupPolicySettingMapping) GetSettingDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayName
    }
}
func (m *GroupPolicySettingMapping) GetSettingDisplayValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayValue
    }
}
func (m *GroupPolicySettingMapping) GetSettingDisplayValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDisplayValueType
    }
}
func (m *GroupPolicySettingMapping) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
func (m *GroupPolicySettingMapping) GetSettingScope()(*GroupPolicySettingScope) {
    if m == nil {
        return nil
    } else {
        return m.settingScope
    }
}
func (m *GroupPolicySettingMapping) GetSettingType()(*GroupPolicySettingType) {
    if m == nil {
        return nil
    } else {
        return m.settingType
    }
}
func (m *GroupPolicySettingMapping) GetSettingValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValue
    }
}
func (m *GroupPolicySettingMapping) GetSettingValueDisplayUnits()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueDisplayUnits
    }
}
func (m *GroupPolicySettingMapping) GetSettingValueType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingValueType
    }
}
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
func (m *GroupPolicySettingMapping) SetAdmxSettingDefinitionId(value *string)() {
    m.admxSettingDefinitionId = value
}
func (m *GroupPolicySettingMapping) SetChildIdList(value []string)() {
    m.childIdList = value
}
func (m *GroupPolicySettingMapping) SetIntuneSettingDefinitionId(value *string)() {
    m.intuneSettingDefinitionId = value
}
func (m *GroupPolicySettingMapping) SetIntuneSettingUriList(value []string)() {
    m.intuneSettingUriList = value
}
func (m *GroupPolicySettingMapping) SetIsMdmSupported(value *bool)() {
    m.isMdmSupported = value
}
func (m *GroupPolicySettingMapping) SetMdmCspName(value *string)() {
    m.mdmCspName = value
}
func (m *GroupPolicySettingMapping) SetMdmMinimumOSVersion(value *int32)() {
    m.mdmMinimumOSVersion = value
}
func (m *GroupPolicySettingMapping) SetMdmSettingUri(value *string)() {
    m.mdmSettingUri = value
}
func (m *GroupPolicySettingMapping) SetMdmSupportedState(value *MdmSupportedState)() {
    m.mdmSupportedState = value
}
func (m *GroupPolicySettingMapping) SetParentId(value *string)() {
    m.parentId = value
}
func (m *GroupPolicySettingMapping) SetSettingCategory(value *string)() {
    m.settingCategory = value
}
func (m *GroupPolicySettingMapping) SetSettingDisplayName(value *string)() {
    m.settingDisplayName = value
}
func (m *GroupPolicySettingMapping) SetSettingDisplayValue(value *string)() {
    m.settingDisplayValue = value
}
func (m *GroupPolicySettingMapping) SetSettingDisplayValueType(value *string)() {
    m.settingDisplayValueType = value
}
func (m *GroupPolicySettingMapping) SetSettingName(value *string)() {
    m.settingName = value
}
func (m *GroupPolicySettingMapping) SetSettingScope(value *GroupPolicySettingScope)() {
    m.settingScope = value
}
func (m *GroupPolicySettingMapping) SetSettingType(value *GroupPolicySettingType)() {
    m.settingType = value
}
func (m *GroupPolicySettingMapping) SetSettingValue(value *string)() {
    m.settingValue = value
}
func (m *GroupPolicySettingMapping) SetSettingValueDisplayUnits(value *string)() {
    m.settingValueDisplayUnits = value
}
func (m *GroupPolicySettingMapping) SetSettingValueType(value *string)() {
    m.settingValueType = value
}
