package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// HardwareInformation hardware information of a given device.
type HardwareInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewHardwareInformation instantiates a new HardwareInformation and sets the default values.
func NewHardwareInformation()(*HardwareInformation) {
    m := &HardwareInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHardwareInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHardwareInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardwareInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HardwareInformation) GetAdditionalData()(map[string]any) {
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
func (m *HardwareInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBatteryChargeCycles gets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *HardwareInformation) GetBatteryChargeCycles()(*int32) {
    val, err := m.GetBackingStore().Get("batteryChargeCycles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBatteryHealthPercentage gets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
// returns a *int32 when successful
func (m *HardwareInformation) GetBatteryHealthPercentage()(*int32) {
    val, err := m.GetBackingStore().Get("batteryHealthPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetBatteryLevelPercentage gets the batteryLevelPercentage property value. The battery level, between 0.0 and 100, or null if the battery level cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 5.0 and later, and is available only when Device Information access right is obtained. Valid values 0 to 100
// returns a *float64 when successful
func (m *HardwareInformation) GetBatteryLevelPercentage()(*float64) {
    val, err := m.GetBackingStore().Get("batteryLevelPercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetBatterySerialNumber gets the batterySerialNumber property value. The serial number of the device’s current battery
// returns a *string when successful
func (m *HardwareInformation) GetBatterySerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("batterySerialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCellularTechnology gets the cellularTechnology property value. Cellular technology of the device
// returns a *string when successful
func (m *HardwareInformation) GetCellularTechnology()(*string) {
    val, err := m.GetBackingStore().Get("cellularTechnology")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceFullQualifiedDomainName gets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
// returns a *string when successful
func (m *HardwareInformation) GetDeviceFullQualifiedDomainName()(*string) {
    val, err := m.GetBackingStore().Get("deviceFullQualifiedDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceGuardLocalSystemAuthorityCredentialGuardState gets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. The deviceGuardLocalSystemAuthorityCredentialGuardState property
// returns a *DeviceGuardLocalSystemAuthorityCredentialGuardState when successful
func (m *HardwareInformation) GetDeviceGuardLocalSystemAuthorityCredentialGuardState()(*DeviceGuardLocalSystemAuthorityCredentialGuardState) {
    val, err := m.GetBackingStore().Get("deviceGuardLocalSystemAuthorityCredentialGuardState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceGuardLocalSystemAuthorityCredentialGuardState)
    }
    return nil
}
// GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState gets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. The deviceGuardVirtualizationBasedSecurityHardwareRequirementState property
// returns a *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState when successful
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) {
    val, err := m.GetBackingStore().Get("deviceGuardVirtualizationBasedSecurityHardwareRequirementState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
    }
    return nil
}
// GetDeviceGuardVirtualizationBasedSecurityState gets the deviceGuardVirtualizationBasedSecurityState property value. The deviceGuardVirtualizationBasedSecurityState property
// returns a *DeviceGuardVirtualizationBasedSecurityState when successful
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityState()(*DeviceGuardVirtualizationBasedSecurityState) {
    val, err := m.GetBackingStore().Get("deviceGuardVirtualizationBasedSecurityState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceGuardVirtualizationBasedSecurityState)
    }
    return nil
}
// GetDeviceLicensingLastErrorCode gets the deviceLicensingLastErrorCode property value. A standard error code indicating the last error, or 0 indicating no error (default). The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *HardwareInformation) GetDeviceLicensingLastErrorCode()(*int32) {
    val, err := m.GetBackingStore().Get("deviceLicensingLastErrorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceLicensingLastErrorDescription gets the deviceLicensingLastErrorDescription property value. Error text message as a descripition for deviceLicensingLastErrorCode. The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing.
// returns a *string when successful
func (m *HardwareInformation) GetDeviceLicensingLastErrorDescription()(*string) {
    val, err := m.GetBackingStore().Get("deviceLicensingLastErrorDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceLicensingStatus gets the deviceLicensingStatus property value. Indicates the device licensing status after Windows device based subscription has been enabled.
// returns a *DeviceLicensingStatus when successful
func (m *HardwareInformation) GetDeviceLicensingStatus()(*DeviceLicensingStatus) {
    val, err := m.GetBackingStore().Get("deviceLicensingStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceLicensingStatus)
    }
    return nil
}
// GetEsimIdentifier gets the esimIdentifier property value. eSIM identifier
// returns a *string when successful
func (m *HardwareInformation) GetEsimIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("esimIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HardwareInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["batteryChargeCycles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryChargeCycles(val)
        }
        return nil
    }
    res["batteryHealthPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthPercentage(val)
        }
        return nil
    }
    res["batteryLevelPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryLevelPercentage(val)
        }
        return nil
    }
    res["batterySerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatterySerialNumber(val)
        }
        return nil
    }
    res["cellularTechnology"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularTechnology(val)
        }
        return nil
    }
    res["deviceFullQualifiedDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFullQualifiedDomainName(val)
        }
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLocalSystemAuthorityCredentialGuardState(val.(*DeviceGuardLocalSystemAuthorityCredentialGuardState))
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityHardwareRequirementState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(val.(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState))
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardVirtualizationBasedSecurityState(val.(*DeviceGuardVirtualizationBasedSecurityState))
        }
        return nil
    }
    res["deviceLicensingLastErrorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLicensingLastErrorCode(val)
        }
        return nil
    }
    res["deviceLicensingLastErrorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLicensingLastErrorDescription(val)
        }
        return nil
    }
    res["deviceLicensingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceLicensingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceLicensingStatus(val.(*DeviceLicensingStatus))
        }
        return nil
    }
    res["esimIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsimIdentifier(val)
        }
        return nil
    }
    res["freeStorageSpace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpace(val)
        }
        return nil
    }
    res["imei"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["ipAddressV4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressV4(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSharedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSharedDevice(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
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
    res["operatingSystemEdition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemEdition(val)
        }
        return nil
    }
    res["operatingSystemLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemLanguage(val)
        }
        return nil
    }
    res["operatingSystemProductType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemProductType(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["residentUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidentUsersCount(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["sharedDeviceCachedUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedAppleDeviceUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedAppleDeviceUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SharedAppleDeviceUserable)
                }
            }
            m.SetSharedDeviceCachedUsers(res)
        }
        return nil
    }
    res["subnetAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetAddress(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["systemManagementBIOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemManagementBIOSVersion(val)
        }
        return nil
    }
    res["totalStorageSpace"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpace(val)
        }
        return nil
    }
    res["tpmManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmManufacturer(val)
        }
        return nil
    }
    res["tpmSpecificationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmSpecificationVersion(val)
        }
        return nil
    }
    res["tpmVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmVersion(val)
        }
        return nil
    }
    res["wifiMac"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiMac(val)
        }
        return nil
    }
    res["wiredIPv4Addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWiredIPv4Addresses(res)
        }
        return nil
    }
    return res
}
// GetFreeStorageSpace gets the freeStorageSpace property value. Free storage space of the device.
// returns a *int64 when successful
func (m *HardwareInformation) GetFreeStorageSpace()(*int64) {
    val, err := m.GetBackingStore().Get("freeStorageSpace")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetImei gets the imei property value. IMEI
// returns a *string when successful
func (m *HardwareInformation) GetImei()(*string) {
    val, err := m.GetBackingStore().Get("imei")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIpAddressV4 gets the ipAddressV4 property value. IPAddressV4
// returns a *string when successful
func (m *HardwareInformation) GetIpAddressV4()(*string) {
    val, err := m.GetBackingStore().Get("ipAddressV4")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsEncrypted gets the isEncrypted property value. Encryption status of the device
// returns a *bool when successful
func (m *HardwareInformation) GetIsEncrypted()(*bool) {
    val, err := m.GetBackingStore().Get("isEncrypted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSharedDevice gets the isSharedDevice property value. Shared iPad
// returns a *bool when successful
func (m *HardwareInformation) GetIsSharedDevice()(*bool) {
    val, err := m.GetBackingStore().Get("isSharedDevice")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSupervised gets the isSupervised property value. Supervised mode of the device
// returns a *bool when successful
func (m *HardwareInformation) GetIsSupervised()(*bool) {
    val, err := m.GetBackingStore().Get("isSupervised")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device
// returns a *string when successful
func (m *HardwareInformation) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMeid gets the meid property value. MEID
// returns a *string when successful
func (m *HardwareInformation) GetMeid()(*string) {
    val, err := m.GetBackingStore().Get("meid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. Model of the device
// returns a *string when successful
func (m *HardwareInformation) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *HardwareInformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystemEdition gets the operatingSystemEdition property value. String that specifies the OS edition.
// returns a *string when successful
func (m *HardwareInformation) GetOperatingSystemEdition()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystemEdition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystemLanguage gets the operatingSystemLanguage property value. Operating system language of the device
// returns a *string when successful
func (m *HardwareInformation) GetOperatingSystemLanguage()(*string) {
    val, err := m.GetBackingStore().Get("operatingSystemLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperatingSystemProductType gets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *HardwareInformation) GetOperatingSystemProductType()(*int32) {
    val, err := m.GetBackingStore().Get("operatingSystemProductType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOsBuildNumber gets the osBuildNumber property value. Operating System Build Number on Android device
// returns a *string when successful
func (m *HardwareInformation) GetOsBuildNumber()(*string) {
    val, err := m.GetBackingStore().Get("osBuildNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device
// returns a *string when successful
func (m *HardwareInformation) GetPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("phoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProductName gets the productName property value. The product name, e.g. iPad8,12 etc. The update frequency of this property is weekly. Note this property is currently supported only on iOS/MacOS devices, and is available only when Device Information access right is obtained.
// returns a *string when successful
func (m *HardwareInformation) GetProductName()(*string) {
    val, err := m.GetBackingStore().Get("productName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResidentUsersCount gets the residentUsersCount property value. The number of users currently on this device, or null (default) if the value of this property cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 13.4 and later, and is available only when Device Information access right is obtained. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *HardwareInformation) GetResidentUsersCount()(*int32) {
    val, err := m.GetBackingStore().Get("residentUsersCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. Serial number.
// returns a *string when successful
func (m *HardwareInformation) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSharedDeviceCachedUsers gets the sharedDeviceCachedUsers property value. All users on the shared Apple device
// returns a []SharedAppleDeviceUserable when successful
func (m *HardwareInformation) GetSharedDeviceCachedUsers()([]SharedAppleDeviceUserable) {
    val, err := m.GetBackingStore().Get("sharedDeviceCachedUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SharedAppleDeviceUserable)
    }
    return nil
}
// GetSubnetAddress gets the subnetAddress property value. SubnetAddress
// returns a *string when successful
func (m *HardwareInformation) GetSubnetAddress()(*string) {
    val, err := m.GetBackingStore().Get("subnetAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber carrier of the device
// returns a *string when successful
func (m *HardwareInformation) GetSubscriberCarrier()(*string) {
    val, err := m.GetBackingStore().Get("subscriberCarrier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSystemManagementBIOSVersion gets the systemManagementBIOSVersion property value. BIOS version as reported by SMBIOS
// returns a *string when successful
func (m *HardwareInformation) GetSystemManagementBIOSVersion()(*string) {
    val, err := m.GetBackingStore().Get("systemManagementBIOSVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalStorageSpace gets the totalStorageSpace property value. Total storage space of the device.
// returns a *int64 when successful
func (m *HardwareInformation) GetTotalStorageSpace()(*int64) {
    val, err := m.GetBackingStore().Get("totalStorageSpace")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTpmManufacturer gets the tpmManufacturer property value. The identifying information that uniquely names the TPM manufacturer
// returns a *string when successful
func (m *HardwareInformation) GetTpmManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("tpmManufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTpmSpecificationVersion gets the tpmSpecificationVersion property value. String that specifies the specification version.
// returns a *string when successful
func (m *HardwareInformation) GetTpmSpecificationVersion()(*string) {
    val, err := m.GetBackingStore().Get("tpmSpecificationVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTpmVersion gets the tpmVersion property value. The version of the TPM, as specified by the manufacturer
// returns a *string when successful
func (m *HardwareInformation) GetTpmVersion()(*string) {
    val, err := m.GetBackingStore().Get("tpmVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWifiMac gets the wifiMac property value. WiFi MAC address of the device
// returns a *string when successful
func (m *HardwareInformation) GetWifiMac()(*string) {
    val, err := m.GetBackingStore().Get("wifiMac")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWiredIPv4Addresses gets the wiredIPv4Addresses property value. A list of wired IPv4 addresses. The update frequency (the maximum delay for the change of property value to be synchronized from the device to the cloud storage) of this property is daily. Note this property is currently supported only on devices running on Windows.
// returns a []string when successful
func (m *HardwareInformation) GetWiredIPv4Addresses()([]string) {
    val, err := m.GetBackingStore().Get("wiredIPv4Addresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *HardwareInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("batteryChargeCycles", m.GetBatteryChargeCycles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("batteryHealthPercentage", m.GetBatteryHealthPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("batteryLevelPercentage", m.GetBatteryLevelPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("batterySerialNumber", m.GetBatterySerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cellularTechnology", m.GetCellularTechnology())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceFullQualifiedDomainName", m.GetDeviceFullQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState() != nil {
        cast := (*m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState()).String()
        err := writer.WriteStringValue("deviceGuardLocalSystemAuthorityCredentialGuardState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() != nil {
        cast := (*m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()).String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityHardwareRequirementState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityState() != nil {
        cast := (*m.GetDeviceGuardVirtualizationBasedSecurityState()).String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("deviceLicensingLastErrorCode", m.GetDeviceLicensingLastErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceLicensingLastErrorDescription", m.GetDeviceLicensingLastErrorDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceLicensingStatus() != nil {
        cast := (*m.GetDeviceLicensingStatus()).String()
        err := writer.WriteStringValue("deviceLicensingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("esimIdentifier", m.GetEsimIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("freeStorageSpace", m.GetFreeStorageSpace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("imei", m.GetImei())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddressV4", m.GetIpAddressV4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSharedDevice", m.GetIsSharedDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("meid", m.GetMeid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
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
        err := writer.WriteStringValue("operatingSystemEdition", m.GetOperatingSystemEdition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemLanguage", m.GetOperatingSystemLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("operatingSystemProductType", m.GetOperatingSystemProductType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("residentUsersCount", m.GetResidentUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetSharedDeviceCachedUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedDeviceCachedUsers()))
        for i, v := range m.GetSharedDeviceCachedUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("sharedDeviceCachedUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetAddress", m.GetSubnetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriberCarrier", m.GetSubscriberCarrier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("systemManagementBIOSVersion", m.GetSystemManagementBIOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalStorageSpace", m.GetTotalStorageSpace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmManufacturer", m.GetTpmManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmSpecificationVersion", m.GetTpmSpecificationVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmVersion", m.GetTpmVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMac", m.GetWifiMac())
        if err != nil {
            return err
        }
    }
    if m.GetWiredIPv4Addresses() != nil {
        err := writer.WriteCollectionOfStringValues("wiredIPv4Addresses", m.GetWiredIPv4Addresses())
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
func (m *HardwareInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *HardwareInformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBatteryChargeCycles sets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
func (m *HardwareInformation) SetBatteryChargeCycles(value *int32)() {
    err := m.GetBackingStore().Set("batteryChargeCycles", value)
    if err != nil {
        panic(err)
    }
}
// SetBatteryHealthPercentage sets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
func (m *HardwareInformation) SetBatteryHealthPercentage(value *int32)() {
    err := m.GetBackingStore().Set("batteryHealthPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetBatteryLevelPercentage sets the batteryLevelPercentage property value. The battery level, between 0.0 and 100, or null if the battery level cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 5.0 and later, and is available only when Device Information access right is obtained. Valid values 0 to 100
func (m *HardwareInformation) SetBatteryLevelPercentage(value *float64)() {
    err := m.GetBackingStore().Set("batteryLevelPercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetBatterySerialNumber sets the batterySerialNumber property value. The serial number of the device’s current battery
func (m *HardwareInformation) SetBatterySerialNumber(value *string)() {
    err := m.GetBackingStore().Set("batterySerialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetCellularTechnology sets the cellularTechnology property value. Cellular technology of the device
func (m *HardwareInformation) SetCellularTechnology(value *string)() {
    err := m.GetBackingStore().Set("cellularTechnology", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceFullQualifiedDomainName sets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
func (m *HardwareInformation) SetDeviceFullQualifiedDomainName(value *string)() {
    err := m.GetBackingStore().Set("deviceFullQualifiedDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardLocalSystemAuthorityCredentialGuardState sets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. The deviceGuardLocalSystemAuthorityCredentialGuardState property
func (m *HardwareInformation) SetDeviceGuardLocalSystemAuthorityCredentialGuardState(value *DeviceGuardLocalSystemAuthorityCredentialGuardState)() {
    err := m.GetBackingStore().Set("deviceGuardLocalSystemAuthorityCredentialGuardState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState sets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. The deviceGuardVirtualizationBasedSecurityHardwareRequirementState property
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(value *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)() {
    err := m.GetBackingStore().Set("deviceGuardVirtualizationBasedSecurityHardwareRequirementState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceGuardVirtualizationBasedSecurityState sets the deviceGuardVirtualizationBasedSecurityState property value. The deviceGuardVirtualizationBasedSecurityState property
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityState(value *DeviceGuardVirtualizationBasedSecurityState)() {
    err := m.GetBackingStore().Set("deviceGuardVirtualizationBasedSecurityState", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLicensingLastErrorCode sets the deviceLicensingLastErrorCode property value. A standard error code indicating the last error, or 0 indicating no error (default). The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing. Valid values 0 to 2147483647
func (m *HardwareInformation) SetDeviceLicensingLastErrorCode(value *int32)() {
    err := m.GetBackingStore().Set("deviceLicensingLastErrorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLicensingLastErrorDescription sets the deviceLicensingLastErrorDescription property value. Error text message as a descripition for deviceLicensingLastErrorCode. The update frequency of this property is daily. Note this property is currently supported only for Windows based Device based subscription licensing.
func (m *HardwareInformation) SetDeviceLicensingLastErrorDescription(value *string)() {
    err := m.GetBackingStore().Set("deviceLicensingLastErrorDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceLicensingStatus sets the deviceLicensingStatus property value. Indicates the device licensing status after Windows device based subscription has been enabled.
func (m *HardwareInformation) SetDeviceLicensingStatus(value *DeviceLicensingStatus)() {
    err := m.GetBackingStore().Set("deviceLicensingStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetEsimIdentifier sets the esimIdentifier property value. eSIM identifier
func (m *HardwareInformation) SetEsimIdentifier(value *string)() {
    err := m.GetBackingStore().Set("esimIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetFreeStorageSpace sets the freeStorageSpace property value. Free storage space of the device.
func (m *HardwareInformation) SetFreeStorageSpace(value *int64)() {
    err := m.GetBackingStore().Set("freeStorageSpace", value)
    if err != nil {
        panic(err)
    }
}
// SetImei sets the imei property value. IMEI
func (m *HardwareInformation) SetImei(value *string)() {
    err := m.GetBackingStore().Set("imei", value)
    if err != nil {
        panic(err)
    }
}
// SetIpAddressV4 sets the ipAddressV4 property value. IPAddressV4
func (m *HardwareInformation) SetIpAddressV4(value *string)() {
    err := m.GetBackingStore().Set("ipAddressV4", value)
    if err != nil {
        panic(err)
    }
}
// SetIsEncrypted sets the isEncrypted property value. Encryption status of the device
func (m *HardwareInformation) SetIsEncrypted(value *bool)() {
    err := m.GetBackingStore().Set("isEncrypted", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSharedDevice sets the isSharedDevice property value. Shared iPad
func (m *HardwareInformation) SetIsSharedDevice(value *bool)() {
    err := m.GetBackingStore().Set("isSharedDevice", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSupervised sets the isSupervised property value. Supervised mode of the device
func (m *HardwareInformation) SetIsSupervised(value *bool)() {
    err := m.GetBackingStore().Set("isSupervised", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device
func (m *HardwareInformation) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetMeid sets the meid property value. MEID
func (m *HardwareInformation) SetMeid(value *string)() {
    err := m.GetBackingStore().Set("meid", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. Model of the device
func (m *HardwareInformation) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HardwareInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemEdition sets the operatingSystemEdition property value. String that specifies the OS edition.
func (m *HardwareInformation) SetOperatingSystemEdition(value *string)() {
    err := m.GetBackingStore().Set("operatingSystemEdition", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemLanguage sets the operatingSystemLanguage property value. Operating system language of the device
func (m *HardwareInformation) SetOperatingSystemLanguage(value *string)() {
    err := m.GetBackingStore().Set("operatingSystemLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetOperatingSystemProductType sets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
func (m *HardwareInformation) SetOperatingSystemProductType(value *int32)() {
    err := m.GetBackingStore().Set("operatingSystemProductType", value)
    if err != nil {
        panic(err)
    }
}
// SetOsBuildNumber sets the osBuildNumber property value. Operating System Build Number on Android device
func (m *HardwareInformation) SetOsBuildNumber(value *string)() {
    err := m.GetBackingStore().Set("osBuildNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device
func (m *HardwareInformation) SetPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("phoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetProductName sets the productName property value. The product name, e.g. iPad8,12 etc. The update frequency of this property is weekly. Note this property is currently supported only on iOS/MacOS devices, and is available only when Device Information access right is obtained.
func (m *HardwareInformation) SetProductName(value *string)() {
    err := m.GetBackingStore().Set("productName", value)
    if err != nil {
        panic(err)
    }
}
// SetResidentUsersCount sets the residentUsersCount property value. The number of users currently on this device, or null (default) if the value of this property cannot be determined. The update frequency of this property is per-checkin. Note this property is currently supported only on devices running iOS 13.4 and later, and is available only when Device Information access right is obtained. Valid values 0 to 2147483647
func (m *HardwareInformation) SetResidentUsersCount(value *int32)() {
    err := m.GetBackingStore().Set("residentUsersCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. Serial number.
func (m *HardwareInformation) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedDeviceCachedUsers sets the sharedDeviceCachedUsers property value. All users on the shared Apple device
func (m *HardwareInformation) SetSharedDeviceCachedUsers(value []SharedAppleDeviceUserable)() {
    err := m.GetBackingStore().Set("sharedDeviceCachedUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetSubnetAddress sets the subnetAddress property value. SubnetAddress
func (m *HardwareInformation) SetSubnetAddress(value *string)() {
    err := m.GetBackingStore().Set("subnetAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber carrier of the device
func (m *HardwareInformation) SetSubscriberCarrier(value *string)() {
    err := m.GetBackingStore().Set("subscriberCarrier", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemManagementBIOSVersion sets the systemManagementBIOSVersion property value. BIOS version as reported by SMBIOS
func (m *HardwareInformation) SetSystemManagementBIOSVersion(value *string)() {
    err := m.GetBackingStore().Set("systemManagementBIOSVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalStorageSpace sets the totalStorageSpace property value. Total storage space of the device.
func (m *HardwareInformation) SetTotalStorageSpace(value *int64)() {
    err := m.GetBackingStore().Set("totalStorageSpace", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmManufacturer sets the tpmManufacturer property value. The identifying information that uniquely names the TPM manufacturer
func (m *HardwareInformation) SetTpmManufacturer(value *string)() {
    err := m.GetBackingStore().Set("tpmManufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmSpecificationVersion sets the tpmSpecificationVersion property value. String that specifies the specification version.
func (m *HardwareInformation) SetTpmSpecificationVersion(value *string)() {
    err := m.GetBackingStore().Set("tpmSpecificationVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetTpmVersion sets the tpmVersion property value. The version of the TPM, as specified by the manufacturer
func (m *HardwareInformation) SetTpmVersion(value *string)() {
    err := m.GetBackingStore().Set("tpmVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiMac sets the wifiMac property value. WiFi MAC address of the device
func (m *HardwareInformation) SetWifiMac(value *string)() {
    err := m.GetBackingStore().Set("wifiMac", value)
    if err != nil {
        panic(err)
    }
}
// SetWiredIPv4Addresses sets the wiredIPv4Addresses property value. A list of wired IPv4 addresses. The update frequency (the maximum delay for the change of property value to be synchronized from the device to the cloud storage) of this property is daily. Note this property is currently supported only on devices running on Windows.
func (m *HardwareInformation) SetWiredIPv4Addresses(value []string)() {
    err := m.GetBackingStore().Set("wiredIPv4Addresses", value)
    if err != nil {
        panic(err)
    }
}
type HardwareInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBatteryChargeCycles()(*int32)
    GetBatteryHealthPercentage()(*int32)
    GetBatteryLevelPercentage()(*float64)
    GetBatterySerialNumber()(*string)
    GetCellularTechnology()(*string)
    GetDeviceFullQualifiedDomainName()(*string)
    GetDeviceGuardLocalSystemAuthorityCredentialGuardState()(*DeviceGuardLocalSystemAuthorityCredentialGuardState)
    GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
    GetDeviceGuardVirtualizationBasedSecurityState()(*DeviceGuardVirtualizationBasedSecurityState)
    GetDeviceLicensingLastErrorCode()(*int32)
    GetDeviceLicensingLastErrorDescription()(*string)
    GetDeviceLicensingStatus()(*DeviceLicensingStatus)
    GetEsimIdentifier()(*string)
    GetFreeStorageSpace()(*int64)
    GetImei()(*string)
    GetIpAddressV4()(*string)
    GetIsEncrypted()(*bool)
    GetIsSharedDevice()(*bool)
    GetIsSupervised()(*bool)
    GetManufacturer()(*string)
    GetMeid()(*string)
    GetModel()(*string)
    GetOdataType()(*string)
    GetOperatingSystemEdition()(*string)
    GetOperatingSystemLanguage()(*string)
    GetOperatingSystemProductType()(*int32)
    GetOsBuildNumber()(*string)
    GetPhoneNumber()(*string)
    GetProductName()(*string)
    GetResidentUsersCount()(*int32)
    GetSerialNumber()(*string)
    GetSharedDeviceCachedUsers()([]SharedAppleDeviceUserable)
    GetSubnetAddress()(*string)
    GetSubscriberCarrier()(*string)
    GetSystemManagementBIOSVersion()(*string)
    GetTotalStorageSpace()(*int64)
    GetTpmManufacturer()(*string)
    GetTpmSpecificationVersion()(*string)
    GetTpmVersion()(*string)
    GetWifiMac()(*string)
    GetWiredIPv4Addresses()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBatteryChargeCycles(value *int32)()
    SetBatteryHealthPercentage(value *int32)()
    SetBatteryLevelPercentage(value *float64)()
    SetBatterySerialNumber(value *string)()
    SetCellularTechnology(value *string)()
    SetDeviceFullQualifiedDomainName(value *string)()
    SetDeviceGuardLocalSystemAuthorityCredentialGuardState(value *DeviceGuardLocalSystemAuthorityCredentialGuardState)()
    SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(value *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)()
    SetDeviceGuardVirtualizationBasedSecurityState(value *DeviceGuardVirtualizationBasedSecurityState)()
    SetDeviceLicensingLastErrorCode(value *int32)()
    SetDeviceLicensingLastErrorDescription(value *string)()
    SetDeviceLicensingStatus(value *DeviceLicensingStatus)()
    SetEsimIdentifier(value *string)()
    SetFreeStorageSpace(value *int64)()
    SetImei(value *string)()
    SetIpAddressV4(value *string)()
    SetIsEncrypted(value *bool)()
    SetIsSharedDevice(value *bool)()
    SetIsSupervised(value *bool)()
    SetManufacturer(value *string)()
    SetMeid(value *string)()
    SetModel(value *string)()
    SetOdataType(value *string)()
    SetOperatingSystemEdition(value *string)()
    SetOperatingSystemLanguage(value *string)()
    SetOperatingSystemProductType(value *int32)()
    SetOsBuildNumber(value *string)()
    SetPhoneNumber(value *string)()
    SetProductName(value *string)()
    SetResidentUsersCount(value *int32)()
    SetSerialNumber(value *string)()
    SetSharedDeviceCachedUsers(value []SharedAppleDeviceUserable)()
    SetSubnetAddress(value *string)()
    SetSubscriberCarrier(value *string)()
    SetSystemManagementBIOSVersion(value *string)()
    SetTotalStorageSpace(value *int64)()
    SetTpmManufacturer(value *string)()
    SetTpmSpecificationVersion(value *string)()
    SetTpmVersion(value *string)()
    SetWifiMac(value *string)()
    SetWiredIPv4Addresses(value []string)()
}
