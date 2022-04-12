package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcAuditActor 
type CloudPcAuditActor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the application.
    applicationDisplayName *string
    // Azure AD application ID.
    applicationId *string
    // IP address.
    ipAddress *string
    // The delegated partner tenant ID.
    remoteTenantId *string
    // The delegated partner user ID.
    remoteUserId *string
    // Service Principal Name (SPN).
    servicePrincipalName *string
    // The actor type. Possible values include ItPro, Application, Partner and Unknown.
    type_escaped *CloudPcAuditActorType
    // Azure AD user ID.
    userId *string
    // List of user permissions and application permissions when the audit event was performed.
    userPermissions []string
    // User Principal Name (UPN).
    userPrincipalName *string
    // List of role scope tags.
    userRoleScopeTags []CloudPcUserRoleScopeTagInfoable
}
// NewCloudPcAuditActor instantiates a new cloudPcAuditActor and sets the default values.
func NewCloudPcAuditActor()(*CloudPcAuditActor) {
    m := &CloudPcAuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcAuditActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcAuditActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcAuditActor(), nil
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
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditActor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationDisplayName(val)
        }
        return nil
    }
    res["applicationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["remoteTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteTenantId(val)
        }
        return nil
    }
    res["remoteUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteUserId(val)
        }
        return nil
    }
    res["servicePrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcAuditActorType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val.(*CloudPcAuditActorType))
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
    res["userPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userRoleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcUserRoleScopeTagInfoable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcUserRoleScopeTagInfoable)
            }
            m.SetUserRoleScopeTags(res)
        }
        return nil
    }
    return res
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
func (m *CloudPcAuditActor) GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfoable) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
// Serialize serializes information the current object
func (m *CloudPcAuditActor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetType_escaped()).String()
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserRoleScopeTags()))
        for i, v := range m.GetUserRoleScopeTags() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *CloudPcAuditActor) SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfoable)() {
    if m != nil {
        m.userRoleScopeTags = value
    }
}
