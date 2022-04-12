package onenote

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i02442f2b4528552778554b6a8d7734fcbcc5c67d05fb38c736fd39e10a087b4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages"
    i21e97cc531e79994ac8f68134cbea3aa92a45effb9959e592cb197af924ec755 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sections"
    i2a34cadbfd73c7d67b9945eae7836418199869d09723bc6ea8f474c361c69476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups"
    i3f6844fa2bdd594acff5a9568bb940d97aee31c03de9263680db2cafec02cd32 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/resources"
    i4d5248421f500fed59a4443906c41e5156b79d60a5dc7eca83763c99b773c741 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks"
    id00a24b6b18f3cd1dcf3862b4f8cc4c98ee4ef1720eea18826d911d1150cf228 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/operations"
    i08bb5592a459ace4070591d9b4083a115e728de9d6fb3517918f3f661f99d4be "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sectiongroups/item"
    i31303fd26db642eb4e6f7dc00a9babd052efccdd57a191a012a948d9d636e966 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/pages/item"
    i7b6a7c655d49ef24caf532d5a7f363e6f1bff778a37ddde30a9ee2aa9a699820 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/operations/item"
    i96a102edf759f9f57dbf2fb12c37a9dd630d0fe2d5f0b61562e23013eb3c9897 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/resources/item"
    ic8c068b58107eb1b9b3bc4b7e43c090c8f2e1ddd8871f128b4089d6f6755e40a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/notebooks/item"
    if48a948123f8134534ad90c988dad247d9e1c7a92254ec0abf9a7a7c5e37a20c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onenote/sections/item"
)

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.user entity.
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteRequestBuilderDeleteOptions options for Delete
type OnenoteRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OnenoteRequestBuilderGetOptions options for Get
type OnenoteRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenoteRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// OnenoteRequestBuilderGetQueryParameters read-only.
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenoteRequestBuilderPatchOptions options for Patch
type OnenoteRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/onenote{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteRequestBuilder instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property onenote for me
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(options *OnenoteRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation read-only.
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(options *OnenoteRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property onenote in me
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(options *OnenoteRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property onenote for me
func (m *OnenoteRequestBuilder) Delete(options *OnenoteRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only.
func (m *OnenoteRequestBuilder) Get(options *OnenoteRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Notebooks the notebooks property
func (m *OnenoteRequestBuilder) Notebooks()(*i4d5248421f500fed59a4443906c41e5156b79d60a5dc7eca83763c99b773c741.NotebooksRequestBuilder) {
    return i4d5248421f500fed59a4443906c41e5156b79d60a5dc7eca83763c99b773c741.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*ic8c068b58107eb1b9b3bc4b7e43c090c8f2e1ddd8871f128b4089d6f6755e40a.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook%2Did"] = id
    }
    return ic8c068b58107eb1b9b3bc4b7e43c090c8f2e1ddd8871f128b4089d6f6755e40a.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*id00a24b6b18f3cd1dcf3862b4f8cc4c98ee4ef1720eea18826d911d1150cf228.OperationsRequestBuilder) {
    return id00a24b6b18f3cd1dcf3862b4f8cc4c98ee4ef1720eea18826d911d1150cf228.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i7b6a7c655d49ef24caf532d5a7f363e6f1bff778a37ddde30a9ee2aa9a699820.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation%2Did"] = id
    }
    return i7b6a7c655d49ef24caf532d5a7f363e6f1bff778a37ddde30a9ee2aa9a699820.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*i02442f2b4528552778554b6a8d7734fcbcc5c67d05fb38c736fd39e10a087b4e.PagesRequestBuilder) {
    return i02442f2b4528552778554b6a8d7734fcbcc5c67d05fb38c736fd39e10a087b4e.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i31303fd26db642eb4e6f7dc00a9babd052efccdd57a191a012a948d9d636e966.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return i31303fd26db642eb4e6f7dc00a9babd052efccdd57a191a012a948d9d636e966.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in me
func (m *OnenoteRequestBuilder) Patch(options *OnenoteRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Resources the resources property
func (m *OnenoteRequestBuilder) Resources()(*i3f6844fa2bdd594acff5a9568bb940d97aee31c03de9263680db2cafec02cd32.ResourcesRequestBuilder) {
    return i3f6844fa2bdd594acff5a9568bb940d97aee31c03de9263680db2cafec02cd32.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i96a102edf759f9f57dbf2fb12c37a9dd630d0fe2d5f0b61562e23013eb3c9897.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource%2Did"] = id
    }
    return i96a102edf759f9f57dbf2fb12c37a9dd630d0fe2d5f0b61562e23013eb3c9897.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*i2a34cadbfd73c7d67b9945eae7836418199869d09723bc6ea8f474c361c69476.SectionGroupsRequestBuilder) {
    return i2a34cadbfd73c7d67b9945eae7836418199869d09723bc6ea8f474c361c69476.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i08bb5592a459ace4070591d9b4083a115e728de9d6fb3517918f3f661f99d4be.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did"] = id
    }
    return i08bb5592a459ace4070591d9b4083a115e728de9d6fb3517918f3f661f99d4be.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*i21e97cc531e79994ac8f68134cbea3aa92a45effb9959e592cb197af924ec755.SectionsRequestBuilder) {
    return i21e97cc531e79994ac8f68134cbea3aa92a45effb9959e592cb197af924ec755.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*if48a948123f8134534ad90c988dad247d9e1c7a92254ec0abf9a7a7c5e37a20c.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return if48a948123f8134534ad90c988dad247d9e1c7a92254ec0abf9a7a7c5e37a20c.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
