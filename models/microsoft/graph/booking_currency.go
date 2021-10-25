package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type BookingCurrency struct {
    Entity
    symbol *string;
}
func NewBookingCurrency()(*BookingCurrency) {
    m := &BookingCurrency{
        Entity: *NewEntity(),
    }
    return m
}
func (m *BookingCurrency) GetSymbol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.symbol
    }
}
func (m *BookingCurrency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["symbol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSymbol(val)
        return nil
    }
    return res
}
func (m *BookingCurrency) IsNil()(bool) {
    return m == nil
}
func (m *BookingCurrency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("symbol", m.GetSymbol())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *BookingCurrency) SetSymbol(value *string)() {
    m.symbol = value
}
