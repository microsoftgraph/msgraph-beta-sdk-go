package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsWorkFromAnywhereDevice struct {
    Entity
    // The user experience analytics work from anywhere intune device's autopilotProfileAssigned.
    autoPilotProfileAssigned *bool;
    // The user experience work from anywhere intune device's autopilotRegistered.
    autoPilotRegistered *bool;
    // The user experience work from anywhere azure Ad device Id.
    azureAdDeviceId *string;
    // The user experience work from anywhere device's azure Ad joinType.
    azureAdJoinType *string;
    // The user experience work from anywhere device's azureAdRegistered.
    azureAdRegistered *bool;
    // The user experience work from anywhere per device cloud identity score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudIdentityScore *float64;
    // The user experience work from anywhere per device cloud management score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudManagementScore *float64;
    // The user experience work from anywhere per device cloud provisioning score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudProvisioningScore *float64;
    // The user experience work from anywhere device's compliancePolicySetToIntune.
    compliancePolicySetToIntune *bool;
    // The user experience work from anywhere device Id.
    deviceId *string;
    // The work from anywhere device's name.
    deviceName *string;
    // The user experience work from anywhere per device health status. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
    healthStatus *UserExperienceAnalyticsHealthState;
    // The user experience work from anywhere device's Cloud Management Gateway for Configuration Manager is enabled.
    isCloudManagedGatewayEnabled *bool;
    // The user experience work from anywhere management agent of the device.
    managedBy *string;
    // The user experience work from anywhere device's manufacturer.
    manufacturer *string;
    // The user experience work from anywhere device's model.
    model *string;
    // The user experience work from anywhere device, Is OS check failed for device to upgrade to the latest version of windows.
    osCheckFailed *bool;
    // The user experience work from anywhere device's OS Description.
    osDescription *string;
    // The user experience work from anywhere device's OS Version.
    osVersion *string;
    // The user experience work from anywhere device's otherWorkloadsSetToIntune.
    otherWorkloadsSetToIntune *bool;
    // The user experience work from anywhere device's ownership.
    ownership *string;
    // The user experience work from anywhere device, Is processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows.
    processor64BitCheckFailed *bool;
    // The user experience work from anywhere device, Is processor hardware core count check failed for device to upgrade to the latest version of windows.
    processorCoreCountCheckFailed *bool;
    // The user experience work from anywhere device, Is processor hardware family check failed for device to upgrade to the latest version of windows.
    processorFamilyCheckFailed *bool;
    // The user experience work from anywhere device, Is processor hardware speed check failed for device to upgrade to the latest version of windows.
    processorSpeedCheckFailed *bool;
    // Is the user experience analytics work from anywhere device RAM hardware check failed for device to upgrade to the latest version of windows
    ramCheckFailed *bool;
    // The user experience work from anywhere device, Is secure boot hardware check failed for device to upgrade to the latest version of windows.
    secureBootCheckFailed *bool;
    // The user experience work from anywhere device's serial number.
    serialNumber *string;
    // The user experience work from anywhere device, Is storage hardware check failed for device to upgrade to the latest version of windows.
    storageCheckFailed *bool;
    // The user experience work from anywhere device's tenantAttached.
    tenantAttached *bool;
    // The user experience work from anywhere device, Is Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows.
    tpmCheckFailed *bool;
    // The user experience work from anywhere windows upgrade eligibility status of device. Possible values are: upgraded, unknown, notCapable, capable.
    upgradeEligibility *OperatingSystemUpgradeEligibility;
    // The user experience work from anywhere per device windows score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    windowsScore *float64;
    // The user experience work from anywhere per device overall score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64;
}
// Instantiates a new userExperienceAnalyticsWorkFromAnywhereDevice and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice()(*UserExperienceAnalyticsWorkFromAnywhereDevice) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the autoPilotProfileAssigned property value. The user experience analytics work from anywhere intune device's autopilotProfileAssigned.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAutoPilotProfileAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoPilotProfileAssigned
    }
}
// Gets the autoPilotRegistered property value. The user experience work from anywhere intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAutoPilotRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoPilotRegistered
    }
}
// Gets the azureAdDeviceId property value. The user experience work from anywhere azure Ad device Id.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdDeviceId
    }
}
// Gets the azureAdJoinType property value. The user experience work from anywhere device's azure Ad joinType.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdJoinType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdJoinType
    }
}
// Gets the azureAdRegistered property value. The user experience work from anywhere device's azureAdRegistered.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureAdRegistered
    }
}
// Gets the cloudIdentityScore property value. The user experience work from anywhere per device cloud identity score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudIdentityScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudIdentityScore
    }
}
// Gets the cloudManagementScore property value. The user experience work from anywhere per device cloud management score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudManagementScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudManagementScore
    }
}
// Gets the cloudProvisioningScore property value. The user experience work from anywhere per device cloud provisioning score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudProvisioningScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudProvisioningScore
    }
}
// Gets the compliancePolicySetToIntune property value. The user experience work from anywhere device's compliancePolicySetToIntune.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCompliancePolicySetToIntune()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicySetToIntune
    }
}
// Gets the deviceId property value. The user experience work from anywhere device Id.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. The work from anywhere device's name.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the healthStatus property value. The user experience work from anywhere per device health status. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// Gets the isCloudManagedGatewayEnabled property value. The user experience work from anywhere device's Cloud Management Gateway for Configuration Manager is enabled.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetIsCloudManagedGatewayEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCloudManagedGatewayEnabled
    }
}
// Gets the managedBy property value. The user experience work from anywhere management agent of the device.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetManagedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedBy
    }
}
// Gets the manufacturer property value. The user experience work from anywhere device's manufacturer.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. The user experience work from anywhere device's model.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the osCheckFailed property value. The user experience work from anywhere device, Is OS check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.osCheckFailed
    }
}
// Gets the osDescription property value. The user experience work from anywhere device's OS Description.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// Gets the osVersion property value. The user experience work from anywhere device's OS Version.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the otherWorkloadsSetToIntune property value. The user experience work from anywhere device's otherWorkloadsSetToIntune.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOtherWorkloadsSetToIntune()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.otherWorkloadsSetToIntune
    }
}
// Gets the ownership property value. The user experience work from anywhere device's ownership.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOwnership()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownership
    }
}
// Gets the processor64BitCheckFailed property value. The user experience work from anywhere device, Is processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessor64BitCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processor64BitCheckFailed
    }
}
// Gets the processorCoreCountCheckFailed property value. The user experience work from anywhere device, Is processor hardware core count check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorCoreCountCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processorCoreCountCheckFailed
    }
}
// Gets the processorFamilyCheckFailed property value. The user experience work from anywhere device, Is processor hardware family check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorFamilyCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processorFamilyCheckFailed
    }
}
// Gets the processorSpeedCheckFailed property value. The user experience work from anywhere device, Is processor hardware speed check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorSpeedCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processorSpeedCheckFailed
    }
}
// Gets the ramCheckFailed property value. Is the user experience analytics work from anywhere device RAM hardware check failed for device to upgrade to the latest version of windows
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetRamCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ramCheckFailed
    }
}
// Gets the secureBootCheckFailed property value. The user experience work from anywhere device, Is secure boot hardware check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetSecureBootCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.secureBootCheckFailed
    }
}
// Gets the serialNumber property value. The user experience work from anywhere device's serial number.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// Gets the storageCheckFailed property value. The user experience work from anywhere device, Is storage hardware check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetStorageCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.storageCheckFailed
    }
}
// Gets the tenantAttached property value. The user experience work from anywhere device's tenantAttached.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetTenantAttached()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.tenantAttached
    }
}
// Gets the tpmCheckFailed property value. The user experience work from anywhere device, Is Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetTpmCheckFailed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.tpmCheckFailed
    }
}
// Gets the upgradeEligibility property value. The user experience work from anywhere windows upgrade eligibility status of device. Possible values are: upgraded, unknown, notCapable, capable.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetUpgradeEligibility()(*OperatingSystemUpgradeEligibility) {
    if m == nil {
        return nil
    } else {
        return m.upgradeEligibility
    }
}
// Gets the windowsScore property value. The user experience work from anywhere per device windows score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetWindowsScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.windowsScore
    }
}
// Gets the workFromAnywhereScore property value. The user experience work from anywhere per device overall score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetWorkFromAnywhereScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.workFromAnywhereScore
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoPilotProfileAssigned(val)
        return nil
    }
    res["autoPilotRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoPilotRegistered(val)
        return nil
    }
    res["azureAdDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureAdDeviceId(val)
        return nil
    }
    res["azureAdJoinType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureAdJoinType(val)
        return nil
    }
    res["azureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAzureAdRegistered(val)
        return nil
    }
    res["cloudIdentityScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCloudIdentityScore(val)
        return nil
    }
    res["cloudManagementScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCloudManagementScore(val)
        return nil
    }
    res["cloudProvisioningScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCloudProvisioningScore(val)
        return nil
    }
    res["compliancePolicySetToIntune"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCompliancePolicySetToIntune(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["healthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        cast := val.(UserExperienceAnalyticsHealthState)
        m.SetHealthStatus(&cast)
        return nil
    }
    res["isCloudManagedGatewayEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCloudManagedGatewayEnabled(val)
        return nil
    }
    res["managedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedBy(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["osCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOsCheckFailed(val)
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsDescription(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["otherWorkloadsSetToIntune"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOtherWorkloadsSetToIntune(val)
        return nil
    }
    res["ownership"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnership(val)
        return nil
    }
    res["processor64BitCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProcessor64BitCheckFailed(val)
        return nil
    }
    res["processorCoreCountCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProcessorCoreCountCheckFailed(val)
        return nil
    }
    res["processorFamilyCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProcessorFamilyCheckFailed(val)
        return nil
    }
    res["processorSpeedCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProcessorSpeedCheckFailed(val)
        return nil
    }
    res["ramCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRamCheckFailed(val)
        return nil
    }
    res["secureBootCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSecureBootCheckFailed(val)
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSerialNumber(val)
        return nil
    }
    res["storageCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetStorageCheckFailed(val)
        return nil
    }
    res["tenantAttached"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTenantAttached(val)
        return nil
    }
    res["tpmCheckFailed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTpmCheckFailed(val)
        return nil
    }
    res["upgradeEligibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperatingSystemUpgradeEligibility)
        if err != nil {
            return err
        }
        cast := val.(OperatingSystemUpgradeEligibility)
        m.SetUpgradeEligibility(&cast)
        return nil
    }
    res["windowsScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetWindowsScore(val)
        return nil
    }
    res["workFromAnywhereScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetWorkFromAnywhereScore(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
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
    {
        err = writer.WriteBoolValue("compliancePolicySetToIntune", m.GetCompliancePolicySetToIntune())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
    if m.GetHealthStatus() != nil {
        cast := m.GetHealthStatus().String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCloudManagedGatewayEnabled", m.GetIsCloudManagedGatewayEnabled())
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
        err = writer.WriteBoolValue("osCheckFailed", m.GetOsCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("otherWorkloadsSetToIntune", m.GetOtherWorkloadsSetToIntune())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownership", m.GetOwnership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processor64BitCheckFailed", m.GetProcessor64BitCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorCoreCountCheckFailed", m.GetProcessorCoreCountCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorFamilyCheckFailed", m.GetProcessorFamilyCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorSpeedCheckFailed", m.GetProcessorSpeedCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ramCheckFailed", m.GetRamCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("secureBootCheckFailed", m.GetSecureBootCheckFailed())
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
    {
        err = writer.WriteBoolValue("storageCheckFailed", m.GetStorageCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tenantAttached", m.GetTenantAttached())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tpmCheckFailed", m.GetTpmCheckFailed())
        if err != nil {
            return err
        }
    }
    if m.GetUpgradeEligibility() != nil {
        cast := m.GetUpgradeEligibility().String()
        err = writer.WriteStringValue("upgradeEligibility", &cast)
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
// Sets the autoPilotProfileAssigned property value. The user experience analytics work from anywhere intune device's autopilotProfileAssigned.
// Parameters:
//  - value : Value to set for the autoPilotProfileAssigned property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAutoPilotProfileAssigned(value *bool)() {
    m.autoPilotProfileAssigned = value
}
// Sets the autoPilotRegistered property value. The user experience work from anywhere intune device's autopilotRegistered.
// Parameters:
//  - value : Value to set for the autoPilotRegistered property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAutoPilotRegistered(value *bool)() {
    m.autoPilotRegistered = value
}
// Sets the azureAdDeviceId property value. The user experience work from anywhere azure Ad device Id.
// Parameters:
//  - value : Value to set for the azureAdDeviceId property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// Sets the azureAdJoinType property value. The user experience work from anywhere device's azure Ad joinType.
// Parameters:
//  - value : Value to set for the azureAdJoinType property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdJoinType(value *string)() {
    m.azureAdJoinType = value
}
// Sets the azureAdRegistered property value. The user experience work from anywhere device's azureAdRegistered.
// Parameters:
//  - value : Value to set for the azureAdRegistered property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdRegistered(value *bool)() {
    m.azureAdRegistered = value
}
// Sets the cloudIdentityScore property value. The user experience work from anywhere per device cloud identity score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the cloudIdentityScore property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudIdentityScore(value *float64)() {
    m.cloudIdentityScore = value
}
// Sets the cloudManagementScore property value. The user experience work from anywhere per device cloud management score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the cloudManagementScore property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudManagementScore(value *float64)() {
    m.cloudManagementScore = value
}
// Sets the cloudProvisioningScore property value. The user experience work from anywhere per device cloud provisioning score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the cloudProvisioningScore property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudProvisioningScore(value *float64)() {
    m.cloudProvisioningScore = value
}
// Sets the compliancePolicySetToIntune property value. The user experience work from anywhere device's compliancePolicySetToIntune.
// Parameters:
//  - value : Value to set for the compliancePolicySetToIntune property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCompliancePolicySetToIntune(value *bool)() {
    m.compliancePolicySetToIntune = value
}
// Sets the deviceId property value. The user experience work from anywhere device Id.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. The work from anywhere device's name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the healthStatus property value. The user experience work from anywhere per device health status. Possible values are: unknown, insufficientData, needsAttention, meetingGoals.
// Parameters:
//  - value : Value to set for the healthStatus property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// Sets the isCloudManagedGatewayEnabled property value. The user experience work from anywhere device's Cloud Management Gateway for Configuration Manager is enabled.
// Parameters:
//  - value : Value to set for the isCloudManagedGatewayEnabled property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetIsCloudManagedGatewayEnabled(value *bool)() {
    m.isCloudManagedGatewayEnabled = value
}
// Sets the managedBy property value. The user experience work from anywhere management agent of the device.
// Parameters:
//  - value : Value to set for the managedBy property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetManagedBy(value *string)() {
    m.managedBy = value
}
// Sets the manufacturer property value. The user experience work from anywhere device's manufacturer.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. The user experience work from anywhere device's model.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetModel(value *string)() {
    m.model = value
}
// Sets the osCheckFailed property value. The user experience work from anywhere device, Is OS check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the osCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsCheckFailed(value *bool)() {
    m.osCheckFailed = value
}
// Sets the osDescription property value. The user experience work from anywhere device's OS Description.
// Parameters:
//  - value : Value to set for the osDescription property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsDescription(value *string)() {
    m.osDescription = value
}
// Sets the osVersion property value. The user experience work from anywhere device's OS Version.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the otherWorkloadsSetToIntune property value. The user experience work from anywhere device's otherWorkloadsSetToIntune.
// Parameters:
//  - value : Value to set for the otherWorkloadsSetToIntune property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOtherWorkloadsSetToIntune(value *bool)() {
    m.otherWorkloadsSetToIntune = value
}
// Sets the ownership property value. The user experience work from anywhere device's ownership.
// Parameters:
//  - value : Value to set for the ownership property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOwnership(value *string)() {
    m.ownership = value
}
// Sets the processor64BitCheckFailed property value. The user experience work from anywhere device, Is processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the processor64BitCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessor64BitCheckFailed(value *bool)() {
    m.processor64BitCheckFailed = value
}
// Sets the processorCoreCountCheckFailed property value. The user experience work from anywhere device, Is processor hardware core count check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the processorCoreCountCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorCoreCountCheckFailed(value *bool)() {
    m.processorCoreCountCheckFailed = value
}
// Sets the processorFamilyCheckFailed property value. The user experience work from anywhere device, Is processor hardware family check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the processorFamilyCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorFamilyCheckFailed(value *bool)() {
    m.processorFamilyCheckFailed = value
}
// Sets the processorSpeedCheckFailed property value. The user experience work from anywhere device, Is processor hardware speed check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the processorSpeedCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorSpeedCheckFailed(value *bool)() {
    m.processorSpeedCheckFailed = value
}
// Sets the ramCheckFailed property value. Is the user experience analytics work from anywhere device RAM hardware check failed for device to upgrade to the latest version of windows
// Parameters:
//  - value : Value to set for the ramCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetRamCheckFailed(value *bool)() {
    m.ramCheckFailed = value
}
// Sets the secureBootCheckFailed property value. The user experience work from anywhere device, Is secure boot hardware check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the secureBootCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetSecureBootCheckFailed(value *bool)() {
    m.secureBootCheckFailed = value
}
// Sets the serialNumber property value. The user experience work from anywhere device's serial number.
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// Sets the storageCheckFailed property value. The user experience work from anywhere device, Is storage hardware check failed for device to upgrade to the latest version of windows.
// Parameters:
//  - value : Value to set for the storageCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetStorageCheckFailed(value *bool)() {
    m.storageCheckFailed = value
}
// Sets the tenantAttached property value. The user experience work from anywhere device's tenantAttached.
// Parameters:
//  - value : Value to set for the tenantAttached property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetTenantAttached(value *bool)() {
    m.tenantAttached = value
}
// Sets the tpmCheckFailed property value. The user experience work from anywhere device, Is Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows.
// Parameters:
//  - value : Value to set for the tpmCheckFailed property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetTpmCheckFailed(value *bool)() {
    m.tpmCheckFailed = value
}
// Sets the upgradeEligibility property value. The user experience work from anywhere windows upgrade eligibility status of device. Possible values are: upgraded, unknown, notCapable, capable.
// Parameters:
//  - value : Value to set for the upgradeEligibility property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetUpgradeEligibility(value *OperatingSystemUpgradeEligibility)() {
    m.upgradeEligibility = value
}
// Sets the windowsScore property value. The user experience work from anywhere per device windows score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the windowsScore property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetWindowsScore(value *float64)() {
    m.windowsScore = value
}
// Sets the workFromAnywhereScore property value. The user experience work from anywhere per device overall score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the workFromAnywhereScore property.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}
