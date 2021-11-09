package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsWorkFromAnywhereDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The value of work from anywhere autopilot devices summary.
    autopilotDevicesSummary *UserExperienceAnalyticsAutopilotDevicesSummary;
    // The user experience analytics work from anywhere Cloud Identity devices summary.
    cloudIdentityDevicesSummary *UserExperienceAnalyticsCloudIdentityDevicesSummary;
    // The user experience work from anywhere Cloud management devices summary.
    cloudManagementDevicesSummary *UserExperienceAnalyticsCloudManagementDevicesSummary;
    // Total number of co-managed devices. Valid values -2147483648 to 2147483647
    coManagedDevices *int32;
    // The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
    devicesNotAutopilotRegistered *int32;
    // The count of intune devices not autopilot profile assigned. Valid values -2147483648 to 2147483647
    devicesWithoutAutopilotProfileAssigned *int32;
    // The count of devices that are not cloud identity. Valid values -2147483648 to 2147483647
    devicesWithoutCloudIdentity *int32;
    // The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
    intuneDevices *int32;
    // Total count of tenant attach devices. Valid values -2147483648 to 2147483647
    tenantAttachDevices *int32;
    // The total count of devices. Valid values -2147483648 to 2147483647
    totalDevices *int32;
    // The count of Windows 10 devices that have unsupported OS versions. Valid values -2147483648 to 2147483647
    unsupportedOSversionDevices *int32;
    // The count of windows 10 devices. Valid values -2147483648 to 2147483647
    windows10Devices *int32;
    // The user experience analytics work from anywhere Windows 10 devices summary.
    windows10DevicesSummary *UserExperienceAnalyticsWindows10DevicesSummary;
    // The count of windows 10 devices that are Intune and Comanaged. Valid values -2147483648 to 2147483647
    windows10DevicesWithoutTenantAttach *int32;
}
// Instantiates a new userExperienceAnalyticsWorkFromAnywhereDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevicesSummary()(*UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the autopilotDevicesSummary property value. The value of work from anywhere autopilot devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetAutopilotDevicesSummary()(*UserExperienceAnalyticsAutopilotDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.autopilotDevicesSummary
    }
}
// Gets the cloudIdentityDevicesSummary property value. The user experience analytics work from anywhere Cloud Identity devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCloudIdentityDevicesSummary()(*UserExperienceAnalyticsCloudIdentityDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.cloudIdentityDevicesSummary
    }
}
// Gets the cloudManagementDevicesSummary property value. The user experience work from anywhere Cloud management devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCloudManagementDevicesSummary()(*UserExperienceAnalyticsCloudManagementDevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.cloudManagementDevicesSummary
    }
}
// Gets the coManagedDevices property value. Total number of co-managed devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetCoManagedDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.coManagedDevices
    }
}
// Gets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetDevicesNotAutopilotRegistered()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesNotAutopilotRegistered
    }
}
// Gets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetDevicesWithoutAutopilotProfileAssigned()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesWithoutAutopilotProfileAssigned
    }
}
// Gets the devicesWithoutCloudIdentity property value. The count of devices that are not cloud identity. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetDevicesWithoutCloudIdentity()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.devicesWithoutCloudIdentity
    }
}
// Gets the intuneDevices property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetIntuneDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.intuneDevices
    }
}
// Gets the tenantAttachDevices property value. Total count of tenant attach devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetTenantAttachDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.tenantAttachDevices
    }
}
// Gets the totalDevices property value. The total count of devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetTotalDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalDevices
    }
}
// Gets the unsupportedOSversionDevices property value. The count of Windows 10 devices that have unsupported OS versions. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetUnsupportedOSversionDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unsupportedOSversionDevices
    }
}
// Gets the windows10Devices property value. The count of windows 10 devices. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10Devices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windows10Devices
    }
}
// Gets the windows10DevicesSummary property value. The user experience analytics work from anywhere Windows 10 devices summary.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10DevicesSummary()(*UserExperienceAnalyticsWindows10DevicesSummary) {
    if m == nil {
        return nil
    } else {
        return m.windows10DevicesSummary
    }
}
// Gets the windows10DevicesWithoutTenantAttach property value. The count of windows 10 devices that are Intune and Comanaged. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetWindows10DevicesWithoutTenantAttach()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windows10DevicesWithoutTenantAttach
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["autopilotDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsAutopilotDevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutopilotDevicesSummary(val.(*UserExperienceAnalyticsAutopilotDevicesSummary))
        }
        return nil
    }
    res["cloudIdentityDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCloudIdentityDevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudIdentityDevicesSummary(val.(*UserExperienceAnalyticsCloudIdentityDevicesSummary))
        }
        return nil
    }
    res["cloudManagementDevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsCloudManagementDevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudManagementDevicesSummary(val.(*UserExperienceAnalyticsCloudManagementDevicesSummary))
        }
        return nil
    }
    res["coManagedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoManagedDevices(val)
        }
        return nil
    }
    res["devicesNotAutopilotRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesNotAutopilotRegistered(val)
        }
        return nil
    }
    res["devicesWithoutAutopilotProfileAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesWithoutAutopilotProfileAssigned(val)
        }
        return nil
    }
    res["devicesWithoutCloudIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesWithoutCloudIdentity(val)
        }
        return nil
    }
    res["intuneDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneDevices(val)
        }
        return nil
    }
    res["tenantAttachDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttachDevices(val)
        }
        return nil
    }
    res["totalDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDevices(val)
        }
        return nil
    }
    res["unsupportedOSversionDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsupportedOSversionDevices(val)
        }
        return nil
    }
    res["windows10Devices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows10Devices(val)
        }
        return nil
    }
    res["windows10DevicesSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserExperienceAnalyticsWindows10DevicesSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows10DevicesSummary(val.(*UserExperienceAnalyticsWindows10DevicesSummary))
        }
        return nil
    }
    res["windows10DevicesWithoutTenantAttach"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows10DevicesWithoutTenantAttach(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the autopilotDevicesSummary property value. The value of work from anywhere autopilot devices summary.
// Parameters:
//  - value : Value to set for the autopilotDevicesSummary property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetAutopilotDevicesSummary(value *UserExperienceAnalyticsAutopilotDevicesSummary)() {
    m.autopilotDevicesSummary = value
}
// Sets the cloudIdentityDevicesSummary property value. The user experience analytics work from anywhere Cloud Identity devices summary.
// Parameters:
//  - value : Value to set for the cloudIdentityDevicesSummary property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCloudIdentityDevicesSummary(value *UserExperienceAnalyticsCloudIdentityDevicesSummary)() {
    m.cloudIdentityDevicesSummary = value
}
// Sets the cloudManagementDevicesSummary property value. The user experience work from anywhere Cloud management devices summary.
// Parameters:
//  - value : Value to set for the cloudManagementDevicesSummary property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCloudManagementDevicesSummary(value *UserExperienceAnalyticsCloudManagementDevicesSummary)() {
    m.cloudManagementDevicesSummary = value
}
// Sets the coManagedDevices property value. Total number of co-managed devices. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the coManagedDevices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetCoManagedDevices(value *int32)() {
    m.coManagedDevices = value
}
// Sets the devicesNotAutopilotRegistered property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the devicesNotAutopilotRegistered property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetDevicesNotAutopilotRegistered(value *int32)() {
    m.devicesNotAutopilotRegistered = value
}
// Sets the devicesWithoutAutopilotProfileAssigned property value. The count of intune devices not autopilot profile assigned. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the devicesWithoutAutopilotProfileAssigned property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetDevicesWithoutAutopilotProfileAssigned(value *int32)() {
    m.devicesWithoutAutopilotProfileAssigned = value
}
// Sets the devicesWithoutCloudIdentity property value. The count of devices that are not cloud identity. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the devicesWithoutCloudIdentity property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetDevicesWithoutCloudIdentity(value *int32)() {
    m.devicesWithoutCloudIdentity = value
}
// Sets the intuneDevices property value. The count of intune devices that are not autopilot registerd. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the intuneDevices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetIntuneDevices(value *int32)() {
    m.intuneDevices = value
}
// Sets the tenantAttachDevices property value. Total count of tenant attach devices. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the tenantAttachDevices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetTenantAttachDevices(value *int32)() {
    m.tenantAttachDevices = value
}
// Sets the totalDevices property value. The total count of devices. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the totalDevices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetTotalDevices(value *int32)() {
    m.totalDevices = value
}
// Sets the unsupportedOSversionDevices property value. The count of Windows 10 devices that have unsupported OS versions. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the unsupportedOSversionDevices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetUnsupportedOSversionDevices(value *int32)() {
    m.unsupportedOSversionDevices = value
}
// Sets the windows10Devices property value. The count of windows 10 devices. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the windows10Devices property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10Devices(value *int32)() {
    m.windows10Devices = value
}
// Sets the windows10DevicesSummary property value. The user experience analytics work from anywhere Windows 10 devices summary.
// Parameters:
//  - value : Value to set for the windows10DevicesSummary property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10DevicesSummary(value *UserExperienceAnalyticsWindows10DevicesSummary)() {
    m.windows10DevicesSummary = value
}
// Sets the windows10DevicesWithoutTenantAttach property value. The count of windows 10 devices that are Intune and Comanaged. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the windows10DevicesWithoutTenantAttach property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) SetWindows10DevicesWithoutTenantAttach(value *int32)() {
    m.windows10DevicesWithoutTenantAttach = value
}
