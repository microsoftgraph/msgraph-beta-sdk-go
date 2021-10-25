package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuditActor struct {
    additionalData map[string]interface{};
    applicationDisplayName *string;
    applicationId *string;
    ipAddress *string;
    remoteTenantId *string;
    remoteUserId *string;
    servicePrincipalName *string;
    type_escaped *string;
    userId *string;
    userPermissions []string;
    userPrincipalName *string;
    userRoleScopeTags []RoleScopeTagInfo;
}
func NewAuditActor()(*AuditActor) {
    m := &AuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AuditActor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AuditActor) GetApplicationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationDisplayName
    }
}
func (m *AuditActor) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
func (m *AuditActor) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
func (m *AuditActor) GetRemoteTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteTenantId
    }
}
func (m *AuditActor) GetRemoteUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteUserId
    }
}
func (m *AuditActor) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
func (m *AuditActor) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *AuditActor) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *AuditActor) GetUserPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userPermissions
    }
}
func (m *AuditActor) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *AuditActor) GetUserRoleScopeTags()([]RoleScopeTagInfo) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
func (m *AuditActor) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationDisplayName(val)
        return nil
    }
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationId(val)
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIpAddress(val)
        return nil
    }
    res["remoteTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemoteTenantId(val)
        return nil
    }
    res["remoteUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemoteUserId(val)
        return nil
    }
    res["servicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePrincipalName(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
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
    res["userPermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetUserPermissions(res)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["userRoleScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTagInfo() })
        if err != nil {
            return err
        }
        res := make([]RoleScopeTagInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*RoleScopeTagInfo))
        }
        m.SetUserRoleScopeTags(res)
        return nil
    }
    return res
}
func (m *AuditActor) IsNil()(bool) {
    return m == nil
}
func (m *AuditActor) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("applicationDisplayName", m.GetApplicationDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteTenantId", m.GetRemoteTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remoteUserId", m.GetRemoteUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePrincipalName", m.GetServicePrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("userPermissions", m.GetUserPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRoleScopeTags()))
        for i, v := range m.GetUserRoleScopeTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("userRoleScopeTags", cast)
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
func (m *AuditActor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AuditActor) SetApplicationDisplayName(value *string)() {
    m.applicationDisplayName = value
}
func (m *AuditActor) SetApplicationId(value *string)() {
    m.applicationId = value
}
func (m *AuditActor) SetIpAddress(value *string)() {
    m.ipAddress = value
}
func (m *AuditActor) SetRemoteTenantId(value *string)() {
    m.remoteTenantId = value
}
func (m *AuditActor) SetRemoteUserId(value *string)() {
    m.remoteUserId = value
}
func (m *AuditActor) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
func (m *AuditActor) SetType_escaped(value *string)() {
    m.type_escaped = value
}
func (m *AuditActor) SetUserId(value *string)() {
    m.userId = value
}
func (m *AuditActor) SetUserPermissions(value []string)() {
    m.userPermissions = value
}
func (m *AuditActor) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *AuditActor) SetUserRoleScopeTags(value []RoleScopeTagInfo)() {
    m.userRoleScopeTags = value
}
