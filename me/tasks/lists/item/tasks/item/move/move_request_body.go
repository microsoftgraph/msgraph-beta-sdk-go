package move

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MoveRequestBody 
type MoveRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    destinationTaskListId *string;
}
// NewMoveRequestBody instantiates a new moveRequestBody and sets the default values.
func NewMoveRequestBody()(*MoveRequestBody) {
    m := &MoveRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MoveRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDestinationTaskListId gets the destinationTaskListId property value. 
func (m *MoveRequestBody) GetDestinationTaskListId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.destinationTaskListId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MoveRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["destinationTaskListId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationTaskListId(val)
        }
        return nil
    }
    return res
}
func (m *MoveRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MoveRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationTaskListId", m.GetDestinationTaskListId())
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
func (m *MoveRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDestinationTaskListId sets the destinationTaskListId property value. 
func (m *MoveRequestBody) SetDestinationTaskListId(value *string)() {
    if m != nil {
        m.destinationTaskListId = value
    }
}
