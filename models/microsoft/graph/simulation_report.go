package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SimulationReport 
type SimulationReport struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Overview of an attack simulation and training campaign.
    overview *SimulationReportOverview;
    // Represents users of a tenant and their online actions in an attack simulation and training campaign.
    simulationUsers []UserSimulationDetails;
}
// NewSimulationReport instantiates a new simulationReport and sets the default values.
func NewSimulationReport()(*SimulationReport) {
    m := &SimulationReport{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReport) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOverview gets the overview property value. Overview of an attack simulation and training campaign.
func (m *SimulationReport) GetOverview()(*SimulationReportOverview) {
    if m == nil {
        return nil
    } else {
        return m.overview
    }
}
// GetSimulationUsers gets the simulationUsers property value. Represents users of a tenant and their online actions in an attack simulation and training campaign.
func (m *SimulationReport) GetSimulationUsers()([]UserSimulationDetails) {
    if m == nil {
        return nil
    } else {
        return m.simulationUsers
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["overview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulationReportOverview() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverview(val.(*SimulationReportOverview))
        }
        return nil
    }
    res["simulationUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSimulationDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSimulationDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserSimulationDetails))
            }
            m.SetSimulationUsers(res)
        }
        return nil
    }
    return res
}
func (m *SimulationReport) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SimulationReport) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("overview", m.GetOverview())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationUsers() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReport) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOverview sets the overview property value. Overview of an attack simulation and training campaign.
func (m *SimulationReport) SetOverview(value *SimulationReportOverview)() {
    if m != nil {
        m.overview = value
    }
}
// SetSimulationUsers sets the simulationUsers property value. Represents users of a tenant and their online actions in an attack simulation and training campaign.
func (m *SimulationReport) SetSimulationUsers(value []UserSimulationDetails)() {
    if m != nil {
        m.simulationUsers = value
    }
}
