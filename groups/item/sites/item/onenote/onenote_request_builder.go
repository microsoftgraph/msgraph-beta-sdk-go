package onenote

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
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
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote{?%24select,%24expand}";
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
// Notebooks provides operations to manage the notebooks property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) Notebooks()(*i827ea0e8f47b335e1d7625bdcee4132d1c39c89e6705bcf8dec5a7ab93e42ca0.NotebooksRequestBuilder) {
    return i827ea0e8f47b335e1d7625bdcee4132d1c39c89e6705bcf8dec5a7ab93e42ca0.NewNotebooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotebooksById provides operations to manage the notebooks property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) NotebooksById(id string)(*i16c9e66d1083ec3a316a178a7ab7797c89164cec84859aa8b653a67213839a46.NotebookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notebook%2Did"] = id
    }
    return i16c9e66d1083ec3a316a178a7ab7797c89164cec84859aa8b653a67213839a46.NewNotebookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) Operations()(*ia90d02a73ceb7e91b48d5ce7fe32b9a89500ca8642e5d65d4101da93b17e5626.OperationsRequestBuilder) {
    return ia90d02a73ceb7e91b48d5ce7fe32b9a89500ca8642e5d65d4101da93b17e5626.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) OperationsById(id string)(*i9ff72b5e8692b7ae8bbf6faa9b5b264f4deb62e9a86157f1efbc1f7194f8b574.OnenoteOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteOperation%2Did"] = id
    }
    return i9ff72b5e8692b7ae8bbf6faa9b5b264f4deb62e9a86157f1efbc1f7194f8b574.NewOnenoteOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Pages provides operations to manage the pages property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) Pages()(*i684e1b1e5e48698390927de8d53cfae3c4f01987c39d29563ef0668bd8f1679e.PagesRequestBuilder) {
    return i684e1b1e5e48698390927de8d53cfae3c4f01987c39d29563ef0668bd8f1679e.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById provides operations to manage the pages property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) PagesById(id string)(*i692555892f99fa2d7786dbc066aac83ffc33935702d8d390dbec130d3317407f.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return i692555892f99fa2d7786dbc066aac83ffc33935702d8d390dbec130d3317407f.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
// Resources provides operations to manage the resources property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) Resources()(*iff417060019bf4ffda2ebf3a821bcc58f8df25c53d1391edb68c749b72dc8aef.ResourcesRequestBuilder) {
    return iff417060019bf4ffda2ebf3a821bcc58f8df25c53d1391edb68c749b72dc8aef.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById provides operations to manage the resources property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) ResourcesById(id string)(*i76052aeda1272044f1ffb535ed14ed55993e8dd9e3ecf2ebea81024d0d368e64.OnenoteResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteResource%2Did"] = id
    }
    return i76052aeda1272044f1ffb535ed14ed55993e8dd9e3ecf2ebea81024d0d368e64.NewOnenoteResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SectionGroups provides operations to manage the sectionGroups property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) SectionGroups()(*ib899460689f03e60e08c6290e962af11bcc826307e556a3e3955e52b07770c9c.SectionGroupsRequestBuilder) {
    return ib899460689f03e60e08c6290e962af11bcc826307e556a3e3955e52b07770c9c.NewSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById provides operations to manage the sectionGroups property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) SectionGroupsById(id string)(*ibacd123a946550e1daaad1fd0fa2b28edcb3193fcb3325c7a0449e2ab7bbaa88.SectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did"] = id
    }
    return ibacd123a946550e1daaad1fd0fa2b28edcb3193fcb3325c7a0449e2ab7bbaa88.NewSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections provides operations to manage the sections property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) Sections()(*i4450981a398fccbc08936fae6552bde46f65e3c3c2c52229d65ba52f2443d712.SectionsRequestBuilder) {
    return i4450981a398fccbc08936fae6552bde46f65e3c3c2c52229d65ba52f2443d712.NewSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById provides operations to manage the sections property of the microsoft.graph.onenote entity.
func (m *OnenoteRequestBuilder) SectionsById(id string)(*i73c3b260138fa11b0198711d81d67f91a5e294a9bd2a8b8703c33b1a03cb23d4.OnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return i73c3b260138fa11b0198711d81d67f91a5e294a9bd2a8b8703c33b1a03cb23d4.NewOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
