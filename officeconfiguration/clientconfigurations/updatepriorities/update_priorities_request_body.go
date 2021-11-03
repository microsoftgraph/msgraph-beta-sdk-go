package updatepriorities

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UpdatePrioritiesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    officeConfigurationPolicyIds []string;
    // 
    officeConfigurationPriorities []int32;
}
// Instantiates a new updatePrioritiesRequestBody and sets the default values.
func NewUpdatePrioritiesRequestBody()(*UpdatePrioritiesRequestBody) {
    m := &UpdatePrioritiesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdatePrioritiesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the officeConfigurationPolicyIds property value. 
func (m *UpdatePrioritiesRequestBody) GetOfficeConfigurationPolicyIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationPolicyIds
    }
}
// Gets the officeConfigurationPriorities property value. 
func (m *UpdatePrioritiesRequestBody) GetOfficeConfigurationPriorities()([]int32) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationPriorities
    }
}
// The deserialization information for the current model
func (m *UpdatePrioritiesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["officeConfigurationPolicyIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetOfficeConfigurationPolicyIds(res)
        return nil
    }
    res["officeConfigurationPriorities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("int32")
        if err != nil {
            return err
        }
        res := make([]int32, len(val))
        for i, v := range val {
            res[i] = *(v.(*int32))
        }
        m.SetOfficeConfigurationPriorities(res)
        return nil
    }
    return res
}
func (m *UpdatePrioritiesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UpdatePrioritiesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("officeConfigurationPolicyIds", m.GetOfficeConfigurationPolicyIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfInt32Values("officeConfigurationPriorities", m.GetOfficeConfigurationPriorities())
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
func (m *UpdatePrioritiesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the officeConfigurationPolicyIds property value. 
// Parameters:
//  - value : Value to set for the officeConfigurationPolicyIds property.
func (m *UpdatePrioritiesRequestBody) SetOfficeConfigurationPolicyIds(value []string)() {
    m.officeConfigurationPolicyIds = value
}
// Sets the officeConfigurationPriorities property value. 
// Parameters:
//  - value : Value to set for the officeConfigurationPriorities property.
func (m *UpdatePrioritiesRequestBody) SetOfficeConfigurationPriorities(value []int32)() {
    m.officeConfigurationPriorities = value
}
