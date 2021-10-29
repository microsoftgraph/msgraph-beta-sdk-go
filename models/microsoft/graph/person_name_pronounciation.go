package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new personNamePronounciation and sets the default values.
func NewPersonNamePronounciation()(*PersonNamePronounciation) {
    m := &PersonNamePronounciation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonNamePronounciation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. 
func (m *PersonNamePronounciation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the first property value. 
func (m *PersonNamePronounciation) GetFirst()(*string) {
    if m == nil {
        return nil
    } else {
        return m.first
    }
}
// Gets the last property value. 
func (m *PersonNamePronounciation) GetLast()(*string) {
    if m == nil {
        return nil
    } else {
        return m.last
    }
}
// Gets the maiden property value. 
func (m *PersonNamePronounciation) GetMaiden()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maiden
    }
}
// Gets the middle property value. 
func (m *PersonNamePronounciation) GetMiddle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middle
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PersonNamePronounciation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonNamePronounciation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the first property value. 
// Parameters:
//  - value : Value to set for the first property.
func (m *PersonNamePronounciation) SetFirst(value *string)() {
    m.first = value
}
// Sets the last property value. 
// Parameters:
//  - value : Value to set for the last property.
func (m *PersonNamePronounciation) SetLast(value *string)() {
    m.last = value
}
// Sets the maiden property value. 
// Parameters:
//  - value : Value to set for the maiden property.
func (m *PersonNamePronounciation) SetMaiden(value *string)() {
    m.maiden = value
}
// Sets the middle property value. 
// Parameters:
//  - value : Value to set for the middle property.
func (m *PersonNamePronounciation) SetMiddle(value *string)() {
    m.middle = value
}
