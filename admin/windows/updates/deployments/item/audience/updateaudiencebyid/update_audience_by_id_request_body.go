package updateaudiencebyid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UpdateAudienceByIdRequestBody struct {
    // 
    addExclusions []string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    addMembers []string;
    // 
    memberEntityType *string;
    // 
    removeExclusions []string;
    // 
    removeMembers []string;
}
// Instantiates a new updateAudienceByIdRequestBody and sets the default values.
func NewUpdateAudienceByIdRequestBody()(*UpdateAudienceByIdRequestBody) {
    m := &UpdateAudienceByIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the addExclusions property value. 
func (m *UpdateAudienceByIdRequestBody) GetAddExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addExclusions
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudienceByIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the addMembers property value. 
func (m *UpdateAudienceByIdRequestBody) GetAddMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addMembers
    }
}
// Gets the memberEntityType property value. 
func (m *UpdateAudienceByIdRequestBody) GetMemberEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberEntityType
    }
}
// Gets the removeExclusions property value. 
func (m *UpdateAudienceByIdRequestBody) GetRemoveExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeExclusions
    }
}
// Gets the removeMembers property value. 
func (m *UpdateAudienceByIdRequestBody) GetRemoveMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeMembers
    }
}
// The deserialization information for the current model
func (m *UpdateAudienceByIdRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addExclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetAddExclusions(res)
        return nil
    }
    res["addMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetAddMembers(res)
        return nil
    }
    res["memberEntityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMemberEntityType(val)
        return nil
    }
    res["removeExclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRemoveExclusions(res)
        return nil
    }
    res["removeMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRemoveMembers(res)
        return nil
    }
    return res
}
func (m *UpdateAudienceByIdRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UpdateAudienceByIdRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("addExclusions", m.GetAddExclusions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("addMembers", m.GetAddMembers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("removeExclusions", m.GetRemoveExclusions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("removeMembers", m.GetRemoveMembers())
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
// Sets the addExclusions property value. 
// Parameters:
//  - value : Value to set for the addExclusions property.
func (m *UpdateAudienceByIdRequestBody) SetAddExclusions(value []string)() {
    m.addExclusions = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UpdateAudienceByIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the addMembers property value. 
// Parameters:
//  - value : Value to set for the addMembers property.
func (m *UpdateAudienceByIdRequestBody) SetAddMembers(value []string)() {
    m.addMembers = value
}
// Sets the memberEntityType property value. 
// Parameters:
//  - value : Value to set for the memberEntityType property.
func (m *UpdateAudienceByIdRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
// Sets the removeExclusions property value. 
// Parameters:
//  - value : Value to set for the removeExclusions property.
func (m *UpdateAudienceByIdRequestBody) SetRemoveExclusions(value []string)() {
    m.removeExclusions = value
}
// Sets the removeMembers property value. 
// Parameters:
//  - value : Value to set for the removeMembers property.
func (m *UpdateAudienceByIdRequestBody) SetRemoveMembers(value []string)() {
    m.removeMembers = value
}
