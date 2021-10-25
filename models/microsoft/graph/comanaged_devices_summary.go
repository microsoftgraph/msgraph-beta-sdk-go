package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ComanagedDevicesSummary struct {
    additionalData map[string]interface{};
    compliancePolicyCount *int32;
    configurationSettingsCount *int32;
    endpointProtectionCount *int32;
    inventoryCount *int32;
    modernAppsCount *int32;
    officeAppsCount *int32;
    resourceAccessCount *int32;
    totalComanagedCount *int32;
    windowsUpdateForBusinessCount *int32;
}
func NewComanagedDevicesSummary()(*ComanagedDevicesSummary) {
    m := &ComanagedDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ComanagedDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ComanagedDevicesSummary) GetCompliancePolicyCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.compliancePolicyCount
    }
}
func (m *ComanagedDevicesSummary) GetConfigurationSettingsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.configurationSettingsCount
    }
}
func (m *ComanagedDevicesSummary) GetEndpointProtectionCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.endpointProtectionCount
    }
}
func (m *ComanagedDevicesSummary) GetInventoryCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.inventoryCount
    }
}
func (m *ComanagedDevicesSummary) GetModernAppsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.modernAppsCount
    }
}
func (m *ComanagedDevicesSummary) GetOfficeAppsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.officeAppsCount
    }
}
func (m *ComanagedDevicesSummary) GetResourceAccessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.resourceAccessCount
    }
}
func (m *ComanagedDevicesSummary) GetTotalComanagedCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalComanagedCount
    }
}
func (m *ComanagedDevicesSummary) GetWindowsUpdateForBusinessCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsUpdateForBusinessCount
    }
}
func (m *ComanagedDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compliancePolicyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCompliancePolicyCount(val)
        return nil
    }
    res["configurationSettingsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetConfigurationSettingsCount(val)
        return nil
    }
    res["endpointProtectionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEndpointProtectionCount(val)
        return nil
    }
    res["inventoryCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetInventoryCount(val)
        return nil
    }
    res["modernAppsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetModernAppsCount(val)
        return nil
    }
    res["officeAppsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOfficeAppsCount(val)
        return nil
    }
    res["resourceAccessCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetResourceAccessCount(val)
        return nil
    }
    res["totalComanagedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalComanagedCount(val)
        return nil
    }
    res["windowsUpdateForBusinessCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWindowsUpdateForBusinessCount(val)
        return nil
    }
    return res
}
func (m *ComanagedDevicesSummary) IsNil()(bool) {
    return m == nil
}
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
func (m *ComanagedDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ComanagedDevicesSummary) SetCompliancePolicyCount(value *int32)() {
    m.compliancePolicyCount = value
}
func (m *ComanagedDevicesSummary) SetConfigurationSettingsCount(value *int32)() {
    m.configurationSettingsCount = value
}
func (m *ComanagedDevicesSummary) SetEndpointProtectionCount(value *int32)() {
    m.endpointProtectionCount = value
}
func (m *ComanagedDevicesSummary) SetInventoryCount(value *int32)() {
    m.inventoryCount = value
}
func (m *ComanagedDevicesSummary) SetModernAppsCount(value *int32)() {
    m.modernAppsCount = value
}
func (m *ComanagedDevicesSummary) SetOfficeAppsCount(value *int32)() {
    m.officeAppsCount = value
}
func (m *ComanagedDevicesSummary) SetResourceAccessCount(value *int32)() {
    m.resourceAccessCount = value
}
func (m *ComanagedDevicesSummary) SetTotalComanagedCount(value *int32)() {
    m.totalComanagedCount = value
}
func (m *ComanagedDevicesSummary) SetWindowsUpdateForBusinessCount(value *int32)() {
    m.windowsUpdateForBusinessCount = value
}
