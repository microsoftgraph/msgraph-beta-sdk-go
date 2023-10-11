package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ItemFindMeetingTimesPostRequestBody 
type ItemFindMeetingTimesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemFindMeetingTimesPostRequestBody instantiates a new ItemFindMeetingTimesPostRequestBody and sets the default values.
func NewItemFindMeetingTimesPostRequestBody()(*ItemFindMeetingTimesPostRequestBody) {
    m := &ItemFindMeetingTimesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemFindMeetingTimesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemFindMeetingTimesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemFindMeetingTimesPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemFindMeetingTimesPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAttendees gets the attendees property value. The attendees property
func (m *ItemFindMeetingTimesPostRequestBody) GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable) {
    val, err := m.GetBackingStore().Get("attendees")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *ItemFindMeetingTimesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemFindMeetingTimesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAttendeeBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)
                }
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
func (m *ItemFindMeetingTimesPostRequestBody) GetIsOrganizerOptional()(*bool) {
    val, err := m.GetBackingStore().Get("isOrganizerOptional")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLocationConstraint gets the locationConstraint property value. The locationConstraint property
func (m *ItemFindMeetingTimesPostRequestBody) GetLocationConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable) {
    val, err := m.GetBackingStore().Get("locationConstraint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)
    }
    return nil
}
// GetMaxCandidates gets the maxCandidates property value. The maxCandidates property
func (m *ItemFindMeetingTimesPostRequestBody) GetMaxCandidates()(*int32) {
    val, err := m.GetBackingStore().Get("maxCandidates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMeetingDuration gets the meetingDuration property value. The meetingDuration property
func (m *ItemFindMeetingTimesPostRequestBody) GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("meetingDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMinimumAttendeePercentage gets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *ItemFindMeetingTimesPostRequestBody) GetMinimumAttendeePercentage()(*float64) {
    val, err := m.GetBackingStore().Get("minimumAttendeePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetReturnSuggestionReasons gets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *ItemFindMeetingTimesPostRequestBody) GetReturnSuggestionReasons()(*bool) {
    val, err := m.GetBackingStore().Get("returnSuggestionReasons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTimeConstraint gets the timeConstraint property value. The timeConstraint property
func (m *ItemFindMeetingTimesPostRequestBody) GetTimeConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable) {
    val, err := m.GetBackingStore().Get("timeConstraint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemFindMeetingTimesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemFindMeetingTimesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttendees sets the attendees property value. The attendees property
func (m *ItemFindMeetingTimesPostRequestBody) SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)() {
    err := m.GetBackingStore().Set("attendees", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemFindMeetingTimesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsOrganizerOptional sets the isOrganizerOptional property value. The isOrganizerOptional property
func (m *ItemFindMeetingTimesPostRequestBody) SetIsOrganizerOptional(value *bool)() {
    err := m.GetBackingStore().Set("isOrganizerOptional", value)
    if err != nil {
        panic(err)
    }
}
// SetLocationConstraint sets the locationConstraint property value. The locationConstraint property
func (m *ItemFindMeetingTimesPostRequestBody) SetLocationConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)() {
    err := m.GetBackingStore().Set("locationConstraint", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxCandidates sets the maxCandidates property value. The maxCandidates property
func (m *ItemFindMeetingTimesPostRequestBody) SetMaxCandidates(value *int32)() {
    err := m.GetBackingStore().Set("maxCandidates", value)
    if err != nil {
        panic(err)
    }
}
// SetMeetingDuration sets the meetingDuration property value. The meetingDuration property
func (m *ItemFindMeetingTimesPostRequestBody) SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("meetingDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumAttendeePercentage sets the minimumAttendeePercentage property value. The minimumAttendeePercentage property
func (m *ItemFindMeetingTimesPostRequestBody) SetMinimumAttendeePercentage(value *float64)() {
    err := m.GetBackingStore().Set("minimumAttendeePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetReturnSuggestionReasons sets the returnSuggestionReasons property value. The returnSuggestionReasons property
func (m *ItemFindMeetingTimesPostRequestBody) SetReturnSuggestionReasons(value *bool)() {
    err := m.GetBackingStore().Set("returnSuggestionReasons", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeConstraint sets the timeConstraint property value. The timeConstraint property
func (m *ItemFindMeetingTimesPostRequestBody) SetTimeConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)() {
    err := m.GetBackingStore().Set("timeConstraint", value)
    if err != nil {
        panic(err)
    }
}
// ItemFindMeetingTimesPostRequestBodyable 
type ItemFindMeetingTimesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttendees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsOrganizerOptional()(*bool)
    GetLocationConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)
    GetMaxCandidates()(*int32)
    GetMeetingDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMinimumAttendeePercentage()(*float64)
    GetReturnSuggestionReasons()(*bool)
    GetTimeConstraint()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)
    SetAttendees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AttendeeBaseable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsOrganizerOptional(value *bool)()
    SetLocationConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LocationConstraintable)()
    SetMaxCandidates(value *int32)()
    SetMeetingDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMinimumAttendeePercentage(value *float64)()
    SetReturnSuggestionReasons(value *bool)()
    SetTimeConstraint(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimeConstraintable)()
}
