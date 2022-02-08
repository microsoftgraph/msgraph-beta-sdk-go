package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnPremisesPublishingSingleSignOn 
type OnPremisesPublishingSingleSignOn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
    kerberosSignOnSettings *KerberosSignOnSettings;
    // The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
    singleSignOnMode *SingleSignOnMode;
}
// NewOnPremisesPublishingSingleSignOn instantiates a new onPremisesPublishingSingleSignOn and sets the default values.
func NewOnPremisesPublishingSingleSignOn()(*OnPremisesPublishingSingleSignOn) {
    m := &OnPremisesPublishingSingleSignOn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishingSingleSignOn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKerberosSignOnSettings gets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) GetKerberosSignOnSettings()(*KerberosSignOnSettings) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnSettings
    }
}
// GetSingleSignOnMode gets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) GetSingleSignOnMode()(*SingleSignOnMode) {
    if m == nil {
        return nil
    } else {
        return m.singleSignOnMode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesPublishingSingleSignOn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["kerberosSignOnSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKerberosSignOnSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnSettings(val.(*KerberosSignOnSettings))
        }
        return nil
    }
    res["singleSignOnMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSingleSignOnMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleSignOnMode(val.(*SingleSignOnMode))
        }
        return nil
    }
    return res
}
func (m *OnPremisesPublishingSingleSignOn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OnPremisesPublishingSingleSignOn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("kerberosSignOnSettings", m.GetKerberosSignOnSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSingleSignOnMode() != nil {
        cast := (*m.GetSingleSignOnMode()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesPublishingSingleSignOn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKerberosSignOnSettings sets the kerberosSignOnSettings property value. The Kerberos Constrained Delegation settings for applications that use Integrated Window Authentication.
func (m *OnPremisesPublishingSingleSignOn) SetKerberosSignOnSettings(value *KerberosSignOnSettings)() {
    if m != nil {
        m.kerberosSignOnSettings = value
    }
}
// SetSingleSignOnMode sets the singleSignOnMode property value. The preferred single-sign on mode for the application. Possible values are: none, onPremisesKerberos, aadHeaderBased,pingHeaderBased.
func (m *OnPremisesPublishingSingleSignOn) SetSingleSignOnMode(value *SingleSignOnMode)() {
    if m != nil {
        m.singleSignOnMode = value
    }
}
