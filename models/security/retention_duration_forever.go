package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionDurationForever 
type RetentionDurationForever struct {
    RetentionDuration
}
// NewRetentionDurationForever instantiates a new retentionDurationForever and sets the default values.
func NewRetentionDurationForever()(*RetentionDurationForever) {
    m := &RetentionDurationForever{
        RetentionDuration: *NewRetentionDuration(),
    }
    odataTypeValue := "#microsoft.graph.security.retentionDurationForever"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRetentionDurationForeverFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionDurationForeverFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionDurationForever(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionDurationForever) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RetentionDuration.GetFieldDeserializers()
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
func (m *RetentionDurationForever) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RetentionDurationForever) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RetentionDuration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetentionDurationForever) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// RetentionDurationForeverable 
type RetentionDurationForeverable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RetentionDurationable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
