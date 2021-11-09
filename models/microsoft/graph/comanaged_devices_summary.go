package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ComanagedDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of devices with CompliancePolicy swung-over. This property is read-only.
    compliancePolicyCount *int32;
    // Number of devices with ConfigurationSettings swung-over. This property is read-only.
    configurationSettingsCount *int32;
    // Number of devices with EndpointProtection swung-over. This property is read-only.
    endpointProtectionCount *int32;
    // Number of devices with Inventory swung-over. This property is read-only.
    inventoryCount *int32;
    // Number of devices with ModernApps swung-over. This property is read-only.
    modernAppsCount *int32;
    // Number of devices with OfficeApps swung-over. This property is read-only.
    officeAppsCount *int32;
    // Number of devices with ResourceAccess swung-over. This property is read-only.
    resourceAccessCount *int32;
    // Number of Co-Managed Devices. This property is read-only.
    totalComanagedCount *int32;
    // Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
    windowsUpdateForBusinessCount *int32;
}
// Instantiates a new comanagedDevicesSummary and sets the default values.
func NewComanagedDevicesSummary()(*ComanagedDevicesSummary) {
    m := &ComanagedDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the compliancePolicyCount property value. Number of devices with CompliancePolicy swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetCompliancePolicyCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyCount
    }
}
// Gets the configurationSettingsCount property value. Number of devices with ConfigurationSettings swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetConfigurationSettingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationSettingsCount
    }
}
// Gets the endpointProtectionCount property value. Number of devices with EndpointProtection swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetEndpointProtectionCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.endpointProtectionCount
    }
}
// Gets the inventoryCount property value. Number of devices with Inventory swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetInventoryCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inventoryCount
    }
}
// Gets the modernAppsCount property value. Number of devices with ModernApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetModernAppsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.modernAppsCount
    }
}
// Gets the officeAppsCount property value. Number of devices with OfficeApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetOfficeAppsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.officeAppsCount
    }
}
// Gets the resourceAccessCount property value. Number of devices with ResourceAccess swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetResourceAccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccessCount
    }
}
// Gets the totalComanagedCount property value. Number of Co-Managed Devices. This property is read-only.
func (m *ComanagedDevicesSummary) GetTotalComanagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalComanagedCount
    }
}
// Gets the windowsUpdateForBusinessCount property value. Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetWindowsUpdateForBusinessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateForBusinessCount
    }
}
// The deserialization information for the current model
func (m *ComanagedDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compliancePolicyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyCount(val)
        }
        return nil
    }
    res["configurationSettingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationSettingsCount(val)
        }
        return nil
    }
    res["endpointProtectionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointProtectionCount(val)
        }
        return nil
    }
    res["inventoryCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventoryCount(val)
        }
        return nil
    }
    res["modernAppsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModernAppsCount(val)
        }
        return nil
    }
    res["officeAppsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeAppsCount(val)
        }
        return nil
    }
    res["resourceAccessCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAccessCount(val)
        }
        return nil
    }
    res["totalComanagedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalComanagedCount(val)
        }
        return nil
    }
    res["windowsUpdateForBusinessCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsUpdateForBusinessCount(val)
        }
        return nil
    }
    return res
}
func (m *ComanagedDevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ComanagedDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("compliancePolicyCount", m.GetCompliancePolicyCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("configurationSettingsCount", m.GetConfigurationSettingsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("endpointProtectionCount", m.GetEndpointProtectionCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inventoryCount", m.GetInventoryCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("modernAppsCount", m.GetModernAppsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("officeAppsCount", m.GetOfficeAppsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("resourceAccessCount", m.GetResourceAccessCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalComanagedCount", m.GetTotalComanagedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsUpdateForBusinessCount", m.GetWindowsUpdateForBusinessCount())
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
func (m *ComanagedDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the compliancePolicyCount property value. Number of devices with CompliancePolicy swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the compliancePolicyCount property.
func (m *ComanagedDevicesSummary) SetCompliancePolicyCount(value *int32)() {
    m.compliancePolicyCount = value
}
// Sets the configurationSettingsCount property value. Number of devices with ConfigurationSettings swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the configurationSettingsCount property.
func (m *ComanagedDevicesSummary) SetConfigurationSettingsCount(value *int32)() {
    m.configurationSettingsCount = value
}
// Sets the endpointProtectionCount property value. Number of devices with EndpointProtection swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the endpointProtectionCount property.
func (m *ComanagedDevicesSummary) SetEndpointProtectionCount(value *int32)() {
    m.endpointProtectionCount = value
}
// Sets the inventoryCount property value. Number of devices with Inventory swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the inventoryCount property.
func (m *ComanagedDevicesSummary) SetInventoryCount(value *int32)() {
    m.inventoryCount = value
}
// Sets the modernAppsCount property value. Number of devices with ModernApps swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the modernAppsCount property.
func (m *ComanagedDevicesSummary) SetModernAppsCount(value *int32)() {
    m.modernAppsCount = value
}
// Sets the officeAppsCount property value. Number of devices with OfficeApps swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the officeAppsCount property.
func (m *ComanagedDevicesSummary) SetOfficeAppsCount(value *int32)() {
    m.officeAppsCount = value
}
// Sets the resourceAccessCount property value. Number of devices with ResourceAccess swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the resourceAccessCount property.
func (m *ComanagedDevicesSummary) SetResourceAccessCount(value *int32)() {
    m.resourceAccessCount = value
}
// Sets the totalComanagedCount property value. Number of Co-Managed Devices. This property is read-only.
// Parameters:
//  - value : Value to set for the totalComanagedCount property.
func (m *ComanagedDevicesSummary) SetTotalComanagedCount(value *int32)() {
    m.totalComanagedCount = value
}
// Sets the windowsUpdateForBusinessCount property value. Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
// Parameters:
//  - value : Value to set for the windowsUpdateForBusinessCount property.
func (m *ComanagedDevicesSummary) SetWindowsUpdateForBusinessCount(value *int32)() {
    m.windowsUpdateForBusinessCount = value
}
