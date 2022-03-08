package verifysignature

import (
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// VerifySignatureResponseable 
type VerifySignatureResponseable interface {
    AdditionalDataHolder
    Parsable
    GetVerificationResult()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VerificationResultable)
    SetVerificationResult(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.VerificationResultable)()
}
