package classifytext

import (
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ClassifyTextResponseable 
type ClassifyTextResponseable interface {
    AdditionalDataHolder
    Parsable
    GetClassificationJobResponse()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ClassificationJobResponseable)
    SetClassificationJobResponse(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ClassificationJobResponseable)()
}
