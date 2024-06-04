package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HardwarePasswordInfo intune will provide customer the ability to configure hardware/bios settings on the enrolled windows 10 Azure Active Directory joined devices. Starting from June, 2024 (Intune Release 2406), this type will no longer be supported and will be marked as deprecated
type HardwarePasswordInfo struct {
    Entity
}
// NewHardwarePasswordInfo instantiates a new HardwarePasswordInfo and sets the default values.
func NewHardwarePasswordInfo()(*HardwarePasswordInfo) {
    m := &HardwarePasswordInfo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHardwarePasswordInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwarePasswordInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwarePasswordInfo(), nil
}
// GetCurrentPassword gets the currentPassword property value. Current device password
// returns a *string when successful
func (m *HardwarePasswordInfo) GetCurrentPassword()(*string) {
    val, err := m.GetBackingStore().Get("currentPassword")
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
func (m *HardwarePasswordInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["currentPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPassword(val)
        }
        return nil
    }
    res["previousPasswords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPreviousPasswords(res)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    return res
}
// GetPreviousPasswords gets the previousPasswords property value. List of previous device passwords
// returns a []string when successful
func (m *HardwarePasswordInfo) GetPreviousPasswords()([]string) {
    val, err := m.GetBackingStore().Get("previousPasswords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. Device serial number
// returns a *string when successful
func (m *HardwarePasswordInfo) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *HardwarePasswordInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCurrentPassword sets the currentPassword property value. Current device password
func (m *HardwarePasswordInfo) SetCurrentPassword(value *string)() {
    err := m.GetBackingStore().Set("currentPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviousPasswords sets the previousPasswords property value. List of previous device passwords
func (m *HardwarePasswordInfo) SetPreviousPasswords(value []string)() {
    err := m.GetBackingStore().Set("previousPasswords", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. Device serial number
func (m *HardwarePasswordInfo) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
type HardwarePasswordInfoable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPassword()(*string)
    GetPreviousPasswords()([]string)
    GetSerialNumber()(*string)
    SetCurrentPassword(value *string)()
    SetPreviousPasswords(value []string)()
    SetSerialNumber(value *string)()
}
