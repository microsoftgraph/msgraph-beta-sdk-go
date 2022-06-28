package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchSessionable 
type ExactMatchSessionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ExactMatchSessionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
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
