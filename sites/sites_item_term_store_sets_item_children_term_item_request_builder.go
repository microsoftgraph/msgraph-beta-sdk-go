package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
)

// SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetQueryParameters children terms of set in term [store].
type SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetQueryParameters
}
// SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) Children()(*SitesItemTermStoreSetsItemChildrenItemChildrenRequestBuilder) {
    return NewSitesItemTermStoreSetsItemChildrenItemChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.termStore.term entity.
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) ChildrenById(id string)(*SitesItemTermStoreSetsItemChildrenItemChildrenTermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did1"] = id
    }
    return NewSitesItemTermStoreSetsItemChildrenItemChildrenTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSitesItemTermStoreSetsItemChildrenTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewSitesItemTermStoreSetsItemChildrenTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) {
    m := &SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site%2Did}/termStore/sets/{set%2Did}/children/{term%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSitesItemTermStoreSetsItemChildrenTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewSitesItemTermStoreSetsItemChildrenTermItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSitesItemTermStoreSetsItemChildrenTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property children for sites
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation children terms of set in term [store].
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property children in sites
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, requestConfiguration *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property children for sites
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get children terms of set in term [store].
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
// Patch update the navigation property children in sites
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) Patch(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, requestConfiguration *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilderPatchRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateTermFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Termable), nil
}
// Relations provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) Relations()(*SitesItemTermStoreSetsItemChildrenItemRelationsRequestBuilder) {
    return NewSitesItemTermStoreSetsItemChildrenItemRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById provides operations to manage the relations property of the microsoft.graph.termStore.term entity.
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) RelationsById(id string)(*SitesItemTermStoreSetsItemChildrenItemRelationsRelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return NewSitesItemTermStoreSetsItemChildrenItemRelationsRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Set provides operations to manage the set property of the microsoft.graph.termStore.term entity.
func (m *SitesItemTermStoreSetsItemChildrenTermItemRequestBuilder) Set()(*SitesItemTermStoreSetsItemChildrenItemSetRequestBuilder) {
    return NewSitesItemTermStoreSetsItemChildrenItemSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
