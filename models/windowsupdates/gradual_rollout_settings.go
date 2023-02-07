package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GradualRolloutSettings 
type GradualRolloutSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The duration between each set of devices being offered the update. The value is represented in ISO 8601 format for duration. Default value is P1D (1 day).
    durationBetweenOffers *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The OdataType property
    odataType *string
}
// NewGradualRolloutSettings instantiates a new gradualRolloutSettings and sets the default values.
func NewGradualRolloutSettings()(*GradualRolloutSettings) {
    m := &GradualRolloutSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGradualRolloutSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGradualRolloutSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsUpdates.dateDrivenRolloutSettings":
                        return NewDateDrivenRolloutSettings(), nil
                    case "#microsoft.graph.windowsUpdates.durationDrivenRolloutSettings":
                        return NewDurationDrivenRolloutSettings(), nil
                    case "#microsoft.graph.windowsUpdates.rateDrivenRolloutSettings":
                        return NewRateDrivenRolloutSettings(), nil
                }
            }
        }
    }
    return NewGradualRolloutSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GradualRolloutSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDurationBetweenOffers gets the durationBetweenOffers property value. The duration between each set of devices being offered the update. The value is represented in ISO 8601 format for duration. Default value is P1D (1 day).
func (m *GradualRolloutSettings) GetDurationBetweenOffers()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.durationBetweenOffers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GradualRolloutSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["durationBetweenOffers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationBetweenOffers(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GradualRolloutSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GradualRolloutSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteISODurationValue("durationBetweenOffers", m.GetDurationBetweenOffers())
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
func (m *GradualRolloutSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDurationBetweenOffers sets the durationBetweenOffers property value. The duration between each set of devices being offered the update. The value is represented in ISO 8601 format for duration. Default value is P1D (1 day).
func (m *GradualRolloutSettings) SetDurationBetweenOffers(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.durationBetweenOffers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GradualRolloutSettings) SetOdataType(value *string)() {
    m.odataType = value
}
