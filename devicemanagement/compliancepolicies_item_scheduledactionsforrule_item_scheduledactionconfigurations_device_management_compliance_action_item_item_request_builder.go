package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
type CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetQueryParameters the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
type CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetQueryParameters
}
// CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal instantiates a new CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder and sets the default values.
func NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    m := &CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations/{deviceManagementComplianceActionItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder instantiates a new CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder and sets the default values.
func NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property scheduledActionConfigurations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
// returns a DeviceManagementComplianceActionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable), nil
}
// Patch update the navigation property scheduledActionConfigurations in deviceManagement
// returns a DeviceManagementComplianceActionItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable), nil
}
// ToDeleteRequestInformation delete navigation property scheduledActionConfigurations for deviceManagement
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property scheduledActionConfigurations in deviceManagement
// returns a *RequestInformation when successful
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, requestConfiguration *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder when successful
func (m *CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) WithUrl(rawUrl string)(*CompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    return NewCompliancepoliciesItemScheduledactionsforruleItemScheduledactionconfigurationsDeviceManagementComplianceActionItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
