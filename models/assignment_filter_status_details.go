package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterStatusDetails represent status details for device and payload and all associated applied filters.
type AssignmentFilterStatusDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Device properties used for filter evaluation during device check-in time.
    deviceProperties []KeyValuePairable
    // Evaluation result summaries for each filter associated to device and payload
    evalutionSummaries []AssignmentFilterEvaluationSummaryable
    // Unique identifier for the device object.
    managedDeviceId *string
    // Unique identifier for payload object.
    payloadId *string
    // Unique identifier for UserId object. Can be null
    userId *string
}
// NewAssignmentFilterStatusDetails instantiates a new assignmentFilterStatusDetails and sets the default values.
func NewAssignmentFilterStatusDetails()(*AssignmentFilterStatusDetails) {
    m := &AssignmentFilterStatusDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignmentFilterStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterStatusDetails(), nil
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
func (m *AssignmentFilterStatusDetails) GetDeviceProperties()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.deviceProperties
    }
}
// GetEvalutionSummaries gets the evalutionSummaries property value. Evaluation result summaries for each filter associated to device and payload
func (m *AssignmentFilterStatusDetails) GetEvalutionSummaries()([]AssignmentFilterEvaluationSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.evalutionSummaries
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterStatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetDeviceProperties(res)
        }
        return nil
    }
    res["evalutionSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterEvaluationSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterEvaluationSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(AssignmentFilterEvaluationSummaryable)
            }
            m.SetEvalutionSummaries(res)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *AssignmentFilterStatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceProperties()))
        for i, v := range m.GetDeviceProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvalutionSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvalutionSummaries()))
        for i, v := range m.GetEvalutionSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *AssignmentFilterStatusDetails) SetDeviceProperties(value []KeyValuePairable)() {
    if m != nil {
        m.deviceProperties = value
    }
}
// SetEvalutionSummaries sets the evalutionSummaries property value. Evaluation result summaries for each filter associated to device and payload
func (m *AssignmentFilterStatusDetails) SetEvalutionSummaries(value []AssignmentFilterEvaluationSummaryable)() {
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
