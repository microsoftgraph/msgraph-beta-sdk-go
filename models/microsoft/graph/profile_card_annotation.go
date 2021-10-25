package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProfileCardAnnotation struct {
    additionalData map[string]interface{};
    displayName *string;
    localizations []DisplayNameLocalization;
}
func NewProfileCardAnnotation()(*ProfileCardAnnotation) {
    m := &ProfileCardAnnotation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ProfileCardAnnotation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ProfileCardAnnotation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ProfileCardAnnotation) GetLocalizations()([]DisplayNameLocalization) {
    if m == nil {
        return nil
    } else {
        return m.localizations
    }
}
func (m *ProfileCardAnnotation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["localizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDisplayNameLocalization() })
        if err != nil {
            return err
        }
        res := make([]DisplayNameLocalization, len(val))
        for i, v := range val {
            res[i] = *(v.(*DisplayNameLocalization))
        }
        m.SetLocalizations(res)
        return nil
    }
    return res
}
func (m *ProfileCardAnnotation) IsNil()(bool) {
    return m == nil
}
func (m *ProfileCardAnnotation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocalizations()))
        for i, v := range m.GetLocalizations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("localizations", cast)
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
func (m *ProfileCardAnnotation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ProfileCardAnnotation) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ProfileCardAnnotation) SetLocalizations(value []DisplayNameLocalization)() {
    m.localizations = value
}
