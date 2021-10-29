package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SimulationReport struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Overview of an attack simulation and training campaign.
    overview *SimulationReportOverview;
    // Represents users of a tenant and their online actions in an attack simulation and training campaign.
    simulationUsers []UserSimulationDetails;
}
// Instantiates a new simulationReport and sets the default values.
func NewSimulationReport()(*SimulationReport) {
    m := &SimulationReport{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReport) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the overview property value. Overview of an attack simulation and training campaign.
func (m *SimulationReport) GetOverview()(*SimulationReportOverview) {
    if m == nil {
        return nil
    } else {
        return m.overview
    }
}
// Gets the simulationUsers property value. Represents users of a tenant and their online actions in an attack simulation and training campaign.
func (m *SimulationReport) GetSimulationUsers()([]UserSimulationDetails) {
    if m == nil {
        return nil
    } else {
        return m.simulationUsers
    }
}
// The deserialization information for the current model
func (m *SimulationReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["overview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulationReportOverview() })
        if err != nil {
            return err
        }
        m.SetOverview(val.(*SimulationReportOverview))
        return nil
    }
    res["simulationUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSimulationDetails() })
        if err != nil {
            return err
        }
        res := make([]UserSimulationDetails, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSimulationDetails))
        }
        m.SetSimulationUsers(res)
        return nil
    }
    return res
}
func (m *SimulationReport) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SimulationReport) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("overview", m.GetOverview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSimulationUsers()))
        for i, v := range m.GetSimulationUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("simulationUsers", cast)
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
func (m *SimulationReport) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the overview property value. Overview of an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the overview property.
func (m *SimulationReport) SetOverview(value *SimulationReportOverview)() {
    m.overview = value
}
// Sets the simulationUsers property value. Represents users of a tenant and their online actions in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the simulationUsers property.
func (m *SimulationReport) SetSimulationUsers(value []UserSimulationDetails)() {
    m.simulationUsers = value
}
