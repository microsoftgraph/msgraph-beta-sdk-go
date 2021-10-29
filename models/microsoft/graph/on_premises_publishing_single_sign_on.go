package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnPremisesPublishingSingleSignOn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
    kerberosSignOnSettings *KerberosSignOnSettings;
    // The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
    singleSignOnMode *SingleSignOnMode;
}
// Instantiates a new onPremisesPublishingSingleSignOn and sets the default values.
func NewOnPremisesPublishingSingleSignOn()(*OnPremisesPublishingSingleSignOn) {
    m := &OnPremisesPublishingSingleSignOn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishingSingleSignOn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) GetKerberosSignOnSettings()(*KerberosSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnSettings
    }
}
// Gets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) GetSingleSignOnMode()(*SingleSignOnMode) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnMode
    }
}
// The deserialization information for the current model
func (m *OnPremisesPublishingSingleSignOn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["kerberosSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKerberosSignOnSettings() })
        if err != nil {
            return err
        }
        m.SetKerberosSignOnSettings(val.(*KerberosSignOnSettings))
        return nil
    }
    res["singleSignOnMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSingleSignOnMode)
        if err != nil {
            return err
        }
        cast := val.(SingleSignOnMode)
        m.SetSingleSignOnMode(&cast)
        return nil
    }
    return res
}
func (m *OnPremisesPublishingSingleSignOn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnPremisesPublishingSingleSignOn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSingleSignOnMode() != nil {
        cast := m.GetSingleSignOnMode().String()
        err := writer.WriteStringValue("singleSignOnMode", &cast)
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
func (m *OnPremisesPublishingSingleSignOn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
// Parameters:
//  - value : Value to set for the kerberosSignOnSettings property.
func (m *OnPremisesPublishingSingleSignOn) SetKerberosSignOnSettings(value *KerberosSignOnSettings)() {
    m.kerberosSignOnSettings = value
}
// Sets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
// Parameters:
//  - value : Value to set for the singleSignOnMode property.
func (m *OnPremisesPublishingSingleSignOn) SetSingleSignOnMode(value *SingleSignOnMode)() {
    m.singleSignOnMode = value
}
