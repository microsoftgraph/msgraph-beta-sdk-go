package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PersonNamePronounciation struct {
    additionalData map[string]interface{};
    displayName *string;
    first *string;
    last *string;
    maiden *string;
    middle *string;
}
func NewPersonNamePronounciation()(*PersonNamePronounciation) {
    m := &PersonNamePronounciation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PersonNamePronounciation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PersonNamePronounciation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *PersonNamePronounciation) GetFirst()(*string) {
    if m == nil {
        return nil
    } else {
        return m.first
    }
}
func (m *PersonNamePronounciation) GetLast()(*string) {
    if m == nil {
        return nil
    } else {
        return m.last
    }
}
func (m *PersonNamePronounciation) GetMaiden()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maiden
    }
}
func (m *PersonNamePronounciation) GetMiddle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middle
    }
}
func (m *PersonNamePronounciation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["first"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFirst(val)
        return nil
    }
    res["last"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLast(val)
        return nil
    }
    res["maiden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMaiden(val)
        return nil
    }
    res["middle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMiddle(val)
        return nil
    }
    return res
}
func (m *PersonNamePronounciation) IsNil()(bool) {
    return m == nil
}
func (m *PersonNamePronounciation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("first", m.GetFirst())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last", m.GetLast())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maiden", m.GetMaiden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("middle", m.GetMiddle())
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
func (m *PersonNamePronounciation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PersonNamePronounciation) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *PersonNamePronounciation) SetFirst(value *string)() {
    m.first = value
}
func (m *PersonNamePronounciation) SetLast(value *string)() {
    m.last = value
}
func (m *PersonNamePronounciation) SetMaiden(value *string)() {
    m.maiden = value
}
func (m *PersonNamePronounciation) SetMiddle(value *string)() {
    m.middle = value
}
