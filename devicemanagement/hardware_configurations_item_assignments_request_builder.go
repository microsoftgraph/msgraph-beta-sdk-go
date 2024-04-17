package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwareConfigurationsItemAssignmentsRequestBuilder provides operations to manage the assignments property of the microsoft.graph.hardwareConfiguration entity.
type HardwareConfigurationsItemAssignmentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwareConfigurationsItemAssignmentsRequestBuilderGetQueryParameters list of the Azure AD user group ids that hardware configuration will be applied to. Only security groups and Office 365 Groups are supported.
type HardwareConfigurationsItemAssignmentsRequestBuilderGetQueryParameters struct {
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
// HardwareConfigurationsItemAssignmentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemAssignmentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwareConfigurationsItemAssignmentsRequestBuilderGetQueryParameters
}
// HardwareConfigurationsItemAssignmentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwareConfigurationsItemAssignmentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHardwareConfigurationAssignmentId provides operations to manage the assignments property of the microsoft.graph.hardwareConfiguration entity.
// returns a *HardwareConfigurationsItemAssignmentsHardwareConfigurationAssignmentItemRequestBuilder when successful
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) ByHardwareConfigurationAssignmentId(hardwareConfigurationAssignmentId string)(*HardwareConfigurationsItemAssignmentsHardwareConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hardwareConfigurationAssignmentId != "" {
        urlTplParams["hardwareConfigurationAssignment%2Did"] = hardwareConfigurationAssignmentId
    }
    return NewHardwareConfigurationsItemAssignmentsHardwareConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHardwareConfigurationsItemAssignmentsRequestBuilderInternal instantiates a new HardwareConfigurationsItemAssignmentsRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemAssignmentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemAssignmentsRequestBuilder) {
    m := &HardwareConfigurationsItemAssignmentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwareConfigurations/{hardwareConfiguration%2Did}/assignments{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHardwareConfigurationsItemAssignmentsRequestBuilder instantiates a new HardwareConfigurationsItemAssignmentsRequestBuilder and sets the default values.
func NewHardwareConfigurationsItemAssignmentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwareConfigurationsItemAssignmentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwareConfigurationsItemAssignmentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HardwareConfigurationsItemAssignmentsCountRequestBuilder when successful
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) Count()(*HardwareConfigurationsItemAssignmentsCountRequestBuilder) {
    return NewHardwareConfigurationsItemAssignmentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of the Azure AD user group ids that hardware configuration will be applied to. Only security groups and Office 365 Groups are supported.
// returns a HardwareConfigurationAssignmentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwareConfigurationsItemAssignmentsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationAssignmentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationAssignmentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationAssignmentCollectionResponseable), nil
}
// Post create new navigation property to assignments for deviceManagement
// returns a HardwareConfigurationAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationAssignmentable, requestConfiguration *HardwareConfigurationsItemAssignmentsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwareConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationAssignmentable), nil
}
// ToGetRequestInformation list of the Azure AD user group ids that hardware configuration will be applied to. Only security groups and Office 365 Groups are supported.
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwareConfigurationsItemAssignmentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to assignments for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwareConfigurationAssignmentable, requestConfiguration *HardwareConfigurationsItemAssignmentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwareConfigurationsItemAssignmentsRequestBuilder when successful
func (m *HardwareConfigurationsItemAssignmentsRequestBuilder) WithUrl(rawUrl string)(*HardwareConfigurationsItemAssignmentsRequestBuilder) {
    return NewHardwareConfigurationsItemAssignmentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
