package changedeploymentstatus

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChangeDeploymentStatusRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    managementActionId *string;
    // 
    managementTemplateId *string;
    // 
    status *string;
    // 
    tenantGroupId *string;
    // 
    tenantId *string;
}
// Instantiates a new changeDeploymentStatusRequestBody and sets the default values.
func NewChangeDeploymentStatusRequestBody()(*ChangeDeploymentStatusRequestBody) {
    m := &ChangeDeploymentStatusRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChangeDeploymentStatusRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the managementActionId property value. 
func (m *ChangeDeploymentStatusRequestBody) GetManagementActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementActionId
    }
}
// Gets the managementTemplateId property value. 
func (m *ChangeDeploymentStatusRequestBody) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// Gets the status property value. 
func (m *ChangeDeploymentStatusRequestBody) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the tenantGroupId property value. 
func (m *ChangeDeploymentStatusRequestBody) GetTenantGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroupId
    }
}
// Gets the tenantId property value. 
func (m *ChangeDeploymentStatusRequestBody) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// The deserialization information for the current model
func (m *ChangeDeploymentStatusRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementActionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
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
    res["tenantGroupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantGroupId(val)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChangeDeploymentStatusRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ChangeDeploymentStatusRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the managementActionId property value. 
// Parameters:
//  - value : Value to set for the managementActionId property.
func (m *ChangeDeploymentStatusRequestBody) SetManagementActionId(value *string)() {
    m.managementActionId = value
}
// Sets the managementTemplateId property value. 
// Parameters:
//  - value : Value to set for the managementTemplateId property.
func (m *ChangeDeploymentStatusRequestBody) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *ChangeDeploymentStatusRequestBody) SetStatus(value *string)() {
    m.status = value
}
// Sets the tenantGroupId property value. 
// Parameters:
//  - value : Value to set for the tenantGroupId property.
func (m *ChangeDeploymentStatusRequestBody) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// Sets the tenantId property value. 
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *ChangeDeploymentStatusRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
