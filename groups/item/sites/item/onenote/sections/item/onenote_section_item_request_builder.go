package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i2743e9948d2e2e087c8d670304dbe53f3382bdeea24ad163ed26809ec18af534 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item/parentsectiongroup"
    i5511861bd51eb3546ac294e8f9a383c464c8abb355b1d9611c4cc24834cf1bb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item/pages"
    i610562aac1ad844650fc6efa744f9ede4702d43144de5c8b4730e67f174393cb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item/copytonotebook"
    i8d5e6d9d540dba8d2cb6bd667b3861d205775760ef29f05e77a4b013c5e73cbe "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item/copytosectiongroup"
    ife51913e77e172b54d9ca822c8dc9205f924b785629cf54d362b3f2952b58134 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item/parentnotebook"
    i8a267b3363b37965398a914920dcc1a776f4a0faffe4fc688a93a1ee255de191 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item/pages/item"
)

// OnenoteSectionItemRequestBuilder provides operations to manage the sections property of the microsoft.graph.onenote entity.
type OnenoteSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionItemRequestBuilderDeleteOptions options for Delete
type OnenoteSectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetOptions options for Get
type OnenoteSectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionItemRequestBuilderGetQueryParameters the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
type OnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionItemRequestBuilderPatchOptions options for Patch
type OnenoteSectionItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSectionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    m := &OnenoteSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/onenote/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionItemRequestBuilder) CopyToNotebook()(*i610562aac1ad844650fc6efa744f9ede4702d43144de5c8b4730e67f174393cb.CopyToNotebookRequestBuilder) {
    return i610562aac1ad844650fc6efa744f9ede4702d43144de5c8b4730e67f174393cb.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*i8d5e6d9d540dba8d2cb6bd667b3861d205775760ef29f05e77a4b013c5e73cbe.CopyToSectionGroupRequestBuilder) {
    return i8d5e6d9d540dba8d2cb6bd667b3861d205775760ef29f05e77a4b013c5e73cbe.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sections for groups
func (m *OnenoteSectionItemRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sections in groups
func (m *OnenoteSectionItemRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property sections for groups
func (m *OnenoteSectionItemRequestBuilder) Delete(options *OnenoteSectionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the sections in all OneNote notebooks that are owned by the user or group.  Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Get(options *OnenoteSectionItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOnenoteSectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSectionable), nil
}
func (m *OnenoteSectionItemRequestBuilder) Pages()(*i5511861bd51eb3546ac294e8f9a383c464c8abb355b1d9611c4cc24834cf1bb4.PagesRequestBuilder) {
    return i5511861bd51eb3546ac294e8f9a383c464c8abb355b1d9611c4cc24834cf1bb4.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.sections.item.pages.item collection
func (m *OnenoteSectionItemRequestBuilder) PagesById(id string)(*i8a267b3363b37965398a914920dcc1a776f4a0faffe4fc688a93a1ee255de191.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i8a267b3363b37965398a914920dcc1a776f4a0faffe4fc688a93a1ee255de191.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) ParentNotebook()(*ife51913e77e172b54d9ca822c8dc9205f924b785629cf54d362b3f2952b58134.ParentNotebookRequestBuilder) {
    return ife51913e77e172b54d9ca822c8dc9205f924b785629cf54d362b3f2952b58134.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionItemRequestBuilder) ParentSectionGroup()(*i2743e9948d2e2e087c8d670304dbe53f3382bdeea24ad163ed26809ec18af534.ParentSectionGroupRequestBuilder) {
    return i2743e9948d2e2e087c8d670304dbe53f3382bdeea24ad163ed26809ec18af534.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sections in groups
func (m *OnenoteSectionItemRequestBuilder) Patch(options *OnenoteSectionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
