package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ResponseTimeSecurityRequirement struct {
    SecurityRequirement
}
// NewResponseTimeSecurityRequirement instantiates a new ResponseTimeSecurityRequirement and sets the default values.
func NewResponseTimeSecurityRequirement()(*ResponseTimeSecurityRequirement) {
    m := &ResponseTimeSecurityRequirement{
        SecurityRequirement: *NewSecurityRequirement(),
    }
    return m
}
// CreateResponseTimeSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateResponseTimeSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResponseTimeSecurityRequirement(), nil
}
// GetAverageResponseTimeInHours gets the averageResponseTimeInHours property value. The average response time for alerts from the past 30 days.
// returns a *float32 when successful
func (m *ResponseTimeSecurityRequirement) GetAverageResponseTimeInHours()(*float32) {
    val, err := m.GetBackingStore().Get("averageResponseTimeInHours")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ResponseTimeSecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecurityRequirement.GetFieldDeserializers()
    res["averageResponseTimeInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageResponseTimeInHours(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ResponseTimeSecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecurityRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat32Value("averageResponseTimeInHours", m.GetAverageResponseTimeInHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageResponseTimeInHours sets the averageResponseTimeInHours property value. The average response time for alerts from the past 30 days.
func (m *ResponseTimeSecurityRequirement) SetAverageResponseTimeInHours(value *float32)() {
    err := m.GetBackingStore().Set("averageResponseTimeInHours", value)
    if err != nil {
        panic(err)
    }
}
type ResponseTimeSecurityRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SecurityRequirementable
    GetAverageResponseTimeInHours()(*float32)
    SetAverageResponseTimeInHours(value *float32)()
}
