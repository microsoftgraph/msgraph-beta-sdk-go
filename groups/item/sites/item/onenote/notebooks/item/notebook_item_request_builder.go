package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0d35278fda68779051f9939ecf54bcfb27428c3efb645e80214a7cd51798bd76 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks/item/sections"
    i422f7d5d1e3b8cc4b2998408c6544417bdcd438dd6a7e5a26814da4387459188 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups"
    id5a0b3ed746613ecd790568cf61176215ad22094b74547751da315f96b83cae9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks/item/copynotebook"
    i39a06b615bbc5e271fcb2f1adb947ddc08e99de7865163932129d1b4941a09d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks/item/sections/item"
    i71254a2ab591ceb1aa1bdb6d86d42a75cef34f7d5c430e4038ecd1a4a45d30ef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks/item/sectiongroups/item"
)

// NotebookItemRequestBuilder provides operations to manage the notebooks property of the microsoft.graph.onenote entity.
type NotebookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NotebookItemRequestBuilderDeleteOptions options for Delete
type NotebookItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotebookItemRequestBuilderGetOptions options for Get
type NotebookItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NotebookItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotebookItemRequestBuilderGetQueryParameters the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
type NotebookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NotebookItemRequestBuilderPatchOptions options for Patch
type NotebookItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebookable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewNotebookItemRequestBuilderInternal instantiates a new NotebookItemRequestBuilder and sets the default values.
func NewNotebookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotebookItemRequestBuilder) {
    m := &NotebookItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/onenote/notebooks/{notebook_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotebookItemRequestBuilder instantiates a new NotebookItemRequestBuilder and sets the default values.
func NewNotebookItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotebookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotebookItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *NotebookItemRequestBuilder) CopyNotebook()(*id5a0b3ed746613ecd790568cf61176215ad22094b74547751da315f96b83cae9.CopyNotebookRequestBuilder) {
    return id5a0b3ed746613ecd790568cf61176215ad22094b74547751da315f96b83cae9.NewCopyNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property notebooks for groups
func (m *NotebookItemRequestBuilder) CreateDeleteRequestInformation(options *NotebookItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookItemRequestBuilder) CreateGetRequestInformation(options *NotebookItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property notebooks in groups
func (m *NotebookItemRequestBuilder) CreatePatchRequestInformation(options *NotebookItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property notebooks for groups
func (m *NotebookItemRequestBuilder) Delete(options *NotebookItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of OneNote notebooks that are owned by the user or group. Read-only. Nullable.
func (m *NotebookItemRequestBuilder) Get(options *NotebookItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebookable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateNotebookFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Notebookable), nil
}
// Patch update the navigation property notebooks in groups
func (m *NotebookItemRequestBuilder) Patch(options *NotebookItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *NotebookItemRequestBuilder) SectionGroups()(*i422f7d5d1e3b8cc4b2998408c6544417bdcd438dd6a7e5a26814da4387459188.SectionGroupsRequestBuilder) {
    return i422f7d5d1e3b8cc4b2998408c6544417bdcd438dd6a7e5a26814da4387459188.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.notebooks.item.sectionGroups.item collection
func (m *NotebookItemRequestBuilder) SectionGroupsById(id string)(*i71254a2ab591ceb1aa1bdb6d86d42a75cef34f7d5c430e4038ecd1a4a45d30ef.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return i71254a2ab591ceb1aa1bdb6d86d42a75cef34f7d5c430e4038ecd1a4a45d30ef.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *NotebookItemRequestBuilder) Sections()(*i0d35278fda68779051f9939ecf54bcfb27428c3efb645e80214a7cd51798bd76.SectionsRequestBuilder) {
    return i0d35278fda68779051f9939ecf54bcfb27428c3efb645e80214a7cd51798bd76.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.notebooks.item.sections.item collection
func (m *NotebookItemRequestBuilder) SectionsById(id string)(*i39a06b615bbc5e271fcb2f1adb947ddc08e99de7865163932129d1b4941a09d5.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i39a06b615bbc5e271fcb2f1adb947ddc08e99de7865163932129d1b4941a09d5.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
