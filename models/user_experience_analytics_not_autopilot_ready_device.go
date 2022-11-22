package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsNotAutopilotReadyDevice the user experience analytics Device not windows autopilot ready.
type UserExperienceAnalyticsNotAutopilotReadyDevice struct {
    Entity
    // The intune device's autopilotProfileAssigned.
    autoPilotProfileAssigned *bool
    // The intune device's autopilotRegistered.
    autoPilotRegistered *bool
    // The intune device's azure Ad joinType.
    azureAdJoinType *string
    // The intune device's azureAdRegistered.
    azureAdRegistered *bool
    // The intune device's name.
    deviceName *string
    // The intune device's managed by.
    managedBy *string
    // The intune device's manufacturer.
    manufacturer *string
    // The intune device's model.
    model *string
    // The intune device's serial number.
    serialNumber *string
}
// NewUserExperienceAnalyticsNotAutopilotReadyDevice instantiates a new userExperienceAnalyticsNotAutopilotReadyDevice and sets the default values.
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
    return m.autoPilotProfileAssigned
}
// GetAutoPilotRegistered gets the autoPilotRegistered property value. The intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotRegistered()(*bool) {
    return m.autoPilotRegistered
}
// GetAzureAdJoinType gets the azureAdJoinType property value. The intune device's azure Ad joinType.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdJoinType()(*string) {
    return m.azureAdJoinType
}
// GetAzureAdRegistered gets the azureAdRegistered property value. The intune device's azureAdRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdRegistered()(*bool) {
    return m.azureAdRegistered
}
// GetDeviceName gets the deviceName property value. The intune device's name.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAutoPilotProfileAssigned)
    res["autoPilotRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAutoPilotRegistered)
    res["azureAdJoinType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureAdJoinType)
    res["azureAdRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAzureAdRegistered)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["managedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedBy)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["serialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSerialNumber)
    return res
}
// GetManagedBy gets the managedBy property value. The intune device's managed by.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManagedBy()(*string) {
    return m.managedBy
}
// GetManufacturer gets the manufacturer property value. The intune device's manufacturer.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The intune device's model.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetModel()(*string) {
    return m.model
}
// GetSerialNumber gets the serialNumber property value. The intune device's serial number.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetSerialNumber()(*string) {
    return m.serialNumber
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
    m.autoPilotProfileAssigned = value
}
// SetAutoPilotRegistered sets the autoPilotRegistered property value. The intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotRegistered(value *bool)() {
    m.autoPilotRegistered = value
}
// SetAzureAdJoinType sets the azureAdJoinType property value. The intune device's azure Ad joinType.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdJoinType(value *string)() {
    m.azureAdJoinType = value
}
// SetAzureAdRegistered sets the azureAdRegistered property value. The intune device's azureAdRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdRegistered(value *bool)() {
    m.azureAdRegistered = value
}
// SetDeviceName sets the deviceName property value. The intune device's name.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetManagedBy sets the managedBy property value. The intune device's managed by.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManagedBy(value *string)() {
    m.managedBy = value
}
// SetManufacturer sets the manufacturer property value. The intune device's manufacturer.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The intune device's model.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetModel(value *string)() {
    m.model = value
}
// SetSerialNumber sets the serialNumber property value. The intune device's serial number.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
