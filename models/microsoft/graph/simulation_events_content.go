package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SimulationEventsContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
    compromisedRate *float64;
    // List of simulation events in an attack simulation and training campaign.
    events []SimulationEvent;
}
// Instantiates a new simulationEventsContent and sets the default values.
func NewSimulationEventsContent()(*SimulationEventsContent) {
    m := &SimulationEventsContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationEventsContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the compromisedRate property value. Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
func (m *SimulationEventsContent) GetCompromisedRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.compromisedRate
    }
}
// Gets the events property value. List of simulation events in an attack simulation and training campaign.
func (m *SimulationEventsContent) GetEvents()([]SimulationEvent) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// The deserialization information for the current model
func (m *SimulationEventsContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compromisedRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCompromisedRate(val)
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulationEvent() })
        if err != nil {
            return err
        }
        res := make([]SimulationEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*SimulationEvent))
        }
        m.SetEvents(res)
        return nil
    }
    return res
}
func (m *SimulationEventsContent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SimulationEventsContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("compromisedRate", m.GetCompromisedRate())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("events", cast)
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
func (m *SimulationEventsContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the compromisedRate property value. Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the compromisedRate property.
func (m *SimulationEventsContent) SetCompromisedRate(value *float64)() {
    m.compromisedRate = value
}
// Sets the events property value. List of simulation events in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the events property.
func (m *SimulationEventsContent) SetEvents(value []SimulationEvent)() {
    m.events = value
}
