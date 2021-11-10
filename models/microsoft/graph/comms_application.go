package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CommsApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    calls []Call;
    // 
    onlineMeetings []OnlineMeeting;
}
// Instantiates a new CommsApplication and sets the default values.
func NewCommsApplication()(*CommsApplication) {
    m := &CommsApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommsApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the calls property value. 
func (m *CommsApplication) GetCalls()([]Call) {
    if m == nil {
        return nil
    } else {
        return m.calls
    }
}
// Gets the onlineMeetings property value. 
func (m *CommsApplication) GetOnlineMeetings()([]OnlineMeeting) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetings
    }
}
// The deserialization information for the current model
func (m *CommsApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCall() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Call, len(val))
            for i, v := range val {
                res[i] = *(v.(*Call))
            }
            m.SetCalls(res)
        }
        return nil
    }
    res["onlineMeetings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnlineMeeting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeeting, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnlineMeeting))
            }
            m.SetOnlineMeetings(res)
        }
        return nil
    }
    return res
}
func (m *CommsApplication) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CommsApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalls()))
        for i, v := range m.GetCalls() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("calls", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnlineMeetings()))
        for i, v := range m.GetOnlineMeetings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
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
func (m *CommsApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the calls property value. 
// Parameters:
//  - value : Value to set for the calls property.
func (m *CommsApplication) SetCalls(value []Call)() {
    m.calls = value
}
// Sets the onlineMeetings property value. 
// Parameters:
//  - value : Value to set for the onlineMeetings property.
func (m *CommsApplication) SetOnlineMeetings(value []OnlineMeeting)() {
    m.onlineMeetings = value
}
