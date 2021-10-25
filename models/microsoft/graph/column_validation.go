package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ColumnValidation struct {
    additionalData map[string]interface{};
    defaultLanguage *string;
    descriptions []DisplayNameLocalization;
    formula *string;
}
func NewColumnValidation()(*ColumnValidation) {
    m := &ColumnValidation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ColumnValidation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ColumnValidation) GetDefaultLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLanguage
    }
}
func (m *ColumnValidation) GetDescriptions()([]DisplayNameLocalization) {
    if m == nil {
        return nil
    } else {
        return m.descriptions
    }
}
func (m *ColumnValidation) GetFormula()(*string) {
    if m == nil {
        return nil
    } else {
        return m.formula
    }
}
func (m *ColumnValidation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["defaultLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDefaultLanguage(val)
        return nil
    }
    res["descriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDisplayNameLocalization() })
        if err != nil {
            return err
        }
        res := make([]DisplayNameLocalization, len(val))
        for i, v := range val {
            res[i] = *(v.(*DisplayNameLocalization))
        }
        m.SetDescriptions(res)
        return nil
    }
    res["formula"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFormula(val)
        return nil
    }
    return res
}
func (m *ColumnValidation) IsNil()(bool) {
    return m == nil
}
func (m *ColumnValidation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDescriptions()))
        for i, v := range m.GetDescriptions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("descriptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ColumnValidation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ColumnValidation) SetDefaultLanguage(value *string)() {
    m.defaultLanguage = value
}
func (m *ColumnValidation) SetDescriptions(value []DisplayNameLocalization)() {
    m.descriptions = value
}
func (m *ColumnValidation) SetFormula(value *string)() {
    m.formula = value
}
