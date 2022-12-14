package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemCalendarGetSchedulePostRequestBody provides operations to call the getSchedule method.
type ItemCalendarGetSchedulePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The AvailabilityViewInterval property
    availabilityViewInterval *int32
    // The EndTime property
    endTime ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable
    // The Schedules property
    schedules []string
    // The StartTime property
    startTime ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable
}
// NewItemCalendarGetSchedulePostRequestBody instantiates a new ItemCalendarGetSchedulePostRequestBody and sets the default values.
func NewItemCalendarGetSchedulePostRequestBody()(*ItemCalendarGetSchedulePostRequestBody) {
    m := &ItemCalendarGetSchedulePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemCalendarGetSchedulePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarGetSchedulePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGetSchedulePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarGetSchedulePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAvailabilityViewInterval gets the availabilityViewInterval property value. The AvailabilityViewInterval property
func (m *ItemCalendarGetSchedulePostRequestBody) GetAvailabilityViewInterval()(*int32) {
    return m.availabilityViewInterval
}
// GetEndTime gets the endTime property value. The EndTime property
func (m *ItemCalendarGetSchedulePostRequestBody) GetEndTime()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable) {
    return m.endTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarGetSchedulePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityViewInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityViewInterval(val)
        }
        return nil
    }
    res["endTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable))
        }
        return nil
    }
    res["schedules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSchedules(res)
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable))
        }
        return nil
    }
    return res
}
// GetSchedules gets the schedules property value. The Schedules property
func (m *ItemCalendarGetSchedulePostRequestBody) GetSchedules()([]string) {
    return m.schedules
}
// GetStartTime gets the startTime property value. The StartTime property
func (m *ItemCalendarGetSchedulePostRequestBody) GetStartTime()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable) {
    return m.startTime
}
// Serialize serializes information the current object
func (m *ItemCalendarGetSchedulePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("availabilityViewInterval", m.GetAvailabilityViewInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("endTime", m.GetEndTime())
        if err != nil {
            return err
        }
    }
    if m.GetSchedules() != nil {
        err := writer.WriteCollectionOfStringValues("schedules", m.GetSchedules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startTime", m.GetStartTime())
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
func (m *ItemCalendarGetSchedulePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAvailabilityViewInterval sets the availabilityViewInterval property value. The AvailabilityViewInterval property
func (m *ItemCalendarGetSchedulePostRequestBody) SetAvailabilityViewInterval(value *int32)() {
    m.availabilityViewInterval = value
}
// SetEndTime sets the endTime property value. The EndTime property
func (m *ItemCalendarGetSchedulePostRequestBody) SetEndTime(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable)() {
    m.endTime = value
}
// SetSchedules sets the schedules property value. The Schedules property
func (m *ItemCalendarGetSchedulePostRequestBody) SetSchedules(value []string)() {
    m.schedules = value
}
// SetStartTime sets the startTime property value. The StartTime property
func (m *ItemCalendarGetSchedulePostRequestBody) SetStartTime(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DateTimeTimeZoneable)() {
    m.startTime = value
}
