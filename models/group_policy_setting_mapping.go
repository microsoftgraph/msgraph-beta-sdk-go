package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicySettingMapping the Group Policy setting to MDM/Intune mapping.
type GroupPolicySettingMapping struct {
    Entity
}
// NewGroupPolicySettingMapping instantiates a new groupPolicySettingMapping and sets the default values.
func NewGroupPolicySettingMapping()(*GroupPolicySettingMapping) {
    m := &GroupPolicySettingMapping{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicySettingMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicySettingMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicySettingMapping(), nil
}
// GetAdmxSettingDefinitionId gets the admxSettingDefinitionId property value. Admx Group Policy Id
func (m *GroupPolicySettingMapping) GetAdmxSettingDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("admxSettingDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetChildIdList gets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) GetChildIdList()([]string) {
    val, err := m.GetBackingStore().Get("childIdList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicySettingMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["admxSettingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmxSettingDefinitionId(val)
        }
        return nil
    }
    res["childIdList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetChildIdList(res)
        }
        return nil
    }
    res["intuneSettingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneSettingDefinitionId(val)
        }
        return nil
    }
    res["intuneSettingUriList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetIntuneSettingUriList(res)
        }
        return nil
    }
    res["isMdmSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMdmSupported(val)
        }
        return nil
    }
    res["mdmCspName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmCspName(val)
        }
        return nil
    }
    res["mdmMinimumOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmMinimumOSVersion(val)
        }
        return nil
    }
    res["mdmSettingUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmSettingUri(val)
        }
        return nil
    }
    res["mdmSupportedState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmSupportedState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmSupportedState(val.(*MdmSupportedState))
        }
        return nil
    }
    res["parentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentId(val)
        }
        return nil
    }
    res["settingCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategory(val)
        }
        return nil
    }
    res["settingDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayName(val)
        }
        return nil
    }
    res["settingDisplayValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayValue(val)
        }
        return nil
    }
    res["settingDisplayValueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayValueType(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["settingScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingScope(val.(*GroupPolicySettingScope))
        }
        return nil
    }
    res["settingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*GroupPolicySettingType))
        }
        return nil
    }
    res["settingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValue(val)
        }
        return nil
    }
    res["settingValueDisplayUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueDisplayUnits(val)
        }
        return nil
    }
    res["settingValueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIntuneSettingDefinitionId gets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) GetIntuneSettingDefinitionId()(*string) {
    val, err := m.GetBackingStore().Get("intuneSettingDefinitionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIntuneSettingUriList gets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) GetIntuneSettingUriList()([]string) {
    val, err := m.GetBackingStore().Get("intuneSettingUriList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetIsMdmSupported gets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) GetIsMdmSupported()(*bool) {
    val, err := m.GetBackingStore().Get("isMdmSupported")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMdmCspName gets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmCspName()(*string) {
    val, err := m.GetBackingStore().Get("mdmCspName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMdmMinimumOSVersion gets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) GetMdmMinimumOSVersion()(*int32) {
    val, err := m.GetBackingStore().Get("mdmMinimumOSVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMdmSettingUri gets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmSettingUri()(*string) {
    val, err := m.GetBackingStore().Get("mdmSettingUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMdmSupportedState gets the mdmSupportedState property value. Mdm Support Status of the setting.
func (m *GroupPolicySettingMapping) GetMdmSupportedState()(*MdmSupportedState) {
    val, err := m.GetBackingStore().Get("mdmSupportedState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MdmSupportedState)
    }
    return nil
}
// GetParentId gets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) GetParentId()(*string) {
    val, err := m.GetBackingStore().Get("parentId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingCategory gets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) GetSettingCategory()(*string) {
    val, err := m.GetBackingStore().Get("settingCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingDisplayName gets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("settingDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingDisplayValue gets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValue()(*string) {
    val, err := m.GetBackingStore().Get("settingDisplayValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingDisplayValueType gets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValueType()(*string) {
    val, err := m.GetBackingStore().Get("settingDisplayValueType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingName gets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingName()(*string) {
    val, err := m.GetBackingStore().Get("settingName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingScope gets the settingScope property value. Scope of the group policy setting.
func (m *GroupPolicySettingMapping) GetSettingScope()(*GroupPolicySettingScope) {
    val, err := m.GetBackingStore().Get("settingScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicySettingScope)
    }
    return nil
}
// GetSettingType gets the settingType property value. Setting type of the group policy.
func (m *GroupPolicySettingMapping) GetSettingType()(*GroupPolicySettingType) {
    val, err := m.GetBackingStore().Get("settingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicySettingType)
    }
    return nil
}
// GetSettingValue gets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValue()(*string) {
    val, err := m.GetBackingStore().Get("settingValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingValueDisplayUnits gets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) GetSettingValueDisplayUnits()(*string) {
    val, err := m.GetBackingStore().Get("settingValueDisplayUnits")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingValueType gets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValueType()(*string) {
    val, err := m.GetBackingStore().Get("settingValueType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GroupPolicySettingMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("admxSettingDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetChildIdList sets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) SetChildIdList(value []string)() {
    err := m.GetBackingStore().Set("childIdList", value)
    if err != nil {
        panic(err)
    }
}
// SetIntuneSettingDefinitionId sets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) SetIntuneSettingDefinitionId(value *string)() {
    err := m.GetBackingStore().Set("intuneSettingDefinitionId", value)
    if err != nil {
        panic(err)
    }
}
// SetIntuneSettingUriList sets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) SetIntuneSettingUriList(value []string)() {
    err := m.GetBackingStore().Set("intuneSettingUriList", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMdmSupported sets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) SetIsMdmSupported(value *bool)() {
    err := m.GetBackingStore().Set("isMdmSupported", value)
    if err != nil {
        panic(err)
    }
}
// SetMdmCspName sets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) SetMdmCspName(value *string)() {
    err := m.GetBackingStore().Set("mdmCspName", value)
    if err != nil {
        panic(err)
    }
}
// SetMdmMinimumOSVersion sets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) SetMdmMinimumOSVersion(value *int32)() {
    err := m.GetBackingStore().Set("mdmMinimumOSVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMdmSettingUri sets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) SetMdmSettingUri(value *string)() {
    err := m.GetBackingStore().Set("mdmSettingUri", value)
    if err != nil {
        panic(err)
    }
}
// SetMdmSupportedState sets the mdmSupportedState property value. Mdm Support Status of the setting.
func (m *GroupPolicySettingMapping) SetMdmSupportedState(value *MdmSupportedState)() {
    err := m.GetBackingStore().Set("mdmSupportedState", value)
    if err != nil {
        panic(err)
    }
}
// SetParentId sets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) SetParentId(value *string)() {
    err := m.GetBackingStore().Set("parentId", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingCategory sets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) SetSettingCategory(value *string)() {
    err := m.GetBackingStore().Set("settingCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDisplayName sets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayName(value *string)() {
    err := m.GetBackingStore().Set("settingDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDisplayValue sets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayValue(value *string)() {
    err := m.GetBackingStore().Set("settingDisplayValue", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingDisplayValueType sets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayValueType(value *string)() {
    err := m.GetBackingStore().Set("settingDisplayValueType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingName sets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingName(value *string)() {
    err := m.GetBackingStore().Set("settingName", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingScope sets the settingScope property value. Scope of the group policy setting.
func (m *GroupPolicySettingMapping) SetSettingScope(value *GroupPolicySettingScope)() {
    err := m.GetBackingStore().Set("settingScope", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingType sets the settingType property value. Setting type of the group policy.
func (m *GroupPolicySettingMapping) SetSettingType(value *GroupPolicySettingType)() {
    err := m.GetBackingStore().Set("settingType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingValue sets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingValue(value *string)() {
    err := m.GetBackingStore().Set("settingValue", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingValueDisplayUnits sets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) SetSettingValueDisplayUnits(value *string)() {
    err := m.GetBackingStore().Set("settingValueDisplayUnits", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingValueType sets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingValueType(value *string)() {
    err := m.GetBackingStore().Set("settingValueType", value)
    if err != nil {
        panic(err)
    }
}
// GroupPolicySettingMappingable 
type GroupPolicySettingMappingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdmxSettingDefinitionId()(*string)
    GetChildIdList()([]string)
    GetIntuneSettingDefinitionId()(*string)
    GetIntuneSettingUriList()([]string)
    GetIsMdmSupported()(*bool)
    GetMdmCspName()(*string)
    GetMdmMinimumOSVersion()(*int32)
    GetMdmSettingUri()(*string)
    GetMdmSupportedState()(*MdmSupportedState)
    GetParentId()(*string)
    GetSettingCategory()(*string)
    GetSettingDisplayName()(*string)
    GetSettingDisplayValue()(*string)
    GetSettingDisplayValueType()(*string)
    GetSettingName()(*string)
    GetSettingScope()(*GroupPolicySettingScope)
    GetSettingType()(*GroupPolicySettingType)
    GetSettingValue()(*string)
    GetSettingValueDisplayUnits()(*string)
    GetSettingValueType()(*string)
    SetAdmxSettingDefinitionId(value *string)()
    SetChildIdList(value []string)()
    SetIntuneSettingDefinitionId(value *string)()
    SetIntuneSettingUriList(value []string)()
    SetIsMdmSupported(value *bool)()
    SetMdmCspName(value *string)()
    SetMdmMinimumOSVersion(value *int32)()
    SetMdmSettingUri(value *string)()
    SetMdmSupportedState(value *MdmSupportedState)()
    SetParentId(value *string)()
    SetSettingCategory(value *string)()
    SetSettingDisplayName(value *string)()
    SetSettingDisplayValue(value *string)()
    SetSettingDisplayValueType(value *string)()
    SetSettingName(value *string)()
    SetSettingScope(value *GroupPolicySettingScope)()
    SetSettingType(value *GroupPolicySettingType)()
    SetSettingValue(value *string)()
    SetSettingValueDisplayUnits(value *string)()
    SetSettingValueType(value *string)()
}
