package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HardwareOathTokenAuthenticationMethodDevice struct {
    AuthenticationMethodDevice
}
// NewHardwareOathTokenAuthenticationMethodDevice instantiates a new HardwareOathTokenAuthenticationMethodDevice and sets the default values.
func NewHardwareOathTokenAuthenticationMethodDevice()(*HardwareOathTokenAuthenticationMethodDevice) {
    m := &HardwareOathTokenAuthenticationMethodDevice{
        AuthenticationMethodDevice: *NewAuthenticationMethodDevice(),
    }
    odataTypeValue := "#microsoft.graph.hardwareOathTokenAuthenticationMethodDevice"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHardwareOathTokenAuthenticationMethodDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareOathTokenAuthenticationMethodDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareOathTokenAuthenticationMethodDevice(), nil
}
// GetAssignedTo gets the assignedTo property value. The assignedTo property
// returns a Identityable when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetAssignedTo()(Identityable) {
    val, err := m.GetBackingStore().Get("assignedTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// GetAssignTo gets the assignTo property value. The assignTo property
// returns a Userable when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetAssignTo()(Userable) {
    val, err := m.GetBackingStore().Get("assignTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Userable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodDevice.GetFieldDeserializers()
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val.(Identityable))
        }
        return nil
    }
    res["assignTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignTo(val.(Userable))
        }
        return nil
    }
    res["hashFunction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHardwareOathTokenHashFunction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashFunction(val.(*HardwareOathTokenHashFunction))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["secretKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretKey(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHardwareOathTokenStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*HardwareOathTokenStatus))
        }
        return nil
    }
    res["timeIntervalInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeIntervalInSeconds(val)
        }
        return nil
    }
    return res
}
// GetHashFunction gets the hashFunction property value. The hashFunction property
// returns a *HardwareOathTokenHashFunction when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetHashFunction()(*HardwareOathTokenHashFunction) {
    val, err := m.GetBackingStore().Get("hashFunction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*HardwareOathTokenHashFunction)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
// returns a *string when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The model property
// returns a *string when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSecretKey gets the secretKey property value. The secretKey property
// returns a *string when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetSecretKey()(*string) {
    val, err := m.GetBackingStore().Get("secretKey")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. The serialNumber property
// returns a *string when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *HardwareOathTokenStatus when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetStatus()(*HardwareOathTokenStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*HardwareOathTokenStatus)
    }
    return nil
}
// GetTimeIntervalInSeconds gets the timeIntervalInSeconds property value. The timeIntervalInSeconds property
// returns a *int32 when successful
func (m *HardwareOathTokenAuthenticationMethodDevice) GetTimeIntervalInSeconds()(*int32) {
    val, err := m.GetBackingStore().Get("timeIntervalInSeconds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *HardwareOathTokenAuthenticationMethodDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodDevice.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("assignedTo", m.GetAssignedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignTo", m.GetAssignTo())
        if err != nil {
            return err
        }
    }
    if m.GetHashFunction() != nil {
        cast := (*m.GetHashFunction()).String()
        err = writer.WriteStringValue("hashFunction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("secretKey", m.GetSecretKey())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("timeIntervalInSeconds", m.GetTimeIntervalInSeconds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The assignedTo property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetAssignedTo(value Identityable)() {
    err := m.GetBackingStore().Set("assignedTo", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignTo sets the assignTo property value. The assignTo property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetAssignTo(value Userable)() {
    err := m.GetBackingStore().Set("assignTo", value)
    if err != nil {
        panic(err)
    }
}
// SetHashFunction sets the hashFunction property value. The hashFunction property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetHashFunction(value *HardwareOathTokenHashFunction)() {
    err := m.GetBackingStore().Set("hashFunction", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The model property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetSecretKey sets the secretKey property value. The secretKey property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetSecretKey(value *string)() {
    err := m.GetBackingStore().Set("secretKey", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. The serialNumber property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetStatus(value *HardwareOathTokenStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeIntervalInSeconds sets the timeIntervalInSeconds property value. The timeIntervalInSeconds property
func (m *HardwareOathTokenAuthenticationMethodDevice) SetTimeIntervalInSeconds(value *int32)() {
    err := m.GetBackingStore().Set("timeIntervalInSeconds", value)
    if err != nil {
        panic(err)
    }
}
type HardwareOathTokenAuthenticationMethodDeviceable interface {
    AuthenticationMethodDeviceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTo()(Identityable)
    GetAssignTo()(Userable)
    GetHashFunction()(*HardwareOathTokenHashFunction)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetSecretKey()(*string)
    GetSerialNumber()(*string)
    GetStatus()(*HardwareOathTokenStatus)
    GetTimeIntervalInSeconds()(*int32)
    SetAssignedTo(value Identityable)()
    SetAssignTo(value Userable)()
    SetHashFunction(value *HardwareOathTokenHashFunction)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetSecretKey(value *string)()
    SetSerialNumber(value *string)()
    SetStatus(value *HardwareOathTokenStatus)()
    SetTimeIntervalInSeconds(value *int32)()
}
