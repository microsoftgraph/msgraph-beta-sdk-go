package parentsection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i45bd5d4bbcce978101ddd641887e4bf5d0ebb11aafe043798797f8d6930b39c8 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/pages/item/parentsection/ref"
    i5ed731e982d63f89e60ac33b0367bd0ac2d85c07a62a3b625666e7fecbc3af13 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/pages/item/parentsection/copytosectiongroup"
    ifdd4f4c2e26f3ac68d4409e0158e92d2015bd42b664dff52914bcd0d0b97abf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/sections/item/pages/item/parentsection/copytonotebook"
)

// ParentSectionRequestBuilder builds and executes requests for operations under \sites\{site-id}\onenote\sections\{onenoteSection-id}\pages\{onenotePage-id}\parentSection
type ParentSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ParentSectionRequestBuilderGetOptions options for Get
type ParentSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ParentSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ParentSectionRequestBuilderGetQueryParameters the section that contains the page. Read-only.
type ParentSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewParentSectionRequestBuilderInternal instantiates a new ParentSectionRequestBuilder and sets the default values.
func NewParentSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    m := &ParentSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/onenote/sections/{onenoteSection_id}/pages/{onenotePage_id}/parentSection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewParentSectionRequestBuilder instantiates a new ParentSectionRequestBuilder and sets the default values.
func NewParentSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ParentSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewParentSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ParentSectionRequestBuilder) CopyToNotebook()(*ifdd4f4c2e26f3ac68d4409e0158e92d2015bd42b664dff52914bcd0d0b97abf5.CopyToNotebookRequestBuilder) {
    return ifdd4f4c2e26f3ac68d4409e0158e92d2015bd42b664dff52914bcd0d0b97abf5.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) CopyToSectionGroup()(*i5ed731e982d63f89e60ac33b0367bd0ac2d85c07a62a3b625666e7fecbc3af13.CopyToSectionGroupRequestBuilder) {
    return i5ed731e982d63f89e60ac33b0367bd0ac2d85c07a62a3b625666e7fecbc3af13.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) CreateGetRequestInformation(options *ParentSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get the section that contains the page. Read-only.
func (m *ParentSectionRequestBuilder) Get(options *ParentSectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenoteSection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection), nil
}
func (m *ParentSectionRequestBuilder) Ref()(*i45bd5d4bbcce978101ddd641887e4bf5d0ebb11aafe043798797f8d6930b39c8.RefRequestBuilder) {
    return i45bd5d4bbcce978101ddd641887e4bf5d0ebb11aafe043798797f8d6930b39c8.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
