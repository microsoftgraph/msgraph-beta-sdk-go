package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Settings struct {
    additionalData map[string]interface{};
    hasGraphMailbox *bool;
    hasLicense *bool;
    hasOptedOut *bool;
}
func NewSettings()(*Settings) {
    m := &Settings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Settings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Settings) GetHasGraphMailbox()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasGraphMailbox
    }
}
func (m *Settings) GetHasLicense()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasLicense
    }
}
func (m *Settings) GetHasOptedOut()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasOptedOut
    }
}
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
func (m *Settings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Settings) SetHasGraphMailbox(value *bool)() {
    m.hasGraphMailbox = value
}
func (m *Settings) SetHasLicense(value *bool)() {
    m.hasLicense = value
}
func (m *Settings) SetHasOptedOut(value *bool)() {
    m.hasOptedOut = value
}
