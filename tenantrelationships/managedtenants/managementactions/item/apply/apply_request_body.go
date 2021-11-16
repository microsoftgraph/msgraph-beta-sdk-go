package apply

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ApplyRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    excludeGroups []string;
    // 
    includeAllUsers *bool;
    // 
    includeGroups []string;
    // 
    managementTemplateId *string;
    // 
    tenantGroupId *string;
    // 
    tenantId *string;
}
// Instantiates a new applyRequestBody and sets the default values.
func NewApplyRequestBody()(*ApplyRequestBody) {
    m := &ApplyRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludeGroups property value. 
func (m *ApplyRequestBody) GetExcludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroups
    }
}
// Gets the includeAllUsers property value. 
func (m *ApplyRequestBody) GetIncludeAllUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.includeAllUsers
    }
}
// Gets the includeGroups property value. 
func (m *ApplyRequestBody) GetIncludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeGroups
    }
}
// Gets the managementTemplateId property value. 
func (m *ApplyRequestBody) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// Gets the tenantGroupId property value. 
func (m *ApplyRequestBody) GetTenantGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantGroupId
    }
}
// Gets the tenantId property value. 
func (m *ApplyRequestBody) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// The deserialization information for the current model
func (m *ApplyRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeGroups(res)
        }
        return nil
    }
    res["includeAllUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllUsers(val)
        }
        return nil
    }
    res["includeGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeGroups(res)
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
func (m *ApplyRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ApplyRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("includeAllUsers", m.GetIncludeAllUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
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
func (m *ApplyRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludeGroups property value. 
// Parameters:
//  - value : Value to set for the excludeGroups property.
func (m *ApplyRequestBody) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// Sets the includeAllUsers property value. 
// Parameters:
//  - value : Value to set for the includeAllUsers property.
func (m *ApplyRequestBody) SetIncludeAllUsers(value *bool)() {
    m.includeAllUsers = value
}
// Sets the includeGroups property value. 
// Parameters:
//  - value : Value to set for the includeGroups property.
func (m *ApplyRequestBody) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// Sets the managementTemplateId property value. 
// Parameters:
//  - value : Value to set for the managementTemplateId property.
func (m *ApplyRequestBody) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// Sets the tenantGroupId property value. 
// Parameters:
//  - value : Value to set for the tenantGroupId property.
func (m *ApplyRequestBody) SetTenantGroupId(value *string)() {
    m.tenantGroupId = value
}
// Sets the tenantId property value. 
// Parameters:
//  - value : Value to set for the tenantId property.
func (m *ApplyRequestBody) SetTenantId(value *string)() {
    m.tenantId = value
}
