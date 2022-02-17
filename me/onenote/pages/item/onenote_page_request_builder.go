package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1daaf4c3ecf2363d24cd0e200d9cedac4f3fb229ddd7776ad17ed3b5fbe2f8ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item/parentnotebook"
    i402753b0167f1a0e92645e87797f47d23e4cfcb5f71767a419f3878578b8018f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item/preview"
    i451c025aa220beeb4fb51b8b093fc5efa7d4e7f52775253eb60e4c5b10a61795 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item/content"
    i5a522d346490d892f5e57f163d1de63ef54efe1de0506b59fae67f6b8cc3c5e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item/copytosection"
    ia74f7418d2a1374784c794d0129acca008734f23015399b132b563e6375baef5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item/onenotepatchcontent"
    if3f03d5b402a9dd599db6efd0076b73a02afef311f53fa655dcc578640033ca7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item/parentsection"
)

// OnenotePageRequestBuilder builds and executes requests for operations under \me\onenote\pages\{onenotePage-id}
type OnenotePageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenotePageRequestBuilderDeleteOptions options for Delete
type OnenotePageRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageRequestBuilderGetOptions options for Get
type OnenotePageRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenotePageRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenotePageRequestBuilderGetQueryParameters the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type OnenotePageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenotePageRequestBuilderPatchOptions options for Patch
type OnenotePageRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenotePageRequestBuilderInternal instantiates a new OnenotePageRequestBuilder and sets the default values.
func NewOnenotePageRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    m := &OnenotePageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/pages/{onenotePage_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenotePageRequestBuilder instantiates a new OnenotePageRequestBuilder and sets the default values.
func NewOnenotePageRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenotePageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenotePageRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenotePageRequestBuilder) Content()(*i451c025aa220beeb4fb51b8b093fc5efa7d4e7f52775253eb60e4c5b10a61795.ContentRequestBuilder) {
    return i451c025aa220beeb4fb51b8b093fc5efa7d4e7f52775253eb60e4c5b10a61795.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) CopyToSection()(*i5a522d346490d892f5e57f163d1de63ef54efe1de0506b59fae67f6b8cc3c5e0.CopyToSectionRequestBuilder) {
    return i5a522d346490d892f5e57f163d1de63ef54efe1de0506b59fae67f6b8cc3c5e0.NewCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreateDeleteRequestInformation(options *OnenotePageRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreateGetRequestInformation(options *OnenotePageRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) CreatePatchRequestInformation(options *OnenotePageRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Delete(options *OnenotePageRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Get(options *OnenotePageRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenotePage() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenotePage), nil
}
func (m *OnenotePageRequestBuilder) OnenotePatchContent()(*ia74f7418d2a1374784c794d0129acca008734f23015399b132b563e6375baef5.OnenotePatchContentRequestBuilder) {
    return ia74f7418d2a1374784c794d0129acca008734f23015399b132b563e6375baef5.NewOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentNotebook()(*i1daaf4c3ecf2363d24cd0e200d9cedac4f3fb229ddd7776ad17ed3b5fbe2f8ad.ParentNotebookRequestBuilder) {
    return i1daaf4c3ecf2363d24cd0e200d9cedac4f3fb229ddd7776ad17ed3b5fbe2f8ad.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenotePageRequestBuilder) ParentSection()(*if3f03d5b402a9dd599db6efd0076b73a02afef311f53fa655dcc578640033ca7.ParentSectionRequestBuilder) {
    return if3f03d5b402a9dd599db6efd0076b73a02afef311f53fa655dcc578640033ca7.NewParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the pages in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenotePageRequestBuilder) Patch(options *OnenotePageRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Preview builds and executes requests for operations under \me\onenote\pages\{onenotePage-id}\microsoft.graph.preview()
func (m *OnenotePageRequestBuilder) Preview()(*i402753b0167f1a0e92645e87797f47d23e4cfcb5f71767a419f3878578b8018f.PreviewRequestBuilder) {
    return i402753b0167f1a0e92645e87797f47d23e4cfcb5f71767a419f3878578b8018f.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
