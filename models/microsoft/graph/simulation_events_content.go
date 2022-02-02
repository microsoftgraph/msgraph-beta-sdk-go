package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SimulationEventsContent 
type SimulationEventsContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
    compromisedRate *float64;
    // List of simulation events in an attack simulation and training campaign.
    events []SimulationEvent;
}
// NewSimulationEventsContent instantiates a new simulationEventsContent and sets the default values.
func NewSimulationEventsContent()(*SimulationEventsContent) {
    m := &SimulationEventsContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationEventsContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompromisedRate gets the compromisedRate property value. Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
func (m *SimulationEventsContent) GetCompromisedRate()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.compromisedRate
    }
}
// GetEvents gets the events property value. List of simulation events in an attack simulation and training campaign.
func (m *SimulationEventsContent) GetEvents()([]SimulationEvent) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationEventsContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["compromisedRate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompromisedRate(val)
        }
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSimulationEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimulationEvent, len(val))
            for i, v := range val {
                res[i] = *(v.(*SimulationEvent))
            }
            m.SetEvents(res)
        }
        return nil
    }
    return res
}
func (m *SimulationEventsContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SimulationEventsContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("compromisedRate", m.GetCompromisedRate())
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationEventsContent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompromisedRate sets the compromisedRate property value. Actual percentage of users who fell for the simulated attack in an attack simulation and training campaign.
func (m *SimulationEventsContent) SetCompromisedRate(value *float64)() {
    if m != nil {
        m.compromisedRate = value
    }
}
// SetEvents sets the events property value. List of simulation events in an attack simulation and training campaign.
func (m *SimulationEventsContent) SetEvents(value []SimulationEvent)() {
    if m != nil {
        m.events = value
    }
}
