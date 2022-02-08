package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcAuditActor 
type CloudPcAuditActor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the application.
    applicationDisplayName *string;
    // Azure AD application ID.
    applicationId *string;
    // IP address.
    ipAddress *string;
    // The delegated partner tenant ID.
    remoteTenantId *string;
    // The delegated partner user ID.
    remoteUserId *string;
    // Service Principal Name (SPN).
    servicePrincipalName *string;
    // The actor type. Possible values include ItPro, Application, Partner and Unknown.
    type_escaped *CloudPcAuditActorType;
    // Azure AD user ID.
    userId *string;
    // List of user permissions and application permissions when the audit event was performed.
    userPermissions []string;
    // User Principal Name (UPN).
    userPrincipalName *string;
    // List of role scope tags.
    userRoleScopeTags []CloudPcUserRoleScopeTagInfo;
}
// NewCloudPcAuditActor instantiates a new cloudPcAuditActor and sets the default values.
func NewCloudPcAuditActor()(*CloudPcAuditActor) {
    m := &CloudPcAuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditActor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationDisplayName gets the applicationDisplayName property value. Name of the application.
func (m *CloudPcAuditActor) GetApplicationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationDisplayName
    }
}
// GetApplicationId gets the applicationId property value. Azure AD application ID.
func (m *CloudPcAuditActor) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// GetIpAddress gets the ipAddress property value. IP address.
func (m *CloudPcAuditActor) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetRemoteTenantId gets the remoteTenantId property value. The delegated partner tenant ID.
func (m *CloudPcAuditActor) GetRemoteTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteTenantId
    }
}
// GetRemoteUserId gets the remoteUserId property value. The delegated partner user ID.
func (m *CloudPcAuditActor) GetRemoteUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteUserId
    }
}
// GetServicePrincipalName gets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *CloudPcAuditActor) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
// GetType gets the type property value. The actor type. Possible values include ItPro, Application, Partner and Unknown.
func (m *CloudPcAuditActor) GetType()(*CloudPcAuditActorType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetUserId gets the userId property value. Azure AD user ID.
func (m *CloudPcAuditActor) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserPermissions gets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
func (m *CloudPcAuditActor) GetUserPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userPermissions
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name (UPN).
func (m *CloudPcAuditActor) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserRoleScopeTags gets the userRoleScopeTags property value. List of role scope tags.
func (m *CloudPcAuditActor) GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfo) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditActor) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseCloudPcAuditActorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*CloudPcAuditActorType))
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcUserRoleScopeTagInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserRoleScopeTagInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*CloudPcUserRoleScopeTagInfo))
            }
            m.SetUserRoleScopeTags(res)
        }
        return nil
    }
    return res
}
func (m *CloudPcAuditActor) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *CloudPcAuditActor) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationDisplayName sets the applicationDisplayName property value. Name of the application.
func (m *CloudPcAuditActor) SetApplicationDisplayName(value *string)() {
    if m != nil {
        m.applicationDisplayName = value
    }
}
// SetApplicationId sets the applicationId property value. Azure AD application ID.
func (m *CloudPcAuditActor) SetApplicationId(value *string)() {
    if m != nil {
        m.applicationId = value
    }
}
// SetIpAddress sets the ipAddress property value. IP address.
func (m *CloudPcAuditActor) SetIpAddress(value *string)() {
    if m != nil {
        m.ipAddress = value
    }
}
// SetRemoteTenantId sets the remoteTenantId property value. The delegated partner tenant ID.
func (m *CloudPcAuditActor) SetRemoteTenantId(value *string)() {
    if m != nil {
        m.remoteTenantId = value
    }
}
// SetRemoteUserId sets the remoteUserId property value. The delegated partner user ID.
func (m *CloudPcAuditActor) SetRemoteUserId(value *string)() {
    if m != nil {
        m.remoteUserId = value
    }
}
// SetServicePrincipalName sets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *CloudPcAuditActor) SetServicePrincipalName(value *string)() {
    if m != nil {
        m.servicePrincipalName = value
    }
}
// SetType sets the type property value. The actor type. Possible values include ItPro, Application, Partner and Unknown.
func (m *CloudPcAuditActor) SetType(value *CloudPcAuditActorType)() {
    if m != nil {
        m.type_escaped = value
    }
}
// SetUserId sets the userId property value. Azure AD user ID.
func (m *CloudPcAuditActor) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserPermissions sets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
func (m *CloudPcAuditActor) SetUserPermissions(value []string)() {
    if m != nil {
        m.userPermissions = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name (UPN).
func (m *CloudPcAuditActor) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserRoleScopeTags sets the userRoleScopeTags property value. List of role scope tags.
func (m *CloudPcAuditActor) SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfo)() {
    if m != nil {
        m.userRoleScopeTags = value
    }
}
