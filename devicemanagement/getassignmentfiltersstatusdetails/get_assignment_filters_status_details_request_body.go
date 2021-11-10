package getassignmentfiltersstatusdetails

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GetAssignmentFiltersStatusDetailsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    assignmentFilterIds []string;
    // 
    managedDeviceId *string;
    // 
    payloadId *string;
    // 
    skip *int32;
    // 
    top *int32;
    // 
    userId *string;
}
// Instantiates a new getAssignmentFiltersStatusDetailsRequestBody and sets the default values.
func NewGetAssignmentFiltersStatusDetailsRequestBody()(*GetAssignmentFiltersStatusDetailsRequestBody) {
    m := &GetAssignmentFiltersStatusDetailsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the assignmentFilterIds property value. 
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetAssignmentFilterIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterIds
    }
}
// Gets the managedDeviceId property value. 
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the payloadId property value. 
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
// Gets the skip property value. 
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetSkip()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.skip
    }
}
// Gets the top property value. 
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetTop()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.top
    }
}
// Gets the userId property value. 
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *GetAssignmentFiltersStatusDetailsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignmentFilterIds(res)
        }
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["payloadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["skip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkip(val)
        }
        return nil
    }
    res["top"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTop(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *GetAssignmentFiltersStatusDetailsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the assignmentFilterIds property value. 
// Parameters:
//  - value : Value to set for the assignmentFilterIds property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetAssignmentFilterIds(value []string)() {
    m.assignmentFilterIds = value
}
// Sets the managedDeviceId property value. 
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the payloadId property value. 
// Parameters:
//  - value : Value to set for the payloadId property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetPayloadId(value *string)() {
    m.payloadId = value
}
// Sets the skip property value. 
// Parameters:
//  - value : Value to set for the skip property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetSkip(value *int32)() {
    m.skip = value
}
// Sets the top property value. 
// Parameters:
//  - value : Value to set for the top property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetTop(value *int32)() {
    m.top = value
}
// Sets the userId property value. 
// Parameters:
//  - value : Value to set for the userId property.
func (m *GetAssignmentFiltersStatusDetailsRequestBody) SetUserId(value *string)() {
    m.userId = value
}
