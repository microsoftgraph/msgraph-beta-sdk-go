package changedeploymentstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeDeploymentStatusPostRequestBody provides operations to call the changeDeploymentStatus method.
type ChangeDeploymentStatusPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The managementTemplateStepId property
    managementTemplateStepId *string
    // The status property
    status *string
    // The tenantId property
    tenantId *string
}
// NewChangeDeploymentStatusPostRequestBody instantiates a new changeDeploymentStatusPostRequestBody and sets the default values.
func NewChangeDeploymentStatusPostRequestBody()(*ChangeDeploymentStatusPostRequestBody) {
    m := &ChangeDeploymentStatusPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChangeDeploymentStatusPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeDeploymentStatusPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeDeploymentStatusPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeDeploymentStatusPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementTemplateStepId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagementTemplateStepId)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetManagementTemplateStepId gets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ChangeDeploymentStatusPostRequestBody) GetManagementTemplateStepId()(*string) {
    return m.managementTemplateStepId
}
// GetStatus gets the status property value. The status property
func (m *ChangeDeploymentStatusPostRequestBody) GetStatus()(*string) {
    return m.status
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ChangeDeploymentStatusPostRequestBody) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ChangeDeploymentStatusPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ChangeDeploymentStatusPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagementTemplateStepId sets the managementTemplateStepId property value. The managementTemplateStepId property
func (m *ChangeDeploymentStatusPostRequestBody) SetManagementTemplateStepId(value *string)() {
    m.managementTemplateStepId = value
}
// SetStatus sets the status property value. The status property
func (m *ChangeDeploymentStatusPostRequestBody) SetStatus(value *string)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ChangeDeploymentStatusPostRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
