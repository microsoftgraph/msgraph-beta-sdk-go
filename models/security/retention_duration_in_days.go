package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RetentionDurationInDays 
type RetentionDurationInDays struct {
    RetentionDuration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the time period in days for which an item with the applied retention label will be retained for.
    days *int32
}
// NewRetentionDurationInDays instantiates a new RetentionDurationInDays and sets the default values.
func NewRetentionDurationInDays()(*RetentionDurationInDays) {
    m := &RetentionDurationInDays{
        RetentionDuration: *NewRetentionDuration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRetentionDurationInDaysFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRetentionDurationInDaysFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRetentionDurationInDays(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetentionDurationInDays) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDays gets the days property value. Specifies the time period in days for which an item with the applied retention label will be retained for.
func (m *RetentionDurationInDays) GetDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.days
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RetentionDurationInDays) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDays sets the days property value. Specifies the time period in days for which an item with the applied retention label will be retained for.
func (m *RetentionDurationInDays) SetDays(value *int32)() {
    if m != nil {
        m.days = value
    }
}
