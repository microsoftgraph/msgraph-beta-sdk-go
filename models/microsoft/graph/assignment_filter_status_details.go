package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterStatusDetails 
type AssignmentFilterStatusDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Device properties used for filter evaluation during device check-in time.
    deviceProperties []KeyValuePair;
    // Evaluation result summaries for each filter associated to device and payload
    evalutionSummaries []AssignmentFilterEvaluationSummary;
    // Unique identifier for the device object.
    managedDeviceId *string;
    // Unique identifier for payload object.
    payloadId *string;
    // Unique identifier for UserId object. Can be null
    userId *string;
}
// NewAssignmentFilterStatusDetails instantiates a new assignmentFilterStatusDetails and sets the default values.
func NewAssignmentFilterStatusDetails()(*AssignmentFilterStatusDetails) {
    m := &AssignmentFilterStatusDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterStatusDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceProperties gets the deviceProperties property value. Device properties used for filter evaluation during device check-in time.
func (m *AssignmentFilterStatusDetails) GetDeviceProperties()([]KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.deviceProperties
    }
}
// GetEvalutionSummaries gets the evalutionSummaries property value. Evaluation result summaries for each filter associated to device and payload
func (m *AssignmentFilterStatusDetails) GetEvalutionSummaries()([]AssignmentFilterEvaluationSummary) {
    if m == nil {
        return nil
    } else {
        return m.evalutionSummaries
    }
}
// GetManagedDeviceId gets the managedDeviceId property value. Unique identifier for the device object.
func (m *AssignmentFilterStatusDetails) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// GetPayloadId gets the payloadId property value. Unique identifier for payload object.
func (m *AssignmentFilterStatusDetails) GetPayloadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadId
    }
}
// GetUserId gets the userId property value. Unique identifier for UserId object. Can be null
func (m *AssignmentFilterStatusDetails) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterStatusDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*KeyValuePair))
            }
            m.SetDeviceProperties(res)
        }
        return nil
    }
    res["evalutionSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignmentFilterEvaluationSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterEvaluationSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignmentFilterEvaluationSummary))
            }
            m.SetEvalutionSummaries(res)
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
func (m *AssignmentFilterStatusDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AssignmentFilterStatusDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDeviceProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceProperties()))
        for i, v := range m.GetDeviceProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("deviceProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvalutionSummaries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvalutionSummaries()))
        for i, v := range m.GetEvalutionSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("evalutionSummaries", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterStatusDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceProperties sets the deviceProperties property value. Device properties used for filter evaluation during device check-in time.
func (m *AssignmentFilterStatusDetails) SetDeviceProperties(value []KeyValuePair)() {
    if m != nil {
        m.deviceProperties = value
    }
}
// SetEvalutionSummaries sets the evalutionSummaries property value. Evaluation result summaries for each filter associated to device and payload
func (m *AssignmentFilterStatusDetails) SetEvalutionSummaries(value []AssignmentFilterEvaluationSummary)() {
    if m != nil {
        m.evalutionSummaries = value
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Unique identifier for the device object.
func (m *AssignmentFilterStatusDetails) SetManagedDeviceId(value *string)() {
    if m != nil {
        m.managedDeviceId = value
    }
}
// SetPayloadId sets the payloadId property value. Unique identifier for payload object.
func (m *AssignmentFilterStatusDetails) SetPayloadId(value *string)() {
    if m != nil {
        m.payloadId = value
    }
}
// SetUserId sets the userId property value. Unique identifier for UserId object. Can be null
func (m *AssignmentFilterStatusDetails) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
