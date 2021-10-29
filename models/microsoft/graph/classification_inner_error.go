package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ClassificationInnerError struct {
    // 
    activityId *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    clientRequestId *string;
    // 
    code *string;
    // 
    errorDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new classificationInnerError and sets the default values.
func NewClassificationInnerError()(*ClassificationInnerError) {
    m := &ClassificationInnerError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the activityId property value. 
func (m *ClassificationInnerError) GetActivityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityId
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationInnerError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the clientRequestId property value. 
func (m *ClassificationInnerError) GetClientRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientRequestId
    }
}
// Gets the code property value. 
func (m *ClassificationInnerError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the errorDateTime property value. 
func (m *ClassificationInnerError) GetErrorDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.errorDateTime
    }
}
// The deserialization information for the current model
func (m *ClassificationInnerError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityId(val)
        return nil
    }
    res["clientRequestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientRequestId(val)
        return nil
    }
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["errorDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetErrorDateTime(val)
        return nil
    }
    return res
}
func (m *ClassificationInnerError) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ClassificationInnerError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityId", m.GetActivityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("clientRequestId", m.GetClientRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("errorDateTime", m.GetErrorDateTime())
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
// Sets the activityId property value. 
// Parameters:
//  - value : Value to set for the activityId property.
func (m *ClassificationInnerError) SetActivityId(value *string)() {
    m.activityId = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ClassificationInnerError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the clientRequestId property value. 
// Parameters:
//  - value : Value to set for the clientRequestId property.
func (m *ClassificationInnerError) SetClientRequestId(value *string)() {
    m.clientRequestId = value
}
// Sets the code property value. 
// Parameters:
//  - value : Value to set for the code property.
func (m *ClassificationInnerError) SetCode(value *string)() {
    m.code = value
}
// Sets the errorDateTime property value. 
// Parameters:
//  - value : Value to set for the errorDateTime property.
func (m *ClassificationInnerError) SetErrorDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.errorDateTime = value
}
