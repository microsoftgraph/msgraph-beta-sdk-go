package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Currency 
type Currency struct {
    Entity
    // 
    amountDecimalPlaces *string;
    // 
    amountRoundingPrecision *float64;
    // 
    code *string;
    // 
    displayName *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    symbol *string;
}
// NewCurrency instantiates a new currency and sets the default values.
func NewCurrency()(*Currency) {
    m := &Currency{
        Entity: *NewEntity(),
    }
    return m
}
// GetAmountDecimalPlaces gets the amountDecimalPlaces property value. 
func (m *Currency) GetAmountDecimalPlaces()(*string) {
    if m == nil {
        return nil
    } else {
        return m.amountDecimalPlaces
    }
}
// GetAmountRoundingPrecision gets the amountRoundingPrecision property value. 
func (m *Currency) GetAmountRoundingPrecision()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.amountRoundingPrecision
    }
}
// GetCode gets the code property value. 
func (m *Currency) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDisplayName gets the displayName property value. 
func (m *Currency) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *Currency) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetSymbol gets the symbol property value. 
func (m *Currency) GetSymbol()(*string) {
    if m == nil {
        return nil
    } else {
        return m.symbol
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Currency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["amountDecimalPlaces"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountDecimalPlaces(val)
        }
        return nil
    }
    res["amountRoundingPrecision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountRoundingPrecision(val)
        }
        return nil
    }
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["symbol"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSymbol(val)
        }
        return nil
    }
    return res
}
func (m *Currency) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Currency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("amountDecimalPlaces", m.GetAmountDecimalPlaces())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amountRoundingPrecision", m.GetAmountRoundingPrecision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("symbol", m.GetSymbol())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAmountDecimalPlaces sets the amountDecimalPlaces property value. 
func (m *Currency) SetAmountDecimalPlaces(value *string)() {
    if m != nil {
        m.amountDecimalPlaces = value
    }
}
// SetAmountRoundingPrecision sets the amountRoundingPrecision property value. 
func (m *Currency) SetAmountRoundingPrecision(value *float64)() {
    if m != nil {
        m.amountRoundingPrecision = value
    }
}
// SetCode sets the code property value. 
func (m *Currency) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *Currency) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *Currency) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetSymbol sets the symbol property value. 
func (m *Currency) SetSymbol(value *string)() {
    if m != nil {
        m.symbol = value
    }
}
