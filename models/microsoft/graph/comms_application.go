package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CommsApplication 
type CommsApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    calls []Call;
    // 
    onlineMeetings []OnlineMeeting;
}
// NewCommsApplication instantiates a new CommsApplication and sets the default values.
func NewCommsApplication()(*CommsApplication) {
    m := &CommsApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommsApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCalls gets the calls property value. 
func (m *CommsApplication) GetCalls()([]Call) {
    if m == nil {
        return nil
    } else {
        return m.calls
    }
}
// GetOnlineMeetings gets the onlineMeetings property value. 
func (m *CommsApplication) GetOnlineMeetings()([]OnlineMeeting) {
    if m == nil {
        return nil
    } else {
        return m.onlineMeetings
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *CommsApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCalls() != nil {
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
    if m.GetOnlineMeetings() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommsApplication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCalls sets the calls property value. 
func (m *CommsApplication) SetCalls(value []Call)() {
    if m != nil {
        m.calls = value
    }
}
// SetOnlineMeetings sets the onlineMeetings property value. 
func (m *CommsApplication) SetOnlineMeetings(value []OnlineMeeting)() {
    if m != nil {
        m.onlineMeetings = value
    }
}
