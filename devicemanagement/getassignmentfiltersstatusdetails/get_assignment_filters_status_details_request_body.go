package getassignmentfiltersstatusdetails

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GetAssignmentFiltersStatusDetailsRequestBody struct {
    additionalData map[string]interface{};
    assignmentFilterIds []string;
    managedDeviceId *string;
    payloadId *string;
    skip *int32;
    top *int32;
    userId *string;
}
func NewGetAssignmentFiltersStatusDetailsRequestBody()(*GetAssignmentFiltersStatusDetailsRequestBody) {
    m := &GetAssignmentFiltersStatusDetailsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetAssignmentFilterIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterIds
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignmentFilterIds(res)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayloadId(val)
        return nil
    }
    res["skip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSkip(val)
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTop(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("assignmentFilterIds", m.GetAssignmentFilterIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("skip", m.GetSkip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("top", m.GetTop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetAssignmentFilterIds(value []string)() {
    m.assignmentFilterIds = value
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetPayloadId(value *string)() {
    m.payloadId = value
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetTop(value *int32)() {
    m.top = value
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetUserId(value *string)() {
    m.userId = value
}
