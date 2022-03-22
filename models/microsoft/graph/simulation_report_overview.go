package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SimulationReportOverview 
type SimulationReportOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
    recommendedActions []RecommendedActionable;
    // Number of valid users in the attack simulation and training campaign.
    resolvedTargetsCount *int32;
    // Summary of simulation events in the attack simulation and training campaign.
    simulationEventsContent SimulationEventsContentable;
    // Summary of assigned trainings in the attack simulation and training campaign.
    trainingEventsContent TrainingEventsContentable;
}
// NewSimulationReportOverview instantiates a new simulationReportOverview and sets the default values.
func NewSimulationReportOverview()(*SimulationReportOverview) {
    m := &SimulationReportOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSimulationReportOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationReportOverviewFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSimulationReportOverview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReportOverview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationReportOverview) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["recommendedActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecommendedActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RecommendedActionable, len(val))
            for i, v := range val {
                res[i] = v.(RecommendedActionable)
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
        val, err := n.GetObjectValue(CreateSimulationEventsContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationEventsContent(val.(SimulationEventsContentable))
        }
        return nil
    }
    res["trainingEventsContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingEventsContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingEventsContent(val.(TrainingEventsContentable))
        }
        return nil
    }
    return res
}
// GetRecommendedActions gets the recommendedActions property value. List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
func (m *SimulationReportOverview) GetRecommendedActions()([]RecommendedActionable) {
    if m == nil {
        return nil
    } else {
        return m.recommendedActions
    }
}
// GetResolvedTargetsCount gets the resolvedTargetsCount property value. Number of valid users in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetResolvedTargetsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.resolvedTargetsCount
    }
}
// GetSimulationEventsContent gets the simulationEventsContent property value. Summary of simulation events in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetSimulationEventsContent()(SimulationEventsContentable) {
    if m == nil {
        return nil
    } else {
        return m.simulationEventsContent
    }
}
// GetTrainingEventsContent gets the trainingEventsContent property value. Summary of assigned trainings in the attack simulation and training campaign.
func (m *SimulationReportOverview) GetTrainingEventsContent()(TrainingEventsContentable) {
    if m == nil {
        return nil
    } else {
        return m.trainingEventsContent
    }
}
// Serialize serializes information the current object
func (m *SimulationReportOverview) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetRecommendedActions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecommendedActions()))
        for i, v := range m.GetRecommendedActions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReportOverview) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetRecommendedActions sets the recommendedActions property value. List of recommended actions for a tenant to improve its security posture based on the attack simulation and training campaign attack type.
func (m *SimulationReportOverview) SetRecommendedActions(value []RecommendedActionable)() {
    if m != nil {
        m.recommendedActions = value
    }
}
// SetResolvedTargetsCount sets the resolvedTargetsCount property value. Number of valid users in the attack simulation and training campaign.
func (m *SimulationReportOverview) SetResolvedTargetsCount(value *int32)() {
    if m != nil {
        m.resolvedTargetsCount = value
    }
}
// SetSimulationEventsContent sets the simulationEventsContent property value. Summary of simulation events in the attack simulation and training campaign.
func (m *SimulationReportOverview) SetSimulationEventsContent(value SimulationEventsContentable)() {
    if m != nil {
        m.simulationEventsContent = value
    }
}
// SetTrainingEventsContent sets the trainingEventsContent property value. Summary of assigned trainings in the attack simulation and training campaign.
func (m *SimulationReportOverview) SetTrainingEventsContent(value TrainingEventsContentable)() {
    if m != nil {
        m.trainingEventsContent = value
    }
}
