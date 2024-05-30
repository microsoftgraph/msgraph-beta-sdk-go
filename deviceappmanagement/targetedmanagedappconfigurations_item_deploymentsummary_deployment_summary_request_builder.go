package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder provides operations to manage the deploymentSummary property of the microsoft.graph.targetedManagedAppConfiguration entity.
type TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetQueryParameters navigation property to deployment summary of the configuration.
type TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetQueryParameters
}
// TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal instantiates a new TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    m := &TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/targetedManagedAppConfigurations/{targetedManagedAppConfiguration%2Did}/deploymentSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder instantiates a new TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder and sets the default values.
func NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deploymentSummary for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigation property to deployment summary of the configuration.
// returns a ManagedAppPolicyDeploymentSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppPolicyDeploymentSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppPolicyDeploymentSummaryable), nil
}
// Patch update the navigation property deploymentSummary in deviceAppManagement
// returns a ManagedAppPolicyDeploymentSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppPolicyDeploymentSummaryable, requestConfiguration *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppPolicyDeploymentSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppPolicyDeploymentSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deploymentSummary for deviceAppManagement
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation navigation property to deployment summary of the configuration.
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deploymentSummary in deviceAppManagement
// returns a *RequestInformation when successful
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppPolicyDeploymentSummaryable, requestConfiguration *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder when successful
func (m *TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) WithUrl(rawUrl string)(*TargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder) {
    return NewTargetedmanagedappconfigurationsItemDeploymentsummaryDeploymentSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
