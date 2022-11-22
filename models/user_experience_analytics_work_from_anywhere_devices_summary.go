package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereDevicesSummary the user experience analytics Work From Anywhere metrics devices summary.
type UserExperienceAnalyticsWorkFromAnywhereDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value of work from anywhere autopilot devices summary.
    autopilotDevicesSummary UserExperienceAnalyticsAutopilotDevicesSummaryable
    // The user experience analytics work from anywhere Cloud Identity devices summary.
    cloudIdentityDevicesSummary UserExperienceAnalyticsCloudIdentityDevicesSummaryable
    // The user experience work from anywhere Cloud management devices summary.
    cloudManagementDevicesSummary UserExperienceAnalyticsCloudManagementDevicesSummaryable
    // Total number of co-managed devices. Valid values -2147483648 to 2147483647
    coManagedDevices *int32
    // The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
    devicesNotAutopilotRegistered *int32
    // The count of intune devices not autopilot profile assigned. Valid values -2147483648 to 2147483647
    devicesWithoutAutopilotProfileAssigned *int32
    // The count of devices that are not cloud identity. Valid values -2147483648 to 2147483647
    devicesWithoutCloudIdentity *int32
    // The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
    intuneDevices *int32
    // The OdataType property
    odataType *string
    // Total count of tenant attach devices. Valid values -2147483648 to 2147483647
    tenantAttachDevices *int32
    // The total count of devices. Valid values -2147483648 to 2147483647
    totalDevices *int32
    // The count of Windows 10 devices that have unsupported OS versions. Valid values -2147483648 to 2147483647
    unsupportedOSversionDevices *int32
    // The count of windows 10 devices. Valid values -2147483648 to 2147483647
    windows10Devices *int32
    // The user experience analytics work from anywhere Windows 10 devices summary.
    windows10DevicesSummary UserExperienceAnalyticsWindows10DevicesSummaryable
    // The count of windows 10 devices that are Intune and Comanaged. Valid values -2147483648 to 2147483647
    windows10DevicesWithoutTenantAttach *int32
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevicesSummary instantiates a new userExperienceAnalyticsWorkFromAnywhereDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevicesSummary()(*UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWorkFromAnywhereDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereDevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAutopilotDevicesSummary gets the autopilotDevicesSummary property value. The value of work from anywhere autopilot devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetAutopilotDevicesSummary()(UserExperienceAnalyticsAutopilotDevicesSummaryable) {
    return m.autopilotDevicesSummary
}
// GetCloudIdentityDevicesSummary gets the cloudIdentityDevicesSummary property value. The user experience analytics work from anywhere Cloud Identity devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCloudIdentityDevicesSummary()(UserExperienceAnalyticsCloudIdentityDevicesSummaryable) {
    return m.cloudIdentityDevicesSummary
}
// GetCloudManagementDevicesSummary gets the cloudManagementDevicesSummary property value. The user experience work from anywhere Cloud management devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCloudManagementDevicesSummary()(UserExperienceAnalyticsCloudManagementDevicesSummaryable) {
    return m.cloudManagementDevicesSummary
}
// GetCoManagedDevices gets the coManagedDevices property value. Total number of co-managed devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCoManagedDevices()(*int32) {
    return m.coManagedDevices
}
// GetDevicesNotAutopilotRegistered gets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetDevicesNotAutopilotRegistered()(*int32) {
    return m.devicesNotAutopilotRegistered
}
// GetDevicesWithoutAutopilotProfileAssigned gets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetDevicesWithoutAutopilotProfileAssigned()(*int32) {
    return m.devicesWithoutAutopilotProfileAssigned
}
// GetDevicesWithoutCloudIdentity gets the devicesWithoutCloudIdentity property value. The count of devices that are not cloud identity. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetDevicesWithoutCloudIdentity()(*int32) {
    return m.devicesWithoutCloudIdentity
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["autopilotDevicesSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsAutopilotDevicesSummaryFromDiscriminatorValue , m.SetAutopilotDevicesSummary)
    res["cloudIdentityDevicesSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCloudIdentityDevicesSummaryFromDiscriminatorValue , m.SetCloudIdentityDevicesSummary)
    res["cloudManagementDevicesSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsCloudManagementDevicesSummaryFromDiscriminatorValue , m.SetCloudManagementDevicesSummary)
    res["coManagedDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetCoManagedDevices)
    res["devicesNotAutopilotRegistered"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDevicesNotAutopilotRegistered)
    res["devicesWithoutAutopilotProfileAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDevicesWithoutAutopilotProfileAssigned)
    res["devicesWithoutCloudIdentity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDevicesWithoutCloudIdentity)
    res["intuneDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetIntuneDevices)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["tenantAttachDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTenantAttachDevices)
    res["totalDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalDevices)
    res["unsupportedOSversionDevices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUnsupportedOSversionDevices)
    res["windows10Devices"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWindows10Devices)
    res["windows10DevicesSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserExperienceAnalyticsWindows10DevicesSummaryFromDiscriminatorValue , m.SetWindows10DevicesSummary)
    res["windows10DevicesWithoutTenantAttach"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetWindows10DevicesWithoutTenantAttach)
    return res
}
// GetIntuneDevices gets the intuneDevices property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetIntuneDevices()(*int32) {
    return m.intuneDevices
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantAttachDevices gets the tenantAttachDevices property value. Total count of tenant attach devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetTenantAttachDevices()(*int32) {
    return m.tenantAttachDevices
}
// GetTotalDevices gets the totalDevices property value. The total count of devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetTotalDevices()(*int32) {
    return m.totalDevices
}
// GetUnsupportedOSversionDevices gets the unsupportedOSversionDevices property value. The count of Windows 10 devices that have unsupported OS versions. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetUnsupportedOSversionDevices()(*int32) {
    return m.unsupportedOSversionDevices
}
// GetWindows10Devices gets the windows10Devices property value. The count of windows 10 devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10Devices()(*int32) {
    return m.windows10Devices
}
// GetWindows10DevicesSummary gets the windows10DevicesSummary property value. The user experience analytics work from anywhere Windows 10 devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10DevicesSummary()(UserExperienceAnalyticsWindows10DevicesSummaryable) {
    return m.windows10DevicesSummary
}
// GetWindows10DevicesWithoutTenantAttach gets the windows10DevicesWithoutTenantAttach property value. The count of windows 10 devices that are Intune and Comanaged. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10DevicesWithoutTenantAttach()(*int32) {
    return m.windows10DevicesWithoutTenantAttach
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("autopilotDevicesSummary", m.GetAutopilotDevicesSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cloudIdentityDevicesSummary", m.GetCloudIdentityDevicesSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cloudManagementDevicesSummary", m.GetCloudManagementDevicesSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("coManagedDevices", m.GetCoManagedDevices())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteInt32Value("devicesWithoutCloudIdentity", m.GetDevicesWithoutCloudIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("intuneDevices", m.GetIntuneDevices())
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
        err := writer.WriteInt32Value("tenantAttachDevices", m.GetTenantAttachDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalDevices", m.GetTotalDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unsupportedOSversionDevices", m.GetUnsupportedOSversionDevices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows10Devices", m.GetWindows10Devices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("windows10DevicesSummary", m.GetWindows10DevicesSummary())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windows10DevicesWithoutTenantAttach", m.GetWindows10DevicesWithoutTenantAttach())
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAutopilotDevicesSummary sets the autopilotDevicesSummary property value. The value of work from anywhere autopilot devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetAutopilotDevicesSummary(value UserExperienceAnalyticsAutopilotDevicesSummaryable)() {
    m.autopilotDevicesSummary = value
}
// SetCloudIdentityDevicesSummary sets the cloudIdentityDevicesSummary property value. The user experience analytics work from anywhere Cloud Identity devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCloudIdentityDevicesSummary(value UserExperienceAnalyticsCloudIdentityDevicesSummaryable)() {
    m.cloudIdentityDevicesSummary = value
}
// SetCloudManagementDevicesSummary sets the cloudManagementDevicesSummary property value. The user experience work from anywhere Cloud management devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCloudManagementDevicesSummary(value UserExperienceAnalyticsCloudManagementDevicesSummaryable)() {
    m.cloudManagementDevicesSummary = value
}
// SetCoManagedDevices sets the coManagedDevices property value. Total number of co-managed devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCoManagedDevices(value *int32)() {
    m.coManagedDevices = value
}
// SetDevicesNotAutopilotRegistered sets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetDevicesNotAutopilotRegistered(value *int32)() {
    m.devicesNotAutopilotRegistered = value
}
// SetDevicesWithoutAutopilotProfileAssigned sets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetDevicesWithoutAutopilotProfileAssigned(value *int32)() {
    m.devicesWithoutAutopilotProfileAssigned = value
}
// SetDevicesWithoutCloudIdentity sets the devicesWithoutCloudIdentity property value. The count of devices that are not cloud identity. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetDevicesWithoutCloudIdentity(value *int32)() {
    m.devicesWithoutCloudIdentity = value
}
// SetIntuneDevices sets the intuneDevices property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetIntuneDevices(value *int32)() {
    m.intuneDevices = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantAttachDevices sets the tenantAttachDevices property value. Total count of tenant attach devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetTenantAttachDevices(value *int32)() {
    m.tenantAttachDevices = value
}
// SetTotalDevices sets the totalDevices property value. The total count of devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetTotalDevices(value *int32)() {
    m.totalDevices = value
}
// SetUnsupportedOSversionDevices sets the unsupportedOSversionDevices property value. The count of Windows 10 devices that have unsupported OS versions. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetUnsupportedOSversionDevices(value *int32)() {
    m.unsupportedOSversionDevices = value
}
// SetWindows10Devices sets the windows10Devices property value. The count of windows 10 devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10Devices(value *int32)() {
    m.windows10Devices = value
}
// SetWindows10DevicesSummary sets the windows10DevicesSummary property value. The user experience analytics work from anywhere Windows 10 devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10DevicesSummary(value UserExperienceAnalyticsWindows10DevicesSummaryable)() {
    m.windows10DevicesSummary = value
}
// SetWindows10DevicesWithoutTenantAttach sets the windows10DevicesWithoutTenantAttach property value. The count of windows 10 devices that are Intune and Comanaged. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10DevicesWithoutTenantAttach(value *int32)() {
    m.windows10DevicesWithoutTenantAttach = value
}
