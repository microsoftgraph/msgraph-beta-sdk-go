package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsNotAutopilotReadyDevice 
type UserExperienceAnalyticsNotAutopilotReadyDevice struct {
    Entity
}
// NewUserExperienceAnalyticsNotAutopilotReadyDevice instantiates a new UserExperienceAnalyticsNotAutopilotReadyDevice and sets the default values.
func NewUserExperienceAnalyticsNotAutopilotReadyDevice()(*UserExperienceAnalyticsNotAutopilotReadyDevice) {
    m := &UserExperienceAnalyticsNotAutopilotReadyDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsNotAutopilotReadyDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsNotAutopilotReadyDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsNotAutopilotReadyDevice(), nil
}
// GetAutoPilotProfileAssigned gets the autoPilotProfileAssigned property value. The intune device's autopilotProfileAssigned.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotProfileAssigned()(*bool) {
    val, err := m.GetBackingStore().Get("autoPilotProfileAssigned")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoPilotRegistered gets the autoPilotRegistered property value. The intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("autoPilotRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAzureAdJoinType gets the azureAdJoinType property value. The intune device's azure Ad joinType.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdJoinType()(*string) {
    val, err := m.GetBackingStore().Get("azureAdJoinType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureAdRegistered gets the azureAdRegistered property value. The intune device's azureAdRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdRegistered()(*bool) {
    val, err := m.GetBackingStore().Get("azureAdRegistered")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The intune device's name.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotProfileAssigned(val)
        }
        return nil
    }
    res["autoPilotRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotRegistered(val)
        }
        return nil
    }
    res["azureAdJoinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdJoinType(val)
        }
        return nil
    }
    res["azureAdRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdRegistered(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["managedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val)
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
// GetManagedBy gets the managedBy property value. The intune device's managed by.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManagedBy()(*string) {
    val, err := m.GetBackingStore().Get("managedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. The intune device's manufacturer.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The intune device's model.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. The intune device's serial number.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetSerialNumber()(*string) {
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
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("autoPilotProfileAssigned", m.GetAutoPilotProfileAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoPilotRegistered", m.GetAutoPilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdJoinType", m.GetAzureAdJoinType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureAdRegistered", m.GetAzureAdRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedBy", m.GetManagedBy())
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
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutoPilotProfileAssigned sets the autoPilotProfileAssigned property value. The intune device's autopilotProfileAssigned.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotProfileAssigned(value *bool)() {
    err := m.GetBackingStore().Set("autoPilotProfileAssigned", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoPilotRegistered sets the autoPilotRegistered property value. The intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotRegistered(value *bool)() {
    err := m.GetBackingStore().Set("autoPilotRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdJoinType sets the azureAdJoinType property value. The intune device's azure Ad joinType.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdJoinType(value *string)() {
    err := m.GetBackingStore().Set("azureAdJoinType", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureAdRegistered sets the azureAdRegistered property value. The intune device's azureAdRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdRegistered(value *bool)() {
    err := m.GetBackingStore().Set("azureAdRegistered", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The intune device's name.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedBy sets the managedBy property value. The intune device's managed by.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManagedBy(value *string)() {
    err := m.GetBackingStore().Set("managedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The intune device's manufacturer.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The intune device's model.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. The intune device's serial number.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsNotAutopilotReadyDeviceable 
type UserExperienceAnalyticsNotAutopilotReadyDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoPilotProfileAssigned()(*bool)
    GetAutoPilotRegistered()(*bool)
    GetAzureAdJoinType()(*string)
    GetAzureAdRegistered()(*bool)
    GetDeviceName()(*string)
    GetManagedBy()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetSerialNumber()(*string)
    SetAutoPilotProfileAssigned(value *bool)()
    SetAutoPilotRegistered(value *bool)()
    SetAzureAdJoinType(value *string)()
    SetAzureAdRegistered(value *bool)()
    SetDeviceName(value *string)()
    SetManagedBy(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetSerialNumber(value *string)()
}
