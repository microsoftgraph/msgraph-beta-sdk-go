package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComanagementEligibleDevice 
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
// NewComanagementEligibleDevice instantiates a new comanagementEligibleDevice and sets the default values.
func NewComanagementEligibleDevice()(*ComanagementEligibleDevice) {
    m := &ComanagementEligibleDevice{
        Entity: *NewEntity(),
    }
    return m
}
// GetClientRegistrationStatus gets the clientRegistrationStatus property value. ClientRegistrationStatus. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ComanagementEligibleDevice) GetClientRegistrationStatus()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.clientRegistrationStatus
    }
}
// GetDeviceName gets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetDeviceType gets the deviceType property value. DeviceType. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *ComanagementEligibleDevice) GetDeviceType()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetEntitySource gets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) GetEntitySource()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.entitySource
    }
}
// GetManagementAgents gets the managementAgents property value. ManagementAgents. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController, microsoft365ManagedMdm, msSense, intuneAosp.
func (m *ComanagementEligibleDevice) GetManagementAgents()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgents
    }
}
// GetManagementState gets the managementState property value. ManagementState. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
func (m *ComanagementEligibleDevice) GetManagementState()(*ManagementState) {
    if m == nil {
        return nil
    } else {
        return m.managementState
    }
}
// GetManufacturer gets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetMdmStatus gets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) GetMdmStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmStatus
    }
}
// GetModel gets the model property value. Model
func (m *ComanagementEligibleDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetOsDescription gets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// GetOsVersion gets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetOwnerType gets the ownerType property value. OwnerType. Possible values are: unknown, company, personal.
func (m *ComanagementEligibleDevice) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
// GetReferenceId gets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) GetReferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceId
    }
}
// GetSerialNumber gets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetStatus gets the status property value. ComanagementEligibleStatus. Possible values are: comanaged, eligible, eligibleButNotAzureAdJoined, needsOsUpdate, ineligible.
func (m *ComanagementEligibleDevice) GetStatus()(*ComanagementEligibleType) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetUpn gets the upn property value. UPN
func (m *ComanagementEligibleDevice) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
// GetUserEmail gets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// GetUserId gets the userId property value. UserId
func (m *ComanagementEligibleDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserName gets the userName property value. UserName
func (m *ComanagementEligibleDevice) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagementEligibleDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientRegistrationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientRegistrationStatus(val.(*DeviceRegistrationState))
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
            m.SetDeviceType(val.(*DeviceType))
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
            m.SetManagementAgents(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementState(val.(*ManagementState))
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
            m.SetOwnerType(val.(*OwnerType))
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
            m.SetStatus(val.(*ComanagementEligibleType))
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
// Serialize serializes information the current object
func (m *ComanagementEligibleDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClientRegistrationStatus() != nil {
        cast := (*m.GetClientRegistrationStatus()).String()
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
        cast := (*m.GetDeviceType()).String()
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
        cast := (*m.GetManagementAgents()).String()
        err = writer.WriteStringValue("managementAgents", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementState() != nil {
        cast := (*m.GetManagementState()).String()
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
        cast := (*m.GetOwnerType()).String()
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
        cast := (*m.GetStatus()).String()
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
// SetClientRegistrationStatus sets the clientRegistrationStatus property value. ClientRegistrationStatus. Possible values are: notRegistered, registered, revoked, keyConflict, approvalPending, certificateReset, notRegisteredPendingEnrollment, unknown.
func (m *ComanagementEligibleDevice) SetClientRegistrationStatus(value *DeviceRegistrationState)() {
    if m != nil {
        m.clientRegistrationStatus = value
    }
}
// SetDeviceName sets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetDeviceType sets the deviceType property value. DeviceType. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.
func (m *ComanagementEligibleDevice) SetDeviceType(value *DeviceType)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetEntitySource sets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) SetEntitySource(value *int32)() {
    if m != nil {
        m.entitySource = value
    }
}
// SetManagementAgents sets the managementAgents property value. ManagementAgents. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController, microsoft365ManagedMdm, msSense, intuneAosp.
func (m *ComanagementEligibleDevice) SetManagementAgents(value *ManagementAgentType)() {
    if m != nil {
        m.managementAgents = value
    }
}
// SetManagementState sets the managementState property value. ManagementState. Possible values are: managed, retirePending, retireFailed, wipePending, wipeFailed, unhealthy, deletePending, retireIssued, wipeIssued, wipeCanceled, retireCanceled, discovered.
func (m *ComanagementEligibleDevice) SetManagementState(value *ManagementState)() {
    if m != nil {
        m.managementState = value
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetMdmStatus sets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) SetMdmStatus(value *string)() {
    if m != nil {
        m.mdmStatus = value
    }
}
// SetModel sets the model property value. Model
func (m *ComanagementEligibleDevice) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetOsDescription sets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) SetOsDescription(value *string)() {
    if m != nil {
        m.osDescription = value
    }
}
// SetOsVersion sets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetOwnerType sets the ownerType property value. OwnerType. Possible values are: unknown, company, personal.
func (m *ComanagementEligibleDevice) SetOwnerType(value *OwnerType)() {
    if m != nil {
        m.ownerType = value
    }
}
// SetReferenceId sets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) SetReferenceId(value *string)() {
    if m != nil {
        m.referenceId = value
    }
}
// SetSerialNumber sets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetStatus sets the status property value. ComanagementEligibleStatus. Possible values are: comanaged, eligible, eligibleButNotAzureAdJoined, needsOsUpdate, ineligible.
func (m *ComanagementEligibleDevice) SetStatus(value *ComanagementEligibleType)() {
    if m != nil {
        m.status = value
    }
}
// SetUpn sets the upn property value. UPN
func (m *ComanagementEligibleDevice) SetUpn(value *string)() {
    if m != nil {
        m.upn = value
    }
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) SetUserEmail(value *string)() {
    if m != nil {
        m.userEmail = value
    }
}
// SetUserId sets the userId property value. UserId
func (m *ComanagementEligibleDevice) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserName sets the userName property value. UserName
func (m *ComanagementEligibleDevice) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
