package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder provides operations to manage the deploySummary property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
type WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetQueryParameters windowsDefenderApplicationControl supplemental policy deployment summary.
type WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetQueryParameters
}
// WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderInternal instantiates a new WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder and sets the default values.
func NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) {
    m := &WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}/deploySummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder instantiates a new WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder and sets the default values.
func NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deploySummary for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get windowsDefenderApplicationControl supplemental policy deployment summary.
// returns a WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable), nil
}
// Patch update the navigation property deploySummary in deviceAppManagement
// returns a WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable, requestConfiguration *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deploySummary for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation windowsDefenderApplicationControl supplemental policy deployment summary.
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deploySummary in deviceAppManagement
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable, requestConfiguration *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder when successful
func (m *WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) WithUrl(rawUrl string)(*WdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder) {
    return NewWdacsupplementalpoliciesItemDeploysummaryDeploySummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
