package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerTaskCreation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Information about the publication process that created this task. null value indicates that the task was not created by a publication process.
    teamsPublicationInfo *PlannerTeamsPublicationInfo;
}
// Instantiates a new plannerTaskCreation and sets the default values.
func NewPlannerTaskCreation()(*PlannerTaskCreation) {
    m := &PlannerTaskCreation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskCreation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the teamsPublicationInfo property value. Information about the publication process that created this task. null value indicates that the task was not created by a publication process.
func (m *PlannerTaskCreation) GetTeamsPublicationInfo()(*PlannerTeamsPublicationInfo) {
    if m == nil {
        return nil
    } else {
        return m.teamsPublicationInfo
    }
}
// The deserialization information for the current model
func (m *PlannerTaskCreation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["teamsPublicationInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerTeamsPublicationInfo() })
        if err != nil {
            return err
        }
        m.SetTeamsPublicationInfo(val.(*PlannerTeamsPublicationInfo))
        return nil
    }
    return res
}
func (m *PlannerTaskCreation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PlannerTaskCreation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the teamsPublicationInfo property value. Information about the publication process that created this task. null value indicates that the task was not created by a publication process.
// Parameters:
//  - value : Value to set for the teamsPublicationInfo property.
func (m *PlannerTaskCreation) SetTeamsPublicationInfo(value *PlannerTeamsPublicationInfo)() {
    m.teamsPublicationInfo = value
}
