package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ComanagementEligibleDevice struct {
    Entity
    // ClientRegistrationStatus. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
    clientRegistrationStatus *DeviceRegistrationState;
    // DeviceName
    deviceName *string;
    // DeviceType. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
    deviceType *DeviceType;
    // EntitySource
    entitySource *int32;
    // ManagementAgents. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController, microsoft365ManagedMdm, msSense, intuneAosp.
    managementAgents *ManagementAgentType;
    // ManagementState. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
    managementState *ManagementState;
    // Manufacturer
    manufacturer *string;
    // MDMStatus
    mdmStatus *string;
    // Model
    model *string;
    // OSDescription
    osDescription *string;
    // OSVersion
    osVersion *string;
    // OwnerType. Possible values are: unknown, company, personal.
    ownerType *OwnerType;
    // ReferenceId
    referenceId *string;
    // SerialNumber
    serialNumber *string;
    // ComanagementEligibleStatus. Possible values are: comanaged, eligible, eligibleButNotAzureAdJoined, needsOsUpdate, ineligible.
    status *ComanagementEligibleType;
    // UPN
    upn *string;
    // UserEmail
    userEmail *string;
    // UserId
    userId *string;
    // UserName
    userName *string;
}
// Instantiates a new comanagementEligibleDevice and sets the default values.
func NewComanagementEligibleDevice()(*ComanagementEligibleDevice) {
    m := &ComanagementEligibleDevice{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the clientRegistrationStatus property value. ClientRegistrationStatus. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ComanagementEligibleDevice) GetClientRegistrationStatus()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.clientRegistrationStatus
    }
}
// Gets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the deviceType property value. DeviceType. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *ComanagementEligibleDevice) GetDeviceType()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// Gets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) GetEntitySource()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.entitySource
    }
}
// Gets the managementAgents property value. ManagementAgents. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController, microsoft365ManagedMdm, msSense, intuneAosp.
func (m *ComanagementEligibleDevice) GetManagementAgents()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgents
    }
}
// Gets the managementState property value. ManagementState. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
func (m *ComanagementEligibleDevice) GetManagementState()(*ManagementState) {
    if m == nil {
        return nil
    } else {
        return m.managementState
    }
}
// Gets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) GetMdmStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmStatus
    }
}
// Gets the model property value. Model
func (m *ComanagementEligibleDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// Gets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the ownerType property value. OwnerType. Possible values are: unknown, company, personal.
func (m *ComanagementEligibleDevice) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// Gets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) GetReferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceId
    }
}
// Gets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// Gets the status property value. ComanagementEligibleStatus. Possible values are: comanaged, eligible, eligibleButNotAzureAdJoined, needsOsUpdate, ineligible.
func (m *ComanagementEligibleDevice) GetStatus()(*ComanagementEligibleType) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the upn property value. UPN
func (m *ComanagementEligibleDevice) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
// Gets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// Gets the userId property value. UserId
func (m *ComanagementEligibleDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userName property value. UserName
func (m *ComanagementEligibleDevice) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// The deserialization information for the current model
func (m *ComanagementEligibleDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientRegistrationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceRegistrationState)
            m.SetClientRegistrationStatus(&cast)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceType)
            m.SetDeviceType(&cast)
        }
        return nil
    }
    res["entitySource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitySource(val)
        }
        return nil
    }
    res["managementAgents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ManagementAgentType)
            m.SetManagementAgents(&cast)
        }
        return nil
    }
    res["managementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ManagementState)
            m.SetManagementState(&cast)
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
    res["mdmStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmStatus(val)
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
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(OwnerType)
            m.SetOwnerType(&cast)
        }
        return nil
    }
    res["referenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceId(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComanagementEligibleType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ComanagementEligibleType)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["upn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    res["userEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
func (m *ComanagementEligibleDevice) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ComanagementEligibleDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClientRegistrationStatus() != nil {
        cast := m.GetClientRegistrationStatus().String()
        err = writer.WriteStringValue("clientRegistrationStatus", &cast)
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
    if m.GetDeviceType() != nil {
        cast := m.GetDeviceType().String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("entitySource", m.GetEntitySource())
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgents() != nil {
        cast := m.GetManagementAgents().String()
        err = writer.WriteStringValue("managementAgents", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementState() != nil {
        cast := m.GetManagementState().String()
        err = writer.WriteStringValue("managementState", &cast)
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
        err = writer.WriteStringValue("mdmStatus", m.GetMdmStatus())
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
    if m.GetOwnerType() != nil {
        cast := m.GetOwnerType().String()
        err = writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceId", m.GetReferenceId())
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upn", m.GetUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the clientRegistrationStatus property value. ClientRegistrationStatus. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
// Parameters:
//  - value : Value to set for the clientRegistrationStatus property.
func (m *ComanagementEligibleDevice) SetClientRegistrationStatus(value *DeviceRegistrationState)() {
    m.clientRegistrationStatus = value
}
// Sets the deviceName property value. DeviceName
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *ComanagementEligibleDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the deviceType property value. DeviceType. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
// Parameters:
//  - value : Value to set for the deviceType property.
func (m *ComanagementEligibleDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
// Sets the entitySource property value. EntitySource
// Parameters:
//  - value : Value to set for the entitySource property.
func (m *ComanagementEligibleDevice) SetEntitySource(value *int32)() {
    m.entitySource = value
}
// Sets the managementAgents property value. ManagementAgents. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController, microsoft365ManagedMdm, msSense, intuneAosp.
// Parameters:
//  - value : Value to set for the managementAgents property.
func (m *ComanagementEligibleDevice) SetManagementAgents(value *ManagementAgentType)() {
    m.managementAgents = value
}
// Sets the managementState property value. ManagementState. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
// Parameters:
//  - value : Value to set for the managementState property.
func (m *ComanagementEligibleDevice) SetManagementState(value *ManagementState)() {
    m.managementState = value
}
// Sets the manufacturer property value. Manufacturer
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *ComanagementEligibleDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the mdmStatus property value. MDMStatus
// Parameters:
//  - value : Value to set for the mdmStatus property.
func (m *ComanagementEligibleDevice) SetMdmStatus(value *string)() {
    m.mdmStatus = value
}
// Sets the model property value. Model
// Parameters:
//  - value : Value to set for the model property.
func (m *ComanagementEligibleDevice) SetModel(value *string)() {
    m.model = value
}
// Sets the osDescription property value. OSDescription
// Parameters:
//  - value : Value to set for the osDescription property.
func (m *ComanagementEligibleDevice) SetOsDescription(value *string)() {
    m.osDescription = value
}
// Sets the osVersion property value. OSVersion
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *ComanagementEligibleDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the ownerType property value. OwnerType. Possible values are: unknown, company, personal.
// Parameters:
//  - value : Value to set for the ownerType property.
func (m *ComanagementEligibleDevice) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// Sets the referenceId property value. ReferenceId
// Parameters:
//  - value : Value to set for the referenceId property.
func (m *ComanagementEligibleDevice) SetReferenceId(value *string)() {
    m.referenceId = value
}
// Sets the serialNumber property value. SerialNumber
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *ComanagementEligibleDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// Sets the status property value. ComanagementEligibleStatus. Possible values are: comanaged, eligible, eligibleButNotAzureAdJoined, needsOsUpdate, ineligible.
// Parameters:
//  - value : Value to set for the status property.
func (m *ComanagementEligibleDevice) SetStatus(value *ComanagementEligibleType)() {
    m.status = value
}
// Sets the upn property value. UPN
// Parameters:
//  - value : Value to set for the upn property.
func (m *ComanagementEligibleDevice) SetUpn(value *string)() {
    m.upn = value
}
// Sets the userEmail property value. UserEmail
// Parameters:
//  - value : Value to set for the userEmail property.
func (m *ComanagementEligibleDevice) SetUserEmail(value *string)() {
    m.userEmail = value
}
// Sets the userId property value. UserId
// Parameters:
//  - value : Value to set for the userId property.
func (m *ComanagementEligibleDevice) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userName property value. UserName
// Parameters:
//  - value : Value to set for the userName property.
func (m *ComanagementEligibleDevice) SetUserName(value *string)() {
    m.userName = value
}
