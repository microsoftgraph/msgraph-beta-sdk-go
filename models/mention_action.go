package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MentionAction 
type MentionAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identities of the users mentioned in this action.
    mentionees []IdentitySetable
    // The OdataType property
    odataType *string
}
// NewMentionAction instantiates a new mentionAction and sets the default values.
func NewMentionAction()(*MentionAction) {
    m := &MentionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMentionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMentionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMentionAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MentionAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MentionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mentionees"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIdentitySetFromDiscriminatorValue , m.SetMentionees)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetMentionees gets the mentionees property value. The identities of the users mentioned in this action.
func (m *MentionAction) GetMentionees()([]IdentitySetable) {
    return m.mentionees
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MentionAction) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MentionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMentionees() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMentionees())
        err := writer.WriteCollectionOfObjectValues("mentionees", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MentionAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMentionees sets the mentionees property value. The identities of the users mentioned in this action.
func (m *MentionAction) SetMentionees(value []IdentitySetable)() {
    m.mentionees = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MentionAction) SetOdataType(value *string)() {
    m.odataType = value
}
