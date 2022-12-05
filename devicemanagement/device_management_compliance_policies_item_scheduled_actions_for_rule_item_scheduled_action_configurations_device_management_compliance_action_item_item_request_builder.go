package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceManagementComplianceScheduledActionForRule entity.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetQueryParameters the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetQueryParameters
}
// DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal instantiates a new DeviceManagementComplianceActionItemItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    m := &DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations/{deviceManagementComplianceActionItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder instantiates a new DeviceManagementComplianceActionItemItemRequestBuilder and sets the default values.
func NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property scheduledActionConfigurations for deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property scheduledActionConfigurations in deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property scheduledActionConfigurations for deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable), nil
}
// Patch update the navigation property scheduledActionConfigurations in deviceManagement
func (m *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, requestConfiguration *DeviceManagementCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceManagementComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementComplianceActionItemable), nil
}
