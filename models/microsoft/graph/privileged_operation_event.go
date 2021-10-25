package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrivilegedOperationEvent struct {
    Entity
    additionalInformation *string;
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    referenceKey *string;
    referenceSystem *string;
    requestorId *string;
    requestorName *string;
    requestType *string;
    roleId *string;
    roleName *string;
    tenantId *string;
    userId *string;
    userMail *string;
    userName *string;
}
func NewPrivilegedOperationEvent()(*PrivilegedOperationEvent) {
    m := &PrivilegedOperationEvent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrivilegedOperationEvent) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
func (m *PrivilegedOperationEvent) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
func (m *PrivilegedOperationEvent) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
func (m *PrivilegedOperationEvent) GetReferenceKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceKey
    }
}
func (m *PrivilegedOperationEvent) GetReferenceSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referenceSystem
    }
}
func (m *PrivilegedOperationEvent) GetRequestorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestorId
    }
}
func (m *PrivilegedOperationEvent) GetRequestorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestorName
    }
}
func (m *PrivilegedOperationEvent) GetRequestType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
func (m *PrivilegedOperationEvent) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
func (m *PrivilegedOperationEvent) GetRoleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleName
    }
}
func (m *PrivilegedOperationEvent) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *PrivilegedOperationEvent) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *PrivilegedOperationEvent) GetUserMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userMail
    }
}
func (m *PrivilegedOperationEvent) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
func (m *PrivilegedOperationEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAdditionalInformation(val)
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreationDateTime(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["referenceKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferenceKey(val)
        return nil
    }
    res["referenceSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferenceSystem(val)
        return nil
    }
    res["requestorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestorId(val)
        return nil
    }
    res["requestorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestorName(val)
        return nil
    }
    res["requestType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestType(val)
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleId(val)
        return nil
    }
    res["roleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
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
    res["userMail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserMail(val)
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
func (m *PrivilegedOperationEvent) IsNil()(bool) {
    return m == nil
}
func (m *PrivilegedOperationEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceKey", m.GetReferenceKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceSystem", m.GetReferenceSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorId", m.GetRequestorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorName", m.GetRequestorName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestType", m.GetRequestType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleName", m.GetRoleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
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
        err = writer.WriteStringValue("userMail", m.GetUserMail())
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
func (m *PrivilegedOperationEvent) SetAdditionalInformation(value *string)() {
    m.additionalInformation = value
}
func (m *PrivilegedOperationEvent) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
func (m *PrivilegedOperationEvent) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
func (m *PrivilegedOperationEvent) SetReferenceKey(value *string)() {
    m.referenceKey = value
}
func (m *PrivilegedOperationEvent) SetReferenceSystem(value *string)() {
    m.referenceSystem = value
}
func (m *PrivilegedOperationEvent) SetRequestorId(value *string)() {
    m.requestorId = value
}
func (m *PrivilegedOperationEvent) SetRequestorName(value *string)() {
    m.requestorName = value
}
func (m *PrivilegedOperationEvent) SetRequestType(value *string)() {
    m.requestType = value
}
func (m *PrivilegedOperationEvent) SetRoleId(value *string)() {
    m.roleId = value
}
func (m *PrivilegedOperationEvent) SetRoleName(value *string)() {
    m.roleName = value
}
func (m *PrivilegedOperationEvent) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *PrivilegedOperationEvent) SetUserId(value *string)() {
    m.userId = value
}
func (m *PrivilegedOperationEvent) SetUserMail(value *string)() {
    m.userMail = value
}
func (m *PrivilegedOperationEvent) SetUserName(value *string)() {
    m.userName = value
}
