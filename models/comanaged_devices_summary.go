package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagedDevicesSummary summary data for co managed devices
type ComanagedDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of devices with CompliancePolicy swung-over. This property is read-only.
    compliancePolicyCount *int32
    // Number of devices with ConfigurationSettings swung-over. This property is read-only.
    configurationSettingsCount *int32
    // Number of devices with EndpointProtection swung-over. This property is read-only.
    endpointProtectionCount *int32
    // Number of devices with Inventory swung-over. This property is read-only.
    inventoryCount *int32
    // Number of devices with ModernApps swung-over. This property is read-only.
    modernAppsCount *int32
    // The OdataType property
    odataType *string
    // Number of devices with OfficeApps swung-over. This property is read-only.
    officeAppsCount *int32
    // Number of devices with ResourceAccess swung-over. This property is read-only.
    resourceAccessCount *int32
    // Number of Co-Managed Devices. This property is read-only.
    totalComanagedCount *int32
    // Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
    windowsUpdateForBusinessCount *int32
}
// NewComanagedDevicesSummary instantiates a new comanagedDevicesSummary and sets the default values.
func NewComanagedDevicesSummary()(*ComanagedDevicesSummary) {
    m := &ComanagedDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateComanagedDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagedDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCompliancePolicyCount gets the compliancePolicyCount property value. Number of devices with CompliancePolicy swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetCompliancePolicyCount()(*int32) {
    return m.compliancePolicyCount
}
// GetConfigurationSettingsCount gets the configurationSettingsCount property value. Number of devices with ConfigurationSettings swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetConfigurationSettingsCount()(*int32) {
    return m.configurationSettingsCount
}
// GetEndpointProtectionCount gets the endpointProtectionCount property value. Number of devices with EndpointProtection swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetEndpointProtectionCount()(*int32) {
    return m.endpointProtectionCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagedDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compliancePolicyCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicyCount(val)
        }
        return nil
    }
    res["configurationSettingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationSettingsCount(val)
        }
        return nil
    }
    res["endpointProtectionCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointProtectionCount(val)
        }
        return nil
    }
    res["inventoryCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventoryCount(val)
        }
        return nil
    }
    res["modernAppsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModernAppsCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["officeAppsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeAppsCount(val)
        }
        return nil
    }
    res["resourceAccessCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceAccessCount(val)
        }
        return nil
    }
    res["totalComanagedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalComanagedCount(val)
        }
        return nil
    }
    res["windowsUpdateForBusinessCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetInventoryCount gets the inventoryCount property value. Number of devices with Inventory swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetInventoryCount()(*int32) {
    return m.inventoryCount
}
// GetModernAppsCount gets the modernAppsCount property value. Number of devices with ModernApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetModernAppsCount()(*int32) {
    return m.modernAppsCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ComanagedDevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetOfficeAppsCount gets the officeAppsCount property value. Number of devices with OfficeApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetOfficeAppsCount()(*int32) {
    return m.officeAppsCount
}
// GetResourceAccessCount gets the resourceAccessCount property value. Number of devices with ResourceAccess swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetResourceAccessCount()(*int32) {
    return m.resourceAccessCount
}
// GetTotalComanagedCount gets the totalComanagedCount property value. Number of Co-Managed Devices. This property is read-only.
func (m *ComanagedDevicesSummary) GetTotalComanagedCount()(*int32) {
    return m.totalComanagedCount
}
// GetWindowsUpdateForBusinessCount gets the windowsUpdateForBusinessCount property value. Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) GetWindowsUpdateForBusinessCount()(*int32) {
    return m.windowsUpdateForBusinessCount
}
// Serialize serializes information the current object
func (m *ComanagedDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *ComanagedDevicesSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCompliancePolicyCount sets the compliancePolicyCount property value. Number of devices with CompliancePolicy swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetCompliancePolicyCount(value *int32)() {
    m.compliancePolicyCount = value
}
// SetConfigurationSettingsCount sets the configurationSettingsCount property value. Number of devices with ConfigurationSettings swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetConfigurationSettingsCount(value *int32)() {
    m.configurationSettingsCount = value
}
// SetEndpointProtectionCount sets the endpointProtectionCount property value. Number of devices with EndpointProtection swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetEndpointProtectionCount(value *int32)() {
    m.endpointProtectionCount = value
}
// SetInventoryCount sets the inventoryCount property value. Number of devices with Inventory swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetInventoryCount(value *int32)() {
    m.inventoryCount = value
}
// SetModernAppsCount sets the modernAppsCount property value. Number of devices with ModernApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetModernAppsCount(value *int32)() {
    m.modernAppsCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ComanagedDevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOfficeAppsCount sets the officeAppsCount property value. Number of devices with OfficeApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetOfficeAppsCount(value *int32)() {
    m.officeAppsCount = value
}
// SetResourceAccessCount sets the resourceAccessCount property value. Number of devices with ResourceAccess swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetResourceAccessCount(value *int32)() {
    m.resourceAccessCount = value
}
// SetTotalComanagedCount sets the totalComanagedCount property value. Number of Co-Managed Devices. This property is read-only.
func (m *ComanagedDevicesSummary) SetTotalComanagedCount(value *int32)() {
    m.totalComanagedCount = value
}
// SetWindowsUpdateForBusinessCount sets the windowsUpdateForBusinessCount property value. Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetWindowsUpdateForBusinessCount(value *int32)() {
    m.windowsUpdateForBusinessCount = value
}
