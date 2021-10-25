package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ComanagementEligibleDevice struct {
    Entity
    clientRegistrationStatus *DeviceRegistrationState;
    deviceName *string;
    deviceType *DeviceType;
    entitySource *int32;
    managementAgents *ManagementAgentType;
    managementState *ManagementState;
    manufacturer *string;
    mdmStatus *string;
    model *string;
    osDescription *string;
    osVersion *string;
    ownerType *OwnerType;
    referenceId *string;
    serialNumber *string;
    status *ComanagementEligibleType;
    upn *string;
    userEmail *string;
    userId *string;
    userName *string;
}
func NewComanagementEligibleDevice()(*ComanagementEligibleDevice) {
    m := &ComanagementEligibleDevice{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ComanagementEligibleDevice) GetClientRegistrationStatus()(*DeviceRegistrationState) {
    if m == nil {
        return nil
    } else {
        return m.clientRegistrationStatus
    }
}
func (m *ComanagementEligibleDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *ComanagementEligibleDevice) GetDeviceType()(*DeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
func (m *ComanagementEligibleDevice) GetEntitySource()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.entitySource
    }
}
func (m *ComanagementEligibleDevice) GetManagementAgents()(*ManagementAgentType) {
    if m == nil {
        return nil
    } else {
        return m.managementAgents
    }
}
func (m *ComanagementEligibleDevice) GetManagementState()(*ManagementState) {
    if m == nil {
        return nil
    } else {
        return m.managementState
    }
}
func (m *ComanagementEligibleDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *ComanagementEligibleDevice) GetMdmStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmStatus
    }
}
func (m *ComanagementEligibleDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *ComanagementEligibleDevice) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
func (m *ComanagementEligibleDevice) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
func (m *ComanagementEligibleDevice) GetOwnerType()(*OwnerType) {
    if m == nil {
        return nil
    } else {
        return m.ownerType
    }
}
func (m *ComanagementEligibleDevice) GetReferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceId
    }
}
func (m *ComanagementEligibleDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
func (m *ComanagementEligibleDevice) GetStatus()(*ComanagementEligibleType) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ComanagementEligibleDevice) GetUpn()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upn
    }
}
func (m *ComanagementEligibleDevice) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
func (m *ComanagementEligibleDevice) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *ComanagementEligibleDevice) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
func (m *ComanagementEligibleDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientRegistrationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        cast := val.(DeviceRegistrationState)
        m.SetClientRegistrationStatus(&cast)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        cast := val.(DeviceType)
        m.SetDeviceType(&cast)
        return nil
    }
    res["entitySource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetEntitySource(val)
        return nil
    }
    res["managementAgents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        cast := val.(ManagementAgentType)
        m.SetManagementAgents(&cast)
        return nil
    }
    res["managementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        cast := val.(ManagementState)
        m.SetManagementState(&cast)
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
    res["mdmStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMdmStatus(val)
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
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsDescription(val)
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsVersion(val)
        return nil
    }
    res["ownerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        cast := val.(OwnerType)
        m.SetOwnerType(&cast)
        return nil
    }
    res["referenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferenceId(val)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseComanagementEligibleType)
        if err != nil {
            return err
        }
        cast := val.(ComanagementEligibleType)
        m.SetStatus(&cast)
        return nil
    }
    res["upn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpn(val)
        return nil
    }
    res["userEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserEmail(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserName(val)
        return nil
    }
    return res
}
func (m *ComanagementEligibleDevice) IsNil()(bool) {
    return m == nil
}
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
func (m *ComanagementEligibleDevice) SetClientRegistrationStatus(value *DeviceRegistrationState)() {
    m.clientRegistrationStatus = value
}
func (m *ComanagementEligibleDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *ComanagementEligibleDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
func (m *ComanagementEligibleDevice) SetEntitySource(value *int32)() {
    m.entitySource = value
}
func (m *ComanagementEligibleDevice) SetManagementAgents(value *ManagementAgentType)() {
    m.managementAgents = value
}
func (m *ComanagementEligibleDevice) SetManagementState(value *ManagementState)() {
    m.managementState = value
}
func (m *ComanagementEligibleDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *ComanagementEligibleDevice) SetMdmStatus(value *string)() {
    m.mdmStatus = value
}
func (m *ComanagementEligibleDevice) SetModel(value *string)() {
    m.model = value
}
func (m *ComanagementEligibleDevice) SetOsDescription(value *string)() {
    m.osDescription = value
}
func (m *ComanagementEligibleDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
func (m *ComanagementEligibleDevice) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
func (m *ComanagementEligibleDevice) SetReferenceId(value *string)() {
    m.referenceId = value
}
func (m *ComanagementEligibleDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
func (m *ComanagementEligibleDevice) SetStatus(value *ComanagementEligibleType)() {
    m.status = value
}
func (m *ComanagementEligibleDevice) SetUpn(value *string)() {
    m.upn = value
}
func (m *ComanagementEligibleDevice) SetUserEmail(value *string)() {
    m.userEmail = value
}
func (m *ComanagementEligibleDevice) SetUserId(value *string)() {
    m.userId = value
}
func (m *ComanagementEligibleDevice) SetUserName(value *string)() {
    m.userName = value
}
