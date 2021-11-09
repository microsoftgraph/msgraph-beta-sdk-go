package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new cloudPcAuditActor and sets the default values.
func NewCloudPcAuditActor()(*CloudPcAuditActor) {
    m := &CloudPcAuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditActor) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applicationDisplayName property value. Name of the application.
func (m *CloudPcAuditActor) GetApplicationDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationDisplayName
    }
}
// Gets the applicationId property value. Azure AD application ID.
func (m *CloudPcAuditActor) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// Gets the ipAddress property value. IP address.
func (m *CloudPcAuditActor) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// Gets the remoteTenantId property value. The delegated partner tenant ID.
func (m *CloudPcAuditActor) GetRemoteTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteTenantId
    }
}
// Gets the remoteUserId property value. The delegated partner user ID.
func (m *CloudPcAuditActor) GetRemoteUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remoteUserId
    }
}
// Gets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *CloudPcAuditActor) GetServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipalName
    }
}
// Gets the type_escaped property value. The actor type. Possible values include ItPro, Application, Partner and Unknown.
func (m *CloudPcAuditActor) GetType_escaped()(*CloudPcAuditActorType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the userId property value. Azure AD user ID.
func (m *CloudPcAuditActor) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// Gets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
func (m *CloudPcAuditActor) GetUserPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.userPermissions
    }
}
// Gets the userPrincipalName property value. User Principal Name (UPN).
func (m *CloudPcAuditActor) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the userRoleScopeTags property value. List of role scope tags.
func (m *CloudPcAuditActor) GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfo) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActorType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcAuditActorType)
            m.SetType_escaped(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CloudPcAuditActor) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applicationDisplayName property value. Name of the application.
// Parameters:
//  - value : Value to set for the applicationDisplayName property.
func (m *CloudPcAuditActor) SetApplicationDisplayName(value *string)() {
    m.applicationDisplayName = value
}
// Sets the applicationId property value. Azure AD application ID.
// Parameters:
//  - value : Value to set for the applicationId property.
func (m *CloudPcAuditActor) SetApplicationId(value *string)() {
    m.applicationId = value
}
// Sets the ipAddress property value. IP address.
// Parameters:
//  - value : Value to set for the ipAddress property.
func (m *CloudPcAuditActor) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// Sets the remoteTenantId property value. The delegated partner tenant ID.
// Parameters:
//  - value : Value to set for the remoteTenantId property.
func (m *CloudPcAuditActor) SetRemoteTenantId(value *string)() {
    m.remoteTenantId = value
}
// Sets the remoteUserId property value. The delegated partner user ID.
// Parameters:
//  - value : Value to set for the remoteUserId property.
func (m *CloudPcAuditActor) SetRemoteUserId(value *string)() {
    m.remoteUserId = value
}
// Sets the servicePrincipalName property value. Service Principal Name (SPN).
// Parameters:
//  - value : Value to set for the servicePrincipalName property.
func (m *CloudPcAuditActor) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// Sets the type_escaped property value. The actor type. Possible values include ItPro, Application, Partner and Unknown.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CloudPcAuditActor) SetType_escaped(value *CloudPcAuditActorType)() {
    m.type_escaped = value
}
// Sets the userId property value. Azure AD user ID.
// Parameters:
//  - value : Value to set for the userId property.
func (m *CloudPcAuditActor) SetUserId(value *string)() {
    m.userId = value
}
// Sets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
// Parameters:
//  - value : Value to set for the userPermissions property.
func (m *CloudPcAuditActor) SetUserPermissions(value []string)() {
    m.userPermissions = value
}
// Sets the userPrincipalName property value. User Principal Name (UPN).
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *CloudPcAuditActor) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the userRoleScopeTags property value. List of role scope tags.
// Parameters:
//  - value : Value to set for the userRoleScopeTags property.
func (m *CloudPcAuditActor) SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfo)() {
    m.userRoleScopeTags = value
}
