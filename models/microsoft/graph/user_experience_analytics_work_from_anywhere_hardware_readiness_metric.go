package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric struct {
    Entity
    // The percentage of devices for which OS check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    osCheckFailedPercentage *float64;
    // The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processor64BitCheckFailedPercentage *float64;
    // The percentage of devices for which processor hardware core count check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processorCoreCountCheckFailedPercentage *float64;
    // The percentage of devices for which processor hardware family check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processorFamilyCheckFailedPercentage *float64;
    // The percentage of devices for which processor hardware speed check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processorSpeedCheckFailedPercentage *float64;
    // The percentage of devices for which RAM hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    ramCheckFailedPercentage *float64;
    // The percentage of devices for which secure boot hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    secureBootCheckFailedPercentage *float64;
    // The percentage of devices for which storage hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    storageCheckFailedPercentage *float64;
    // The count of total devices in an organization. Valid values -2147483648 to 2147483647
    totalDeviceCount *int32;
    // The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    tpmCheckFailedPercentage *float64;
    // The count of devices in an organization eligible for windows upgrade. Valid values -2147483648 to 2147483647
    upgradeEligibleDeviceCount *int32;
}
// Instantiates a new userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the osCheckFailedPercentage property value. The percentage of devices for which OS check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetOsCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.osCheckFailedPercentage
    }
}
// Gets the processor64BitCheckFailedPercentage property value. The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessor64BitCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.processor64BitCheckFailedPercentage
    }
}
// Gets the processorCoreCountCheckFailedPercentage property value. The percentage of devices for which processor hardware core count check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorCoreCountCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.processorCoreCountCheckFailedPercentage
    }
}
// Gets the processorFamilyCheckFailedPercentage property value. The percentage of devices for which processor hardware family check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorFamilyCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.processorFamilyCheckFailedPercentage
    }
}
// Gets the processorSpeedCheckFailedPercentage property value. The percentage of devices for which processor hardware speed check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorSpeedCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.processorSpeedCheckFailedPercentage
    }
}
// Gets the ramCheckFailedPercentage property value. The percentage of devices for which RAM hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetRamCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.ramCheckFailedPercentage
    }
}
// Gets the secureBootCheckFailedPercentage property value. The percentage of devices for which secure boot hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetSecureBootCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.secureBootCheckFailedPercentage
    }
}
// Gets the storageCheckFailedPercentage property value. The percentage of devices for which storage hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetStorageCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.storageCheckFailedPercentage
    }
}
// Gets the totalDeviceCount property value. The count of total devices in an organization. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetTotalDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalDeviceCount
    }
}
// Gets the tpmCheckFailedPercentage property value. The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetTpmCheckFailedPercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.tpmCheckFailedPercentage
    }
}
// Gets the upgradeEligibleDeviceCount property value. The count of devices in an organization eligible for windows upgrade. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetUpgradeEligibleDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.upgradeEligibleDeviceCount
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["osCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsCheckFailedPercentage(val)
        }
        return nil
    }
    res["processor64BitCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessor64BitCheckFailedPercentage(val)
        }
        return nil
    }
    res["processorCoreCountCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCoreCountCheckFailedPercentage(val)
        }
        return nil
    }
    res["processorFamilyCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorFamilyCheckFailedPercentage(val)
        }
        return nil
    }
    res["processorSpeedCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSpeedCheckFailedPercentage(val)
        }
        return nil
    }
    res["ramCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamCheckFailedPercentage(val)
        }
        return nil
    }
    res["secureBootCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootCheckFailedPercentage(val)
        }
        return nil
    }
    res["storageCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageCheckFailedPercentage(val)
        }
        return nil
    }
    res["totalDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDeviceCount(val)
        }
        return nil
    }
    res["tpmCheckFailedPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmCheckFailedPercentage(val)
        }
        return nil
    }
    res["upgradeEligibleDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeEligibleDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("osCheckFailedPercentage", m.GetOsCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processor64BitCheckFailedPercentage", m.GetProcessor64BitCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processorCoreCountCheckFailedPercentage", m.GetProcessorCoreCountCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processorFamilyCheckFailedPercentage", m.GetProcessorFamilyCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processorSpeedCheckFailedPercentage", m.GetProcessorSpeedCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("ramCheckFailedPercentage", m.GetRamCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("secureBootCheckFailedPercentage", m.GetSecureBootCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("storageCheckFailedPercentage", m.GetStorageCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalDeviceCount", m.GetTotalDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("tpmCheckFailedPercentage", m.GetTpmCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("upgradeEligibleDeviceCount", m.GetUpgradeEligibleDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the osCheckFailedPercentage property value. The percentage of devices for which OS check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the osCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetOsCheckFailedPercentage(value *float64)() {
    m.osCheckFailedPercentage = value
}
// Sets the processor64BitCheckFailedPercentage property value. The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the processor64BitCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessor64BitCheckFailedPercentage(value *float64)() {
    m.processor64BitCheckFailedPercentage = value
}
// Sets the processorCoreCountCheckFailedPercentage property value. The percentage of devices for which processor hardware core count check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the processorCoreCountCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorCoreCountCheckFailedPercentage(value *float64)() {
    m.processorCoreCountCheckFailedPercentage = value
}
// Sets the processorFamilyCheckFailedPercentage property value. The percentage of devices for which processor hardware family check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the processorFamilyCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorFamilyCheckFailedPercentage(value *float64)() {
    m.processorFamilyCheckFailedPercentage = value
}
// Sets the processorSpeedCheckFailedPercentage property value. The percentage of devices for which processor hardware speed check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the processorSpeedCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorSpeedCheckFailedPercentage(value *float64)() {
    m.processorSpeedCheckFailedPercentage = value
}
// Sets the ramCheckFailedPercentage property value. The percentage of devices for which RAM hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the ramCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetRamCheckFailedPercentage(value *float64)() {
    m.ramCheckFailedPercentage = value
}
// Sets the secureBootCheckFailedPercentage property value. The percentage of devices for which secure boot hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the secureBootCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetSecureBootCheckFailedPercentage(value *float64)() {
    m.secureBootCheckFailedPercentage = value
}
// Sets the storageCheckFailedPercentage property value. The percentage of devices for which storage hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the storageCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetStorageCheckFailedPercentage(value *float64)() {
    m.storageCheckFailedPercentage = value
}
// Sets the totalDeviceCount property value. The count of total devices in an organization. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the totalDeviceCount property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetTotalDeviceCount(value *int32)() {
    m.totalDeviceCount = value
}
// Sets the tpmCheckFailedPercentage property value. The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the tpmCheckFailedPercentage property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetTpmCheckFailedPercentage(value *float64)() {
    m.tpmCheckFailedPercentage = value
}
// Sets the upgradeEligibleDeviceCount property value. The count of devices in an organization eligible for windows upgrade. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the upgradeEligibleDeviceCount property.
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetUpgradeEligibleDeviceCount(value *int32)() {
    m.upgradeEligibleDeviceCount = value
}
