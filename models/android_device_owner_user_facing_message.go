package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerUserFacingMessage represents a user-facing message with locale information as well as a default message to be used if the user's locale doesn't match with any of the localized messages
type AndroidDeviceOwnerUserFacingMessage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The default message displayed if the user's locale doesn't match with any of the localized messages
    defaultMessage *string
    // The list of <locale, message> pairs. This collection can contain a maximum of 500 elements.
    localizedMessages []KeyValuePairable
    // The OdataType property
    odataType *string
}
// NewAndroidDeviceOwnerUserFacingMessage instantiates a new androidDeviceOwnerUserFacingMessage and sets the default values.
func NewAndroidDeviceOwnerUserFacingMessage()(*AndroidDeviceOwnerUserFacingMessage) {
    m := &AndroidDeviceOwnerUserFacingMessage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerUserFacingMessage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerUserFacingMessage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerUserFacingMessage) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultMessage gets the defaultMessage property value. The default message displayed if the user's locale doesn't match with any of the localized messages
func (m *AndroidDeviceOwnerUserFacingMessage) GetDefaultMessage()(*string) {
    return m.defaultMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerUserFacingMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultMessage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultMessage)
    res["localizedMessages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue , m.SetLocalizedMessages)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetLocalizedMessages gets the localizedMessages property value. The list of <locale, message> pairs. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerUserFacingMessage) GetLocalizedMessages()([]KeyValuePairable) {
    return m.localizedMessages
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerUserFacingMessage) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerUserFacingMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultMessage", m.GetDefaultMessage())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizedMessages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLocalizedMessages())
        err := writer.WriteCollectionOfObjectValues("localizedMessages", cast)
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
func (m *AndroidDeviceOwnerUserFacingMessage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultMessage sets the defaultMessage property value. The default message displayed if the user's locale doesn't match with any of the localized messages
func (m *AndroidDeviceOwnerUserFacingMessage) SetDefaultMessage(value *string)() {
    m.defaultMessage = value
}
// SetLocalizedMessages sets the localizedMessages property value. The list of <locale, message> pairs. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerUserFacingMessage) SetLocalizedMessages(value []KeyValuePairable)() {
    m.localizedMessages = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerUserFacingMessage) SetOdataType(value *string)() {
    m.odataType = value
}
