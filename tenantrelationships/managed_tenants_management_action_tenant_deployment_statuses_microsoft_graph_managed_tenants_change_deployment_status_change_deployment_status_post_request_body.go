package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody 
type ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
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
// NewManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody instantiates a new ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody and sets the default values.
func NewManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody()(*ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) {
    m := &ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetManagementActionId()(*string) {
    return m.managementActionId
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetManagementTemplateId()(*string) {
    return m.managementTemplateId
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetManagementTemplateVersion()(*int32) {
    return m.managementTemplateVersion
}
// GetStatus gets the status property value. The status property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetStatus()(*string) {
    return m.status
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetTenantGroupId()(*string) {
    return m.tenantGroupId
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagementActionId sets the managementActionId property value. The managementActionId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetManagementActionId(value *string)() {
    m.managementActionId = value
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetManagementTemplateVersion(value *int32)() {
    m.managementTemplateVersion = value
}
// SetStatus sets the status property value. The status property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetStatus(value *string)() {
    m.status = value
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
