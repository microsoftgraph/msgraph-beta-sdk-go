package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SimulationReportOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
    recommendedActions []RecommendedAction;
    // Number of valid users in the attack simulation and training campaign.
    resolvedTargetsCount *int32;
    // Summary of simulation events in the attack simulation and training campaign.
    simulationEventsContent *SimulationEventsContent;
    // Summary of assigned trainings in the attack simulation and training campaign.
    trainingEventsContent *TrainingEventsContent;
}
// Instantiates a new simulationReportOverview and sets the default values.
func NewSimulationReportOverview()(*SimulationReportOverview) {
    m := &SimulationReportOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReportOverview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the recommendedActions property value. List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
func (m *SimulationReportOverview) GetRecommendedActions()([]RecommendedAction) {
    if m == nil {
        return nil
    } else {
        return m.recommendedActions
    }
}
// Gets the resolvedTargetsCount property value. Number of valid users in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetResolvedTargetsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.resolvedTargetsCount
    }
}
// Gets the simulationEventsContent property value. Summary of simulation events in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetSimulationEventsContent()(*SimulationEventsContent) {
    if m == nil {
        return nil
    } else {
        return m.simulationEventsContent
    }
}
// Gets the trainingEventsContent property value. Summary of assigned trainings in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetTrainingEventsContent()(*TrainingEventsContent) {
    if m == nil {
        return nil
    } else {
        return m.trainingEventsContent
    }
}
// The deserialization information for the current model
func (m *SimulationReportOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["recommendedActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecommendedAction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendedAction, len(val))
            for i, v := range val {
                res[i] = *(v.(*RecommendedAction))
            }
            m.SetRecommendedActions(res)
        }
        return nil
    }
    res["resolvedTargetsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResolvedTargetsCount(val)
        }
        return nil
    }
    res["simulationEventsContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulationEventsContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationEventsContent(val.(*SimulationEventsContent))
        }
        return nil
    }
    res["trainingEventsContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTrainingEventsContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingEventsContent(val.(*TrainingEventsContent))
        }
        return nil
    }
    return res
}
func (m *SimulationReportOverview) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SimulationReportOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecommendedActions()))
        for i, v := range m.GetRecommendedActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("recommendedActions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("resolvedTargetsCount", m.GetResolvedTargetsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("simulationEventsContent", m.GetSimulationEventsContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingEventsContent", m.GetTrainingEventsContent())
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
func (m *SimulationReportOverview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the recommendedActions property value. List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
// Parameters:
//  - value : Value to set for the recommendedActions property.
func (m *SimulationReportOverview) SetRecommendedActions(value []RecommendedAction)() {
    m.recommendedActions = value
}
// Sets the resolvedTargetsCount property value. Number of valid users in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the resolvedTargetsCount property.
func (m *SimulationReportOverview) SetResolvedTargetsCount(value *int32)() {
    m.resolvedTargetsCount = value
}
// Sets the simulationEventsContent property value. Summary of simulation events in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the simulationEventsContent property.
func (m *SimulationReportOverview) SetSimulationEventsContent(value *SimulationEventsContent)() {
    m.simulationEventsContent = value
}
// Sets the trainingEventsContent property value. Summary of assigned trainings in the attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the trainingEventsContent property.
func (m *SimulationReportOverview) SetTrainingEventsContent(value *TrainingEventsContent)() {
    m.trainingEventsContent = value
}
