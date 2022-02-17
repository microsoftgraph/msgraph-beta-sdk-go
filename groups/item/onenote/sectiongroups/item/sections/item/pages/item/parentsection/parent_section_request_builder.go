package parentsection

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i425cfeae4a3cf0636113ce8b181be2722ef7324c025c460fce1d3e1376c0c10a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sectiongroups/item/sections/item/pages/item/parentsection/copytosectiongroup"
    i48f68d00460058bb822da1bc795f9b5b79e45e739956fb8c10933bd787f8ecb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sectiongroups/item/sections/item/pages/item/parentsection/ref"
    ifcefbf911839944609fbd8d9bb2b130182f21a416b527a0e65563bd283b96d57 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sectiongroups/item/sections/item/pages/item/parentsection/copytonotebook"
)

// ParentSectionRequestBuilder builds and executes requests for operations under \groups\{group-id}\onenote\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}\pages\{onenotePage-id}\parentSection
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
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}/pages/{onenotePage_id}/parentSection{?select,expand}";
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
func (m *ParentSectionRequestBuilder) CopyToNotebook()(*ifcefbf911839944609fbd8d9bb2b130182f21a416b527a0e65563bd283b96d57.CopyToNotebookRequestBuilder) {
    return ifcefbf911839944609fbd8d9bb2b130182f21a416b527a0e65563bd283b96d57.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ParentSectionRequestBuilder) CopyToSectionGroup()(*i425cfeae4a3cf0636113ce8b181be2722ef7324c025c460fce1d3e1376c0c10a.CopyToSectionGroupRequestBuilder) {
    return i425cfeae4a3cf0636113ce8b181be2722ef7324c025c460fce1d3e1376c0c10a.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ParentSectionRequestBuilder) Ref()(*i48f68d00460058bb822da1bc795f9b5b79e45e739956fb8c10933bd787f8ecb2.RefRequestBuilder) {
    return i48f68d00460058bb822da1bc795f9b5b79e45e739956fb8c10933bd787f8ecb2.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
