package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingStateSummary property of the microsoft.graph.deviceManagement entity.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters the summary state of ATP onboarding state for this account.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters
}
// AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedThreatProtectionOnboardingDeviceSettingStates provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates()(*AdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedthreatprotectiononboardingdevicesettingstatesAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal instantiates a new AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    m := &AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder instantiates a new AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the summary state of ATP onboarding state for this account.
// returns a AdvancedThreatProtectionOnboardingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable), nil
}
// Patch update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
// returns a AdvancedThreatProtectionOnboardingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the summary state of ATP onboarding state for this account.
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
// returns a *RequestInformation when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingStateSummaryable, requestConfiguration *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder when successful
func (m *AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) WithUrl(rawUrl string)(*AdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    return NewAdvancedthreatprotectiononboardingstatesummaryAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
