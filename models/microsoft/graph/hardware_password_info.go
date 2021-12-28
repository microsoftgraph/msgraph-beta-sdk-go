package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HardwarePasswordInfo 
type HardwarePasswordInfo struct {
    Entity
    // Current device password
    currentPassword *string;
    // List of previous device passwords
    previousPasswords []string;
    // Device serial number
    serialNumber *string;
}
// NewHardwarePasswordInfo instantiates a new hardwarePasswordInfo and sets the default values.
func NewHardwarePasswordInfo()(*HardwarePasswordInfo) {
    m := &HardwarePasswordInfo{
        Entity: *NewEntity(),
    }
    return m
}
// GetCurrentPassword gets the currentPassword property value. Current device password
func (m *HardwarePasswordInfo) GetCurrentPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.currentPassword
    }
}
// GetPreviousPasswords gets the previousPasswords property value. List of previous device passwords
func (m *HardwarePasswordInfo) GetPreviousPasswords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.previousPasswords
    }
}
// GetSerialNumber gets the serialNumber property value. Device serial number
func (m *HardwarePasswordInfo) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwarePasswordInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["currentPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPassword(val)
        }
        return nil
    }
    res["previousPasswords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPreviousPasswords(res)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *HardwarePasswordInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HardwarePasswordInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("currentPassword", m.GetCurrentPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("previousPasswords", m.GetPreviousPasswords())
        if err != nil {
            return err
        }
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
    if m != nil {
        m.currentPassword = value
    }
}
// SetPreviousPasswords sets the previousPasswords property value. List of previous device passwords
func (m *HardwarePasswordInfo) SetPreviousPasswords(value []string)() {
    if m != nil {
        m.previousPasswords = value
    }
}
// SetSerialNumber sets the serialNumber property value. Device serial number
func (m *HardwarePasswordInfo) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
