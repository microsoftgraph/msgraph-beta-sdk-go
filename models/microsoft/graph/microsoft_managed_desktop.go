package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftManagedDesktop 
type MicrosoftManagedDesktop struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    profile *string;
    // 
    type *MicrosoftManagedDesktopType;
}
// NewMicrosoftManagedDesktop instantiates a new microsoftManagedDesktop and sets the default values.
func NewMicrosoftManagedDesktop()(*MicrosoftManagedDesktop) {
    m := &MicrosoftManagedDesktop{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftManagedDesktop) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetProfile gets the profile property value. 
func (m *MicrosoftManagedDesktop) GetProfile()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profile
    }
}
// GetType gets the type property value. 
func (m *MicrosoftManagedDesktop) GetType()(*MicrosoftManagedDesktopType) {
    if m == nil {
        return nil
    } else {
        return m.type
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftManagedDesktop) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["profile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftManagedDesktopType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MicrosoftManagedDesktopType)
            m.SetType(&cast)
        }
        return nil
    }
    return res
}
func (m *MicrosoftManagedDesktop) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MicrosoftManagedDesktop) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := m.GetType().String()
        err := writer.WriteStringValue("type", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftManagedDesktop) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetProfile sets the profile property value. 
func (m *MicrosoftManagedDesktop) SetProfile(value *string)() {
    if m != nil {
        m.profile = value
    }
}
// SetType sets the type property value. 
func (m *MicrosoftManagedDesktop) SetType(value *MicrosoftManagedDesktopType)() {
    if m != nil {
        m.type = value
    }
}
