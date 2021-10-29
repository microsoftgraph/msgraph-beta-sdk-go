package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new embeddedSIMActivationCode and sets the default values.
func NewEmbeddedSIMActivationCode()(*EmbeddedSIMActivationCode) {
    m := &EmbeddedSIMActivationCode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmbeddedSIMActivationCode) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) GetIntegratedCircuitCardIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.integratedCircuitCardIdentifier
    }
}
// Gets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) GetMatchingIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.matchingIdentifier
    }
}
// Gets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) GetSmdpPlusServerAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.smdpPlusServerAddress
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EmbeddedSIMActivationCode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
// Parameters:
//  - value : Value to set for the integratedCircuitCardIdentifier property.
func (m *EmbeddedSIMActivationCode) SetIntegratedCircuitCardIdentifier(value *string)() {
    m.integratedCircuitCardIdentifier = value
}
// Sets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
// Parameters:
//  - value : Value to set for the matchingIdentifier property.
func (m *EmbeddedSIMActivationCode) SetMatchingIdentifier(value *string)() {
    m.matchingIdentifier = value
}
// Sets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
// Parameters:
//  - value : Value to set for the smdpPlusServerAddress property.
func (m *EmbeddedSIMActivationCode) SetSmdpPlusServerAddress(value *string)() {
    m.smdpPlusServerAddress = value
}
