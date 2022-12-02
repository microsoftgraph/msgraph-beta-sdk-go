package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody provides operations to call the changeDeploymentStatus method.
type TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managementActionId property
    managementActionId *string
    // The managementTemplateId property
    managementTemplateId *string
    // The managementTemplateVersion property
    managementTemplateVersion *int32
    // The status property
    status *string
    // The tenantGroupId property
    tenantGroupId *string
    // The tenantId property
    tenantId *string
}
// NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody instantiates a new TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody and sets the default values.
func NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody()(*TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) {
    m := &TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementActionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateVersion(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["tenantGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantGroupId(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetManagementActionId gets the managementActionId property value. The managementActionId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetManagementActionId()(*string) {
    return m.managementActionId
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetManagementTemplateId()(*string) {
    return m.managementTemplateId
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetManagementTemplateVersion()(*int32) {
    return m.managementTemplateVersion
}
// GetStatus gets the status property value. The status property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetStatus()(*string) {
    return m.status
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetTenantGroupId()(*string) {
    return m.tenantGroupId
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementActionId", m.GetManagementActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("managementTemplateVersion", m.GetManagementTemplateVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantGroupId", m.GetTenantGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
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
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagementActionId sets the managementActionId property value. The managementActionId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetManagementActionId(value *string)() {
    m.managementActionId = value
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetManagementTemplateVersion(value *int32)() {
    m.managementTemplateVersion = value
}
// SetStatus sets the status property value. The status property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetStatus(value *string)() {
    m.status = value
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *TenantRelationshipsManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
