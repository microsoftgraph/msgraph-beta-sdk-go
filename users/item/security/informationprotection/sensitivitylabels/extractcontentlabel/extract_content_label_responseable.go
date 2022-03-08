package extractcontentlabel

import (
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ExtractContentLabelResponseable 
type ExtractContentLabelResponseable interface {
    AdditionalDataHolder
    Parsable
    GetContentLabel()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentLabelable)
    SetContentLabel(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ContentLabelable)()
}
