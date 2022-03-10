package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserTrainingStatusInfo provides operations to call the getAttackSimulationTrainingUserCoverage method.
type UserTrainingStatusInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Date and time of assignment of the training to the user.
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date and time of completion of the training by the user.
    completionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Display name of the assigned training.
    displayName *string;
    // Status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, notCompleted, unknownFutureValue.
    trainingStatus *TrainingStatus;
}
// NewUserTrainingStatusInfo instantiates a new userTrainingStatusInfo and sets the default values.
func NewUserTrainingStatusInfo()(*UserTrainingStatusInfo) {
    m := &UserTrainingStatusInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserTrainingStatusInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserTrainingStatusInfoFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserTrainingStatusInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserTrainingStatusInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignedDateTime gets the assignedDateTime property value. Date and time of assignment of the training to the user.
func (m *UserTrainingStatusInfo) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.assignedDateTime
    }
}
// GetCompletionDateTime gets the completionDateTime property value. Date and time of completion of the training by the user.
func (m *UserTrainingStatusInfo) GetCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completionDateTime
    }
}
// GetDisplayName gets the displayName property value. Display name of the assigned training.
func (m *UserTrainingStatusInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserTrainingStatusInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedDateTime(val)
        }
        return nil
    }
    res["completionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionDateTime(val)
        }
        return nil
    }
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
    res["trainingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrainingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingStatus(val.(*TrainingStatus))
        }
        return nil
    }
    return res
}
// GetTrainingStatus gets the trainingStatus property value. Status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, notCompleted, unknownFutureValue.
func (m *UserTrainingStatusInfo) GetTrainingStatus()(*TrainingStatus) {
    if m == nil {
        return nil
    } else {
        return m.trainingStatus
    }
}
func (m *UserTrainingStatusInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserTrainingStatusInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("assignedDateTime", m.GetAssignedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("completionDateTime", m.GetCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetTrainingStatus() != nil {
        cast := (*m.GetTrainingStatus()).String()
        err := writer.WriteStringValue("trainingStatus", &cast)
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
func (m *UserTrainingStatusInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignedDateTime sets the assignedDateTime property value. Date and time of assignment of the training to the user.
func (m *UserTrainingStatusInfo) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.assignedDateTime = value
    }
}
// SetCompletionDateTime sets the completionDateTime property value. Date and time of completion of the training by the user.
func (m *UserTrainingStatusInfo) SetCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completionDateTime = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the assigned training.
func (m *UserTrainingStatusInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetTrainingStatus sets the trainingStatus property value. Status of the training assigned to the user. Possible values are: unknown, assigned, inProgress, completed, overdue, notCompleted, unknownFutureValue.
func (m *UserTrainingStatusInfo) SetTrainingStatus(value *TrainingStatus)() {
    if m != nil {
        m.trainingStatus = value
    }
}
