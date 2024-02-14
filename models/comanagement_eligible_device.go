package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDevice device Co-Management eligibility state
type ComanagementEligibleDevice struct {
    Entity
}
// NewComanagementEligibleDevice instantiates a new ComanagementEligibleDevice and sets the default values.
func NewComanagementEligibleDevice()(*ComanagementEligibleDevice) {
    m := &ComanagementEligibleDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateComanagementEligibleDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateComanagementEligibleDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagementEligibleDevice(), nil
}
// GetClientRegistrationStatus gets the clientRegistrationStatus property value. Device registration status.
// returns a *DeviceRegistrationState when successful
func (m *ComanagementEligibleDevice) GetClientRegistrationStatus()(*DeviceRegistrationState) {
    val, err := m.GetBackingStore().Get("clientRegistrationStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceRegistrationState)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. DeviceName
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceType gets the deviceType property value. Device type.
// returns a *DeviceType when successful
func (m *ComanagementEligibleDevice) GetDeviceType()(*DeviceType) {
    val, err := m.GetBackingStore().Get("deviceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceType)
    }
    return nil
}
// GetEntitySource gets the entitySource property value. EntitySource
// returns a *int32 when successful
func (m *ComanagementEligibleDevice) GetEntitySource()(*int32) {
    val, err := m.GetBackingStore().Get("entitySource")
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
func (m *ComanagementEligibleDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientRegistrationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientRegistrationStatus(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceType))
        }
        return nil
    }
    res["entitySource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitySource(val)
        }
        return nil
    }
    res["managementAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgents(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementState(val.(*ManagementState))
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
    res["mdmStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmStatus(val)
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
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["referenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceId(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComanagementEligibleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ComanagementEligibleType))
        }
        return nil
    }
    res["upn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    res["userEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetManagementAgents gets the managementAgents property value. Management agent type.
// returns a *ManagementAgentType when successful
func (m *ComanagementEligibleDevice) GetManagementAgents()(*ManagementAgentType) {
    val, err := m.GetBackingStore().Get("managementAgents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementAgentType)
    }
    return nil
}
// GetManagementState gets the managementState property value. Management state of device in Microsoft Intune.
// returns a *ManagementState when successful
func (m *ComanagementEligibleDevice) GetManagementState()(*ManagementState) {
    val, err := m.GetBackingStore().Get("managementState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagementState)
    }
    return nil
}
// GetManufacturer gets the manufacturer property value. Manufacturer
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMdmStatus gets the mdmStatus property value. MDMStatus
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetMdmStatus()(*string) {
    val, err := m.GetBackingStore().Get("mdmStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. Model
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsDescription gets the osDescription property value. OSDescription
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetOsDescription()(*string) {
    val, err := m.GetBackingStore().Get("osDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. OSVersion
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerType gets the ownerType property value. Owner type of device.
// returns a *OwnerType when successful
func (m *ComanagementEligibleDevice) GetOwnerType()(*OwnerType) {
    val, err := m.GetBackingStore().Get("ownerType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OwnerType)
    }
    return nil
}
// GetReferenceId gets the referenceId property value. ReferenceId
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetReferenceId()(*string) {
    val, err := m.GetBackingStore().Get("referenceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. SerialNumber
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
// returns a *ComanagementEligibleType when successful
func (m *ComanagementEligibleDevice) GetStatus()(*ComanagementEligibleType) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ComanagementEligibleType)
    }
    return nil
}
// GetUpn gets the upn property value. UPN
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetUpn()(*string) {
    val, err := m.GetBackingStore().Get("upn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserEmail gets the userEmail property value. UserEmail
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetUserEmail()(*string) {
    val, err := m.GetBackingStore().Get("userEmail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. UserId
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserName gets the userName property value. UserName
// returns a *string when successful
func (m *ComanagementEligibleDevice) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("clientRegistrationStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ComanagementEligibleDevice) SetDeviceType(value *DeviceType)() {
    err := m.GetBackingStore().Set("deviceType", value)
    if err != nil {
        panic(err)
    }
}
// SetEntitySource sets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) SetEntitySource(value *int32)() {
    err := m.GetBackingStore().Set("entitySource", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementAgents sets the managementAgents property value. Management agent type.
func (m *ComanagementEligibleDevice) SetManagementAgents(value *ManagementAgentType)() {
    err := m.GetBackingStore().Set("managementAgents", value)
    if err != nil {
        panic(err)
    }
}
// SetManagementState sets the managementState property value. Management state of device in Microsoft Intune.
func (m *ComanagementEligibleDevice) SetManagementState(value *ManagementState)() {
    err := m.GetBackingStore().Set("managementState", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetMdmStatus sets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) SetMdmStatus(value *string)() {
    err := m.GetBackingStore().Set("mdmStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. Model
func (m *ComanagementEligibleDevice) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetOsDescription sets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) SetOsDescription(value *string)() {
    err := m.GetBackingStore().Set("osDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *ComanagementEligibleDevice) SetOwnerType(value *OwnerType)() {
    err := m.GetBackingStore().Set("ownerType", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceId sets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) SetReferenceId(value *string)() {
    err := m.GetBackingStore().Set("referenceId", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *ComanagementEligibleDevice) SetStatus(value *ComanagementEligibleType)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUpn sets the upn property value. UPN
func (m *ComanagementEligibleDevice) SetUpn(value *string)() {
    err := m.GetBackingStore().Set("upn", value)
    if err != nil {
        panic(err)
    }
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) SetUserEmail(value *string)() {
    err := m.GetBackingStore().Set("userEmail", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. UserId
func (m *ComanagementEligibleDevice) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. UserName
func (m *ComanagementEligibleDevice) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
type ComanagementEligibleDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientRegistrationStatus()(*DeviceRegistrationState)
    GetDeviceName()(*string)
    GetDeviceType()(*DeviceType)
    GetEntitySource()(*int32)
    GetManagementAgents()(*ManagementAgentType)
    GetManagementState()(*ManagementState)
    GetManufacturer()(*string)
    GetMdmStatus()(*string)
    GetModel()(*string)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetOwnerType()(*OwnerType)
    GetReferenceId()(*string)
    GetSerialNumber()(*string)
    GetStatus()(*ComanagementEligibleType)
    GetUpn()(*string)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    SetClientRegistrationStatus(value *DeviceRegistrationState)()
    SetDeviceName(value *string)()
    SetDeviceType(value *DeviceType)()
    SetEntitySource(value *int32)()
    SetManagementAgents(value *ManagementAgentType)()
    SetManagementState(value *ManagementState)()
    SetManufacturer(value *string)()
    SetMdmStatus(value *string)()
    SetModel(value *string)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetOwnerType(value *OwnerType)()
    SetReferenceId(value *string)()
    SetSerialNumber(value *string)()
    SetStatus(value *ComanagementEligibleType)()
    SetUpn(value *string)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
}
