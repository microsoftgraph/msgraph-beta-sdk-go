package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody 
type ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The atApprovedLocation property
    atApprovedLocation *bool
    // The notes property
    notes ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable
}
// NewItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody instantiates a new ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody and sets the default values.
func NewItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody()(*ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) {
    m := &ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAtApprovedLocation gets the atApprovedLocation property value. The atApprovedLocation property
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) GetAtApprovedLocation()(*bool) {
    return m.atApprovedLocation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["atApprovedLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAtApprovedLocation(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable))
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. The notes property
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) GetNotes()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable) {
    return m.notes
}
// Serialize serializes information the current object
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("atApprovedLocation", m.GetAtApprovedLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notes", m.GetNotes())
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
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAtApprovedLocation sets the atApprovedLocation property value. The atApprovedLocation property
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) SetAtApprovedLocation(value *bool)() {
    m.atApprovedLocation = value
}
// SetNotes sets the notes property value. The notes property
func (m *ItemScheduleTimeCardsItemMicrosoftGraphStartBreakStartBreakPostRequestBody) SetNotes(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemBodyable)() {
    m.notes = value
}
