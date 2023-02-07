package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentState 
type DeploymentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The effectiveValue property
    effectiveValue *DeploymentStateValue
    // The OdataType property
    odataType *string
    // Specifies the reasons the deployment has its state value. Read-only.
    reasons []DeploymentStateReasonable
    // The requestedValue property
    requestedValue *RequestedDeploymentStateValue
}
// NewDeploymentState instantiates a new deploymentState and sets the default values.
func NewDeploymentState()(*DeploymentState) {
    m := &DeploymentState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEffectiveValue gets the effectiveValue property value. The effectiveValue property
func (m *DeploymentState) GetEffectiveValue()(*DeploymentStateValue) {
    return m.effectiveValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["effectiveValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEffectiveValue(val.(*DeploymentStateValue))
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
    res["reasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeploymentStateReasonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeploymentStateReasonable, len(val))
            for i, v := range val {
                res[i] = v.(DeploymentStateReasonable)
            }
            m.SetReasons(res)
        }
        return nil
    }
    res["requestedValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestedDeploymentStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedValue(val.(*RequestedDeploymentStateValue))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeploymentState) GetOdataType()(*string) {
    return m.odataType
}
// GetReasons gets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) GetReasons()([]DeploymentStateReasonable) {
    return m.reasons
}
// GetRequestedValue gets the requestedValue property value. The requestedValue property
func (m *DeploymentState) GetRequestedValue()(*RequestedDeploymentStateValue) {
    return m.requestedValue
}
// Serialize serializes information the current object
func (m *DeploymentState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEffectiveValue() != nil {
        cast := (*m.GetEffectiveValue()).String()
        err := writer.WriteStringValue("effectiveValue", &cast)
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
    if m.GetReasons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReasons()))
        for i, v := range m.GetReasons() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("reasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestedValue() != nil {
        cast := (*m.GetRequestedValue()).String()
        err := writer.WriteStringValue("requestedValue", &cast)
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
func (m *DeploymentState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEffectiveValue sets the effectiveValue property value. The effectiveValue property
func (m *DeploymentState) SetEffectiveValue(value *DeploymentStateValue)() {
    m.effectiveValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeploymentState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetReasons sets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) SetReasons(value []DeploymentStateReasonable)() {
    m.reasons = value
}
// SetRequestedValue sets the requestedValue property value. The requestedValue property
func (m *DeploymentState) SetRequestedValue(value *RequestedDeploymentStateValue)() {
    m.requestedValue = value
}
