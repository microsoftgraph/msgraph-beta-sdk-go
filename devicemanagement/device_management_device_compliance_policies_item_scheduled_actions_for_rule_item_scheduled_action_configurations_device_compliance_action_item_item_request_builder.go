package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder provides operations to manage the scheduledActionConfigurations property of the microsoft.graph.deviceComplianceScheduledActionForRule entity.
type DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetQueryParameters the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
type DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderInternal instantiates a new DeviceComplianceActionItemItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    m := &DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/scheduledActionsForRule/{deviceComplianceScheduledActionForRule%2Did}/scheduledActionConfigurations/{deviceComplianceActionItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder instantiates a new DeviceComplianceActionItemItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property scheduledActionConfigurations for deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, requestConfiguration *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of scheduled action configurations for this compliance policy. Compliance policy must have one and only one block scheduled action.
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable), nil
}
// Patch update the navigation property scheduledActionConfigurations in deviceManagement
func (m *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, requestConfiguration *DeviceManagementDeviceCompliancePoliciesItemScheduledActionsForRuleItemScheduledActionConfigurationsDeviceComplianceActionItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceComplianceActionItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceComplianceActionItemable), nil
}
