package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonNamePronounciation 
type PersonNamePronounciation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    displayName *string;
    // 
    first *string;
    // 
    last *string;
    // 
    maiden *string;
    // 
    middle *string;
}
// NewPersonNamePronounciation instantiates a new personNamePronounciation and sets the default values.
func NewPersonNamePronounciation()(*PersonNamePronounciation) {
    m := &PersonNamePronounciation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonNamePronounciation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. 
func (m *PersonNamePronounciation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFirst gets the first property value. 
func (m *PersonNamePronounciation) GetFirst()(*string) {
    if m == nil {
        return nil
    } else {
        return m.first
    }
}
// GetLast gets the last property value. 
func (m *PersonNamePronounciation) GetLast()(*string) {
    if m == nil {
        return nil
    } else {
        return m.last
    }
}
// GetMaiden gets the maiden property value. 
func (m *PersonNamePronounciation) GetMaiden()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maiden
    }
}
// GetMiddle gets the middle property value. 
func (m *PersonNamePronounciation) GetMiddle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middle
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonNamePronounciation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["first"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirst(val)
        }
        return nil
    }
    res["last"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLast(val)
        }
        return nil
    }
    res["maiden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaiden(val)
        }
        return nil
    }
    res["middle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddle(val)
        }
        return nil
    }
    return res
}
func (m *PersonNamePronounciation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonNamePronounciation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *PersonNamePronounciation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFirst sets the first property value. 
func (m *PersonNamePronounciation) SetFirst(value *string)() {
    if m != nil {
        m.first = value
    }
}
// SetLast sets the last property value. 
func (m *PersonNamePronounciation) SetLast(value *string)() {
    if m != nil {
        m.last = value
    }
}
// SetMaiden sets the maiden property value. 
func (m *PersonNamePronounciation) SetMaiden(value *string)() {
    if m != nil {
        m.maiden = value
    }
}
// SetMiddle sets the middle property value. 
func (m *PersonNamePronounciation) SetMiddle(value *string)() {
    if m != nil {
        m.middle = value
    }
}
