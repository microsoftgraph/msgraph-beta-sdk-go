package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new microsoftTunnelHealthThreshold and sets the default values.
func NewMicrosoftTunnelHealthThreshold()(*MicrosoftTunnelHealthThreshold) {
    m := &MicrosoftTunnelHealthThreshold{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the defaultHealthyThreshold property value. The default threshold for being healthy
func (m *MicrosoftTunnelHealthThreshold) GetDefaultHealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultHealthyThreshold
    }
}
// Gets the defaultUnhealthyThreshold property value. The default threshold for being unhealthy
func (m *MicrosoftTunnelHealthThreshold) GetDefaultUnhealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultUnhealthyThreshold
    }
}
// Gets the healthyThreshold property value. The threshold for being healthy
func (m *MicrosoftTunnelHealthThreshold) GetHealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.healthyThreshold
    }
}
// Gets the unhealthyThreshold property value. The threshold for being unhealthy
func (m *MicrosoftTunnelHealthThreshold) GetUnhealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unhealthyThreshold
    }
}
// The deserialization information for the current model
func (m *MicrosoftTunnelHealthThreshold) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultHealthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDefaultHealthyThreshold(val)
        return nil
    }
    res["defaultUnhealthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDefaultUnhealthyThreshold(val)
        return nil
    }
    res["healthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetHealthyThreshold(val)
        return nil
    }
    res["unhealthyThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUnhealthyThreshold(val)
        return nil
    }
    return res
}
func (m *MicrosoftTunnelHealthThreshold) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the defaultHealthyThreshold property value. The default threshold for being healthy
// Parameters:
//  - value : Value to set for the defaultHealthyThreshold property.
func (m *MicrosoftTunnelHealthThreshold) SetDefaultHealthyThreshold(value *int64)() {
    m.defaultHealthyThreshold = value
}
// Sets the defaultUnhealthyThreshold property value. The default threshold for being unhealthy
// Parameters:
//  - value : Value to set for the defaultUnhealthyThreshold property.
func (m *MicrosoftTunnelHealthThreshold) SetDefaultUnhealthyThreshold(value *int64)() {
    m.defaultUnhealthyThreshold = value
}
// Sets the healthyThreshold property value. The threshold for being healthy
// Parameters:
//  - value : Value to set for the healthyThreshold property.
func (m *MicrosoftTunnelHealthThreshold) SetHealthyThreshold(value *int64)() {
    m.healthyThreshold = value
}
// Sets the unhealthyThreshold property value. The threshold for being unhealthy
// Parameters:
//  - value : Value to set for the unhealthyThreshold property.
func (m *MicrosoftTunnelHealthThreshold) SetUnhealthyThreshold(value *int64)() {
    m.unhealthyThreshold = value
}
