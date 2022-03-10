package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchSessionable 
type ExactMatchSessionable interface {
    ExactMatchSessionBaseable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetChecksum()(*string)
    GetDataUploadURI()(*string)
    GetFields()([]string)
    GetFileName()(*string)
    GetRowsPerBlock()(*int32)
    GetSalt()(*string)
    GetUploadAgent()(ExactMatchUploadAgentable)
    GetUploadAgentId()(*string)
    SetChecksum(value *string)()
    SetDataUploadURI(value *string)()
    SetFields(value []string)()
    SetFileName(value *string)()
    SetRowsPerBlock(value *int32)()
    SetSalt(value *string)()
    SetUploadAgent(value ExactMatchUploadAgentable)()
    SetUploadAgentId(value *string)()
}
