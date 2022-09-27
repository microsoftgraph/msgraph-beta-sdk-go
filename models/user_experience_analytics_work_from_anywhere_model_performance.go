package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereModelPerformance the user experience analytics work from anywhere model performance.
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance struct {
    Entity
    // The user experience work from anywhere's cloud identity score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudIdentityScore *float64
    // The user experience work from anywhere's cloud management score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudManagementScore *float64
    // The user experience work from anywhere's cloud provisioning score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudProvisioningScore *float64
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The user experience work from anywhere's manufacturer name of the devices.
    manufacturer *string
    // The user experience work from anywhere's model name of the devices.
    model *string
    // The user experience work from anywhere's devices count for the model. Valid values -2147483648 to 2147483647
    modelDeviceCount *int32
    // The user experience work from anywhere windows score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    windowsScore *float64
    // The user experience work from anywhere overall score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance instantiates a new userExperienceAnalyticsWorkFromAnywhereModelPerformance and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsWorkFromAnywhereModelPerformance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance(), nil
}
// GetCloudIdentityScore gets the cloudIdentityScore property value. The user experience work from anywhere's cloud identity score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudIdentityScore()(*float64) {
    return m.cloudIdentityScore
}
// GetCloudManagementScore gets the cloudManagementScore property value. The user experience work from anywhere's cloud management score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudManagementScore()(*float64) {
    return m.cloudManagementScore
}
// GetCloudProvisioningScore gets the cloudProvisioningScore property value. The user experience work from anywhere's cloud provisioning score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudProvisioningScore()(*float64) {
    return m.cloudProvisioningScore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudIdentityScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCloudIdentityScore)
    res["cloudManagementScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCloudManagementScore)
    res["cloudProvisioningScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCloudProvisioningScore)
    res["healthStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserExperienceAnalyticsHealthState , m.SetHealthStatus)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["modelDeviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetModelDeviceCount)
    res["windowsScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetWindowsScore)
    res["workFromAnywhereScore"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetWorkFromAnywhereScore)
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetManufacturer gets the manufacturer property value. The user experience work from anywhere's manufacturer name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The user experience work from anywhere's model name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetModel()(*string) {
    return m.model
}
// GetModelDeviceCount gets the modelDeviceCount property value. The user experience work from anywhere's devices count for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetModelDeviceCount()(*int32) {
    return m.modelDeviceCount
}
// GetWindowsScore gets the windowsScore property value. The user experience work from anywhere windows score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetWindowsScore()(*float64) {
    return m.windowsScore
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. The user experience work from anywhere overall score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetWorkFromAnywhereScore()(*float64) {
    return m.workFromAnywhereScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("cloudIdentityScore", m.GetCloudIdentityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cloudManagementScore", m.GetCloudManagementScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cloudProvisioningScore", m.GetCloudProvisioningScore())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
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
        err = writer.WriteInt32Value("modelDeviceCount", m.GetModelDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("windowsScore", m.GetWindowsScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("workFromAnywhereScore", m.GetWorkFromAnywhereScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudIdentityScore sets the cloudIdentityScore property value. The user experience work from anywhere's cloud identity score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudIdentityScore(value *float64)() {
    m.cloudIdentityScore = value
}
// SetCloudManagementScore sets the cloudManagementScore property value. The user experience work from anywhere's cloud management score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudManagementScore(value *float64)() {
    m.cloudManagementScore = value
}
// SetCloudProvisioningScore sets the cloudProvisioningScore property value. The user experience work from anywhere's cloud provisioning score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudProvisioningScore(value *float64)() {
    m.cloudProvisioningScore = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetManufacturer sets the manufacturer property value. The user experience work from anywhere's manufacturer name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The user experience work from anywhere's model name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetModel(value *string)() {
    m.model = value
}
// SetModelDeviceCount sets the modelDeviceCount property value. The user experience work from anywhere's devices count for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetModelDeviceCount(value *int32)() {
    m.modelDeviceCount = value
}
// SetWindowsScore sets the windowsScore property value. The user experience work from anywhere windows score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetWindowsScore(value *float64)() {
    m.windowsScore = value
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. The user experience work from anywhere overall score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}
