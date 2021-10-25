package updateaudiencebyid

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UpdateAudienceByIdRequestBody struct {
    addExclusions []string;
    additionalData map[string]interface{};
    addMembers []string;
    memberEntityType *string;
    removeExclusions []string;
    removeMembers []string;
}
func NewUpdateAudienceByIdRequestBody()(*UpdateAudienceByIdRequestBody) {
    m := &UpdateAudienceByIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UpdateAudienceByIdRequestBody) GetAddExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addExclusions
    }
}
func (m *UpdateAudienceByIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UpdateAudienceByIdRequestBody) GetAddMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.addMembers
    }
}
func (m *UpdateAudienceByIdRequestBody) GetMemberEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberEntityType
    }
}
func (m *UpdateAudienceByIdRequestBody) GetRemoveExclusions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeExclusions
    }
}
func (m *UpdateAudienceByIdRequestBody) GetRemoveMembers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeMembers
    }
}
func (m *UpdateAudienceByIdRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["addExclusions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
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
            res[i] = v.(string)
        }
        m.SetRemoveMembers(res)
        return nil
    }
    return res
}
func (m *UpdateAudienceByIdRequestBody) IsNil()(bool) {
    return m == nil
}
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
func (m *UpdateAudienceByIdRequestBody) SetAddExclusions(value []string)() {
    m.addExclusions = value
}
func (m *UpdateAudienceByIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UpdateAudienceByIdRequestBody) SetAddMembers(value []string)() {
    m.addMembers = value
}
func (m *UpdateAudienceByIdRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
func (m *UpdateAudienceByIdRequestBody) SetRemoveExclusions(value []string)() {
    m.removeExclusions = value
}
func (m *UpdateAudienceByIdRequestBody) SetRemoveMembers(value []string)() {
    m.removeMembers = value
}
