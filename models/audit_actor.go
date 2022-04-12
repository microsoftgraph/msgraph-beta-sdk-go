package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditActor a class containing the properties for Audit Actor.
type AuditActor struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the Application.
    applicationDisplayName *string
    // AAD Application Id.
    applicationId *string
    // IPAddress.
    ipAddress *string
    // Remote Tenant Id
    remoteTenantId *string
    // Remote User Id
    remoteUserId *string
    // Service Principal Name (SPN).
    servicePrincipalName *string
    // Actor Type.
    type_escaped *string
    // User Id.
    userId *string
    // List of user permissions when the audit was performed.
    userPermissions []string
    // User Principal Name (UPN).
    userPrincipalName *string
    // List of user scope tags when the audit was performed.
    userRoleScopeTags []RoleScopeTagInfoable
}
// NewAuditActor instantiates a new auditActor and sets the default values.
func NewAuditActor()(*AuditActor) {
    m := &AuditActor{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuditActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuditActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditActor(), nil
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
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditActor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
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
        val, err := n.GetCollectionOfObjectValues(CreateRoleScopeTagInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RoleScopeTagInfoable, len(val))
            for i, v := range val {
                res[i] = v.(RoleScopeTagInfoable)
            }
            m.SetUserRoleScopeTags(res)
        }
        return nil
    }
    return res
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
func (m *AuditActor) GetUserRoleScopeTags()([]RoleScopeTagInfoable) {
    if m == nil {
        return nil
    } else {
        return m.userRoleScopeTags
    }
}
// Serialize serializes information the current object
func (m *AuditActor) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("type", m.GetType_escaped())
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
func (m *AuditActor) SetUserRoleScopeTags(value []RoleScopeTagInfoable)() {
    if m != nil {
        m.userRoleScopeTags = value
    }
}
