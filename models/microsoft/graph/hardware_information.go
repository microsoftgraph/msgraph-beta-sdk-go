package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// HardwareInformation 
type HardwareInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
    batteryChargeCycles *int32;
    // The device’s current battery’s health percentage. Valid values 0 to 100
    batteryHealthPercentage *int32;
    // The serial number of the device’s current battery
    batterySerialNumber *string;
    // Cellular technology of the device
    cellularTechnology *string;
    // Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
    deviceFullQualifiedDomainName *string;
    // Local System Authority (LSA) credential guard status. . Possible values are: running, rebootRequired, notLicensed, notConfigured, virtualizationBasedSecurityNotRunning.
    deviceGuardLocalSystemAuthorityCredentialGuardState *DeviceGuardLocalSystemAuthorityCredentialGuardState;
    // Virtualization-based security hardware requirement status. Possible values are: meetHardwareRequirements, secureBootRequired, dmaProtectionRequired, hyperVNotSupportedForGuestVM, hyperVNotAvailable.
    deviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState;
    // Virtualization-based security status. . Possible values are: running, rebootRequired, require64BitArchitecture, notLicensed, notConfigured, doesNotMeetHardwareRequirements, other.
    deviceGuardVirtualizationBasedSecurityState *DeviceGuardVirtualizationBasedSecurityState;
    // eSIM identifier
    esimIdentifier *string;
    // Free storage space of the device.
    freeStorageSpace *int64;
    // IMEI
    imei *string;
    // IPAddressV4
    ipAddressV4 *string;
    // Encryption status of the device
    isEncrypted *bool;
    // Shared iPad
    isSharedDevice *bool;
    // Supervised mode of the device
    isSupervised *bool;
    // Manufacturer of the device
    manufacturer *string;
    // MEID
    meid *string;
    // Model of the device
    model *string;
    // String that specifies the OS edition.
    operatingSystemEdition *string;
    // Operating system language of the device
    operatingSystemLanguage *string;
    // Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
    operatingSystemProductType *int32;
    // Operating System Build Number on Android device
    osBuildNumber *string;
    // Phone number of the device
    phoneNumber *string;
    // Serial number.
    serialNumber *string;
    // All users on the shared Apple device
    sharedDeviceCachedUsers []SharedAppleDeviceUser;
    // SubnetAddress
    subnetAddress *string;
    // Subscriber carrier of the device
    subscriberCarrier *string;
    // BIOS version as reported by SMBIOS
    systemManagementBIOSVersion *string;
    // Total storage space of the device.
    totalStorageSpace *int64;
    // The identifying information that uniquely names the TPM manufacturer
    tpmManufacturer *string;
    // String that specifies the specification version.
    tpmSpecificationVersion *string;
    // The version of the TPM, as specified by the manufacturer
    tpmVersion *string;
    // WiFi MAC address of the device
    wifiMac *string;
}
// NewHardwareInformation instantiates a new hardwareInformation and sets the default values.
func NewHardwareInformation()(*HardwareInformation) {
    m := &HardwareInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HardwareInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBatteryChargeCycles gets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
func (m *HardwareInformation) GetBatteryChargeCycles()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryChargeCycles
    }
}
// GetBatteryHealthPercentage gets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
func (m *HardwareInformation) GetBatteryHealthPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryHealthPercentage
    }
}
// GetBatterySerialNumber gets the batterySerialNumber property value. The serial number of the device’s current battery
func (m *HardwareInformation) GetBatterySerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.batterySerialNumber
    }
}
// GetCellularTechnology gets the cellularTechnology property value. Cellular technology of the device
func (m *HardwareInformation) GetCellularTechnology()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cellularTechnology
    }
}
// GetDeviceFullQualifiedDomainName gets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
func (m *HardwareInformation) GetDeviceFullQualifiedDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceFullQualifiedDomainName
    }
}
// GetDeviceGuardLocalSystemAuthorityCredentialGuardState gets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. Local System Authority (LSA) credential guard status. . Possible values are: running, rebootRequired, notLicensed, notConfigured, virtualizationBasedSecurityNotRunning.
func (m *HardwareInformation) GetDeviceGuardLocalSystemAuthorityCredentialGuardState()(*DeviceGuardLocalSystemAuthorityCredentialGuardState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardLocalSystemAuthorityCredentialGuardState
    }
}
// GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState gets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. Virtualization-based security hardware requirement status. Possible values are: meetHardwareRequirements, secureBootRequired, dmaProtectionRequired, hyperVNotSupportedForGuestVM, hyperVNotAvailable.
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState
    }
}
// GetDeviceGuardVirtualizationBasedSecurityState gets the deviceGuardVirtualizationBasedSecurityState property value. Virtualization-based security status. . Possible values are: running, rebootRequired, require64BitArchitecture, notLicensed, notConfigured, doesNotMeetHardwareRequirements, other.
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityState()(*DeviceGuardVirtualizationBasedSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardVirtualizationBasedSecurityState
    }
}
// GetEsimIdentifier gets the esimIdentifier property value. eSIM identifier
func (m *HardwareInformation) GetEsimIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.esimIdentifier
    }
}
// GetFreeStorageSpace gets the freeStorageSpace property value. Free storage space of the device.
func (m *HardwareInformation) GetFreeStorageSpace()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpace
    }
}
// GetImei gets the imei property value. IMEI
func (m *HardwareInformation) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
// GetIpAddressV4 gets the ipAddressV4 property value. IPAddressV4
func (m *HardwareInformation) GetIpAddressV4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressV4
    }
}
// GetIsEncrypted gets the isEncrypted property value. Encryption status of the device
func (m *HardwareInformation) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// GetIsSharedDevice gets the isSharedDevice property value. Shared iPad
func (m *HardwareInformation) GetIsSharedDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSharedDevice
    }
}
// GetIsSupervised gets the isSupervised property value. Supervised mode of the device
func (m *HardwareInformation) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device
func (m *HardwareInformation) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetMeid gets the meid property value. MEID
func (m *HardwareInformation) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
// GetModel gets the model property value. Model of the device
func (m *HardwareInformation) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetOperatingSystemEdition gets the operatingSystemEdition property value. String that specifies the OS edition.
func (m *HardwareInformation) GetOperatingSystemEdition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemEdition
    }
}
// GetOperatingSystemLanguage gets the operatingSystemLanguage property value. Operating system language of the device
func (m *HardwareInformation) GetOperatingSystemLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemLanguage
    }
}
// GetOperatingSystemProductType gets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
func (m *HardwareInformation) GetOperatingSystemProductType()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemProductType
    }
}
// GetOsBuildNumber gets the osBuildNumber property value. Operating System Build Number on Android device
func (m *HardwareInformation) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device
func (m *HardwareInformation) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// GetSerialNumber gets the serialNumber property value. Serial number.
func (m *HardwareInformation) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetSharedDeviceCachedUsers gets the sharedDeviceCachedUsers property value. All users on the shared Apple device
func (m *HardwareInformation) GetSharedDeviceCachedUsers()([]SharedAppleDeviceUser) {
    if m == nil {
        return nil
    } else {
        return m.sharedDeviceCachedUsers
    }
}
// GetSubnetAddress gets the subnetAddress property value. SubnetAddress
func (m *HardwareInformation) GetSubnetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetAddress
    }
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber carrier of the device
func (m *HardwareInformation) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
// GetSystemManagementBIOSVersion gets the systemManagementBIOSVersion property value. BIOS version as reported by SMBIOS
func (m *HardwareInformation) GetSystemManagementBIOSVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.systemManagementBIOSVersion
    }
}
// GetTotalStorageSpace gets the totalStorageSpace property value. Total storage space of the device.
func (m *HardwareInformation) GetTotalStorageSpace()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpace
    }
}
// GetTpmManufacturer gets the tpmManufacturer property value. The identifying information that uniquely names the TPM manufacturer
func (m *HardwareInformation) GetTpmManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmManufacturer
    }
}
// GetTpmSpecificationVersion gets the tpmSpecificationVersion property value. String that specifies the specification version.
func (m *HardwareInformation) GetTpmSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmSpecificationVersion
    }
}
// GetTpmVersion gets the tpmVersion property value. The version of the TPM, as specified by the manufacturer
func (m *HardwareInformation) GetTpmVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmVersion
    }
}
// GetWifiMac gets the wifiMac property value. WiFi MAC address of the device
func (m *HardwareInformation) GetWifiMac()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMac
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardwareInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["batteryChargeCycles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryChargeCycles(val)
        }
        return nil
    }
    res["batteryHealthPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthPercentage(val)
        }
        return nil
    }
    res["batterySerialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatterySerialNumber(val)
        }
        return nil
    }
    res["cellularTechnology"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularTechnology(val)
        }
        return nil
    }
    res["deviceFullQualifiedDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFullQualifiedDomainName(val)
        }
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardLocalSystemAuthorityCredentialGuardState(val.(*DeviceGuardLocalSystemAuthorityCredentialGuardState))
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityHardwareRequirementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(val.(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState))
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceGuardVirtualizationBasedSecurityState(val.(*DeviceGuardVirtualizationBasedSecurityState))
        }
        return nil
    }
    res["esimIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsimIdentifier(val)
        }
        return nil
    }
    res["freeStorageSpace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpace(val)
        }
        return nil
    }
    res["imei"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["ipAddressV4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressV4(val)
        }
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSharedDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSharedDevice(val)
        }
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["operatingSystemEdition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemEdition(val)
        }
        return nil
    }
    res["operatingSystemLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemLanguage(val)
        }
        return nil
    }
    res["operatingSystemProductType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemProductType(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["sharedDeviceCachedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharedAppleDeviceUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedAppleDeviceUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*SharedAppleDeviceUser))
            }
            m.SetSharedDeviceCachedUsers(res)
        }
        return nil
    }
    res["subnetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetAddress(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["systemManagementBIOSVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemManagementBIOSVersion(val)
        }
        return nil
    }
    res["totalStorageSpace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpace(val)
        }
        return nil
    }
    res["tpmManufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmManufacturer(val)
        }
        return nil
    }
    res["tpmSpecificationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmSpecificationVersion(val)
        }
        return nil
    }
    res["tpmVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmVersion(val)
        }
        return nil
    }
    res["wifiMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiMac(val)
        }
        return nil
    }
    return res
}
func (m *HardwareInformation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *HardwareInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetSharedDeviceCachedUsers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharedDeviceCachedUsers()))
        for i, v := range m.GetSharedDeviceCachedUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HardwareInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBatteryChargeCycles sets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
