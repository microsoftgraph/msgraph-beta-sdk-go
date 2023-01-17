package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadCoachmark 
type PayloadCoachmark struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The coachmarkLocation property
    coachmarkLocation CoachmarkLocationable
    // The description property
    description *string
    // The indicator property
    indicator *string
    // The isValid property
    isValid *bool
    // The language property
    language *string
    // The OdataType property
    odataType *string
    // The order property
    order *string
}
// NewPayloadCoachmark instantiates a new payloadCoachmark and sets the default values.
func NewPayloadCoachmark()(*PayloadCoachmark) {
    m := &PayloadCoachmark{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreatePayloadCoachmarkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadCoachmarkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayloadCoachmark(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadCoachmark) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCoachmarkLocation gets the coachmarkLocation property value. The coachmarkLocation property
func (m *PayloadCoachmark) GetCoachmarkLocation()(CoachmarkLocationable) {
    return m.coachmarkLocation
}
// GetDescription gets the description property value. The description property
func (m *PayloadCoachmark) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadCoachmark) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["coachmarkLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCoachmarkLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoachmarkLocation(val.(CoachmarkLocationable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["indicator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndicator(val)
        }
        return nil
    }
    res["isValid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
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
    res["order"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrder(val)
        }
        return nil
    }
    return res
}
// GetIndicator gets the indicator property value. The indicator property
func (m *PayloadCoachmark) GetIndicator()(*string) {
    return m.indicator
}
// GetIsValid gets the isValid property value. The isValid property
func (m *PayloadCoachmark) GetIsValid()(*bool) {
    return m.isValid
}
// GetLanguage gets the language property value. The language property
func (m *PayloadCoachmark) GetLanguage()(*string) {
    return m.language
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PayloadCoachmark) GetOdataType()(*string) {
    return m.odataType
}
// GetOrder gets the order property value. The order property
func (m *PayloadCoachmark) GetOrder()(*string) {
    return m.order
}
// Serialize serializes information the current object
func (m *PayloadCoachmark) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("coachmarkLocation", m.GetCoachmarkLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("indicator", m.GetIndicator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isValid", m.GetIsValid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
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
        err := writer.WriteStringValue("order", m.GetOrder())
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
func (m *PayloadCoachmark) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCoachmarkLocation sets the coachmarkLocation property value. The coachmarkLocation property
func (m *PayloadCoachmark) SetCoachmarkLocation(value CoachmarkLocationable)() {
    m.coachmarkLocation = value
}
// SetDescription sets the description property value. The description property
func (m *PayloadCoachmark) SetDescription(value *string)() {
    m.description = value
}
// SetIndicator sets the indicator property value. The indicator property
func (m *PayloadCoachmark) SetIndicator(value *string)() {
    m.indicator = value
}
// SetIsValid sets the isValid property value. The isValid property
func (m *PayloadCoachmark) SetIsValid(value *bool)() {
    m.isValid = value
}
// SetLanguage sets the language property value. The language property
func (m *PayloadCoachmark) SetLanguage(value *string)() {
    m.language = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PayloadCoachmark) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrder sets the order property value. The order property
func (m *PayloadCoachmark) SetOrder(value *string)() {
    m.order = value
}
