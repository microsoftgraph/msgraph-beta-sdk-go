package changedeploymentstatus

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeDeploymentStatusRequestBody provides operations to call the changeDeploymentStatus method.
type ChangeDeploymentStatusRequestBody struct {
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
// NewChangeDeploymentStatusRequestBody instantiates a new changeDeploymentStatusRequestBody and sets the default values.
func NewChangeDeploymentStatusRequestBody()(*ChangeDeploymentStatusRequestBody) {
    m := &ChangeDeploymentStatusRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChangeDeploymentStatusRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChangeDeploymentStatusRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChangeDeploymentStatusRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeDeploymentStatusRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeDeploymentStatusRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ChangeDeploymentStatusRequestBody) GetManagementActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementActionId
    }
}
// GetManagementTemplateId gets the managementTemplateId property value. The managementTemplateId property
func (m *ChangeDeploymentStatusRequestBody) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ChangeDeploymentStatusRequestBody) GetManagementTemplateVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateVersion
    }
}
// GetStatus gets the status property value. The status property
func (m *ChangeDeploymentStatusRequestBody) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTenantGroupId gets the tenantGroupId property value. The tenantGroupId property
func (m *ChangeDeploymentStatusRequestBody) GetTenantGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroupId
    }
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ChangeDeploymentStatusRequestBody) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *ChangeDeploymentStatusRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ChangeDeploymentStatusRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementActionId sets the managementActionId property value. The managementActionId property
func (m *ChangeDeploymentStatusRequestBody) SetManagementActionId(value *string)() {
    if m != nil {
        m.managementActionId = value
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The managementTemplateId property
func (m *ChangeDeploymentStatusRequestBody) SetManagementTemplateId(value *string)() {
    if m != nil {
        m.managementTemplateId = value
    }
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ChangeDeploymentStatusRequestBody) SetManagementTemplateVersion(value *int32)() {
    if m != nil {
        m.managementTemplateVersion = value
    }
}
// SetStatus sets the status property value. The status property
func (m *ChangeDeploymentStatusRequestBody) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTenantGroupId sets the tenantGroupId property value. The tenantGroupId property
func (m *ChangeDeploymentStatusRequestBody) SetTenantGroupId(value *string)() {
    if m != nil {
        m.tenantGroupId = value
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ChangeDeploymentStatusRequestBody) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
