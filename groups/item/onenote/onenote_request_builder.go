package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i64b24ffe316f3722f215f78bae12ade40453c483129ad7fc06625f65bc4a90ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections"
    ia2351d1cdc3c79fba6288e7e74277c73ecd3558bb33b94740769e6650cbef5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/operations"
    ia546d15a045a5ffd8c2159899a9c2281cddc3a248537676e7e8c5782e9f2dfb5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/pages"
    ia81554816758a58901e38207656295c0f31e53f3c927a7b0998d2e590a935cd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks"
    ibaf9d3fc8edad28847225337bdcf7c119635455e424ad62bb63c7aea19a15fd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/resources"
    idb9ca44db7e4bb6917817740ec023009919aa158a3726ac72592155e5b66d8b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sectiongroups"
    i1b3fafc9d65a1513998934b94cfdb803b12e0573061793461c080db5ea06927d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sectiongroups/item"
    i480af79e3917198296c86a1af7f50fdaf55227d7044cc056fe7d549f796a793b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/notebooks/item"
    i51c40c3051772d274b9e4c783b749920d188bab082d98ca9d8e370e423cfb21f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/resources/item"
    i6c43559146cead075fda5951961149a049336212534c9a0a3d64cee92a61448b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/sections/item"
    i6db3c66e6f512160cc54be7500f683c71d6c2119d05f16a8dd1d6cb33c50f2b7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/pages/item"
    i9f9f088768842079632d1a5cfd365532c04905907271a753e482d2a749f0e074 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/onenote/operations/item"
)

// OnenoteRequestBuilder builds and executes requests for operations under \groups\{group-id}\onenote
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnenoteRequestBuilderDeleteOptions options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnenoteRequestBuilderGetOptions options for Get
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
// OnenoteRequestBuilderGetQueryParameters read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteRequestBuilderPatchOptions options for Patch
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
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation read-only.
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
// CreateGetRequestInformation read-only.
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation read-only.
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
// Delete read-only.
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
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
// Get read-only.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnenote() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenote), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*ia81554816758a58901e38207656295c0f31e53f3c927a7b0998d2e590a935cd0.NotebooksRequestBuilder) {
    return ia81554816758a58901e38207656295c0f31e53f3c927a7b0998d2e590a935cd0.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i480af79e3917198296c86a1af7f50fdaf55227d7044cc056fe7d549f796a793b.NotebookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i480af79e3917198296c86a1af7f50fdaf55227d7044cc056fe7d549f796a793b.NewNotebookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ia2351d1cdc3c79fba6288e7e74277c73ecd3558bb33b94740769e6650cbef5ac.OperationsRequestBuilder) {
    return ia2351d1cdc3c79fba6288e7e74277c73ecd3558bb33b94740769e6650cbef5ac.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i9f9f088768842079632d1a5cfd365532c04905907271a753e482d2a749f0e074.OnenoteOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i9f9f088768842079632d1a5cfd365532c04905907271a753e482d2a749f0e074.NewOnenoteOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*ia546d15a045a5ffd8c2159899a9c2281cddc3a248537676e7e8c5782e9f2dfb5.PagesRequestBuilder) {
    return ia546d15a045a5ffd8c2159899a9c2281cddc3a248537676e7e8c5782e9f2dfb5.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i6db3c66e6f512160cc54be7500f683c71d6c2119d05f16a8dd1d6cb33c50f2b7.OnenotePageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i6db3c66e6f512160cc54be7500f683c71d6c2119d05f16a8dd1d6cb33c50f2b7.NewOnenotePageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch read-only.
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
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
func (m *OnenoteRequestBuilder) Resources()(*ibaf9d3fc8edad28847225337bdcf7c119635455e424ad62bb63c7aea19a15fd0.ResourcesRequestBuilder) {
    return ibaf9d3fc8edad28847225337bdcf7c119635455e424ad62bb63c7aea19a15fd0.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i51c40c3051772d274b9e4c783b749920d188bab082d98ca9d8e370e423cfb21f.OnenoteResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return i51c40c3051772d274b9e4c783b749920d188bab082d98ca9d8e370e423cfb21f.NewOnenoteResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*idb9ca44db7e4bb6917817740ec023009919aa158a3726ac72592155e5b66d8b0.SectionGroupsRequestBuilder) {
    return idb9ca44db7e4bb6917817740ec023009919aa158a3726ac72592155e5b66d8b0.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i1b3fafc9d65a1513998934b94cfdb803b12e0573061793461c080db5ea06927d.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i1b3fafc9d65a1513998934b94cfdb803b12e0573061793461c080db5ea06927d.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*i64b24ffe316f3722f215f78bae12ade40453c483129ad7fc06625f65bc4a90ea.SectionsRequestBuilder) {
    return i64b24ffe316f3722f215f78bae12ade40453c483129ad7fc06625f65bc4a90ea.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i6c43559146cead075fda5951961149a049336212534c9a0a3d64cee92a61448b.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i6c43559146cead075fda5951961149a049336212534c9a0a3d64cee92a61448b.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
