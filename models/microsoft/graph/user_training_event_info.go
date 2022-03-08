package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserTrainingEventInfo provides operations to manage the attackSimulation property of the microsoft.graph.security entity.
type UserTrainingEventInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name of the training.
    displayName *string;
    // Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, notCompleted, unknownFutureValue.
    latestTrainingStatus *TrainingStatus;
    // Event details of the training when it was assigned to the user.
    trainingAssignedProperties UserTrainingContentEventInfoable;
    // Event details of the training when it was completed by the user.
    trainingCompletedProperties UserTrainingContentEventInfoable;
    // Event details of the training when it was updated/in-progress by the user.
    trainingUpdatedProperties UserTrainingContentEventInfoable;
}
// NewUserTrainingEventInfo instantiates a new userTrainingEventInfo and sets the default values.
func NewUserTrainingEventInfo()(*UserTrainingEventInfo) {
    m := &UserTrainingEventInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserTrainingEventInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserTrainingEventInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserTrainingEventInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingEventInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name of the training.
func (m *UserTrainingEventInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTrainingEventInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["latestTrainingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatestTrainingStatus(val.(*TrainingStatus))
        }
        return nil
    }
    res["trainingAssignedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingContentEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingAssignedProperties(val.(UserTrainingContentEventInfoable))
        }
        return nil
    }
    res["trainingCompletedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingContentEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingCompletedProperties(val.(UserTrainingContentEventInfoable))
        }
        return nil
    }
    res["trainingUpdatedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingContentEventInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingUpdatedProperties(val.(UserTrainingContentEventInfoable))
        }
        return nil
    }
    return res
}
// GetLatestTrainingStatus gets the latestTrainingStatus property value. Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, notCompleted, unknownFutureValue.
func (m *UserTrainingEventInfo) GetLatestTrainingStatus()(*TrainingStatus) {
    if m == nil {
        return nil
    } else {
        return m.latestTrainingStatus
    }
}
// GetTrainingAssignedProperties gets the trainingAssignedProperties property value. Event details of the training when it was assigned to the user.
func (m *UserTrainingEventInfo) GetTrainingAssignedProperties()(UserTrainingContentEventInfoable) {
    if m == nil {
        return nil
    } else {
        return m.trainingAssignedProperties
    }
}
// GetTrainingCompletedProperties gets the trainingCompletedProperties property value. Event details of the training when it was completed by the user.
func (m *UserTrainingEventInfo) GetTrainingCompletedProperties()(UserTrainingContentEventInfoable) {
    if m == nil {
        return nil
    } else {
        return m.trainingCompletedProperties
    }
}
// GetTrainingUpdatedProperties gets the trainingUpdatedProperties property value. Event details of the training when it was updated/in-progress by the user.
func (m *UserTrainingEventInfo) GetTrainingUpdatedProperties()(UserTrainingContentEventInfoable) {
    if m == nil {
        return nil
    } else {
        return m.trainingUpdatedProperties
    }
}
func (m *UserTrainingEventInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserTrainingEventInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetLatestTrainingStatus() != nil {
        cast := (*m.GetLatestTrainingStatus()).String()
        err := writer.WriteStringValue("latestTrainingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingAssignedProperties", m.GetTrainingAssignedProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingCompletedProperties", m.GetTrainingCompletedProperties())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingUpdatedProperties", m.GetTrainingUpdatedProperties())
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
func (m *UserTrainingEventInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the training.
func (m *UserTrainingEventInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLatestTrainingStatus sets the latestTrainingStatus property value. Latest status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, notCompleted, unknownFutureValue.
func (m *UserTrainingEventInfo) SetLatestTrainingStatus(value *TrainingStatus)() {
    if m != nil {
        m.latestTrainingStatus = value
    }
}
// SetTrainingAssignedProperties sets the trainingAssignedProperties property value. Event details of the training when it was assigned to the user.
func (m *UserTrainingEventInfo) SetTrainingAssignedProperties(value UserTrainingContentEventInfoable)() {
    if m != nil {
        m.trainingAssignedProperties = value
    }
}
// SetTrainingCompletedProperties sets the trainingCompletedProperties property value. Event details of the training when it was completed by the user.
func (m *UserTrainingEventInfo) SetTrainingCompletedProperties(value UserTrainingContentEventInfoable)() {
    if m != nil {
        m.trainingCompletedProperties = value
    }
}
// SetTrainingUpdatedProperties sets the trainingUpdatedProperties property value. Event details of the training when it was updated/in-progress by the user.
func (m *UserTrainingEventInfo) SetTrainingUpdatedProperties(value UserTrainingContentEventInfoable)() {
    if m != nil {
        m.trainingUpdatedProperties = value
    }
}
