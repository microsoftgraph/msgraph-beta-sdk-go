package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsanalyticsPermissionsAnalyticsRequestBuilder provides operations to manage the permissionsAnalytics property of the microsoft.graph.identityGovernance entity.
type PermissionsanalyticsPermissionsAnalyticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsanalyticsPermissionsAnalyticsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsPermissionsAnalyticsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetQueryParameters get permissionsAnalytics from identityGovernance
type PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetQueryParameters
}
// PermissionsanalyticsPermissionsAnalyticsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsanalyticsPermissionsAnalyticsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Aws provides operations to manage the aws property of the microsoft.graph.permissionsAnalyticsAggregation entity.
// returns a *PermissionsanalyticsAwsRequestBuilder when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) Aws()(*PermissionsanalyticsAwsRequestBuilder) {
    return NewPermissionsanalyticsAwsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Azure provides operations to manage the azure property of the microsoft.graph.permissionsAnalyticsAggregation entity.
// returns a *PermissionsanalyticsAzureRequestBuilder when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) Azure()(*PermissionsanalyticsAzureRequestBuilder) {
    return NewPermissionsanalyticsAzureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsanalyticsPermissionsAnalyticsRequestBuilderInternal instantiates a new PermissionsanalyticsPermissionsAnalyticsRequestBuilder and sets the default values.
func NewPermissionsanalyticsPermissionsAnalyticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsPermissionsAnalyticsRequestBuilder) {
    m := &PermissionsanalyticsPermissionsAnalyticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPermissionsanalyticsPermissionsAnalyticsRequestBuilder instantiates a new PermissionsanalyticsPermissionsAnalyticsRequestBuilder and sets the default values.
func NewPermissionsanalyticsPermissionsAnalyticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsanalyticsPermissionsAnalyticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsanalyticsPermissionsAnalyticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsAnalytics for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsanalyticsPermissionsAnalyticsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Gcp provides operations to manage the gcp property of the microsoft.graph.permissionsAnalyticsAggregation entity.
// returns a *PermissionsanalyticsGcpRequestBuilder when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) Gcp()(*PermissionsanalyticsGcpRequestBuilder) {
    return NewPermissionsanalyticsGcpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get permissionsAnalytics from identityGovernance
// returns a PermissionsAnalyticsAggregationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsAnalyticsAggregationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsAnalyticsAggregationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsAnalyticsAggregationable), nil
}
// Patch update the navigation property permissionsAnalytics in identityGovernance
// returns a PermissionsAnalyticsAggregationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsAnalyticsAggregationable, requestConfiguration *PermissionsanalyticsPermissionsAnalyticsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsAnalyticsAggregationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsAnalyticsAggregationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsAnalyticsAggregationable), nil
}
// ToDeleteRequestInformation delete navigation property permissionsAnalytics for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsPermissionsAnalyticsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get permissionsAnalytics from identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsanalyticsPermissionsAnalyticsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property permissionsAnalytics in identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsAnalyticsAggregationable, requestConfiguration *PermissionsanalyticsPermissionsAnalyticsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PermissionsanalyticsPermissionsAnalyticsRequestBuilder when successful
func (m *PermissionsanalyticsPermissionsAnalyticsRequestBuilder) WithUrl(rawUrl string)(*PermissionsanalyticsPermissionsAnalyticsRequestBuilder) {
    return NewPermissionsanalyticsPermissionsAnalyticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
