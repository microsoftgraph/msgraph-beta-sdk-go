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
    // The OdataType property
    odataType *string
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
    odataTypeValue := "#microsoft.graph.assignmentFilterStatusDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAssignmentFilterStatusDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterStatusDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterStatusDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterStatusDetails) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceProperties gets the deviceProperties property value. Device properties used for filter evaluation during device check-in time.
func (m *AssignmentFilterStatusDetails) GetDeviceProperties()([]KeyValuePairable) {
    return m.deviceProperties
}
// GetEvalutionSummaries gets the evalutionSummaries property value. Evaluation result summaries for each filter associated to device and payload
func (m *AssignmentFilterStatusDetails) GetEvalutionSummaries()([]AssignmentFilterEvaluationSummaryable) {
    return m.evalutionSummaries
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterStatusDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetDeviceProperties)
    res["evalutionSummaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAssignmentFilterEvaluationSummaryFromDiscriminatorValue , m.SetEvalutionSummaries)
    res["managedDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["payloadId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPayloadId)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. Unique identifier for the device object.
func (m *AssignmentFilterStatusDetails) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterStatusDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetPayloadId gets the payloadId property value. Unique identifier for payload object.
func (m *AssignmentFilterStatusDetails) GetPayloadId()(*string) {
    return m.payloadId
}
// GetUserId gets the userId property value. Unique identifier for UserId object. Can be null
func (m *AssignmentFilterStatusDetails) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *AssignmentFilterStatusDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceProperties())
        err := writer.WriteCollectionOfObjectValues("deviceProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvalutionSummaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEvalutionSummaries())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetDeviceProperties sets the deviceProperties property value. Device properties used for filter evaluation during device check-in time.
func (m *AssignmentFilterStatusDetails) SetDeviceProperties(value []KeyValuePairable)() {
    m.deviceProperties = value
}
// SetEvalutionSummaries sets the evalutionSummaries property value. Evaluation result summaries for each filter associated to device and payload
func (m *AssignmentFilterStatusDetails) SetEvalutionSummaries(value []AssignmentFilterEvaluationSummaryable)() {
    m.evalutionSummaries = value
}
// SetManagedDeviceId sets the managedDeviceId property value. Unique identifier for the device object.
func (m *AssignmentFilterStatusDetails) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterStatusDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPayloadId sets the payloadId property value. Unique identifier for payload object.
func (m *AssignmentFilterStatusDetails) SetPayloadId(value *string)() {
    m.payloadId = value
}
// SetUserId sets the userId property value. Unique identifier for UserId object. Can be null
func (m *AssignmentFilterStatusDetails) SetUserId(value *string)() {
    m.userId = value
}
