package changedeploymentstatus

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChangeDeploymentStatusRequestBody 
type ChangeDeploymentStatusRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    managementTemplateStepId *string;
    // 
    status *string;
    // 
    tenantId *string;
}
// NewChangeDeploymentStatusRequestBody instantiates a new changeDeploymentStatusRequestBody and sets the default values.
func NewChangeDeploymentStatusRequestBody()(*ChangeDeploymentStatusRequestBody) {
    m := &ChangeDeploymentStatusRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeDeploymentStatusRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetManagementTemplateStepId gets the managementTemplateStepId property value. 
func (m *ChangeDeploymentStatusRequestBody) GetManagementTemplateStepId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateStepId
    }
}
// GetStatus gets the status property value. 
func (m *ChangeDeploymentStatusRequestBody) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTenantId gets the tenantId property value. 
func (m *ChangeDeploymentStatusRequestBody) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChangeDeploymentStatusRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementTemplateStepId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateStepId(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *ChangeDeploymentStatusRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChangeDeploymentStatusRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ChangeDeploymentStatusRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementTemplateStepId sets the managementTemplateStepId property value. 
func (m *ChangeDeploymentStatusRequestBody) SetManagementTemplateStepId(value *string)() {
    if m != nil {
        m.managementTemplateStepId = value
    }
}
// SetStatus sets the status property value. 
func (m *ChangeDeploymentStatusRequestBody) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetTenantId sets the tenantId property value. 
func (m *ChangeDeploymentStatusRequestBody) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
