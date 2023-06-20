package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityLabel 
type SensitivityLabel struct {
    Entity
}
// NewSensitivityLabel instantiates a new sensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
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
// GetAssignedPolicies gets the assignedPolicies property value. The assignedPolicies property
func (m *SensitivityLabel) GetAssignedPolicies()([]LabelPolicyable) {
    val, err := m.GetBackingStore().Get("assignedPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LabelPolicyable)
    }
    return nil
}
// GetAutoLabeling gets the autoLabeling property value. The autoLabeling property
func (m *SensitivityLabel) GetAutoLabeling()(AutoLabelingable) {
    val, err := m.GetBackingStore().Get("autoLabeling")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AutoLabelingable)
    }
    return nil
}
// GetDescription gets the description property value. The description property
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
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["assignedPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LabelPolicyable)
                }
            }
            m.SetAssignedPolicies(res)
        }
        return nil
    }
    res["autoLabeling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutoLabelingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoLabeling(val.(AutoLabelingable))
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
    res["labelActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLabelActionBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LabelActionBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LabelActionBaseable)
                }
            }
            m.SetLabelActions(res)
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
// GetIsEndpointProtectionEnabled gets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
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
// GetLabelActions gets the labelActions property value. The labelActions property
func (m *SensitivityLabel) GetLabelActions()([]LabelActionBaseable) {
    val, err := m.GetBackingStore().Get("labelActions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LabelActionBaseable)
    }
    return nil
}
// GetName gets the name property value. The name property
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
// GetSublabels gets the sublabels property value. The sublabels property
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
    if m.GetAssignedPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedPolicies()))
        for i, v := range m.GetAssignedPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("autoLabeling", m.GetAutoLabeling())
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
        err = writer.WriteBoolValue("isEndpointProtectionEnabled", m.GetIsEndpointProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetLabelActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLabelActions()))
        for i, v := range m.GetLabelActions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("labelActions", cast)
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
// SetAssignedPolicies sets the assignedPolicies property value. The assignedPolicies property
func (m *SensitivityLabel) SetAssignedPolicies(value []LabelPolicyable)() {
    err := m.GetBackingStore().Set("assignedPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoLabeling sets the autoLabeling property value. The autoLabeling property
func (m *SensitivityLabel) SetAutoLabeling(value AutoLabelingable)() {
    err := m.GetBackingStore().Set("autoLabeling", value)
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
// SetIsEndpointProtectionEnabled sets the isEndpointProtectionEnabled property value. The isEndpointProtectionEnabled property
func (m *SensitivityLabel) SetIsEndpointProtectionEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isEndpointProtectionEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetLabelActions sets the labelActions property value. The labelActions property
func (m *SensitivityLabel) SetLabelActions(value []LabelActionBaseable)() {
    err := m.GetBackingStore().Set("labelActions", value)
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
// SensitivityLabelable 
type SensitivityLabelable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableTo()(*SensitivityLabelTarget)
    GetApplicationMode()(*ApplicationMode)
    GetAssignedPolicies()([]LabelPolicyable)
    GetAutoLabeling()(AutoLabelingable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsDefault()(*bool)
    GetIsEndpointProtectionEnabled()(*bool)
    GetLabelActions()([]LabelActionBaseable)
    GetName()(*string)
    GetPriority()(*int32)
    GetSublabels()([]SensitivityLabelable)
    GetToolTip()(*string)
    SetApplicableTo(value *SensitivityLabelTarget)()
    SetApplicationMode(value *ApplicationMode)()
    SetAssignedPolicies(value []LabelPolicyable)()
    SetAutoLabeling(value AutoLabelingable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsDefault(value *bool)()
    SetIsEndpointProtectionEnabled(value *bool)()
    SetLabelActions(value []LabelActionBaseable)()
    SetName(value *string)()
    SetPriority(value *int32)()
    SetSublabels(value []SensitivityLabelable)()
    SetToolTip(value *string)()
}
