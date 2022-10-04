package onenote

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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

// OnenoteRequestBuilder provides operations to manage the onenote property of the microsoft.graph.group entity.
type OnenoteRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnenoteRequestBuilderGetQueryParameters get onenote from groups
type OnenoteRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenoteRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenoteRequestBuilderGetQueryParameters
}
// OnenoteRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnenoteRequestBuilderInternal instantiates a new OnenoteRequestBuilder and sets the default values.
func NewOnenoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteRequestBuilder) {
    m := &OnenoteRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/onenote{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property onenote for groups
func (m *OnenoteRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get onenote from groups
func (m *OnenoteRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property onenote in groups
func (m *OnenoteRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property onenote for groups
func (m *OnenoteRequestBuilder) Delete(ctx context.Context, requestConfiguration *OnenoteRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get onenote from groups
func (m *OnenoteRequestBuilder) Get(ctx context.Context, requestConfiguration *OnenoteRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Notebooks the notebooks property
func (m *OnenoteRequestBuilder) Notebooks()(*ia81554816758a58901e38207656295c0f31e53f3c927a7b0998d2e590a935cd0.NotebooksRequestBuilder) {
    return ia81554816758a58901e38207656295c0f31e53f3c927a7b0998d2e590a935cd0.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.notebooks.item collection
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i480af79e3917198296c86a1af7f50fdaf55227d7044cc056fe7d549f796a793b.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook%2Did"] = id
    }
    return i480af79e3917198296c86a1af7f50fdaf55227d7044cc056fe7d549f796a793b.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *OnenoteRequestBuilder) Operations()(*ia2351d1cdc3c79fba6288e7e74277c73ecd3558bb33b94740769e6650cbef5ac.OperationsRequestBuilder) {
    return ia2351d1cdc3c79fba6288e7e74277c73ecd3558bb33b94740769e6650cbef5ac.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.operations.item collection
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i9f9f088768842079632d1a5cfd365532c04905907271a753e482d2a749f0e074.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation%2Did"] = id
    }
    return i9f9f088768842079632d1a5cfd365532c04905907271a753e482d2a749f0e074.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages the pages property
func (m *OnenoteRequestBuilder) Pages()(*ia546d15a045a5ffd8c2159899a9c2281cddc3a248537676e7e8c5782e9f2dfb5.PagesRequestBuilder) {
    return ia546d15a045a5ffd8c2159899a9c2281cddc3a248537676e7e8c5782e9f2dfb5.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.pages.item collection
func (m *OnenoteRequestBuilder) PagesById(id string)(*i6db3c66e6f512160cc54be7500f683c71d6c2119d05f16a8dd1d6cb33c50f2b7.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return i6db3c66e6f512160cc54be7500f683c71d6c2119d05f16a8dd1d6cb33c50f2b7.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property onenote in groups
func (m *OnenoteRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, requestConfiguration *OnenoteRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Onenoteable), nil
}
// Resources the resources property
func (m *OnenoteRequestBuilder) Resources()(*ibaf9d3fc8edad28847225337bdcf7c119635455e424ad62bb63c7aea19a15fd0.ResourcesRequestBuilder) {
    return ibaf9d3fc8edad28847225337bdcf7c119635455e424ad62bb63c7aea19a15fd0.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.resources.item collection
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i51c40c3051772d274b9e4c783b749920d188bab082d98ca9d8e370e423cfb21f.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource%2Did"] = id
    }
    return i51c40c3051772d274b9e4c783b749920d188bab082d98ca9d8e370e423cfb21f.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups the sectionGroups property
func (m *OnenoteRequestBuilder) SectionGroups()(*idb9ca44db7e4bb6917817740ec023009919aa158a3726ac72592155e5b66d8b0.SectionGroupsRequestBuilder) {
    return idb9ca44db7e4bb6917817740ec023009919aa158a3726ac72592155e5b66d8b0.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.sectionGroups.item collection
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*i1b3fafc9d65a1513998934b94cfdb803b12e0573061793461c080db5ea06927d.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did"] = id
    }
    return i1b3fafc9d65a1513998934b94cfdb803b12e0573061793461c080db5ea06927d.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections the sections property
func (m *OnenoteRequestBuilder) Sections()(*i64b24ffe316f3722f215f78bae12ade40453c483129ad7fc06625f65bc4a90ea.SectionsRequestBuilder) {
    return i64b24ffe316f3722f215f78bae12ade40453c483129ad7fc06625f65bc4a90ea.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.onenote.sections.item collection
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i6c43559146cead075fda5951961149a049336212534c9a0a3d64cee92a61448b.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return i6c43559146cead075fda5951961149a049336212534c9a0a3d64cee92a61448b.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
