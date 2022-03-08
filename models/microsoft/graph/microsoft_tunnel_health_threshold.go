package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelHealthThreshold provides operations to manage the deviceManagement singleton.
type MicrosoftTunnelHealthThreshold struct {
    Entity
    // The default threshold for being healthy
    defaultHealthyThreshold *int64;
    // The default threshold for being unhealthy
    defaultUnhealthyThreshold *int64;
    // The threshold for being healthy
    healthyThreshold *int64;
    // The threshold for being unhealthy
    unhealthyThreshold *int64;
}
// NewMicrosoftTunnelHealthThreshold instantiates a new microsoftTunnelHealthThreshold and sets the default values.
func NewMicrosoftTunnelHealthThreshold()(*MicrosoftTunnelHealthThreshold) {
    m := &MicrosoftTunnelHealthThreshold{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelHealthThresholdFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *MicrosoftTunnelHealthThreshold) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultHealthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultHealthyThreshold(val)
        }
        return nil
    }
    res["defaultUnhealthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultUnhealthyThreshold(val)
        }
        return nil
    }
    res["healthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthyThreshold(val)
        }
        return nil
    }
    res["unhealthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *MicrosoftTunnelHealthThreshold) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MicrosoftTunnelHealthThreshold) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
