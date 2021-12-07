package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EmbeddedSIMActivationCode 
type EmbeddedSIMActivationCode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
    integratedCircuitCardIdentifier *string;
    // The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
    matchingIdentifier *string;
    // The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
    smdpPlusServerAddress *string;
}
// NewEmbeddedSIMActivationCode instantiates a new embeddedSIMActivationCode and sets the default values.
func NewEmbeddedSIMActivationCode()(*EmbeddedSIMActivationCode) {
    m := &EmbeddedSIMActivationCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmbeddedSIMActivationCode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIntegratedCircuitCardIdentifier gets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) GetIntegratedCircuitCardIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.integratedCircuitCardIdentifier
    }
}
// GetMatchingIdentifier gets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) GetMatchingIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.matchingIdentifier
    }
}
// GetSmdpPlusServerAddress gets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) GetSmdpPlusServerAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.smdpPlusServerAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCode) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["integratedCircuitCardIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegratedCircuitCardIdentifier(val)
        }
        return nil
    }
    res["matchingIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingIdentifier(val)
        }
        return nil
    }
    res["smdpPlusServerAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmdpPlusServerAddress(val)
        }
        return nil
    }
    return res
}
func (m *EmbeddedSIMActivationCode) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmbeddedSIMActivationCode) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIntegratedCircuitCardIdentifier sets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) SetIntegratedCircuitCardIdentifier(value *string)() {
    if m != nil {
        m.integratedCircuitCardIdentifier = value
    }
}
// SetMatchingIdentifier sets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) SetMatchingIdentifier(value *string)() {
    if m != nil {
        m.matchingIdentifier = value
    }
}
// SetSmdpPlusServerAddress sets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) SetSmdpPlusServerAddress(value *string)() {
    if m != nil {
        m.smdpPlusServerAddress = value
    }
}
