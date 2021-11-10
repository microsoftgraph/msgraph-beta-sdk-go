package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SimulationEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Count of occurence of the simulation event in an attack simulation and training campaign.
    count *int32;
    // Name of the simulation event in an attack simulation and training campaign.
    eventName *string;
}
// Instantiates a new simulationEvent and sets the default values.
func NewSimulationEvent()(*SimulationEvent) {
    m := &SimulationEvent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationEvent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the count property value. Count of occurence of the simulation event in an attack simulation and training campaign.
func (m *SimulationEvent) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// Gets the eventName property value. Name of the simulation event in an attack simulation and training campaign.
func (m *SimulationEvent) GetEventName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.eventName
    }
}
// The deserialization information for the current model
func (m *SimulationEvent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["eventName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    return res
}
func (m *SimulationEvent) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SimulationEvent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventName", m.GetEventName())
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
func (m *SimulationEvent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the count property value. Count of occurence of the simulation event in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the count property.
func (m *SimulationEvent) SetCount(value *int32)() {
    m.count = value
}
// Sets the eventName property value. Name of the simulation event in an attack simulation and training campaign.
// Parameters:
//  - value : Value to set for the eventName property.
func (m *SimulationEvent) SetEventName(value *string)() {
    m.eventName = value
}
