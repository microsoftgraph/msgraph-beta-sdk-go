package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
type ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetQueryParameters a list of ConfigManagerCollection
type ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetQueryParameters
}
// ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByConfigManagerCollectionId provides operations to manage the configManagerCollections property of the microsoft.graph.deviceManagement entity.
// returns a *ConfigmanagercollectionsConfigManagerCollectionItemRequestBuilder when successful
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) ByConfigManagerCollectionId(configManagerCollectionId string)(*ConfigmanagercollectionsConfigManagerCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if configManagerCollectionId != "" {
        urlTplParams["configManagerCollection%2Did"] = configManagerCollectionId
    }
    return NewConfigmanagercollectionsConfigManagerCollectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilderInternal instantiates a new ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder and sets the default values.
func NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) {
    m := &ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/configManagerCollections{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilder instantiates a new ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder and sets the default values.
func NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConfigmanagercollectionsCountRequestBuilder when successful
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) Count()(*ConfigmanagercollectionsCountRequestBuilder) {
    return NewConfigmanagercollectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of ConfigManagerCollection
// returns a ConfigManagerCollectionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerCollectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigManagerCollectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerCollectionCollectionResponseable), nil
}
// GetPolicySummaryWithPolicyId provides operations to call the getPolicySummary method.
// returns a *ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder when successful
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) GetPolicySummaryWithPolicyId(policyId *string)(*ConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilder) {
    return NewConfigmanagercollectionsGetpolicysummarywithpolicyidGetPolicySummaryWithPolicyIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, policyId)
}
// Post create new navigation property to configManagerCollections for deviceManagement
// returns a ConfigManagerCollectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerCollectionable, requestConfiguration *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerCollectionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConfigManagerCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerCollectionable), nil
}
// ToGetRequestInformation a list of ConfigManagerCollection
// returns a *RequestInformation when successful
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to configManagerCollections for deviceManagement
// returns a *RequestInformation when successful
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConfigManagerCollectionable, requestConfiguration *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder when successful
func (m *ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) WithUrl(rawUrl string)(*ConfigmanagercollectionsConfigManagerCollectionsRequestBuilder) {
    return NewConfigmanagercollectionsConfigManagerCollectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
