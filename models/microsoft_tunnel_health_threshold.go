package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelHealthThreshold 
type MicrosoftTunnelHealthThreshold struct {
    Entity
    // The default threshold for being healthy
    defaultHealthyThreshold *int64
    // The default threshold for being unhealthy
    defaultUnhealthyThreshold *int64
    // The threshold for being healthy
    healthyThreshold *int64
    // The threshold for being unhealthy
    unhealthyThreshold *int64
}
// NewMicrosoftTunnelHealthThreshold instantiates a new MicrosoftTunnelHealthThreshold and sets the default values.
func NewMicrosoftTunnelHealthThreshold()(*MicrosoftTunnelHealthThreshold) {
    m := &MicrosoftTunnelHealthThreshold{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.microsoftTunnelHealthThreshold";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelHealthThreshold(), nil
}
// GetDefaultHealthyThreshold gets the defaultHealthyThreshold property value. The default threshold for being healthy
func (m *MicrosoftTunnelHealthThreshold) GetDefaultHealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultHealthyThreshold
    }
}
// GetDefaultUnhealthyThreshold gets the defaultUnhealthyThreshold property value. The default threshold for being unhealthy
func (m *MicrosoftTunnelHealthThreshold) GetDefaultUnhealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultUnhealthyThreshold
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelHealthThreshold) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultHealthyThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultHealthyThreshold(val)
        }
        return nil
    }
    res["defaultUnhealthyThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultUnhealthyThreshold(val)
        }
        return nil
    }
    res["healthyThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthyThreshold(val)
        }
        return nil
    }
    res["unhealthyThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnhealthyThreshold(val)
        }
        return nil
    }
    return res
}
// GetHealthyThreshold gets the healthyThreshold property value. The threshold for being healthy
func (m *MicrosoftTunnelHealthThreshold) GetHealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.healthyThreshold
    }
}
// GetUnhealthyThreshold gets the unhealthyThreshold property value. The threshold for being unhealthy
func (m *MicrosoftTunnelHealthThreshold) GetUnhealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unhealthyThreshold
    }
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelHealthThreshold) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("defaultHealthyThreshold", m.GetDefaultHealthyThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("defaultUnhealthyThreshold", m.GetDefaultUnhealthyThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("healthyThreshold", m.GetHealthyThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unhealthyThreshold", m.GetUnhealthyThreshold())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultHealthyThreshold sets the defaultHealthyThreshold property value. The default threshold for being healthy
func (m *MicrosoftTunnelHealthThreshold) SetDefaultHealthyThreshold(value *int64)() {
    if m != nil {
        m.defaultHealthyThreshold = value
    }
}
// SetDefaultUnhealthyThreshold sets the defaultUnhealthyThreshold property value. The default threshold for being unhealthy
func (m *MicrosoftTunnelHealthThreshold) SetDefaultUnhealthyThreshold(value *int64)() {
    if m != nil {
        m.defaultUnhealthyThreshold = value
    }
}
// SetHealthyThreshold sets the healthyThreshold property value. The threshold for being healthy
func (m *MicrosoftTunnelHealthThreshold) SetHealthyThreshold(value *int64)() {
    if m != nil {
        m.healthyThreshold = value
    }
}
// SetUnhealthyThreshold sets the unhealthyThreshold property value. The threshold for being unhealthy
func (m *MicrosoftTunnelHealthThreshold) SetUnhealthyThreshold(value *int64)() {
    if m != nil {
        m.unhealthyThreshold = value
    }
}
