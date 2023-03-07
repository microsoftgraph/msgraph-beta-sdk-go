package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OemWarrantyInformationOnboarding warranty status entity for a given OEM
type OemWarrantyInformationOnboarding struct {
    Entity
}
// NewOemWarrantyInformationOnboarding instantiates a new oemWarrantyInformationOnboarding and sets the default values.
func NewOemWarrantyInformationOnboarding()(*OemWarrantyInformationOnboarding) {
    m := &OemWarrantyInformationOnboarding{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOemWarrantyInformationOnboardingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOemWarrantyInformationOnboardingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOemWarrantyInformationOnboarding(), nil
}
// GetAvailable gets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetAvailable()(*bool) {
    val, err := m.GetBackingStore().Get("available")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetEnabled gets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarrantyInformationOnboarding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["available"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailable(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["oemName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOemName(val)
        }
        return nil
    }
    return res
}
// GetOemName gets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) GetOemName()(*string) {
    val, err := m.GetBackingStore().Get("oemName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OemWarrantyInformationOnboarding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetAvailable sets the available property value. Specifies whether warranty API is available. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetAvailable(value *bool)() {
    err := m.GetBackingStore().Set("available", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. Specifies whether warranty query is enabled for given OEM. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOemName sets the oemName property value. OEM name. This property is read-only.
func (m *OemWarrantyInformationOnboarding) SetOemName(value *string)() {
    err := m.GetBackingStore().Set("oemName", value)
    if err != nil {
        panic(err)
    }
}
// OemWarrantyInformationOnboardingable 
type OemWarrantyInformationOnboardingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailable()(*bool)
    GetEnabled()(*bool)
    GetOemName()(*string)
    SetAvailable(value *bool)()
    SetEnabled(value *bool)()
    SetOemName(value *string)()
}
