package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BookingWorkTimeSlot provides operations to manage the collection of bookingBusiness entities.
type BookingWorkTimeSlot struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The time of the day when work stops. For example, 17:00:00.0000000.
    end *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The time of the day when work starts. For example, 08:00:00.0000000.
    start *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
}
// NewBookingWorkTimeSlot instantiates a new bookingWorkTimeSlot and sets the default values.
func NewBookingWorkTimeSlot()(*BookingWorkTimeSlot) {
    m := &BookingWorkTimeSlot{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBookingWorkTimeSlotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingWorkTimeSlotFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewBookingWorkTimeSlot(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkTimeSlot) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnd gets the end property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) GetEnd()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.end
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingWorkTimeSlot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["end"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["start"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    return res
}
// GetStart gets the start property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) GetStart()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.start
    }
}
func (m *BookingWorkTimeSlot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BookingWorkTimeSlot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeOnlyValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("start", m.GetStart())
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
func (m *BookingWorkTimeSlot) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnd sets the end property value. The time of the day when work stops. For example, 17:00:00.0000000.
func (m *BookingWorkTimeSlot) SetEnd(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.end = value
    }
}
// SetStart sets the start property value. The time of the day when work starts. For example, 08:00:00.0000000.
func (m *BookingWorkTimeSlot) SetStart(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.start = value
    }
}
