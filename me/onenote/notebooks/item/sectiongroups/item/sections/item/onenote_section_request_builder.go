package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i446908705c1ddee4882afd2e4611611621abaa084248970a1b001f9a4c3f93c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item/sectiongroups/item/sections/item/copytosectiongroup"
    i6b88ea813bf4804160ecb63495a5d65b7c5c3d81998515d99dc1b8189593fa74 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item/sectiongroups/item/sections/item/parentnotebook"
    i873554b7b847d1d15d4a54e90f68c2d1e8e0e6c46101a06a7008eedc4dd8cb63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item/sectiongroups/item/sections/item/copytonotebook"
    ic50e85cfbfe1e68ccdf2beec3955ddd7c2ce5ebb6956cb05fa5a04e731b2190a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item/sectiongroups/item/sections/item/parentsectiongroup"
    if33c50562d8903d2dc99274fc537bcb09ad5988398de7108aa512b0390c432e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item/sectiongroups/item/sections/item/pages"
    i53277344ee337bfe414f01155cbb7c88e6b1d015b2391b56e0cba731399ceeda "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item/sectiongroups/item/sections/item/pages/item"
)

// OnenoteSectionRequestBuilder builds and executes requests for operations under \me\onenote\notebooks\{notebook-id}\sectionGroups\{sectionGroup-id}\sections\{onenoteSection-id}
type OnenoteSectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteSectionRequestBuilderDeleteOptions options for Delete
type OnenoteSectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionRequestBuilderGetOptions options for Get
type OnenoteSectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnenoteSectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteSectionRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type OnenoteSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteSectionRequestBuilderPatchOptions options for Patch
type OnenoteSectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewOnenoteSectionRequestBuilderInternal instantiates a new OnenoteSectionRequestBuilder and sets the default values.
func NewOnenoteSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    m := &OnenoteSectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote/notebooks/{notebook_id}/sectionGroups/{sectionGroup_id}/sections/{onenoteSection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionRequestBuilder instantiates a new OnenoteSectionRequestBuilder and sets the default values.
func NewOnenoteSectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *OnenoteSectionRequestBuilder) CopyToNotebook()(*i873554b7b847d1d15d4a54e90f68c2d1e8e0e6c46101a06a7008eedc4dd8cb63.CopyToNotebookRequestBuilder) {
    return i873554b7b847d1d15d4a54e90f68c2d1e8e0e6c46101a06a7008eedc4dd8cb63.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) CopyToSectionGroup()(*i446908705c1ddee4882afd2e4611611621abaa084248970a1b001f9a4c3f93c1.CopyToSectionGroupRequestBuilder) {
    return i446908705c1ddee4882afd2e4611611621abaa084248970a1b001f9a4c3f93c1.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreateDeleteRequestInformation(options *OnenoteSectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreateGetRequestInformation(options *OnenoteSectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) CreatePatchRequestInformation(options *OnenoteSectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Delete(options *OnenoteSectionRequestBuilderDeleteOptions)(error) {
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
// Get the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Get(options *OnenoteSectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnenoteSection, error) {
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
func (m *OnenoteSectionRequestBuilder) Pages()(*if33c50562d8903d2dc99274fc537bcb09ad5988398de7108aa512b0390c432e2.PagesRequestBuilder) {
    return if33c50562d8903d2dc99274fc537bcb09ad5988398de7108aa512b0390c432e2.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.notebooks.item.sectionGroups.item.sections.item.pages.item collection
func (m *OnenoteSectionRequestBuilder) PagesById(id string)(*i53277344ee337bfe414f01155cbb7c88e6b1d015b2391b56e0cba731399ceeda.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i53277344ee337bfe414f01155cbb7c88e6b1d015b2391b56e0cba731399ceeda.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentNotebook()(*i6b88ea813bf4804160ecb63495a5d65b7c5c3d81998515d99dc1b8189593fa74.ParentNotebookRequestBuilder) {
    return i6b88ea813bf4804160ecb63495a5d65b7c5c3d81998515d99dc1b8189593fa74.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *OnenoteSectionRequestBuilder) ParentSectionGroup()(*ic50e85cfbfe1e68ccdf2beec3955ddd7c2ce5ebb6956cb05fa5a04e731b2190a.ParentSectionGroupRequestBuilder) {
    return ic50e85cfbfe1e68ccdf2beec3955ddd7c2ce5ebb6956cb05fa5a04e731b2190a.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionRequestBuilder) Patch(options *OnenoteSectionRequestBuilderPatchOptions)(error) {
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
