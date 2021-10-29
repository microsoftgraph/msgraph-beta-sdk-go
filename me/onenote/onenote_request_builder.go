package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i02442f2b4528552778554b6a8d7734fcbcc5c67d05fb38c736fd39e10a087b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i21e97cc531e79994ac8f68134cbea3aa92a45effb9959e592cb197af924ec755 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sections"
    i2a34cadbfd73c7d67b9945eae7836418199869d09723bc6ea8f474c361c69476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups"
    i3f6844fa2bdd594acff5a9568bb940d97aee31c03de9263680db2cafec02cd32 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/resources"
    i4d5248421f500fed59a4443906c41e5156b79d60a5dc7eca83763c99b773c741 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    id00a24b6b18f3cd1dcf3862b4f8cc4c98ee4ef1720eea18826d911d1150cf228 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/operations"
    i08bb5592a459ace4070591d9b4083a115e728de9d6fb3517918f3f661f99d4be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item"
    i31303fd26db642eb4e6f7dc00a9babd052efccdd57a191a012a948d9d636e966 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item"
    i7b6a7c655d49ef24caf532d5a7f363e6f1bff778a37ddde30a9ee2aa9a699820 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/operations/item"
    i96a102edf759f9f57dbf2fb12c37a9dd630d0fe2d5f0b61562e23013eb3c9897 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/resources/item"
    ic8c068b58107eb1b9b3bc4b7e43c090c8f2e1ddd8871f128b4089d6f6755e40a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item"
    if48a948123f8134534ad90c988dad247d9e1c7a92254ec0abf9a7a7c5e37a20c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sections/item"
)

// Builds and executes requests for operations under \me\onenote
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type OnenoteRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type OnenoteRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new OnenoteRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new OnenoteRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(options *OnenoteRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(options *OnenoteRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenote() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*i4d5248421f500fed59a4443906c41e5156b79d60a5dc7eca83763c99b773c741.NotebooksRequestBuilder) {
    return i4d5248421f500fed59a4443906c41e5156b79d60a5dc7eca83763c99b773c741.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onenote.notebooks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*ic8c068b58107eb1b9b3bc4b7e43c090c8f2e1ddd8871f128b4089d6f6755e40a.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return ic8c068b58107eb1b9b3bc4b7e43c090c8f2e1ddd8871f128b4089d6f6755e40a.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*id00a24b6b18f3cd1dcf3862b4f8cc4c98ee4ef1720eea18826d911d1150cf228.OperationsRequestBuilder) {
    return id00a24b6b18f3cd1dcf3862b4f8cc4c98ee4ef1720eea18826d911d1150cf228.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onenote.operations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i7b6a7c655d49ef24caf532d5a7f363e6f1bff778a37ddde30a9ee2aa9a699820.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i7b6a7c655d49ef24caf532d5a7f363e6f1bff778a37ddde30a9ee2aa9a699820.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*i02442f2b4528552778554b6a8d7734fcbcc5c67d05fb38c736fd39e10a087b4e.PagesRequestBuilder) {
    return i02442f2b4528552778554b6a8d7734fcbcc5c67d05fb38c736fd39e10a087b4e.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onenote.pages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) PagesById(id string)(*i31303fd26db642eb4e6f7dc00a9babd052efccdd57a191a012a948d9d636e966.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i31303fd26db642eb4e6f7dc00a9babd052efccdd57a191a012a948d9d636e966.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Read-only.
// Parameters:
//  - options : Options for the request
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnenoteRequestBuilder) Resources()(*i3f6844fa2bdd594acff5a9568bb940d97aee31c03de9263680db2cafec02cd32.ResourcesRequestBuilder) {
    return i3f6844fa2bdd594acff5a9568bb940d97aee31c03de9263680db2cafec02cd32.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onenote.resources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i96a102edf759f9f57dbf2fb12c37a9dd630d0fe2d5f0b61562e23013eb3c9897.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return i96a102edf759f9f57dbf2fb12c37a9dd630d0fe2d5f0b61562e23013eb3c9897.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*i2a34cadbfd73c7d67b9945eae7836418199869d09723bc6ea8f474c361c69476.SectionGroupsRequestBuilder) {
    return i2a34cadbfd73c7d67b9945eae7836418199869d09723bc6ea8f474c361c69476.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onenote.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i08bb5592a459ace4070591d9b4083a115e728de9d6fb3517918f3f661f99d4be.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i08bb5592a459ace4070591d9b4083a115e728de9d6fb3517918f3f661f99d4be.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*i21e97cc531e79994ac8f68134cbea3aa92a45effb9959e592cb197af924ec755.SectionsRequestBuilder) {
    return i21e97cc531e79994ac8f68134cbea3aa92a45effb9959e592cb197af924ec755.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onenote.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *OnenoteRequestBuilder) SectionsById(id string)(*if48a948123f8134534ad90c988dad247d9e1c7a92254ec0abf9a7a7c5e37a20c.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return if48a948123f8134534ad90c988dad247d9e1c7a92254ec0abf9a7a7c5e37a20c.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
