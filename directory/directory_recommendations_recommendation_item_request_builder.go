package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectoryRecommendationsRecommendationItemRequestBuilder provides operations to manage the recommendations property of the microsoft.graph.directory entity.
type DirectoryRecommendationsRecommendationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectoryRecommendationsRecommendationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRecommendationsRecommendationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRecommendationsRecommendationItemRequestBuilderGetQueryParameters get recommendations from directory
type DirectoryRecommendationsRecommendationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRecommendationsRecommendationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRecommendationsRecommendationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRecommendationsRecommendationItemRequestBuilderGetQueryParameters
}
// DirectoryRecommendationsRecommendationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRecommendationsRecommendationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Complete provides operations to call the complete method.
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Complete()(*DirectoryRecommendationsItemCompleteRequestBuilder) {
    return NewDirectoryRecommendationsItemCompleteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRecommendationsRecommendationItemRequestBuilderInternal instantiates a new RecommendationItemRequestBuilder and sets the default values.
func NewDirectoryRecommendationsRecommendationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRecommendationsRecommendationItemRequestBuilder) {
    m := &DirectoryRecommendationsRecommendationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/recommendations/{recommendation%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRecommendationsRecommendationItemRequestBuilder instantiates a new RecommendationItemRequestBuilder and sets the default values.
func NewDirectoryRecommendationsRecommendationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRecommendationsRecommendationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRecommendationsRecommendationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property recommendations for directory
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRecommendationsRecommendationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get recommendations from directory
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRecommendationsRecommendationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property recommendations in directory
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable, requestConfiguration *DirectoryRecommendationsRecommendationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property recommendations for directory
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRecommendationsRecommendationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Dismiss provides operations to call the dismiss method.
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Dismiss()(*DirectoryRecommendationsItemDismissRequestBuilder) {
    return NewDirectoryRecommendationsItemDismissRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get recommendations from directory
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRecommendationsRecommendationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRecommendationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable), nil
}
// ImpactedResources provides operations to manage the impactedResources property of the microsoft.graph.recommendation entity.
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) ImpactedResources()(*DirectoryRecommendationsItemImpactedResourcesRequestBuilder) {
    return NewDirectoryRecommendationsItemImpactedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImpactedResourcesById provides operations to manage the impactedResources property of the microsoft.graph.recommendation entity.
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) ImpactedResourcesById(id string)(*DirectoryRecommendationsItemImpactedResourcesRecommendationResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["recommendationResource%2Did"] = id
    }
    return NewDirectoryRecommendationsItemImpactedResourcesRecommendationResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property recommendations in directory
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable, requestConfiguration *DirectoryRecommendationsRecommendationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRecommendationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Recommendationable), nil
}
// Postpone provides operations to call the postpone method.
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Postpone()(*DirectoryRecommendationsItemPostponeRequestBuilder) {
    return NewDirectoryRecommendationsItemPostponeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reactivate provides operations to call the reactivate method.
func (m *DirectoryRecommendationsRecommendationItemRequestBuilder) Reactivate()(*DirectoryRecommendationsItemReactivateRequestBuilder) {
    return NewDirectoryRecommendationsItemReactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