func (m *HardwareInformation) SetBatteryChargeCycles(value *int32)() {
    if m != nil {
        m.batteryChargeCycles = value
    }
}
// SetBatteryHealthPercentage sets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
func (m *HardwareInformation) SetBatteryHealthPercentage(value *int32)() {
    if m != nil {
        m.batteryHealthPercentage = value
    }
}
// SetBatterySerialNumber sets the batterySerialNumber property value. The serial number of the device’s current battery
func (m *HardwareInformation) SetBatterySerialNumber(value *string)() {
    if m != nil {
        m.batterySerialNumber = value
    }
}
// SetCellularTechnology sets the cellularTechnology property value. Cellular technology of the device
func (m *HardwareInformation) SetCellularTechnology(value *string)() {
    if m != nil {
        m.cellularTechnology = value
    }
}
// SetDeviceFullQualifiedDomainName sets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
func (m *HardwareInformation) SetDeviceFullQualifiedDomainName(value *string)() {
    if m != nil {
        m.deviceFullQualifiedDomainName = value
    }
}
// SetDeviceGuardLocalSystemAuthorityCredentialGuardState sets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. Local System Authority (LSA) credential guard status. . Possible values are: running, rebootRequired, notLicensed, notConfigured, virtualizationBasedSecurityNotRunning.
func (m *HardwareInformation) SetDeviceGuardLocalSystemAuthorityCredentialGuardState(value *DeviceGuardLocalSystemAuthorityCredentialGuardState)() {
    if m != nil {
        m.deviceGuardLocalSystemAuthorityCredentialGuardState = value
    }
}
// SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState sets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. Virtualization-based security hardware requirement status. Possible values are: meetHardwareRequirements, secureBootRequired, dmaProtectionRequired, hyperVNotSupportedForGuestVM, hyperVNotAvailable.
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(value *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)() {
    if m != nil {
        m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState = value
    }
}
// SetDeviceGuardVirtualizationBasedSecurityState sets the deviceGuardVirtualizationBasedSecurityState property value. Virtualization-based security status. . Possible values are: running, rebootRequired, require64BitArchitecture, notLicensed, notConfigured, doesNotMeetHardwareRequirements, other.
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityState(value *DeviceGuardVirtualizationBasedSecurityState)() {
    if m != nil {
        m.deviceGuardVirtualizationBasedSecurityState = value
    }
}
// SetEsimIdentifier sets the esimIdentifier property value. eSIM identifier
func (m *HardwareInformation) SetEsimIdentifier(value *string)() {
    if m != nil {
        m.esimIdentifier = value
    }
}
// SetFreeStorageSpace sets the freeStorageSpace property value. Free storage space of the device.
func (m *HardwareInformation) SetFreeStorageSpace(value *int64)() {
    if m != nil {
        m.freeStorageSpace = value
    }
}
// SetImei sets the imei property value. IMEI
func (m *HardwareInformation) SetImei(value *string)() {
    if m != nil {
        m.imei = value
    }
}
// SetIpAddressV4 sets the ipAddressV4 property value. IPAddressV4
func (m *HardwareInformation) SetIpAddressV4(value *string)() {
    if m != nil {
        m.ipAddressV4 = value
    }
}
// SetIsEncrypted sets the isEncrypted property value. Encryption status of the device
func (m *HardwareInformation) SetIsEncrypted(value *bool)() {
    if m != nil {
        m.isEncrypted = value
    }
}
// SetIsSharedDevice sets the isSharedDevice property value. Shared iPad
func (m *HardwareInformation) SetIsSharedDevice(value *bool)() {
    if m != nil {
        m.isSharedDevice = value
    }
}
// SetIsSupervised sets the isSupervised property value. Supervised mode of the device
func (m *HardwareInformation) SetIsSupervised(value *bool)() {
    if m != nil {
        m.isSupervised = value
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device
func (m *HardwareInformation) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetMeid sets the meid property value. MEID
func (m *HardwareInformation) SetMeid(value *string)() {
    if m != nil {
        m.meid = value
    }
}
// SetModel sets the model property value. Model of the device
func (m *HardwareInformation) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetOperatingSystemEdition sets the operatingSystemEdition property value. String that specifies the OS edition.
func (m *HardwareInformation) SetOperatingSystemEdition(value *string)() {
    if m != nil {
        m.operatingSystemEdition = value
    }
}
// SetOperatingSystemLanguage sets the operatingSystemLanguage property value. Operating system language of the device
func (m *HardwareInformation) SetOperatingSystemLanguage(value *string)() {
    if m != nil {
        m.operatingSystemLanguage = value
    }
}
// SetOperatingSystemProductType sets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
func (m *HardwareInformation) SetOperatingSystemProductType(value *int32)() {
    if m != nil {
        m.operatingSystemProductType = value
    }
}
// SetOsBuildNumber sets the osBuildNumber property value. Operating System Build Number on Android device
func (m *HardwareInformation) SetOsBuildNumber(value *string)() {
    if m != nil {
        m.osBuildNumber = value
    }
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device
func (m *HardwareInformation) SetPhoneNumber(value *string)() {
    if m != nil {
        m.phoneNumber = value
    }
}
// SetSerialNumber sets the serialNumber property value. Serial number.
func (m *HardwareInformation) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetSharedDeviceCachedUsers sets the sharedDeviceCachedUsers property value. All users on the shared Apple device
func (m *HardwareInformation) SetSharedDeviceCachedUsers(value []SharedAppleDeviceUser)() {
    if m != nil {
        m.sharedDeviceCachedUsers = value
    }
}
// SetSubnetAddress sets the subnetAddress property value. SubnetAddress
func (m *HardwareInformation) SetSubnetAddress(value *string)() {
    if m != nil {
        m.subnetAddress = value
    }
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber carrier of the device
func (m *HardwareInformation) SetSubscriberCarrier(value *string)() {
    if m != nil {
        m.subscriberCarrier = value
    }
}
// SetSystemManagementBIOSVersion sets the systemManagementBIOSVersion property value. BIOS version as reported by SMBIOS
func (m *HardwareInformation) SetSystemManagementBIOSVersion(value *string)() {
    if m != nil {
        m.systemManagementBIOSVersion = value
    }
}
// SetTotalStorageSpace sets the totalStorageSpace property value. Total storage space of the device.
func (m *HardwareInformation) SetTotalStorageSpace(value *int64)() {
    if m != nil {
        m.totalStorageSpace = value
    }
}
// SetTpmManufacturer sets the tpmManufacturer property value. The identifying information that uniquely names the TPM manufacturer
func (m *HardwareInformation) SetTpmManufacturer(value *string)() {
    if m != nil {
        m.tpmManufacturer = value
    }
}
// SetTpmSpecificationVersion sets the tpmSpecificationVersion property value. String that specifies the specification version.
func (m *HardwareInformation) SetTpmSpecificationVersion(value *string)() {
    if m != nil {
        m.tpmSpecificationVersion = value
    }
}
// SetTpmVersion sets the tpmVersion property value. The version of the TPM, as specified by the manufacturer
func (m *HardwareInformation) SetTpmVersion(value *string)() {
    if m != nil {
        m.tpmVersion = value
    }
}
// SetWifiMac sets the wifiMac property value. WiFi MAC address of the device
func (m *HardwareInformation) SetWifiMac(value *string)() {
    if m != nil {
        m.wifiMac = value
    }
}
