package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSAppleEventReceiver represents a process that can receive an Apple Event notification.
type MacOSAppleEventReceiver struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Allow or block this app from receiving Apple events.
    allowed *bool
    // Code requirement for the app or binary that receives the Apple Event.
    codeRequirement *string
    // Bundle ID of the app or file path of the process or executable that receives the Apple Event.
    identifier *string
    // Process identifier types for MacOS Privacy Preferences
    identifierType *MacOSProcessIdentifierType
    // The OdataType property
    odataType *string
}
// NewMacOSAppleEventReceiver instantiates a new macOSAppleEventReceiver and sets the default values.
func NewMacOSAppleEventReceiver()(*MacOSAppleEventReceiver) {
    m := &MacOSAppleEventReceiver{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.macOSAppleEventReceiver";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSAppleEventReceiverFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSAppleEventReceiverFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSAppleEventReceiver(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSAppleEventReceiver) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowed gets the allowed property value. Allow or block this app from receiving Apple events.
func (m *MacOSAppleEventReceiver) GetAllowed()(*bool) {
    return m.allowed
}
// GetCodeRequirement gets the codeRequirement property value. Code requirement for the app or binary that receives the Apple Event.
func (m *MacOSAppleEventReceiver) GetCodeRequirement()(*string) {
    return m.codeRequirement
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSAppleEventReceiver) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowed)
    res["codeRequirement"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCodeRequirement)
    res["identifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIdentifier)
    res["identifierType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMacOSProcessIdentifierType , m.SetIdentifierType)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetIdentifier gets the identifier property value. Bundle ID of the app or file path of the process or executable that receives the Apple Event.
func (m *MacOSAppleEventReceiver) GetIdentifier()(*string) {
    return m.identifier
}
// GetIdentifierType gets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSAppleEventReceiver) GetIdentifierType()(*MacOSProcessIdentifierType) {
    return m.identifierType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSAppleEventReceiver) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MacOSAppleEventReceiver) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowed", m.GetAllowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeRequirement", m.GetCodeRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierType() != nil {
        cast := (*m.GetIdentifierType()).String()
        err := writer.WriteStringValue("identifierType", &cast)
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
func (m *MacOSAppleEventReceiver) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowed sets the allowed property value. Allow or block this app from receiving Apple events.
func (m *MacOSAppleEventReceiver) SetAllowed(value *bool)() {
    m.allowed = value
}
// SetCodeRequirement sets the codeRequirement property value. Code requirement for the app or binary that receives the Apple Event.
func (m *MacOSAppleEventReceiver) SetCodeRequirement(value *string)() {
    m.codeRequirement = value
}
// SetIdentifier sets the identifier property value. Bundle ID of the app or file path of the process or executable that receives the Apple Event.
func (m *MacOSAppleEventReceiver) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetIdentifierType sets the identifierType property value. Process identifier types for MacOS Privacy Preferences
func (m *MacOSAppleEventReceiver) SetIdentifierType(value *MacOSProcessIdentifierType)() {
    m.identifierType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSAppleEventReceiver) SetOdataType(value *string)() {
    m.odataType = value
}
