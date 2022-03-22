package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlannerTaskCreation 
type PlannerTaskCreation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Information about the publication process that created this task. null value indicates that the task was not created by a publication process.
    teamsPublicationInfo PlannerTeamsPublicationInfoable;
}
// NewPlannerTaskCreation instantiates a new plannerTaskCreation and sets the default values.
func NewPlannerTaskCreation()(*PlannerTaskCreation) {
    m := &PlannerTaskCreation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerTaskCreationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskCreationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlannerTaskCreation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskCreation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskCreation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["teamsPublicationInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTeamsPublicationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsPublicationInfo(val.(PlannerTeamsPublicationInfoable))
        }
        return nil
    }
    return res
}
// GetTeamsPublicationInfo gets the teamsPublicationInfo property value. Information about the publication process that created this task. null value indicates that the task was not created by a publication process.
func (m *PlannerTaskCreation) GetTeamsPublicationInfo()(PlannerTeamsPublicationInfoable) {
    if m == nil {
        return nil
    } else {
        return m.teamsPublicationInfo
    }
}
// Serialize serializes information the current object
func (m *PlannerTaskCreation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("teamsPublicationInfo", m.GetTeamsPublicationInfo())
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
func (m *PlannerTaskCreation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTeamsPublicationInfo sets the teamsPublicationInfo property value. Information about the publication process that created this task. null value indicates that the task was not created by a publication process.
func (m *PlannerTaskCreation) SetTeamsPublicationInfo(value PlannerTeamsPublicationInfoable)() {
    if m != nil {
        m.teamsPublicationInfo = value
    }
}
