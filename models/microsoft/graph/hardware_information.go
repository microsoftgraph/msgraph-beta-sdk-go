package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type HardwareInformation struct {
    additionalData map[string]interface{};
    batteryChargeCycles *int32;
    batteryHealthPercentage *int32;
    batterySerialNumber *string;
    cellularTechnology *string;
    deviceFullQualifiedDomainName *string;
    deviceGuardLocalSystemAuthorityCredentialGuardState *DeviceGuardLocalSystemAuthorityCredentialGuardState;
    deviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState;
    deviceGuardVirtualizationBasedSecurityState *DeviceGuardVirtualizationBasedSecurityState;
    esimIdentifier *string;
    freeStorageSpace *int64;
    imei *string;
    ipAddressV4 *string;
    isEncrypted *bool;
    isSharedDevice *bool;
    isSupervised *bool;
    manufacturer *string;
    meid *string;
    model *string;
    operatingSystemEdition *string;
    operatingSystemLanguage *string;
    operatingSystemProductType *int32;
    osBuildNumber *string;
    phoneNumber *string;
    serialNumber *string;
    sharedDeviceCachedUsers []SharedAppleDeviceUser;
    subnetAddress *string;
    subscriberCarrier *string;
    totalStorageSpace *int64;
    tpmSpecificationVersion *string;
    wifiMac *string;
}
func NewHardwareInformation()(*HardwareInformation) {
    m := &HardwareInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *HardwareInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *HardwareInformation) GetBatteryChargeCycles()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryChargeCycles
    }
}
func (m *HardwareInformation) GetBatteryHealthPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryHealthPercentage
    }
}
func (m *HardwareInformation) GetBatterySerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.batterySerialNumber
    }
}
func (m *HardwareInformation) GetCellularTechnology()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cellularTechnology
    }
}
func (m *HardwareInformation) GetDeviceFullQualifiedDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceFullQualifiedDomainName
    }
}
func (m *HardwareInformation) GetDeviceGuardLocalSystemAuthorityCredentialGuardState()(*DeviceGuardLocalSystemAuthorityCredentialGuardState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardLocalSystemAuthorityCredentialGuardState
    }
}
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState
    }
}
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityState()(*DeviceGuardVirtualizationBasedSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardVirtualizationBasedSecurityState
    }
}
func (m *HardwareInformation) GetEsimIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.esimIdentifier
    }
}
func (m *HardwareInformation) GetFreeStorageSpace()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpace
    }
}
func (m *HardwareInformation) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
func (m *HardwareInformation) GetIpAddressV4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressV4
    }
}
func (m *HardwareInformation) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
func (m *HardwareInformation) GetIsSharedDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSharedDevice
    }
}
func (m *HardwareInformation) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
func (m *HardwareInformation) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *HardwareInformation) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
func (m *HardwareInformation) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *HardwareInformation) GetOperatingSystemEdition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemEdition
    }
}
func (m *HardwareInformation) GetOperatingSystemLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemLanguage
    }
}
func (m *HardwareInformation) GetOperatingSystemProductType()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemProductType
    }
}
func (m *HardwareInformation) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
func (m *HardwareInformation) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
func (m *HardwareInformation) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
func (m *HardwareInformation) GetSharedDeviceCachedUsers()([]SharedAppleDeviceUser) {
    if m == nil {
        return nil
    } else {
        return m.sharedDeviceCachedUsers
    }
}
func (m *HardwareInformation) GetSubnetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetAddress
    }
}
func (m *HardwareInformation) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
func (m *HardwareInformation) GetTotalStorageSpace()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpace
    }
}
func (m *HardwareInformation) GetTpmSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmSpecificationVersion
    }
}
func (m *HardwareInformation) GetWifiMac()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMac
    }
}
func (m *HardwareInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["batteryChargeCycles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBatteryChargeCycles(val)
        return nil
    }
    res["batteryHealthPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBatteryHealthPercentage(val)
        return nil
    }
    res["batterySerialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBatterySerialNumber(val)
        return nil
    }
    res["cellularTechnology"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCellularTechnology(val)
        return nil
    }
    res["deviceFullQualifiedDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceFullQualifiedDomainName(val)
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardState)
        if err != nil {
            return err
        }
        cast := val.(DeviceGuardLocalSystemAuthorityCredentialGuardState)
        m.SetDeviceGuardLocalSystemAuthorityCredentialGuardState(&cast)
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityHardwareRequirementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
        if err != nil {
            return err
        }
        cast := val.(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
        m.SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(&cast)
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityState)
        if err != nil {
            return err
        }
        cast := val.(DeviceGuardVirtualizationBasedSecurityState)
        m.SetDeviceGuardVirtualizationBasedSecurityState(&cast)
        return nil
    }
    res["esimIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEsimIdentifier(val)
        return nil
    }
    res["freeStorageSpace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetFreeStorageSpace(val)
        return nil
    }
    res["imei"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImei(val)
        return nil
    }
    res["ipAddressV4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddressV4(val)
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEncrypted(val)
        return nil
    }
    res["isSharedDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSharedDevice(val)
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSupervised(val)
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
    res["meid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMeid(val)
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
    res["operatingSystemEdition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystemEdition(val)
        return nil
    }
    res["operatingSystemLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystemLanguage(val)
        return nil
    }
    res["operatingSystemProductType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOperatingSystemProductType(val)
        return nil
    }
    res["osBuildNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsBuildNumber(val)
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhoneNumber(val)
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
    res["sharedDeviceCachedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharedAppleDeviceUser() })
        if err != nil {
            return err
        }
        res := make([]SharedAppleDeviceUser, len(val))
        for i, v := range val {
            res[i] = *(v.(*SharedAppleDeviceUser))
        }
        m.SetSharedDeviceCachedUsers(res)
        return nil
    }
    res["subnetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubnetAddress(val)
        return nil
    }
    res["subscriberCarrier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubscriberCarrier(val)
        return nil
    }
    res["totalStorageSpace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalStorageSpace(val)
        return nil
    }
    res["tpmSpecificationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTpmSpecificationVersion(val)
        return nil
    }
    res["wifiMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWifiMac(val)
        return nil
    }
    return res
}
func (m *HardwareInformation) IsNil()(bool) {
    return m == nil
}
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
        cast := m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState().String()
        err := writer.WriteStringValue("deviceGuardLocalSystemAuthorityCredentialGuardState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() != nil {
        cast := m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState().String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityHardwareRequirementState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityState() != nil {
        cast := m.GetDeviceGuardVirtualizationBasedSecurityState().String()
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
    {
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
        err := writer.WriteInt64Value("totalStorageSpace", m.GetTotalStorageSpace())
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
func (m *HardwareInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *HardwareInformation) SetBatteryChargeCycles(value *int32)() {
    m.batteryChargeCycles = value
}
func (m *HardwareInformation) SetBatteryHealthPercentage(value *int32)() {
    m.batteryHealthPercentage = value
}
func (m *HardwareInformation) SetBatterySerialNumber(value *string)() {
    m.batterySerialNumber = value
}
func (m *HardwareInformation) SetCellularTechnology(value *string)() {
    m.cellularTechnology = value
}
func (m *HardwareInformation) SetDeviceFullQualifiedDomainName(value *string)() {
    m.deviceFullQualifiedDomainName = value
}
func (m *HardwareInformation) SetDeviceGuardLocalSystemAuthorityCredentialGuardState(value *DeviceGuardLocalSystemAuthorityCredentialGuardState)() {
    m.deviceGuardLocalSystemAuthorityCredentialGuardState = value
}
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(value *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)() {
    m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState = value
}
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityState(value *DeviceGuardVirtualizationBasedSecurityState)() {
    m.deviceGuardVirtualizationBasedSecurityState = value
}
func (m *HardwareInformation) SetEsimIdentifier(value *string)() {
    m.esimIdentifier = value
}
func (m *HardwareInformation) SetFreeStorageSpace(value *int64)() {
    m.freeStorageSpace = value
}
func (m *HardwareInformation) SetImei(value *string)() {
    m.imei = value
}
func (m *HardwareInformation) SetIpAddressV4(value *string)() {
    m.ipAddressV4 = value
}
func (m *HardwareInformation) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
func (m *HardwareInformation) SetIsSharedDevice(value *bool)() {
    m.isSharedDevice = value
}
func (m *HardwareInformation) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
func (m *HardwareInformation) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *HardwareInformation) SetMeid(value *string)() {
    m.meid = value
}
func (m *HardwareInformation) SetModel(value *string)() {
    m.model = value
}
func (m *HardwareInformation) SetOperatingSystemEdition(value *string)() {
    m.operatingSystemEdition = value
}
func (m *HardwareInformation) SetOperatingSystemLanguage(value *string)() {
    m.operatingSystemLanguage = value
}
func (m *HardwareInformation) SetOperatingSystemProductType(value *int32)() {
    m.operatingSystemProductType = value
}
func (m *HardwareInformation) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
func (m *HardwareInformation) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
func (m *HardwareInformation) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
func (m *HardwareInformation) SetSharedDeviceCachedUsers(value []SharedAppleDeviceUser)() {
    m.sharedDeviceCachedUsers = value
}
func (m *HardwareInformation) SetSubnetAddress(value *string)() {
    m.subnetAddress = value
}
func (m *HardwareInformation) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
func (m *HardwareInformation) SetTotalStorageSpace(value *int64)() {
    m.totalStorageSpace = value
}
func (m *HardwareInformation) SetTpmSpecificationVersion(value *string)() {
    m.tpmSpecificationVersion = value
}
func (m *HardwareInformation) SetWifiMac(value *string)() {
    m.wifiMac = value
}
