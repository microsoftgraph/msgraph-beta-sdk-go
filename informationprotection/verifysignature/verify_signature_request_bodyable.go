package verifysignature

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VerifySignatureRequestBodyable 
type VerifySignatureRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDigest()([]byte)
    GetSignature()([]byte)
    GetSigningKeyId()(*string)
    SetDigest(value []byte)()
    SetSignature(value []byte)()
    SetSigningKeyId(value *string)()
}
