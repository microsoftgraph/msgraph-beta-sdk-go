package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MicrosoftTunnelHealthThreshold struct {
    Entity
    defaultHealthyThreshold *int64;
    defaultUnhealthyThreshold *int64;
    healthyThreshold *int64;
    unhealthyThreshold *int64;
}
func NewMicrosoftTunnelHealthThreshold()(*MicrosoftTunnelHealthThreshold) {
    m := &MicrosoftTunnelHealthThreshold{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MicrosoftTunnelHealthThreshold) GetDefaultHealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultHealthyThreshold
    }
}
func (m *MicrosoftTunnelHealthThreshold) GetDefaultUnhealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.defaultUnhealthyThreshold
    }
}
func (m *MicrosoftTunnelHealthThreshold) GetHealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.healthyThreshold
    }
}
func (m *MicrosoftTunnelHealthThreshold) GetUnhealthyThreshold()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.unhealthyThreshold
    }
}
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
func (m *MicrosoftTunnelHealthThreshold) SetDefaultHealthyThreshold(value *int64)() {
    m.defaultHealthyThreshold = value
}
func (m *MicrosoftTunnelHealthThreshold) SetDefaultUnhealthyThreshold(value *int64)() {
    m.defaultUnhealthyThreshold = value
}
func (m *MicrosoftTunnelHealthThreshold) SetHealthyThreshold(value *int64)() {
    m.healthyThreshold = value
}
func (m *MicrosoftTunnelHealthThreshold) SetUnhealthyThreshold(value *int64)() {
    m.unhealthyThreshold = value
}
