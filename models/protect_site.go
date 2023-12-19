package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProtectSite 
type ProtectSite struct {
    LabelActionBase
}
// NewProtectSite instantiates a new protectSite and sets the default values.
func NewProtectSite()(*ProtectSite) {
    m := &ProtectSite{
        LabelActionBase: *NewLabelActionBase(),
    }
    odataTypeValue := "#microsoft.graph.protectSite"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProtectSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProtectSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProtectSite(), nil
}
// GetAccessType gets the accessType property value. The accessType property
func (m *ProtectSite) GetAccessType()(*ProtectSite_accessType) {
    val, err := m.GetBackingStore().Get("accessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ProtectSite_accessType)
    }
    return nil
}
// GetConditionalAccessProtectionLevelId gets the conditionalAccessProtectionLevelId property value. The conditionalAccessProtectionLevelId property
func (m *ProtectSite) GetConditionalAccessProtectionLevelId()(*string) {
    val, err := m.GetBackingStore().Get("conditionalAccessProtectionLevelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProtectSite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["accessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProtectSite_accessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessType(val.(*ProtectSite_accessType))
        }
        return nil
    }
    res["conditionalAccessProtectionLevelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessProtectionLevelId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProtectSite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessType() != nil {
        cast := (*m.GetAccessType()).String()
        err = writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conditionalAccessProtectionLevelId", m.GetConditionalAccessProtectionLevelId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessType sets the accessType property value. The accessType property
func (m *ProtectSite) SetAccessType(value *ProtectSite_accessType)() {
    err := m.GetBackingStore().Set("accessType", value)
    if err != nil {
        panic(err)
    }
}
// SetConditionalAccessProtectionLevelId sets the conditionalAccessProtectionLevelId property value. The conditionalAccessProtectionLevelId property
func (m *ProtectSite) SetConditionalAccessProtectionLevelId(value *string)() {
    err := m.GetBackingStore().Set("conditionalAccessProtectionLevelId", value)
    if err != nil {
        panic(err)
    }
}
// ProtectSiteable 
type ProtectSiteable interface {
    LabelActionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessType()(*ProtectSite_accessType)
    GetConditionalAccessProtectionLevelId()(*string)
    SetAccessType(value *ProtectSite_accessType)()
    SetConditionalAccessProtectionLevelId(value *string)()
}
