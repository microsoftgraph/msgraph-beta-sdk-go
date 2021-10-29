package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Settings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
    hasGraphMailbox *bool;
    // Specifies if the user has a MyAnalytics license assigned.
    hasLicense *bool;
    // Specifies if the user opted out of MyAnalytics.
    hasOptedOut *bool;
}
// Instantiates a new settings and sets the default values.
func NewSettings()(*Settings) {
    m := &Settings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Settings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the hasGraphMailbox property value. Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
func (m *Settings) GetHasGraphMailbox()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasGraphMailbox
    }
}
// Gets the hasLicense property value. Specifies if the user has a MyAnalytics license assigned.
func (m *Settings) GetHasLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasLicense
    }
}
// Gets the hasOptedOut property value. Specifies if the user opted out of MyAnalytics.
func (m *Settings) GetHasOptedOut()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOptedOut
    }
}
// The deserialization information for the current model
func (m *Settings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hasGraphMailbox"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasGraphMailbox(val)
        return nil
    }
    res["hasLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasLicense(val)
        return nil
    }
    res["hasOptedOut"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasOptedOut(val)
        return nil
    }
    return res
}
func (m *Settings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Settings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hasGraphMailbox", m.GetHasGraphMailbox())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasLicense", m.GetHasLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasOptedOut", m.GetHasOptedOut())
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
func (m *Settings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the hasGraphMailbox property value. Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
// Parameters:
//  - value : Value to set for the hasGraphMailbox property.
func (m *Settings) SetHasGraphMailbox(value *bool)() {
    m.hasGraphMailbox = value
}
// Sets the hasLicense property value. Specifies if the user has a MyAnalytics license assigned.
// Parameters:
//  - value : Value to set for the hasLicense property.
func (m *Settings) SetHasLicense(value *bool)() {
    m.hasLicense = value
}
// Sets the hasOptedOut property value. Specifies if the user opted out of MyAnalytics.
// Parameters:
//  - value : Value to set for the hasOptedOut property.
func (m *Settings) SetHasOptedOut(value *bool)() {
    m.hasOptedOut = value
}
