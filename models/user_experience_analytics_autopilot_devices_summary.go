package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAutopilotDevicesSummary the user experience analytics summary of Devices not windows autopilot ready.
type UserExperienceAnalyticsAutopilotDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The count of intune devices that are not autopilot registerd.
    devicesNotAutopilotRegistered *int32
    // The count of intune devices not autopilot profile assigned.
    devicesWithoutAutopilotProfileAssigned *int32
    // The OdataType property
    odataType *string
    // The count of windows 10 devices that are Intune and Comanaged.
    totalWindows10DevicesWithoutTenantAttached *int32
}
// NewUserExperienceAnalyticsAutopilotDevicesSummary instantiates a new userExperienceAnalyticsAutopilotDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsAutopilotDevicesSummary()(*UserExperienceAnalyticsAutopilotDevicesSummary) {
    m := &UserExperienceAnalyticsAutopilotDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsAutopilotDevicesSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsAutopilotDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAutopilotDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAutopilotDevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDevicesNotAutopilotRegistered gets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetDevicesNotAutopilotRegistered()(*int32) {
    return m.devicesNotAutopilotRegistered
}
// GetDevicesWithoutAutopilotProfileAssigned gets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetDevicesWithoutAutopilotProfileAssigned()(*int32) {
    return m.devicesWithoutAutopilotProfileAssigned
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["devicesNotAutopilotRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDevicesNotAutopilotRegistered)
    res["devicesWithoutAutopilotProfileAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDevicesWithoutAutopilotProfileAssigned)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["totalWindows10DevicesWithoutTenantAttached"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalWindows10DevicesWithoutTenantAttached)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetTotalWindows10DevicesWithoutTenantAttached gets the totalWindows10DevicesWithoutTenantAttached property value. The count of windows 10 devices that are Intune and Comanaged.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) GetTotalWindows10DevicesWithoutTenantAttached()(*int32) {
    return m.totalWindows10DevicesWithoutTenantAttached
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("devicesNotAutopilotRegistered", m.GetDevicesNotAutopilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("devicesWithoutAutopilotProfileAssigned", m.GetDevicesWithoutAutopilotProfileAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalWindows10DevicesWithoutTenantAttached", m.GetTotalWindows10DevicesWithoutTenantAttached())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDevicesNotAutopilotRegistered sets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetDevicesNotAutopilotRegistered(value *int32)() {
    m.devicesNotAutopilotRegistered = value
}
// SetDevicesWithoutAutopilotProfileAssigned sets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetDevicesWithoutAutopilotProfileAssigned(value *int32)() {
    m.devicesWithoutAutopilotProfileAssigned = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTotalWindows10DevicesWithoutTenantAttached sets the totalWindows10DevicesWithoutTenantAttached property value. The count of windows 10 devices that are Intune and Comanaged.
func (m *UserExperienceAnalyticsAutopilotDevicesSummary) SetTotalWindows10DevicesWithoutTenantAttached(value *int32)() {
    m.totalWindows10DevicesWithoutTenantAttached = value
}
