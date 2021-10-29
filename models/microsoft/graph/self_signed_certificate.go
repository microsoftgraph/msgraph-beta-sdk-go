package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SelfSignedCertificate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    customKeyIdentifier []byte;
    // 
    displayName *string;
    // 
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    key []byte;
    // 
    keyId *string;
    // 
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    thumbprint *string;
    // 
    type_escaped *string;
    // 
    usage *string;
}
// Instantiates a new SelfSignedCertificate and sets the default values.
func NewSelfSignedCertificate()(*SelfSignedCertificate) {
    m := &SelfSignedCertificate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfSignedCertificate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the customKeyIdentifier property value. 
func (m *SelfSignedCertificate) GetCustomKeyIdentifier()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.customKeyIdentifier
    }
}
// Gets the displayName property value. 
func (m *SelfSignedCertificate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endDateTime property value. 
func (m *SelfSignedCertificate) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the key property value. 
func (m *SelfSignedCertificate) GetKey()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Gets the keyId property value. 
func (m *SelfSignedCertificate) GetKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyId
    }
}
// Gets the startDateTime property value. 
func (m *SelfSignedCertificate) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the thumbprint property value. 
func (m *SelfSignedCertificate) GetThumbprint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbprint
    }
}
// Gets the type_escaped property value. 
func (m *SelfSignedCertificate) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the usage property value. 
func (m *SelfSignedCertificate) GetUsage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.usage
    }
}
// The deserialization information for the current model
func (m *SelfSignedCertificate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["customKeyIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetCustomKeyIdentifier(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["keyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKeyId(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["thumbprint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbprint(val)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    res["usage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUsage(val)
        return nil
    }
    return res
}
func (m *SelfSignedCertificate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SelfSignedCertificate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("customKeyIdentifier", m.GetCustomKeyIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbprint", m.GetThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("usage", m.GetUsage())
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
func (m *SelfSignedCertificate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the customKeyIdentifier property value. 
// Parameters:
//  - value : Value to set for the customKeyIdentifier property.
func (m *SelfSignedCertificate) SetCustomKeyIdentifier(value []byte)() {
    m.customKeyIdentifier = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SelfSignedCertificate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endDateTime property value. 
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *SelfSignedCertificate) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the key property value. 
// Parameters:
//  - value : Value to set for the key property.
func (m *SelfSignedCertificate) SetKey(value []byte)() {
    m.key = value
}
// Sets the keyId property value. 
// Parameters:
//  - value : Value to set for the keyId property.
func (m *SelfSignedCertificate) SetKeyId(value *string)() {
    m.keyId = value
}
// Sets the startDateTime property value. 
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *SelfSignedCertificate) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the thumbprint property value. 
// Parameters:
//  - value : Value to set for the thumbprint property.
func (m *SelfSignedCertificate) SetThumbprint(value *string)() {
    m.thumbprint = value
}
// Sets the type_escaped property value. 
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *SelfSignedCertificate) SetType_escaped(value *string)() {
    m.type_escaped = value
}
// Sets the usage property value. 
// Parameters:
//  - value : Value to set for the usage property.
func (m *SelfSignedCertificate) SetUsage(value *string)() {
    m.usage = value
}
