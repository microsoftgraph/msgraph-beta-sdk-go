package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageHistoryItem 
type ChatMessageHistoryItem struct {
    // The actions property
    actions *ChatMessageActions
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The date and time when the message was modified.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The reaction in the modified message.
    reaction ChatMessageReactionable
}
// NewChatMessageHistoryItem instantiates a new chatMessageHistoryItem and sets the default values.
func NewChatMessageHistoryItem()(*ChatMessageHistoryItem) {
    m := &ChatMessageHistoryItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatMessageHistoryItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageHistoryItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessageHistoryItem(), nil
}
// GetActions gets the actions property value. The actions property
func (m *ChatMessageHistoryItem) GetActions()(*ChatMessageActions) {
    return m.actions
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessageHistoryItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageHistoryItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseChatMessageActions , m.SetActions)
    res["modifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetModifiedDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["reaction"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateChatMessageReactionFromDiscriminatorValue , m.SetReaction)
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The date and time when the message was modified.
func (m *ChatMessageHistoryItem) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChatMessageHistoryItem) GetOdataType()(*string) {
    return m.odataType
}
// GetReaction gets the reaction property value. The reaction in the modified message.
func (m *ChatMessageHistoryItem) GetReaction()(ChatMessageReactionable) {
    return m.reaction
}
// Serialize serializes information the current object
func (m *ChatMessageHistoryItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        cast := (*m.GetActions()).String()
        err := writer.WriteStringValue("actions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
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
        err := writer.WriteObjectValue("reaction", m.GetReaction())
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
// SetActions sets the actions property value. The actions property
func (m *ChatMessageHistoryItem) SetActions(value *ChatMessageActions)() {
    m.actions = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessageHistoryItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The date and time when the message was modified.
func (m *ChatMessageHistoryItem) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatMessageHistoryItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReaction sets the reaction property value. The reaction in the modified message.
func (m *ChatMessageHistoryItem) SetReaction(value ChatMessageReactionable)() {
    m.reaction = value
}
