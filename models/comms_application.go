package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommsApplication 
type CommsApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The calls property
    calls []Callable
    // The OdataType property
    odataType *string
    // The onlineMeetings property
    onlineMeetings []OnlineMeetingable
}
// NewCommsApplication instantiates a new CommsApplication and sets the default values.
func NewCommsApplication()(*CommsApplication) {
    m := &CommsApplication{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCommsApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommsApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommsApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CommsApplication) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalls gets the calls property value. The calls property
func (m *CommsApplication) GetCalls()([]Callable) {
    return m.calls
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommsApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCallFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Callable, len(val))
            for i, v := range val {
                res[i] = v.(Callable)
            }
            m.SetCalls(res)
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
    res["onlineMeetings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnlineMeetingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeetingable, len(val))
            for i, v := range val {
                res[i] = v.(OnlineMeetingable)
            }
            m.SetOnlineMeetings(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CommsApplication) GetOdataType()(*string) {
    return m.odataType
}
// GetOnlineMeetings gets the onlineMeetings property value. The onlineMeetings property
func (m *CommsApplication) GetOnlineMeetings()([]OnlineMeetingable) {
    return m.onlineMeetings
}
// Serialize serializes information the current object
func (m *CommsApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCalls() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalls()))
        for i, v := range m.GetCalls() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("calls", cast)
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
    if m.GetOnlineMeetings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnlineMeetings()))
        for i, v := range m.GetOnlineMeetings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("onlineMeetings", cast)
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
func (m *CommsApplication) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalls sets the calls property value. The calls property
func (m *CommsApplication) SetCalls(value []Callable)() {
    m.calls = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CommsApplication) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnlineMeetings sets the onlineMeetings property value. The onlineMeetings property
func (m *CommsApplication) SetOnlineMeetings(value []OnlineMeetingable)() {
    m.onlineMeetings = value
}
