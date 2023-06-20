package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeploymentState 
type DeploymentState struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeploymentState instantiates a new deploymentState and sets the default values.
func NewDeploymentState()(*DeploymentState) {
    m := &DeploymentState{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeploymentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentState) GetAdditionalData()(map[string]any) {
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
func (m *DeploymentState) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEffectiveValue gets the effectiveValue property value. The effectiveValue property
func (m *DeploymentState) GetEffectiveValue()(*DeploymentStateValue) {
    val, err := m.GetBackingStore().Get("effectiveValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeploymentStateValue)
    }
    return nil
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
                if v != nil {
                    res[i] = v.(DeploymentStateReasonable)
                }
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
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReasons gets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) GetReasons()([]DeploymentStateReasonable) {
    val, err := m.GetBackingStore().Get("reasons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeploymentStateReasonable)
    }
    return nil
}
// GetRequestedValue gets the requestedValue property value. The requestedValue property
func (m *DeploymentState) GetRequestedValue()(*RequestedDeploymentStateValue) {
    val, err := m.GetBackingStore().Get("requestedValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RequestedDeploymentStateValue)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeploymentState) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEffectiveValue sets the effectiveValue property value. The effectiveValue property
func (m *DeploymentState) SetEffectiveValue(value *DeploymentStateValue)() {
    err := m.GetBackingStore().Set("effectiveValue", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeploymentState) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetReasons sets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) SetReasons(value []DeploymentStateReasonable)() {
    err := m.GetBackingStore().Set("reasons", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedValue sets the requestedValue property value. The requestedValue property
func (m *DeploymentState) SetRequestedValue(value *RequestedDeploymentStateValue)() {
    err := m.GetBackingStore().Set("requestedValue", value)
    if err != nil {
        panic(err)
    }
}
// DeploymentStateable 
type DeploymentStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEffectiveValue()(*DeploymentStateValue)
    GetOdataType()(*string)
    GetReasons()([]DeploymentStateReasonable)
    GetRequestedValue()(*RequestedDeploymentStateValue)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEffectiveValue(value *DeploymentStateValue)()
    SetOdataType(value *string)()
    SetReasons(value []DeploymentStateReasonable)()
    SetRequestedValue(value *RequestedDeploymentStateValue)()
}
