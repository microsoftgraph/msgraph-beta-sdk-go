package findmeetingtimes

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// FindMeetingTimesRequestBody provides operations to call the findMeetingTimes method.
type FindMeetingTimesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The attendees property
    attendees []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable
    // The isOrganizerOptional property
    isOrganizerOptional *bool
    // The locationConstraint property
    locationConstraint ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable
    // The maxCandidates property
    maxCandidates *int32
    // The meetingDuration property
    meetingDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The minimumAttendeePercentage property
    minimumAttendeePercentage *float64
    // The returnSuggestionReasons property
    returnSuggestionReasons *bool
    // The timeConstraint property
    timeConstraint ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable
}
// NewFindMeetingTimesRequestBody instantiates a new findMeetingTimesRequestBody and sets the default values.
func NewFindMeetingTimesRequestBody()(*FindMeetingTimesRequestBody) {
    m := &FindMeetingTimesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFindMeetingTimesRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFindMeetingTimesRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFindMeetingTimesRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FindMeetingTimesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttendees gets the attendees property value. The attendees property
func (m *FindMeetingTimesRequestBody) GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FindMeetingTimesRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendeeBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["isOrganizerOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizerOptional(val)
        }
        return nil
    }
    res["locationConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLocationConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationConstraint(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable))
        }
        return nil
    }
    res["maxCandidates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCandidates(val)
        }
        return nil
    }
    res["meetingDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingDuration(val)
        }
        return nil
    }
    res["minimumAttendeePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumAttendeePercentage(val)
        }
        return nil
    }
    res["returnSuggestionReasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnSuggestionReasons(val)
        }
        return nil
    }
    res["timeConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimeConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeConstraint(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable))
        }
        return nil
    }
    return res
}
// GetIsOrganizerOptional gets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *FindMeetingTimesRequestBody) GetIsOrganizerOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizerOptional
    }
}
// GetLocationConstraint gets the locationConstraint property value. The locationConstraint property
func (m *FindMeetingTimesRequestBody) GetLocationConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable) {
    if m == nil {
        return nil
    } else {
        return m.locationConstraint
    }
}
// GetMaxCandidates gets the maxCandidates property value. The maxCandidates property
func (m *FindMeetingTimesRequestBody) GetMaxCandidates()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxCandidates
    }
}
// GetMeetingDuration gets the meetingDuration property value. The meetingDuration property
func (m *FindMeetingTimesRequestBody) GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.meetingDuration
    }
}
// GetMinimumAttendeePercentage gets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *FindMeetingTimesRequestBody) GetMinimumAttendeePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.minimumAttendeePercentage
    }
}
// GetReturnSuggestionReasons gets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *FindMeetingTimesRequestBody) GetReturnSuggestionReasons()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.returnSuggestionReasons
    }
}
// GetTimeConstraint gets the timeConstraint property value. The timeConstraint property
func (m *FindMeetingTimesRequestBody) GetTimeConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable) {
    if m == nil {
        return nil
    } else {
        return m.timeConstraint
    }
}
// Serialize serializes information the current object
func (m *FindMeetingTimesRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOrganizerOptional", m.GetIsOrganizerOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locationConstraint", m.GetLocationConstraint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxCandidates", m.GetMaxCandidates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("meetingDuration", m.GetMeetingDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("minimumAttendeePercentage", m.GetMinimumAttendeePercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("returnSuggestionReasons", m.GetReturnSuggestionReasons())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeConstraint", m.GetTimeConstraint())
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
func (m *FindMeetingTimesRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttendees sets the attendees property value. The attendees property
func (m *FindMeetingTimesRequestBody) SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)() {
    if m != nil {
        m.attendees = value
    }
}
// SetIsOrganizerOptional sets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *FindMeetingTimesRequestBody) SetIsOrganizerOptional(value *bool)() {
    if m != nil {
        m.isOrganizerOptional = value
    }
}
// SetLocationConstraint sets the locationConstraint property value. The locationConstraint property
func (m *FindMeetingTimesRequestBody) SetLocationConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)() {
    if m != nil {
        m.locationConstraint = value
    }
}
// SetMaxCandidates sets the maxCandidates property value. The maxCandidates property
func (m *FindMeetingTimesRequestBody) SetMaxCandidates(value *int32)() {
    if m != nil {
        m.maxCandidates = value
    }
}
// SetMeetingDuration sets the meetingDuration property value. The meetingDuration property
func (m *FindMeetingTimesRequestBody) SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.meetingDuration = value
    }
}
// SetMinimumAttendeePercentage sets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *FindMeetingTimesRequestBody) SetMinimumAttendeePercentage(value *float64)() {
    if m != nil {
        m.minimumAttendeePercentage = value
    }
}
// SetReturnSuggestionReasons sets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *FindMeetingTimesRequestBody) SetReturnSuggestionReasons(value *bool)() {
    if m != nil {
        m.returnSuggestionReasons = value
    }
}
// SetTimeConstraint sets the timeConstraint property value. The timeConstraint property
func (m *FindMeetingTimesRequestBody) SetTimeConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)() {
    if m != nil {
        m.timeConstraint = value
    }
}
