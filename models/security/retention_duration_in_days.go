package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionDurationInDays 
type RetentionDurationInDays struct {
    RetentionDuration
    // Specifies the time period in days for which an item with the applied retention label will be retained for.
    days *int32
}
// NewRetentionDurationInDays instantiates a new RetentionDurationInDays and sets the default values.
func NewRetentionDurationInDays()(*RetentionDurationInDays) {
    m := &RetentionDurationInDays{
        RetentionDuration: *NewRetentionDuration(),
    }
    odataTypeValue := "#microsoft.graph.security.retentionDurationInDays";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRetentionDurationInDaysFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionDurationInDaysFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionDurationInDays(), nil
}
// GetDays gets the days property value. Specifies the time period in days for which an item with the applied retention label will be retained for.
func (m *RetentionDurationInDays) GetDays()(*int32) {
    return m.days
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RetentionDurationInDays) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RetentionDuration.GetFieldDeserializers()
    res["days"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDays)
    return res
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
    return nil
}
// SetDays sets the days property value. Specifies the time period in days for which an item with the applied retention label will be retained for.
func (m *RetentionDurationInDays) SetDays(value *int32)() {
    m.days = value
}
