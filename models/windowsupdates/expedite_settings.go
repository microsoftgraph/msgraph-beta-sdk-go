package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpediteSettings 
type ExpediteSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // True indicates that the deployment of the content is expedited.
    isExpedited *bool
    // The OdataType property
    odataType *string
}
// NewExpediteSettings instantiates a new expediteSettings and sets the default values.
func NewExpediteSettings()(*ExpediteSettings) {
    m := &ExpediteSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateExpediteSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpediteSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExpediteSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpediteSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpediteSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isExpedited"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsExpedited(val)
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
    return res
}
// GetIsExpedited gets the isExpedited property value. True indicates that the deployment of the content is expedited.
func (m *ExpediteSettings) GetIsExpedited()(*bool) {
    return m.isExpedited
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ExpediteSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ExpediteSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isExpedited", m.GetIsExpedited())
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
func (m *ExpediteSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsExpedited sets the isExpedited property value. True indicates that the deployment of the content is expedited.
func (m *ExpediteSettings) SetIsExpedited(value *bool)() {
    m.isExpedited = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ExpediteSettings) SetOdataType(value *string)() {
    m.odataType = value
}
