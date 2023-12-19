package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectGroup 
type ProtectGroup struct {
    LabelActionBase
}
// NewProtectGroup instantiates a new protectGroup and sets the default values.
func NewProtectGroup()(*ProtectGroup) {
    m := &ProtectGroup{
        LabelActionBase: *NewLabelActionBase(),
    }
    odataTypeValue := "#microsoft.graph.protectGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProtectGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectGroup(), nil
}
// GetAllowEmailFromGuestUsers gets the allowEmailFromGuestUsers property value. The allowEmailFromGuestUsers property
func (m *ProtectGroup) GetAllowEmailFromGuestUsers()(*bool) {
    val, err := m.GetBackingStore().Get("allowEmailFromGuestUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAllowGuestUsers gets the allowGuestUsers property value. The allowGuestUsers property
func (m *ProtectGroup) GetAllowGuestUsers()(*bool) {
    val, err := m.GetBackingStore().Get("allowGuestUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["allowEmailFromGuestUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEmailFromGuestUsers(val)
        }
        return nil
    }
    res["allowGuestUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowGuestUsers(val)
        }
        return nil
    }
    res["privacy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtectGroup_privacy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacy(val.(*ProtectGroup_privacy))
        }
        return nil
    }
    return res
}
// GetPrivacy gets the privacy property value. The privacy property
func (m *ProtectGroup) GetPrivacy()(*ProtectGroup_privacy) {
    val, err := m.GetBackingStore().Get("privacy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ProtectGroup_privacy)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProtectGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowEmailFromGuestUsers", m.GetAllowEmailFromGuestUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowGuestUsers", m.GetAllowGuestUsers())
        if err != nil {
            return err
        }
    }
    if m.GetPrivacy() != nil {
        cast := (*m.GetPrivacy()).String()
        err = writer.WriteStringValue("privacy", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowEmailFromGuestUsers sets the allowEmailFromGuestUsers property value. The allowEmailFromGuestUsers property
func (m *ProtectGroup) SetAllowEmailFromGuestUsers(value *bool)() {
    err := m.GetBackingStore().Set("allowEmailFromGuestUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowGuestUsers sets the allowGuestUsers property value. The allowGuestUsers property
func (m *ProtectGroup) SetAllowGuestUsers(value *bool)() {
    err := m.GetBackingStore().Set("allowGuestUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivacy sets the privacy property value. The privacy property
func (m *ProtectGroup) SetPrivacy(value *ProtectGroup_privacy)() {
    err := m.GetBackingStore().Set("privacy", value)
    if err != nil {
        panic(err)
    }
}
// ProtectGroupable 
type ProtectGroupable interface {
    LabelActionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowEmailFromGuestUsers()(*bool)
    GetAllowGuestUsers()(*bool)
    GetPrivacy()(*ProtectGroup_privacy)
    SetAllowEmailFromGuestUsers(value *bool)()
    SetAllowGuestUsers(value *bool)()
    SetPrivacy(value *ProtectGroup_privacy)()
}
