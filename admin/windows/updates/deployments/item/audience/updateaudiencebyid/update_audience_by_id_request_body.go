package updateaudiencebyid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateAudienceByIdRequestBody 
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
// NewUpdateAudienceByIdRequestBody instantiates a new updateAudienceByIdRequestBody and sets the default values.
func NewUpdateAudienceByIdRequestBody()(*UpdateAudienceByIdRequestBody) {
    m := &UpdateAudienceByIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAddExclusions gets the addExclusions property value. 
func (m *UpdateAudienceByIdRequestBody) GetAddExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addExclusions
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudienceByIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddMembers gets the addMembers property value. 
func (m *UpdateAudienceByIdRequestBody) GetAddMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addMembers
    }
}
// GetMemberEntityType gets the memberEntityType property value. 
func (m *UpdateAudienceByIdRequestBody) GetMemberEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberEntityType
    }
}
// GetRemoveExclusions gets the removeExclusions property value. 
func (m *UpdateAudienceByIdRequestBody) GetRemoveExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeExclusions
    }
}
// GetRemoveMembers gets the removeMembers property value. 
func (m *UpdateAudienceByIdRequestBody) GetRemoveMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeMembers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAudienceByIdRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addExclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAddExclusions(res)
        }
        return nil
    }
    res["addMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAddMembers(res)
        }
        return nil
    }
    res["memberEntityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    res["removeExclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveExclusions(res)
        }
        return nil
    }
    res["removeMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveMembers(res)
        }
        return nil
    }
    return res
}
func (m *UpdateAudienceByIdRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateAudienceByIdRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAddExclusions() != nil {
        err := writer.WriteCollectionOfStringValues("addExclusions", m.GetAddExclusions())
        if err != nil {
            return err
        }
    }
    if m.GetAddMembers() != nil {
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
    if m.GetRemoveExclusions() != nil {
        err := writer.WriteCollectionOfStringValues("removeExclusions", m.GetRemoveExclusions())
        if err != nil {
            return err
        }
    }
    if m.GetRemoveMembers() != nil {
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
// SetAddExclusions sets the addExclusions property value. 
func (m *UpdateAudienceByIdRequestBody) SetAddExclusions(value []string)() {
    if m != nil {
        m.addExclusions = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAudienceByIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddMembers sets the addMembers property value. 
func (m *UpdateAudienceByIdRequestBody) SetAddMembers(value []string)() {
    if m != nil {
        m.addMembers = value
    }
}
// SetMemberEntityType sets the memberEntityType property value. 
func (m *UpdateAudienceByIdRequestBody) SetMemberEntityType(value *string)() {
    if m != nil {
        m.memberEntityType = value
    }
}
// SetRemoveExclusions sets the removeExclusions property value. 
func (m *UpdateAudienceByIdRequestBody) SetRemoveExclusions(value []string)() {
    if m != nil {
        m.removeExclusions = value
    }
}
// SetRemoveMembers sets the removeMembers property value. 
func (m *UpdateAudienceByIdRequestBody) SetRemoveMembers(value []string)() {
    if m != nil {
        m.removeMembers = value
    }
}
