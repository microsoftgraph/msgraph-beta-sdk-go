package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateWindow 
type UpdateWindow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // End of a time window during which agents can receive updates
    updateWindowEndTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // Start of a time window during which agents can receive updates
    updateWindowStartTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
}
// NewUpdateWindow instantiates a new updateWindow and sets the default values.
func NewUpdateWindow()(*UpdateWindow) {
    m := &UpdateWindow{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindow) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetUpdateWindowEndTime gets the updateWindowEndTime property value. End of a time window during which agents can receive updates
func (m *UpdateWindow) GetUpdateWindowEndTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.updateWindowEndTime
    }
}
// GetUpdateWindowStartTime gets the updateWindowStartTime property value. Start of a time window during which agents can receive updates
func (m *UpdateWindow) GetUpdateWindowStartTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.updateWindowStartTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateWindow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["updateWindowEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowEndTime(val)
        }
        return nil
    }
    res["updateWindowStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowStartTime(val)
        }
        return nil
    }
    return res
}
func (m *UpdateWindow) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateWindow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeOnlyValue("updateWindowEndTime", m.GetUpdateWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("updateWindowStartTime", m.GetUpdateWindowStartTime())
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
func (m *UpdateWindow) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUpdateWindowEndTime sets the updateWindowEndTime property value. End of a time window during which agents can receive updates
func (m *UpdateWindow) SetUpdateWindowEndTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.updateWindowEndTime = value
    }
}
// SetUpdateWindowStartTime sets the updateWindowStartTime property value. Start of a time window during which agents can receive updates
func (m *UpdateWindow) SetUpdateWindowStartTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.updateWindowStartTime = value
    }
}
