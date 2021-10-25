package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EmbeddedSIMActivationCode struct {
    additionalData map[string]interface{};
    integratedCircuitCardIdentifier *string;
    matchingIdentifier *string;
    smdpPlusServerAddress *string;
}
func NewEmbeddedSIMActivationCode()(*EmbeddedSIMActivationCode) {
    m := &EmbeddedSIMActivationCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EmbeddedSIMActivationCode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EmbeddedSIMActivationCode) GetIntegratedCircuitCardIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.integratedCircuitCardIdentifier
    }
}
func (m *EmbeddedSIMActivationCode) GetMatchingIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.matchingIdentifier
    }
}
func (m *EmbeddedSIMActivationCode) GetSmdpPlusServerAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.smdpPlusServerAddress
    }
}
func (m *EmbeddedSIMActivationCode) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["integratedCircuitCardIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIntegratedCircuitCardIdentifier(val)
        return nil
    }
    res["matchingIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMatchingIdentifier(val)
        return nil
    }
    res["smdpPlusServerAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSmdpPlusServerAddress(val)
        return nil
    }
    return res
}
func (m *EmbeddedSIMActivationCode) IsNil()(bool) {
    return m == nil
}
func (m *EmbeddedSIMActivationCode) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("integratedCircuitCardIdentifier", m.GetIntegratedCircuitCardIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("matchingIdentifier", m.GetMatchingIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("smdpPlusServerAddress", m.GetSmdpPlusServerAddress())
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
func (m *EmbeddedSIMActivationCode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EmbeddedSIMActivationCode) SetIntegratedCircuitCardIdentifier(value *string)() {
    m.integratedCircuitCardIdentifier = value
}
func (m *EmbeddedSIMActivationCode) SetMatchingIdentifier(value *string)() {
    m.matchingIdentifier = value
}
func (m *EmbeddedSIMActivationCode) SetSmdpPlusServerAddress(value *string)() {
    m.smdpPlusServerAddress = value
}
