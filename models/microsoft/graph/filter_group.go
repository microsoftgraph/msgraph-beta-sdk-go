package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FilterGroup struct {
    additionalData map[string]interface{};
    clauses []FilterClause;
    name *string;
}
func NewFilterGroup()(*FilterGroup) {
    m := &FilterGroup{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FilterGroup) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FilterGroup) GetClauses()([]FilterClause) {
    if m == nil {
        return nil
    } else {
        return m.clauses
    }
}
func (m *FilterGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *FilterGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["clauses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFilterClause() })
        if err != nil {
            return err
        }
        res := make([]FilterClause, len(val))
        for i, v := range val {
            res[i] = *(v.(*FilterClause))
        }
        m.SetClauses(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    return res
}
func (m *FilterGroup) IsNil()(bool) {
    return m == nil
}
func (m *FilterGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetClauses()))
        for i, v := range m.GetClauses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("clauses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *FilterGroup) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FilterGroup) SetClauses(value []FilterClause)() {
    m.clauses = value
}
func (m *FilterGroup) SetName(value *string)() {
    m.name = value
}
