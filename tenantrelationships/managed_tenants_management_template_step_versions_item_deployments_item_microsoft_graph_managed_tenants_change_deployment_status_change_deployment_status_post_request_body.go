package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody 
type ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The managementTemplateStepId property
    managementTemplateStepId *string
    // The status property
    status *string
    // The tenantId property
    tenantId *string
}
// NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody instantiates a new ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody()(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) {
    m := &ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementTemplateStepId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepId(val)
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
// GetManagementTemplateStepId gets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetManagementTemplateStepId()(*string) {
    return m.managementTemplateStepId
}
// GetStatus gets the status property value. The status property
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetStatus()(*string) {
    return m.status
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementTemplateStepId", m.GetManagementTemplateStepId())
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
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetManagementTemplateStepId sets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetManagementTemplateStepId(value *string)() {
    m.managementTemplateStepId = value
}
// SetStatus sets the status property value. The status property
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetStatus(value *string)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagedTenantsManagementTemplateStepVersionsItemDeploymentsItemMicrosoftGraphManagedTenantsChangeDeploymentStatusChangeDeploymentStatusPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
