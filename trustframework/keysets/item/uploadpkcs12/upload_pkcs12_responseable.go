package uploadpkcs12

import (
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UploadPkcs12Responseable 
type UploadPkcs12Responseable interface {
    AdditionalDataHolder
    Parsable
    GetTrustFrameworkKey()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeyable)
    SetTrustFrameworkKey(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TrustFrameworkKeyable)()
}
