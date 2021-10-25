package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UsageRight struct {
    Entity
    catalogId *string;
    serviceIdentifier *string;
    state *UsageRightState;
}
func NewUsageRight()(*UsageRight) {
    m := &UsageRight{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UsageRight) GetCatalogId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.catalogId
    }
}
func (m *UsageRight) GetServiceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serviceIdentifier
    }
}
func (m *UsageRight) GetState()(*UsageRightState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *UsageRight) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalogId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCatalogId(val)
        return nil
    }
    res["serviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServiceIdentifier(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageRightState)
        if err != nil {
            return err
        }
        cast := val.(UsageRightState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *UsageRight) IsNil()(bool) {
    return m == nil
}
func (m *UsageRight) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("catalogId", m.GetCatalogId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceIdentifier", m.GetServiceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UsageRight) SetCatalogId(value *string)() {
    m.catalogId = value
}
func (m *UsageRight) SetServiceIdentifier(value *string)() {
    m.serviceIdentifier = value
}
func (m *UsageRight) SetState(value *UsageRightState)() {
    m.state = value
}
