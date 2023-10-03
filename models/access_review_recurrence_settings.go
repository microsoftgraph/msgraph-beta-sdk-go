package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AccessReviewRecurrenceSettings 
type AccessReviewRecurrenceSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAccessReviewRecurrenceSettings instantiates a new accessReviewRecurrenceSettings and sets the default values.
func NewAccessReviewRecurrenceSettings()(*AccessReviewRecurrenceSettings) {
    m := &AccessReviewRecurrenceSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewRecurrenceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewRecurrenceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewRecurrenceSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewRecurrenceSettings) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AccessReviewRecurrenceSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDurationInDays gets the durationInDays property value. The duration in days for recurrence.
func (m *AccessReviewRecurrenceSettings) GetDurationInDays()(*int32) {
    val, err := m.GetBackingStore().Get("durationInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewRecurrenceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInDays(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["recurrenceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceCount(val)
        }
        return nil
    }
    res["recurrenceEndType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceEndType(val)
        }
        return nil
    }
    res["recurrenceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewRecurrenceSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecurrenceCount gets the recurrenceCount property value. The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceCount()(*int32) {
    val, err := m.GetBackingStore().Get("recurrenceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRecurrenceEndType gets the recurrenceEndType property value. How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it's never, then there's no explicit end of the recurrence series. If it's endBy, then the recurrence ends at a certain date. If it's occurrences, then the series ends after recurrenceCount instances of the review have completed.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceEndType()(*string) {
    val, err := m.GetBackingStore().Get("recurrenceEndType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecurrenceType gets the recurrenceType property value. The recurrence interval. Possible values: onetime, weekly, monthly, quarterly, halfyearly or annual.
func (m *AccessReviewRecurrenceSettings) GetRecurrenceType()(*string) {
    val, err := m.GetBackingStore().Get("recurrenceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewRecurrenceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewRecurrenceSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AccessReviewRecurrenceSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDurationInDays sets the durationInDays property value. The duration in days for recurrence.
func (m *AccessReviewRecurrenceSettings) SetDurationInDays(value *int32)() {
    err := m.GetBackingStore().Set("durationInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewRecurrenceSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrenceCount sets the recurrenceCount property value. The count of recurrences, if the value of recurrenceEndType is occurrences, or 0 otherwise.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceCount(value *int32)() {
    err := m.GetBackingStore().Set("recurrenceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrenceEndType sets the recurrenceEndType property value. How the recurrence ends. Possible values: never, endBy, occurrences, or recurrenceCount. If it's never, then there's no explicit end of the recurrence series. If it's endBy, then the recurrence ends at a certain date. If it's occurrences, then the series ends after recurrenceCount instances of the review have completed.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceEndType(value *string)() {
    err := m.GetBackingStore().Set("recurrenceEndType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrenceType sets the recurrenceType property value. The recurrence interval. Possible values: onetime, weekly, monthly, quarterly, halfyearly or annual.
func (m *AccessReviewRecurrenceSettings) SetRecurrenceType(value *string)() {
    err := m.GetBackingStore().Set("recurrenceType", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewRecurrenceSettingsable 
type AccessReviewRecurrenceSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDurationInDays()(*int32)
    GetOdataType()(*string)
    GetRecurrenceCount()(*int32)
    GetRecurrenceEndType()(*string)
    GetRecurrenceType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDurationInDays(value *int32)()
    SetOdataType(value *string)()
    SetRecurrenceCount(value *int32)()
    SetRecurrenceEndType(value *string)()
    SetRecurrenceType(value *string)()
}
