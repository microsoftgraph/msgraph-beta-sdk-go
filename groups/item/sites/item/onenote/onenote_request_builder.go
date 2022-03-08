package onenote

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i4450981a398fccbc08936fae6552bde46f65e3c3c2c52229d65ba52f2443d712 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections"
    i684e1b1e5e48698390927de8d53cfae3c4f01987c39d29563ef0668bd8f1679e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/pages"
    i827ea0e8f47b335e1d7625bdcee4132d1c39c89e6705bcf8dec5a7ab93e42ca0 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks"
    ia90d02a73ceb7e91b48d5ce7fe32b9a89500ca8642e5d65d4101da93b17e5626 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/operations"
    ib899460689f03e60e08c6290e962af11bcc826307e556a3e3955e52b07770c9c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sectiongroups"
    iff417060019bf4ffda2ebf3a821bcc58f8df25c53d1391edb68c749b72dc8aef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/resources"
    i16c9e66d1083ec3a316a178a7ab7797c89164cec84859aa8b653a67213839a46 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/notebooks/item"
    i692555892f99fa2d7786dbc066aac83ffc33935702d8d390dbec130d3317407f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/pages/item"
    i73c3b260138fa11b0198711d81d67f91a5e294a9bd2a8b8703c33b1a03cb23d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sections/item"
    i76052aeda1272044f1ffb535ed14ed55993e8dd9e3ecf2ebea81024d0d368e64 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/resources/item"
    i9ff72b5e8692b7ae8bbf6faa9b5b264f4deb62e9a86157f1efbc1f7194f8b574 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/operations/item"
    ibacd123a946550e1daaad1fd0fa2b28edcb3193fcb3325c7a0449e2ab7bbaa88 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/onenote/sectiongroups/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.site entity.
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
// OnenoteRequestBuilderGetQueryParameters calls the OneNote service for notebook related operations.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnenoteRequestBuilderPatchOptions options for Patch
type OnenoteRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenoteable;
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
    m.urlTemplate = "{+baseurl}/groups/{group_id}/sites/{site_id}/onenote{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onenote for groups
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
// CreateGetRequestInformation calls the OneNote service for notebook related operations.
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
// CreatePatchRequestInformation update the navigation property onenote in groups
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
// Delete delete navigation property onenote for groups
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
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
// Get calls the OneNote service for notebook related operations.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateOnenoteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Onenoteable), nil
}
func (m *OnenoteRequestBuilder) Notebooks()(*i827ea0e8f47b335e1d7625bdcee4132d1c39c89e6705bcf8dec5a7ab93e42ca0.NotebooksRequestBuilder) {
    return i827ea0e8f47b335e1d7625bdcee4132d1c39c89e6705bcf8dec5a7ab93e42ca0.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i16c9e66d1083ec3a316a178a7ab7797c89164cec84859aa8b653a67213839a46.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook_id"] = id
    }
    return i16c9e66d1083ec3a316a178a7ab7797c89164cec84859aa8b653a67213839a46.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Operations()(*ia90d02a73ceb7e91b48d5ce7fe32b9a89500ca8642e5d65d4101da93b17e5626.OperationsRequestBuilder) {
    return ia90d02a73ceb7e91b48d5ce7fe32b9a89500ca8642e5d65d4101da93b17e5626.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i9ff72b5e8692b7ae8bbf6faa9b5b264f4deb62e9a86157f1efbc1f7194f8b574.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation_id"] = id
    }
    return i9ff72b5e8692b7ae8bbf6faa9b5b264f4deb62e9a86157f1efbc1f7194f8b574.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Pages()(*i684e1b1e5e48698390927de8d53cfae3c4f01987c39d29563ef0668bd8f1679e.PagesRequestBuilder) {
    return i684e1b1e5e48698390927de8d53cfae3c4f01987c39d29563ef0668bd8f1679e.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i692555892f99fa2d7786dbc066aac83ffc33935702d8d390dbec130d3317407f.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage_id"] = id
    }
    return i692555892f99fa2d7786dbc066aac83ffc33935702d8d390dbec130d3317407f.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in groups
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
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
func (m *OnenoteRequestBuilder) Resources()(*iff417060019bf4ffda2ebf3a821bcc58f8df25c53d1391edb68c749b72dc8aef.ResourcesRequestBuilder) {
    return iff417060019bf4ffda2ebf3a821bcc58f8df25c53d1391edb68c749b72dc8aef.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i76052aeda1272044f1ffb535ed14ed55993e8dd9e3ecf2ebea81024d0d368e64.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource_id"] = id
    }
    return i76052aeda1272044f1ffb535ed14ed55993e8dd9e3ecf2ebea81024d0d368e64.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) SectionGroups()(*ib899460689f03e60e08c6290e962af11bcc826307e556a3e3955e52b07770c9c.SectionGroupsRequestBuilder) {
    return ib899460689f03e60e08c6290e962af11bcc826307e556a3e3955e52b07770c9c.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*ibacd123a946550e1daaad1fd0fa2b28edcb3193fcb3325c7a0449e2ab7bbaa88.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup_id"] = id
    }
    return ibacd123a946550e1daaad1fd0fa2b28edcb3193fcb3325c7a0449e2ab7bbaa88.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnenoteRequestBuilder) Sections()(*i4450981a398fccbc08936fae6552bde46f65e3c3c2c52229d65ba52f2443d712.SectionsRequestBuilder) {
    return i4450981a398fccbc08936fae6552bde46f65e3c3c2c52229d65ba52f2443d712.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.sites.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i73c3b260138fa11b0198711d81d67f91a5e294a9bd2a8b8703c33b1a03cb23d4.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection_id"] = id
    }
    return i73c3b260138fa11b0198711d81d67f91a5e294a9bd2a8b8703c33b1a03cb23d4.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
