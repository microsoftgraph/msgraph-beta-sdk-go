// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SensitivityLabel struct {
    Entity
}
// NewSensitivityLabel instantiates a new SensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetActionSource gets the actionSource property value. The actionSource property
// returns a *LabelActionSource when successful
func (m *SensitivityLabel) GetActionSource()(*LabelActionSource) {
    val, err := m.GetBackingStore().Get("actionSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LabelActionSource)
    }
    return nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
// returns a *SensitivityLabelTarget when successful
func (m *SensitivityLabel) GetApplicableTo()(*SensitivityLabelTarget) {
    val, err := m.GetBackingStore().Get("applicableTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensitivityLabelTarget)
    }
    return nil
}
// GetApplicationMode gets the applicationMode property value. The applicationMode property
// returns a *ApplicationMode when successful
func (m *SensitivityLabel) GetApplicationMode()(*ApplicationMode) {
    val, err := m.GetBackingStore().Get("applicationMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApplicationMode)
    }
    return nil
}
// GetAutoTooltip gets the autoTooltip property value. The autoTooltip property
// returns a *string when successful
func (m *SensitivityLabel) GetAutoTooltip()(*string) {
    val, err := m.GetBackingStore().Get("autoTooltip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetColor gets the color property value. The color property
// returns a *string when successful
func (m *SensitivityLabel) GetColor()(*string) {
    val, err := m.GetBackingStore().Get("color")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *SensitivityLabel) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *SensitivityLabel) GetDisplayName()(*string) {
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
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLabelActionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSource(val.(*LabelActionSource))
        }
        return nil
    }
    res["applicableTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableTo(val.(*SensitivityLabelTarget))
        }
        return nil
    }
    res["applicationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationMode(val.(*ApplicationMode))
        }
        return nil
    }
    res["autoTooltip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoTooltip(val)
        }
        return nil
    }
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
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
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isEndpointProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEndpointProtectionEnabled(val)
        }
        return nil
    }
    res["isScopedToUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsScopedToUser(val)
        }
        return nil
    }
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["rights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUsageRightsIncludedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRights(val.(UsageRightsIncludedable))
        }
        return nil
    }
    res["sublabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SensitivityLabelable)
                }
            }
            m.SetSublabels(res)
        }
        return nil
    }
    res["toolTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToolTip(val)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. The isDefault property
// returns a *bool when successful
func (m *SensitivityLabel) GetIsDefault()(*bool) {
    val, err := m.GetBackingStore().Get("isDefault")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
// returns a *bool when successful
func (m *SensitivityLabel) GetIsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsEndpointProtectionEnabled gets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
// returns a *bool when successful
func (m *SensitivityLabel) GetIsEndpointProtectionEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isEndpointProtectionEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsScopedToUser gets the isScopedToUser property value. The isScopedToUser property
// returns a *bool when successful
func (m *SensitivityLabel) GetIsScopedToUser()(*bool) {
    val, err := m.GetBackingStore().Get("isScopedToUser")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocale gets the locale property value. The locale property
// returns a *string when successful
func (m *SensitivityLabel) GetLocale()(*string) {
    val, err := m.GetBackingStore().Get("locale")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *SensitivityLabel) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPriority gets the priority property value. The priority property
// returns a *int32 when successful
func (m *SensitivityLabel) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRights gets the rights property value. The rights property
// returns a UsageRightsIncludedable when successful
func (m *SensitivityLabel) GetRights()(UsageRightsIncludedable) {
    val, err := m.GetBackingStore().Get("rights")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UsageRightsIncludedable)
    }
    return nil
}
// GetSublabels gets the sublabels property value. The sublabels property
// returns a []SensitivityLabelable when successful
func (m *SensitivityLabel) GetSublabels()([]SensitivityLabelable) {
    val, err := m.GetBackingStore().Get("sublabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SensitivityLabelable)
    }
    return nil
}
// GetToolTip gets the toolTip property value. The toolTip property
// returns a *string when successful
func (m *SensitivityLabel) GetToolTip()(*string) {
    val, err := m.GetBackingStore().Get("toolTip")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SensitivityLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionSource() != nil {
        cast := (*m.GetActionSource()).String()
        err = writer.WriteStringValue("actionSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicableTo() != nil {
        cast := (*m.GetApplicableTo()).String()
        err = writer.WriteStringValue("applicableTo", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationMode() != nil {
        cast := (*m.GetApplicationMode()).String()
        err = writer.WriteStringValue("applicationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("autoTooltip", m.GetAutoTooltip())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEndpointProtectionEnabled", m.GetIsEndpointProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isScopedToUser", m.GetIsScopedToUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rights", m.GetRights())
        if err != nil {
            return err
        }
    }
    if m.GetSublabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSublabels()))
        for i, v := range m.GetSublabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sublabels", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("toolTip", m.GetToolTip())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionSource sets the actionSource property value. The actionSource property
func (m *SensitivityLabel) SetActionSource(value *LabelActionSource)() {
    err := m.GetBackingStore().Set("actionSource", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicableTo sets the applicableTo property value. The applicableTo property
func (m *SensitivityLabel) SetApplicableTo(value *SensitivityLabelTarget)() {
    err := m.GetBackingStore().Set("applicableTo", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicationMode sets the applicationMode property value. The applicationMode property
func (m *SensitivityLabel) SetApplicationMode(value *ApplicationMode)() {
    err := m.GetBackingStore().Set("applicationMode", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoTooltip sets the autoTooltip property value. The autoTooltip property
func (m *SensitivityLabel) SetAutoTooltip(value *string)() {
    err := m.GetBackingStore().Set("autoTooltip", value)
    if err != nil {
        panic(err)
    }
}
// SetColor sets the color property value. The color property
func (m *SensitivityLabel) SetColor(value *string)() {
    err := m.GetBackingStore().Set("color", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *SensitivityLabel) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *SensitivityLabel) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *SensitivityLabel) SetIsDefault(value *bool)() {
    err := m.GetBackingStore().Set("isDefault", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *SensitivityLabel) SetIsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEndpointProtectionEnabled sets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *SensitivityLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEndpointProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsScopedToUser sets the isScopedToUser property value. The isScopedToUser property
func (m *SensitivityLabel) SetIsScopedToUser(value *bool)() {
    err := m.GetBackingStore().Set("isScopedToUser", value)
    if err != nil {
        panic(err)
    }
}
// SetLocale sets the locale property value. The locale property
func (m *SensitivityLabel) SetLocale(value *string)() {
    err := m.GetBackingStore().Set("locale", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *SensitivityLabel) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority property
func (m *SensitivityLabel) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetRights sets the rights property value. The rights property
func (m *SensitivityLabel) SetRights(value UsageRightsIncludedable)() {
    err := m.GetBackingStore().Set("rights", value)
    if err != nil {
        panic(err)
    }
}
// SetSublabels sets the sublabels property value. The sublabels property
func (m *SensitivityLabel) SetSublabels(value []SensitivityLabelable)() {
    err := m.GetBackingStore().Set("sublabels", value)
    if err != nil {
        panic(err)
    }
}
// SetToolTip sets the toolTip property value. The toolTip property
func (m *SensitivityLabel) SetToolTip(value *string)() {
    err := m.GetBackingStore().Set("toolTip", value)
    if err != nil {
        panic(err)
    }
}
type SensitivityLabelable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSource()(*LabelActionSource)
    GetApplicableTo()(*SensitivityLabelTarget)
    GetApplicationMode()(*ApplicationMode)
    GetAutoTooltip()(*string)
    GetColor()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsDefault()(*bool)
    GetIsEnabled()(*bool)
    GetIsEndpointProtectionEnabled()(*bool)
    GetIsScopedToUser()(*bool)
    GetLocale()(*string)
    GetName()(*string)
    GetPriority()(*int32)
    GetRights()(UsageRightsIncludedable)
    GetSublabels()([]SensitivityLabelable)
    GetToolTip()(*string)
    SetActionSource(value *LabelActionSource)()
    SetApplicableTo(value *SensitivityLabelTarget)()
    SetApplicationMode(value *ApplicationMode)()
    SetAutoTooltip(value *string)()
    SetColor(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsDefault(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsEndpointProtectionEnabled(value *bool)()
    SetIsScopedToUser(value *bool)()
    SetLocale(value *string)()
    SetName(value *string)()
    SetPriority(value *int32)()
    SetRights(value UsageRightsIncludedable)()
    SetSublabels(value []SensitivityLabelable)()
    SetToolTip(value *string)()
}
