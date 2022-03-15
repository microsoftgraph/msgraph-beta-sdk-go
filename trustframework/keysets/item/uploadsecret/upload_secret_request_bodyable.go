package uploadsecret

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UploadSecretRequestBodyable 
type UploadSecretRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExp()(*int32)
    GetK()(*string)
    GetNbf()(*int32)
    GetUse()(*string)
    SetExp(value *int32)()
    SetK(value *string)()
    SetNbf(value *int32)()
    SetUse(value *string)()
}
