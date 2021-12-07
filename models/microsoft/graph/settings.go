package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Settings 
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
// NewSettings instantiates a new settings and sets the default values.
func NewSettings()(*Settings) {
    m := &Settings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Settings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHasGraphMailbox gets the hasGraphMailbox property value. Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
func (m *Settings) GetHasGraphMailbox()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasGraphMailbox
    }
}
// GetHasLicense gets the hasLicense property value. Specifies if the user has a MyAnalytics license assigned.
func (m *Settings) GetHasLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasLicense
    }
}
// GetHasOptedOut gets the hasOptedOut property value. Specifies if the user opted out of MyAnalytics.
func (m *Settings) GetHasOptedOut()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOptedOut
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Settings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hasGraphMailbox"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasGraphMailbox(val)
        }
        return nil
    }
    res["hasLicense"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasLicense(val)
        }
        return nil
    }
    res["hasOptedOut"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOptedOut(val)
        }
        return nil
    }
    return res
}
func (m *Settings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Settings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHasGraphMailbox sets the hasGraphMailbox property value. Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
func (m *Settings) SetHasGraphMailbox(value *bool)() {
    if m != nil {
        m.hasGraphMailbox = value
    }
}
// SetHasLicense sets the hasLicense property value. Specifies if the user has a MyAnalytics license assigned.
func (m *Settings) SetHasLicense(value *bool)() {
    if m != nil {
        m.hasLicense = value
    }
}
// SetHasOptedOut sets the hasOptedOut property value. Specifies if the user opted out of MyAnalytics.
func (m *Settings) SetHasOptedOut(value *bool)() {
    if m != nil {
        m.hasOptedOut = value
    }
}
