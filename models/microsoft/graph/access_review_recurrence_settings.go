package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessReviewRecurrenceSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The duration in days for recurrence.
    durationInDays *int32;
    // The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
    recurrenceCount *int32;
    // How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it is never, then there is no explicit end of the recurrence series. If it is endBy, then the recurrence ends at a certain date. If it is occurrences, then the series ends after recurrenceCount instances of the review have completed.
    recurrenceEndType *string;
    // The recurrence interval. Possible vaules: onetime, weekly, monthly, quarterly, halfyearly or annual.
    recurrenceType *string;
}
// Instantiates a new accessReviewRecurrenceSettings and sets the default values.
func NewAccessReviewRecurrenceSettings()(*AccessReviewRecurrenceSettings) {
    m := &AccessReviewRecurrenceSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewRecurrenceSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the durationInDays property value. The duration in days for recurrence.
func (m *AccessReviewRecurrenceSettings) GetDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInDays
    }
}
// Gets the recurrenceCount property value. The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceCount
    }
}
// Gets the recurrenceEndType property value. How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it is never, then there is no explicit end of the recurrence series. If it is endBy, then the recurrence ends at a certain date. If it is occurrences, then the series ends after recurrenceCount instances of the review have completed.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceEndType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceEndType
    }
}
// Gets the recurrenceType property value. The recurrence interval. Possible vaules: onetime, weekly, monthly, quarterly, halfyearly or annual.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceType
    }
}
// The deserialization information for the current model
func (m *AccessReviewRecurrenceSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["durationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    res["recurrenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceCount(val)
        }
        return nil
    }
    res["recurrenceEndType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceEndType(val)
        }
        return nil
    }
    res["recurrenceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceType(val)
        }
        return nil
    }
    return res
}
func (m *AccessReviewRecurrenceSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessReviewRecurrenceSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("recurrenceCount", m.GetRecurrenceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceEndType", m.GetRecurrenceEndType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("recurrenceType", m.GetRecurrenceType())
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
func (m *AccessReviewRecurrenceSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the durationInDays property value. The duration in days for recurrence.
// Parameters:
//  - value : Value to set for the durationInDays property.
func (m *AccessReviewRecurrenceSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// Sets the recurrenceCount property value. The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
// Parameters:
//  - value : Value to set for the recurrenceCount property.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceCount(value *int32)() {
    m.recurrenceCount = value
}
// Sets the recurrenceEndType property value. How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it is never, then there is no explicit end of the recurrence series. If it is endBy, then the recurrence ends at a certain date. If it is occurrences, then the series ends after recurrenceCount instances of the review have completed.
// Parameters:
//  - value : Value to set for the recurrenceEndType property.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceEndType(value *string)() {
    m.recurrenceEndType = value
}
// Sets the recurrenceType property value. The recurrence interval. Possible vaules: onetime, weekly, monthly, quarterly, halfyearly or annual.
// Parameters:
//  - value : Value to set for the recurrenceType property.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceType(value *string)() {
    m.recurrenceType = value
}
