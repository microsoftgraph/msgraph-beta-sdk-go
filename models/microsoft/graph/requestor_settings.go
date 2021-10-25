package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RequestorSettings struct {
    acceptRequests *bool;
    additionalData map[string]interface{};
    allowedRequestors []UserSet;
    scopeType *string;
}
func NewRequestorSettings()(*RequestorSettings) {
    m := &RequestorSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RequestorSettings) GetAcceptRequests()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.acceptRequests
    }
}
func (m *RequestorSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RequestorSettings) GetAllowedRequestors()([]UserSet) {
    if m == nil {
        return nil
    } else {
        return m.allowedRequestors
    }
}
func (m *RequestorSettings) GetScopeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.scopeType
    }
}
func (m *RequestorSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["acceptRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAcceptRequests(val)
        return nil
    }
    res["allowedRequestors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserSet() })
        if err != nil {
            return err
        }
        res := make([]UserSet, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserSet))
        }
        m.SetAllowedRequestors(res)
        return nil
    }
    res["scopeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetScopeType(val)
        return nil
    }
    return res
}
func (m *RequestorSettings) IsNil()(bool) {
    return m == nil
}
func (m *RequestorSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("acceptRequests", m.GetAcceptRequests())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAllowedRequestors()))
        for i, v := range m.GetAllowedRequestors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("allowedRequestors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scopeType", m.GetScopeType())
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
func (m *RequestorSettings) SetAcceptRequests(value *bool)() {
    m.acceptRequests = value
}
func (m *RequestorSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RequestorSettings) SetAllowedRequestors(value []UserSet)() {
    m.allowedRequestors = value
}
func (m *RequestorSettings) SetScopeType(value *string)() {
    m.scopeType = value
}
