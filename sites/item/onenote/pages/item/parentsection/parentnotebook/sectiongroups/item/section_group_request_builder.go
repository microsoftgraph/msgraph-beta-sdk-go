package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1c3c278ab06005814a40cf8bff5a0757d146f17dd5971c5a4834b4abfbd0bb25 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item/parentsectiongroup"
    id78eca3e88dc790585b7014ee591b5f2a1546291469de8ef59c9992f9f454a38 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item/sections"
    ie6a7e2e458f75b01eb1446d7061a6fd23bf6228d71c926a9a97958f46ec942ed "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item/sectiongroups"
    iff627328332e0f91eac98e17cb3263aad3814b09ff8f1c18283fa4d6f3cea224 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item/parentnotebook"
    i2507eae957f509c3ae8379a207cb8cc962de06b63632be5d085bd6aa64159ab6 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item/sections/item"
    ib42c6219694d4e104591f8ccfdc2a330d8eae2e28e44b660e5e6557a8a6cf7a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/onenote/pages/item/parentsection/parentnotebook/sectiongroups/item/sectiongroups/item"
)

// Builds and executes requests for operations under \sites\{site-id}\onenote\pages\{onenotePage-id}\parentSection\parentNotebook\sectionGroups\{sectionGroup-id}
type SectionGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SectionGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SectionGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SectionGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The section groups in the notebook. Read-only. Nullable.
type SectionGroupRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SectionGroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new SectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSectionGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    m := &SectionGroupRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/onenote/pages/{onenotePage_id}/parentSection/parentNotebook/sectionGroups/{sectionGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SectionGroupRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSectionGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SectionGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSectionGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) CreateDeleteRequestInformation(options *SectionGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) CreateGetRequestInformation(options *SectionGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) CreatePatchRequestInformation(options *SectionGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) Delete(options *SectionGroupRequestBuilderDeleteOptions)(error) {
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
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) Get(options *SectionGroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSectionGroup() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SectionGroup), nil
}
func (m *SectionGroupRequestBuilder) ParentNotebook()(*iff627328332e0f91eac98e17cb3263aad3814b09ff8f1c18283fa4d6f3cea224.ParentNotebookRequestBuilder) {
    return iff627328332e0f91eac98e17cb3263aad3814b09ff8f1c18283fa4d6f3cea224.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) ParentSectionGroup()(*i1c3c278ab06005814a40cf8bff5a0757d146f17dd5971c5a4834b4abfbd0bb25.ParentSectionGroupRequestBuilder) {
    return i1c3c278ab06005814a40cf8bff5a0757d146f17dd5971c5a4834b4abfbd0bb25.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The section groups in the notebook. Read-only. Nullable.
// Parameters:
//  - options : Options for the request
func (m *SectionGroupRequestBuilder) Patch(options *SectionGroupRequestBuilderPatchOptions)(error) {
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
func (m *SectionGroupRequestBuilder) SectionGroups()(*ie6a7e2e458f75b01eb1446d7061a6fd23bf6228d71c926a9a97958f46ec942ed.SectionGroupsRequestBuilder) {
    return ie6a7e2e458f75b01eb1446d7061a6fd23bf6228d71c926a9a97958f46ec942ed.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.sites.item.onenote.pages.item.parentSection.parentNotebook.sectionGroups.item.sectionGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SectionGroupRequestBuilder) SectionGroupsById(id string)(*ib42c6219694d4e104591f8ccfdc2a330d8eae2e28e44b660e5e6557a8a6cf7a4.SectionGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id1"] = id
    }
    return ib42c6219694d4e104591f8ccfdc2a330d8eae2e28e44b660e5e6557a8a6cf7a4.NewSectionGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SectionGroupRequestBuilder) Sections()(*id78eca3e88dc790585b7014ee591b5f2a1546291469de8ef59c9992f9f454a38.SectionsRequestBuilder) {
    return id78eca3e88dc790585b7014ee591b5f2a1546291469de8ef59c9992f9f454a38.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.sites.item.onenote.pages.item.parentSection.parentNotebook.sectionGroups.item.sections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SectionGroupRequestBuilder) SectionsById(id string)(*i2507eae957f509c3ae8379a207cb8cc962de06b63632be5d085bd6aa64159ab6.OnenoteSectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i2507eae957f509c3ae8379a207cb8cc962de06b63632be5d085bd6aa64159ab6.NewOnenoteSectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
