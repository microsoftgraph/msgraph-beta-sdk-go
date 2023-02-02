package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody 
type ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
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
// NewItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody instantiates a new ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody and sets the default values.
func NewItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody()(*ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) {
    m := &ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttendees gets the attendees property value. The attendees property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable) {
    return m.attendees
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetIsOrganizerOptional()(*bool) {
    return m.isOrganizerOptional
}
// GetLocationConstraint gets the locationConstraint property value. The locationConstraint property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetLocationConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable) {
    return m.locationConstraint
}
// GetMaxCandidates gets the maxCandidates property value. The maxCandidates property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetMaxCandidates()(*int32) {
    return m.maxCandidates
}
// GetMeetingDuration gets the meetingDuration property value. The meetingDuration property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.meetingDuration
}
// GetMinimumAttendeePercentage gets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetMinimumAttendeePercentage()(*float64) {
    return m.minimumAttendeePercentage
}
// GetReturnSuggestionReasons gets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetReturnSuggestionReasons()(*bool) {
    return m.returnSuggestionReasons
}
// GetTimeConstraint gets the timeConstraint property value. The timeConstraint property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) GetTimeConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable) {
    return m.timeConstraint
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttendees sets the attendees property value. The attendees property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)() {
    m.attendees = value
}
// SetIsOrganizerOptional sets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetIsOrganizerOptional(value *bool)() {
    m.isOrganizerOptional = value
}
// SetLocationConstraint sets the locationConstraint property value. The locationConstraint property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetLocationConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)() {
    m.locationConstraint = value
}
// SetMaxCandidates sets the maxCandidates property value. The maxCandidates property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetMaxCandidates(value *int32)() {
    m.maxCandidates = value
}
// SetMeetingDuration sets the meetingDuration property value. The meetingDuration property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.meetingDuration = value
}
// SetMinimumAttendeePercentage sets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetMinimumAttendeePercentage(value *float64)() {
    m.minimumAttendeePercentage = value
}
// SetReturnSuggestionReasons sets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetReturnSuggestionReasons(value *bool)() {
    m.returnSuggestionReasons = value
}
// SetTimeConstraint sets the timeConstraint property value. The timeConstraint property
func (m *ItemMicrosoftGraphFindMeetingTimesFindMeetingTimesPostRequestBody) SetTimeConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)() {
    m.timeConstraint = value
}
