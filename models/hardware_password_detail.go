package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HardwarePasswordDetail device BIOS password information for devices with managed BIOS and firmware configuration, which provides device serial number, list of previous passwords, and current password.
type HardwarePasswordDetail struct {
    Entity
}
// NewHardwarePasswordDetail instantiates a new HardwarePasswordDetail and sets the default values.
func NewHardwarePasswordDetail()(*HardwarePasswordDetail) {
    m := &HardwarePasswordDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateHardwarePasswordDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwarePasswordDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwarePasswordDetail(), nil
}
// GetCurrentPassword gets the currentPassword property value. The current device's BIOS password. Supports: $filter, $select, $top, $OrderBy, $skip. This property is read-only.
// returns a *string when successful
func (m *HardwarePasswordDetail) GetCurrentPassword()(*string) {
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
func (m *HardwarePasswordDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetPreviousPasswords gets the previousPasswords property value. The list of all the previous device BIOS passwords. Supports: $filter, $select, $top, $skip. This property is read-only.
// returns a []string when successful
func (m *HardwarePasswordDetail) GetPreviousPasswords()([]string) {
    val, err := m.GetBackingStore().Get("previousPasswords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. The device serial number as defined by the device manufacturer. Supports: $filter, $select, $top, $OrderBy, $skip. This property is read-only.
// returns a *string when successful
func (m *HardwarePasswordDetail) GetSerialNumber()(*string) {
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
func (m *HardwarePasswordDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetCurrentPassword sets the currentPassword property value. The current device's BIOS password. Supports: $filter, $select, $top, $OrderBy, $skip. This property is read-only.
func (m *HardwarePasswordDetail) SetCurrentPassword(value *string)() {
    err := m.GetBackingStore().Set("currentPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviousPasswords sets the previousPasswords property value. The list of all the previous device BIOS passwords. Supports: $filter, $select, $top, $skip. This property is read-only.
func (m *HardwarePasswordDetail) SetPreviousPasswords(value []string)() {
    err := m.GetBackingStore().Set("previousPasswords", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. The device serial number as defined by the device manufacturer. Supports: $filter, $select, $top, $OrderBy, $skip. This property is read-only.
func (m *HardwarePasswordDetail) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
type HardwarePasswordDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPassword()(*string)
    GetPreviousPasswords()([]string)
    GetSerialNumber()(*string)
    SetCurrentPassword(value *string)()
    SetPreviousPasswords(value []string)()
    SetSerialNumber(value *string)()
}
