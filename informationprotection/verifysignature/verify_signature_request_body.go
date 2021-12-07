package verifysignature

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VerifySignatureRequestBody 
type VerifySignatureRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    digest []byte;
    // 
    signature []byte;
    // 
    signingKeyId *string;
}
// NewVerifySignatureRequestBody instantiates a new verifySignatureRequestBody and sets the default values.
func NewVerifySignatureRequestBody()(*VerifySignatureRequestBody) {
    m := &VerifySignatureRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifySignatureRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDigest gets the digest property value. 
func (m *VerifySignatureRequestBody) GetDigest()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.digest
    }
}
// GetSignature gets the signature property value. 
func (m *VerifySignatureRequestBody) GetSignature()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.signature
    }
}
// GetSigningKeyId gets the signingKeyId property value. 
func (m *VerifySignatureRequestBody) GetSigningKeyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.signingKeyId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifySignatureRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["digest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
    res["signature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val)
        }
        return nil
    }
    res["signingKeyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningKeyId(val)
        }
        return nil
    }
    return res
}
func (m *VerifySignatureRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *VerifySignatureRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("signature", m.GetSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("signingKeyId", m.GetSigningKeyId())
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
func (m *VerifySignatureRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDigest sets the digest property value. 
func (m *VerifySignatureRequestBody) SetDigest(value []byte)() {
    if m != nil {
        m.digest = value
    }
}
// SetSignature sets the signature property value. 
func (m *VerifySignatureRequestBody) SetSignature(value []byte)() {
    if m != nil {
        m.signature = value
    }
}
// SetSigningKeyId sets the signingKeyId property value. 
func (m *VerifySignatureRequestBody) SetSigningKeyId(value *string)() {
    if m != nil {
        m.signingKeyId = value
    }
}
