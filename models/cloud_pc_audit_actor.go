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
    // The OdataType property
    odataType *string
    // The delegated partner tenant ID.
    remoteTenantId *string
    // The delegated partner user ID.
    remoteUserId *string
    // Service Principal Name (SPN).
    servicePrincipalName *string
    // The type property
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
    odataTypeValue := "#microsoft.graph.cloudPcAuditActor";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcAuditActorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcAuditActorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcAuditActor(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcAuditActor) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetApplicationDisplayName gets the applicationDisplayName property value. Name of the application.
func (m *CloudPcAuditActor) GetApplicationDisplayName()(*string) {
    return m.applicationDisplayName
}
// GetApplicationId gets the applicationId property value. Azure AD application ID.
func (m *CloudPcAuditActor) GetApplicationId()(*string) {
    return m.applicationId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcAuditActor) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicationDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationDisplayName)
    res["applicationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationId)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["remoteTenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemoteTenantId)
    res["remoteUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRemoteUserId)
    res["servicePrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetServicePrincipalName)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcAuditActorType , m.SetType)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userPermissions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetUserPermissions)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["userRoleScopeTags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCloudPcUserRoleScopeTagInfoFromDiscriminatorValue , m.SetUserRoleScopeTags)
    return res
}
// GetIpAddress gets the ipAddress property value. IP address.
func (m *CloudPcAuditActor) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcAuditActor) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoteTenantId gets the remoteTenantId property value. The delegated partner tenant ID.
func (m *CloudPcAuditActor) GetRemoteTenantId()(*string) {
    return m.remoteTenantId
}
// GetRemoteUserId gets the remoteUserId property value. The delegated partner user ID.
func (m *CloudPcAuditActor) GetRemoteUserId()(*string) {
    return m.remoteUserId
}
// GetServicePrincipalName gets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *CloudPcAuditActor) GetServicePrincipalName()(*string) {
    return m.servicePrincipalName
}
// GetType gets the type property value. The type property
func (m *CloudPcAuditActor) GetType()(*CloudPcAuditActorType) {
    return m.type_escaped
}
// GetUserId gets the userId property value. Azure AD user ID.
func (m *CloudPcAuditActor) GetUserId()(*string) {
    return m.userId
}
// GetUserPermissions gets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
func (m *CloudPcAuditActor) GetUserPermissions()([]string) {
    return m.userPermissions
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name (UPN).
func (m *CloudPcAuditActor) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserRoleScopeTags gets the userRoleScopeTags property value. List of role scope tags.
func (m *CloudPcAuditActor) GetUserRoleScopeTags()([]CloudPcUserRoleScopeTagInfoable) {
    return m.userRoleScopeTags
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserRoleScopeTags())
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
    m.additionalData = value
}
// SetApplicationDisplayName sets the applicationDisplayName property value. Name of the application.
func (m *CloudPcAuditActor) SetApplicationDisplayName(value *string)() {
    m.applicationDisplayName = value
}
// SetApplicationId sets the applicationId property value. Azure AD application ID.
func (m *CloudPcAuditActor) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetIpAddress sets the ipAddress property value. IP address.
func (m *CloudPcAuditActor) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcAuditActor) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoteTenantId sets the remoteTenantId property value. The delegated partner tenant ID.
func (m *CloudPcAuditActor) SetRemoteTenantId(value *string)() {
    m.remoteTenantId = value
}
// SetRemoteUserId sets the remoteUserId property value. The delegated partner user ID.
func (m *CloudPcAuditActor) SetRemoteUserId(value *string)() {
    m.remoteUserId = value
}
// SetServicePrincipalName sets the servicePrincipalName property value. Service Principal Name (SPN).
func (m *CloudPcAuditActor) SetServicePrincipalName(value *string)() {
    m.servicePrincipalName = value
}
// SetType sets the type property value. The type property
func (m *CloudPcAuditActor) SetType(value *CloudPcAuditActorType)() {
    m.type_escaped = value
}
// SetUserId sets the userId property value. Azure AD user ID.
func (m *CloudPcAuditActor) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPermissions sets the userPermissions property value. List of user permissions and application permissions when the audit event was performed.
func (m *CloudPcAuditActor) SetUserPermissions(value []string)() {
    m.userPermissions = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name (UPN).
func (m *CloudPcAuditActor) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserRoleScopeTags sets the userRoleScopeTags property value. List of role scope tags.
func (m *CloudPcAuditActor) SetUserRoleScopeTags(value []CloudPcUserRoleScopeTagInfoable)() {
    m.userRoleScopeTags = value
}
