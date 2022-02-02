package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TrainingEventsContent 
type TrainingEventsContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of assigned trainings and their information in an attack simulation and training campaign.
    assignedTrainingsInfos []AssignedTrainingInfo;
    // Number of users who were assigned trainings in an attack simulation and training campaign.
    trainingsAssignedUserCount *int32;
}
// NewTrainingEventsContent instantiates a new trainingEventsContent and sets the default values.
func NewTrainingEventsContent()(*TrainingEventsContent) {
    m := &TrainingEventsContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TrainingEventsContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAssignedTrainingsInfos gets the assignedTrainingsInfos property value. List of assigned trainings and their information in an attack simulation and training campaign.
func (m *TrainingEventsContent) GetAssignedTrainingsInfos()([]AssignedTrainingInfo) {
    if m == nil {
        return nil
    } else {
        return m.assignedTrainingsInfos
    }
}
// GetTrainingsAssignedUserCount gets the trainingsAssignedUserCount property value. Number of users who were assigned trainings in an attack simulation and training campaign.
func (m *TrainingEventsContent) GetTrainingsAssignedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.trainingsAssignedUserCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TrainingEventsContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignedTrainingsInfos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedTrainingInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedTrainingInfo, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedTrainingInfo))
            }
            m.SetAssignedTrainingsInfos(res)
        }
        return nil
    }
    res["trainingsAssignedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingsAssignedUserCount(val)
        }
        return nil
    }
    return res
}
func (m *TrainingEventsContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TrainingEventsContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssignedTrainingsInfos() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedTrainingsInfos()))
        for i, v := range m.GetAssignedTrainingsInfos() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("assignedTrainingsInfos", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("trainingsAssignedUserCount", m.GetTrainingsAssignedUserCount())
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
func (m *TrainingEventsContent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAssignedTrainingsInfos sets the assignedTrainingsInfos property value. List of assigned trainings and their information in an attack simulation and training campaign.
func (m *TrainingEventsContent) SetAssignedTrainingsInfos(value []AssignedTrainingInfo)() {
    if m != nil {
        m.assignedTrainingsInfos = value
    }
}
// SetTrainingsAssignedUserCount sets the trainingsAssignedUserCount property value. Number of users who were assigned trainings in an attack simulation and training campaign.
func (m *TrainingEventsContent) SetTrainingsAssignedUserCount(value *int32)() {
    if m != nil {
        m.trainingsAssignedUserCount = value
    }
}
