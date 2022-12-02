package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBodyable 
type ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureBlobContainer()(*string)
    GetAzureBlobToken()(*string)
    GetDescription()(*string)
    GetOutputName()(*string)
    SetAzureBlobContainer(value *string)()
    SetAzureBlobToken(value *string)()
    SetDescription(value *string)()
    SetOutputName(value *string)()
}
