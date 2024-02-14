package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ComanagedDevicesSummary summary data for co managed devices
type ComanagedDevicesSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewComanagedDevicesSummary instantiates a new ComanagedDevicesSummary and sets the default values.
func NewComanagedDevicesSummary()(*ComanagedDevicesSummary) {
    m := &ComanagedDevicesSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateComanagedDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComanagedDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagedDevicesSummary(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ComanagedDevicesSummary) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ComanagedDevicesSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompliancePolicyCount gets the compliancePolicyCount property value. Number of devices with CompliancePolicy swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetCompliancePolicyCount()(*int32) {
    val, err := m.GetBackingStore().Get("compliancePolicyCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetConfigurationSettingsCount gets the configurationSettingsCount property value. Number of devices with ConfigurationSettings swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetConfigurationSettingsCount()(*int32) {
    val, err := m.GetBackingStore().Get("configurationSettingsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEndpointProtectionCount gets the endpointProtectionCount property value. Number of devices with EndpointProtection swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetEndpointProtectionCount()(*int32) {
    val, err := m.GetBackingStore().Get("endpointProtectionCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetInventoryCount()(*int32) {
    val, err := m.GetBackingStore().Get("inventoryCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetModernAppsCount gets the modernAppsCount property value. Number of devices with ModernApps swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetModernAppsCount()(*int32) {
    val, err := m.GetBackingStore().Get("modernAppsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ComanagedDevicesSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOfficeAppsCount gets the officeAppsCount property value. Number of devices with OfficeApps swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetOfficeAppsCount()(*int32) {
    val, err := m.GetBackingStore().Get("officeAppsCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetResourceAccessCount gets the resourceAccessCount property value. Number of devices with ResourceAccess swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetResourceAccessCount()(*int32) {
    val, err := m.GetBackingStore().Get("resourceAccessCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalComanagedCount gets the totalComanagedCount property value. Number of Co-Managed Devices. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetTotalComanagedCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalComanagedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetWindowsUpdateForBusinessCount gets the windowsUpdateForBusinessCount property value. Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
// returns a *int32 when successful
func (m *ComanagedDevicesSummary) GetWindowsUpdateForBusinessCount()(*int32) {
    val, err := m.GetBackingStore().Get("windowsUpdateForBusinessCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagedDevicesSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ComanagedDevicesSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompliancePolicyCount sets the compliancePolicyCount property value. Number of devices with CompliancePolicy swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetCompliancePolicyCount(value *int32)() {
    err := m.GetBackingStore().Set("compliancePolicyCount", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigurationSettingsCount sets the configurationSettingsCount property value. Number of devices with ConfigurationSettings swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetConfigurationSettingsCount(value *int32)() {
    err := m.GetBackingStore().Set("configurationSettingsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetEndpointProtectionCount sets the endpointProtectionCount property value. Number of devices with EndpointProtection swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetEndpointProtectionCount(value *int32)() {
    err := m.GetBackingStore().Set("endpointProtectionCount", value)
    if err != nil {
        panic(err)
    }
}
// SetInventoryCount sets the inventoryCount property value. Number of devices with Inventory swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetInventoryCount(value *int32)() {
    err := m.GetBackingStore().Set("inventoryCount", value)
    if err != nil {
        panic(err)
    }
}
// SetModernAppsCount sets the modernAppsCount property value. Number of devices with ModernApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetModernAppsCount(value *int32)() {
    err := m.GetBackingStore().Set("modernAppsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ComanagedDevicesSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOfficeAppsCount sets the officeAppsCount property value. Number of devices with OfficeApps swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetOfficeAppsCount(value *int32)() {
    err := m.GetBackingStore().Set("officeAppsCount", value)
    if err != nil {
        panic(err)
    }
}
// SetResourceAccessCount sets the resourceAccessCount property value. Number of devices with ResourceAccess swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetResourceAccessCount(value *int32)() {
    err := m.GetBackingStore().Set("resourceAccessCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalComanagedCount sets the totalComanagedCount property value. Number of Co-Managed Devices. This property is read-only.
func (m *ComanagedDevicesSummary) SetTotalComanagedCount(value *int32)() {
    err := m.GetBackingStore().Set("totalComanagedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetWindowsUpdateForBusinessCount sets the windowsUpdateForBusinessCount property value. Number of devices with WindowsUpdateForBusiness swung-over. This property is read-only.
func (m *ComanagedDevicesSummary) SetWindowsUpdateForBusinessCount(value *int32)() {
    err := m.GetBackingStore().Set("windowsUpdateForBusinessCount", value)
    if err != nil {
        panic(err)
    }
}
type ComanagedDevicesSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompliancePolicyCount()(*int32)
    GetConfigurationSettingsCount()(*int32)
    GetEndpointProtectionCount()(*int32)
    GetInventoryCount()(*int32)
    GetModernAppsCount()(*int32)
    GetOdataType()(*string)
    GetOfficeAppsCount()(*int32)
    GetResourceAccessCount()(*int32)
    GetTotalComanagedCount()(*int32)
    GetWindowsUpdateForBusinessCount()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompliancePolicyCount(value *int32)()
    SetConfigurationSettingsCount(value *int32)()
    SetEndpointProtectionCount(value *int32)()
    SetInventoryCount(value *int32)()
    SetModernAppsCount(value *int32)()
    SetOdataType(value *string)()
    SetOfficeAppsCount(value *int32)()
    SetResourceAccessCount(value *int32)()
    SetTotalComanagedCount(value *int32)()
    SetWindowsUpdateForBusinessCount(value *int32)()
}
