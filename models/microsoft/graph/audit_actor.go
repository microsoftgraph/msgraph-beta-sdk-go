package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuditActor 
type AuditActor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the Application.
    applicationDisplayName *string;
    // AAD Application Id.
    applicationId *string;
    // IPAddress.
    ipAddress *string;
    // Remote Tenant Id
    remoteTenantId *string;
    // Remote User Id
    remoteUserId *string;
    // Service Principal Name (SPN).
    servicePrincipalName *string;
    // Actor Type.
    type_escaped *string;
    // User Id.
    userId *string;
    // List of user permissions when the audit was performed.
    userPermissions []string;
    // User Principal Name (UPN).
    userPrincipalName *string;
    // List of user scope tags when the audit was performed.
    userRoleScopeTags []RoleScopeTagInfo;
}
// NewAuditActor instantiates a new auditActor and sets the default values.
func NewAuditActor()(*AuditActor) {
    m := &AuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditActor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationDisplayName gets the applicationDisplayName property value. Name of the Application.
func (m *AuditActor) GetApplicationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationDisplayName
    }
}
// GetApplicationId gets the applicationId property value. AAD Application Id.
func (m *AuditActor) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// GetIpAddress gets the ipAddress property value. IPAddress.
func (m *AuditActor) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetRemoteTenantId gets the remoteTenantId property value. Remote Tenant Id
func (m *AuditActor) GetRemoteTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteTenantId
    }
}
// GetRemoteUserId gets the remoteUserId property value. Remote User Id
func (m *AuditActor) GetRemoteUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteUserId
    }
}
// GetServicePrincipalName gets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *AuditActor) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
// GetType gets the type property value. Actor Type.
func (m *AuditActor) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUserId gets the userId property value. User Id.
func (m *AuditActor) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPermissions gets the userPermissions property value. List of user permissions when the audit was performed.
func (m *AuditActor) GetUserPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userPermissions
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name (UPN).
func (m *AuditActor) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserRoleScopeTags gets the userRoleScopeTags property value. List of user scope tags when the audit was performed.
func (m *AuditActor) GetUserRoleScopeTags()([]RoleScopeTagInfo) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditActor) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationDisplayName(val)
        }
        return nil
    }
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["remoteTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteTenantId(val)
        }
        return nil
    }
    res["remoteUserId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteUserId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
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
    res["userPermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserPermissions(res)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userRoleScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScopeTagInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*RoleScopeTagInfo))
            }
            m.SetUserRoleScopeTags(res)
        }
        return nil
    }
    return res
}
func (m *AuditActor) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        err := writer.WriteStringValue("type", m.GetType())
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
    if m.GetUserPermissions() != nil {
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
    if m.GetUserRoleScopeTags() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuditActor) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationDisplayName sets the applicationDisplayName property value. Name of the Application.
func (m *AuditActor) SetApplicationDisplayName(value *string)() {
    if m != nil {
        m.applicationDisplayName = value
    }
}
// SetApplicationId sets the applicationId property value. AAD Application Id.
func (m *AuditActor) SetApplicationId(value *string)() {
    if m != nil {
        m.applicationId = value
    }
}
// SetIpAddress sets the ipAddress property value. IPAddress.
func (m *AuditActor) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetRemoteTenantId sets the remoteTenantId property value. Remote Tenant Id
func (m *AuditActor) SetRemoteTenantId(value *string)() {
    if m != nil {
        m.remoteTenantId = value
    }
}
// SetRemoteUserId sets the remoteUserId property value. Remote User Id
func (m *AuditActor) SetRemoteUserId(value *string)() {
    if m != nil {
        m.remoteUserId = value
    }
}
// SetServicePrincipalName sets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *AuditActor) SetServicePrincipalName(value *string)() {
    if m != nil {
        m.servicePrincipalName = value
    }
}
// SetType sets the type property value. Actor Type.
func (m *AuditActor) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUserId sets the userId property value. User Id.
func (m *AuditActor) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPermissions sets the userPermissions property value. List of user permissions when the audit was performed.
func (m *AuditActor) SetUserPermissions(value []string)() {
    if m != nil {
        m.userPermissions = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name (UPN).
func (m *AuditActor) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserRoleScopeTags sets the userRoleScopeTags property value. List of user scope tags when the audit was performed.
func (m *AuditActor) SetUserRoleScopeTags(value []RoleScopeTagInfo)() {
    if m != nil {
        m.userRoleScopeTags = value
    }
}
