package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Financials struct {
    additionalData map[string]interface{};
    companies []Company;
}
func NewFinancials()(*Financials) {
    m := &Financials{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Financials) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Financials) GetCompanies()([]Company) {
    if m == nil {
        return nil
    } else {
        return m.companies
    }
}
func (m *Financials) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["companies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCompany() })
        if err != nil {
            return err
        }
        res := make([]Company, len(val))
        for i, v := range val {
            res[i] = *(v.(*Company))
        }
        m.SetCompanies(res)
        return nil
    }
    return res
}
func (m *Financials) IsNil()(bool) {
    return m == nil
}
func (m *Financials) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCompanies()))
        for i, v := range m.GetCompanies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("companies", cast)
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
func (m *Financials) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Financials) SetCompanies(value []Company)() {
    m.companies = value
}
