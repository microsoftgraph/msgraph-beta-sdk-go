package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CloudPcAuditActor struct {
    additionalData map[string]interface{};
    applicationDisplayName *string;
    applicationId *string;
    ipAddress *string;
    remoteTenantId *string;
    remoteUserId *string;
    servicePrincipalName *string;
    type_escaped *CloudPcAuditActorType;
    userId *string;
    userPermissions []string;
    userPrincipalName *string;
    userRoleScopeTags []CloudPcUserRoleScopeTagInfo;
}
func NewCloudPcAuditActor()(*CloudPcAuditActor) {
    m := &CloudPcAuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CloudPcAuditActor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CloudPcAuditActor) GetApplicationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationDisplayName
    }
}
func (m *CloudPcAuditActor) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
func (m *CloudPcAuditActor) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
func (m *CloudPcAuditActor) GetRemoteTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteTenantId
    }
}
func (m *CloudPcAuditActor) GetRemoteUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteUserId
    }
}
func (m *CloudPcAuditActor) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
func (m *CloudPcAuditActor) GetType_escaped()(*CloudPcAuditActorType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *CloudPcAuditActor) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *CloudPcAuditActor) GetUserPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userPermissions
    }
}
func (m *CloudPcAuditActor) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *CloudPcAuditActor) GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfo) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
func (m *CloudPcAuditActor) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseCloudPcAuditActorType)
        if err != nil {
            return err
        }
        cast := val.(CloudPcAuditActorType)
        m.SetType_escaped(&cast)
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcUserRoleScopeTagInfo() })
        if err != nil {
            return err
        }
        res := make([]CloudPcUserRoleScopeTagInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*CloudPcUserRoleScopeTagInfo))
        }
        m.SetUserRoleScopeTags(res)
        return nil
    }
    return res
}
func (m *CloudPcAuditActor) IsNil()(bool) {
    return m == nil
}
func (m *CloudPcAuditActor) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *CloudPcAuditActor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CloudPcAuditActor) SetApplicationDisplayName(value *string)() {
    m.applicationDisplayName = value
}
func (m *CloudPcAuditActor) SetApplicationId(value *string)() {
    m.applicationId = value
}
func (m *CloudPcAuditActor) SetIpAddress(value *string)() {
    m.ipAddress = value
}
func (m *CloudPcAuditActor) SetRemoteTenantId(value *string)() {
    m.remoteTenantId = value
}
func (m *CloudPcAuditActor) SetRemoteUserId(value *string)() {
    m.remoteUserId = value
}
func (m *CloudPcAuditActor) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
func (m *CloudPcAuditActor) SetType_escaped(value *CloudPcAuditActorType)() {
    m.type_escaped = value
}
func (m *CloudPcAuditActor) SetUserId(value *string)() {
    m.userId = value
}
func (m *CloudPcAuditActor) SetUserPermissions(value []string)() {
    m.userPermissions = value
}
func (m *CloudPcAuditActor) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
func (m *CloudPcAuditActor) SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfo)() {
    m.userRoleScopeTags = value
}
