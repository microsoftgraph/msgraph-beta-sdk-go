package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDevice device Co-Management eligibility state
type ComanagementEligibleDevice struct {
    Entity
    // Device registration status.
    clientRegistrationStatus *DeviceRegistrationState
    // DeviceName
    deviceName *string
    // Device type.
    deviceType *DeviceType
    // EntitySource
    entitySource *int32
    // Management agent type.
    managementAgents *ManagementAgentType
    // Management state of device in Microsoft Intune.
    managementState *ManagementState
    // Manufacturer
    manufacturer *string
    // MDMStatus
    mdmStatus *string
    // Model
    model *string
    // OSDescription
    osDescription *string
    // OSVersion
    osVersion *string
    // Owner type of device.
    ownerType *OwnerType
    // ReferenceId
    referenceId *string
    // SerialNumber
    serialNumber *string
    // The status property
    status *ComanagementEligibleType
    // UPN
    upn *string
    // UserEmail
    userEmail *string
    // UserId
    userId *string
    // UserName
    userName *string
}
// NewComanagementEligibleDevice instantiates a new comanagementEligibleDevice and sets the default values.
func NewComanagementEligibleDevice()(*ComanagementEligibleDevice) {
    m := &ComanagementEligibleDevice{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.comanagementEligibleDevice";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateComanagementEligibleDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagementEligibleDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagementEligibleDevice(), nil
}
// GetClientRegistrationStatus gets the clientRegistrationStatus property value. Device registration status.
func (m *ComanagementEligibleDevice) GetClientRegistrationStatus()(*DeviceRegistrationState) {
    return m.clientRegistrationStatus
}
// GetDeviceName gets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ComanagementEligibleDevice) GetDeviceType()(*DeviceType) {
    return m.deviceType
}
// GetEntitySource gets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) GetEntitySource()(*int32) {
    return m.entitySource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagementEligibleDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientRegistrationStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceRegistrationState , m.SetClientRegistrationStatus)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["deviceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceType , m.SetDeviceType)
    res["entitySource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetEntitySource)
    res["managementAgents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementAgentType , m.SetManagementAgents)
    res["managementState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementState , m.SetManagementState)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["mdmStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMdmStatus)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["osDescription"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsDescription)
    res["osVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsVersion)
    res["ownerType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseOwnerType , m.SetOwnerType)
    res["referenceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetReferenceId)
    res["serialNumber"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSerialNumber)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseComanagementEligibleType , m.SetStatus)
    res["upn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUpn)
    res["userEmail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserEmail)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserName)
    return res
}
// GetManagementAgents gets the managementAgents property value. Management agent type.
func (m *ComanagementEligibleDevice) GetManagementAgents()(*ManagementAgentType) {
    return m.managementAgents
}
// GetManagementState gets the managementState property value. Management state of device in Microsoft Intune.
func (m *ComanagementEligibleDevice) GetManagementState()(*ManagementState) {
    return m.managementState
}
// GetManufacturer gets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMdmStatus gets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) GetMdmStatus()(*string) {
    return m.mdmStatus
}
// GetModel gets the model property value. Model
func (m *ComanagementEligibleDevice) GetModel()(*string) {
    return m.model
}
// GetOsDescription gets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) GetOsDescription()(*string) {
    return m.osDescription
}
// GetOsVersion gets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *ComanagementEligibleDevice) GetOwnerType()(*OwnerType) {
    return m.ownerType
}
// GetReferenceId gets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) GetReferenceId()(*string) {
    return m.referenceId
}
// GetSerialNumber gets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetStatus gets the status property value. The status property
func (m *ComanagementEligibleDevice) GetStatus()(*ComanagementEligibleType) {
    return m.status
}
// GetUpn gets the upn property value. UPN
func (m *ComanagementEligibleDevice) GetUpn()(*string) {
    return m.upn
}
// GetUserEmail gets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) GetUserEmail()(*string) {
    return m.userEmail
}
// GetUserId gets the userId property value. UserId
func (m *ComanagementEligibleDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. UserName
func (m *ComanagementEligibleDevice) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *ComanagementEligibleDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetClientRegistrationStatus sets the clientRegistrationStatus property value. Device registration status.
func (m *ComanagementEligibleDevice) SetClientRegistrationStatus(value *DeviceRegistrationState)() {
    m.clientRegistrationStatus = value
}
// SetDeviceName sets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ComanagementEligibleDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
// SetEntitySource sets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) SetEntitySource(value *int32)() {
    m.entitySource = value
}
// SetManagementAgents sets the managementAgents property value. Management agent type.
func (m *ComanagementEligibleDevice) SetManagementAgents(value *ManagementAgentType)() {
    m.managementAgents = value
}
// SetManagementState sets the managementState property value. Management state of device in Microsoft Intune.
func (m *ComanagementEligibleDevice) SetManagementState(value *ManagementState)() {
    m.managementState = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMdmStatus sets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) SetMdmStatus(value *string)() {
    m.mdmStatus = value
}
// SetModel sets the model property value. Model
func (m *ComanagementEligibleDevice) SetModel(value *string)() {
    m.model = value
}
// SetOsDescription sets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) SetOsDescription(value *string)() {
    m.osDescription = value
}
// SetOsVersion sets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *ComanagementEligibleDevice) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// SetReferenceId sets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) SetReferenceId(value *string)() {
    m.referenceId = value
}
// SetSerialNumber sets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetStatus sets the status property value. The status property
func (m *ComanagementEligibleDevice) SetStatus(value *ComanagementEligibleType)() {
    m.status = value
}
// SetUpn sets the upn property value. UPN
func (m *ComanagementEligibleDevice) SetUpn(value *string)() {
    m.upn = value
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) SetUserEmail(value *string)() {
    m.userEmail = value
}
// SetUserId sets the userId property value. UserId
func (m *ComanagementEligibleDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. UserName
func (m *ComanagementEligibleDevice) SetUserName(value *string)() {
    m.userName = value
}
