package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionDurationInDays 
type RetentionDurationInDays struct {
    RetentionDuration
}
// NewRetentionDurationInDays instantiates a new retentionDurationInDays and sets the default values.
func NewRetentionDurationInDays()(*RetentionDurationInDays) {
    m := &RetentionDurationInDays{
        RetentionDuration: *NewRetentionDuration(),
    }
    odataTypeValue := "#microsoft.graph.security.retentionDurationInDays"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRetentionDurationInDaysFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionDurationInDaysFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionDurationInDays(), nil
}
// GetDays gets the days property value. Specifies the time period in days for which an item with the applied retention label will be retained for.
func (m *RetentionDurationInDays) GetDays()(*int32) {
    val, err := m.GetBackingStore().Get("days")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionDurationInDays) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RetentionDuration.GetFieldDeserializers()
    res["days"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDays(val)
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
func (m *RetentionDurationInDays) GetOdataType()(*string) {
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
func (m *RetentionDurationInDays) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RetentionDuration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("days", m.GetDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDays sets the days property value. Specifies the time period in days for which an item with the applied retention label will be retained for.
func (m *RetentionDurationInDays) SetDays(value *int32)() {
    err := m.GetBackingStore().Set("days", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RetentionDurationInDays) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// RetentionDurationInDaysable 
type RetentionDurationInDaysable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RetentionDurationable
    GetDays()(*int32)
    GetOdataType()(*string)
    SetDays(value *int32)()
    SetOdataType(value *string)()
}
