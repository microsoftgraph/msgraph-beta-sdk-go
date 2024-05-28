package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
type WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
type WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters struct {
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
// WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters
}
// WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusId provides operations to manage the deviceStatuses property of the microsoft.graph.windowsDefenderApplicationControlSupplementalPolicy entity.
// returns a *WdacsupplementalpoliciesItemDevicestatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder when successful
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) ByWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusId(windowsDefenderApplicationControlSupplementalPolicyDeploymentStatusId string)(*WdacsupplementalpoliciesItemDevicestatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsDefenderApplicationControlSupplementalPolicyDeploymentStatusId != "" {
        urlTplParams["windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus%2Did"] = windowsDefenderApplicationControlSupplementalPolicyDeploymentStatusId
    }
    return NewWdacsupplementalpoliciesItemDevicestatusesWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal instantiates a new WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    m := &WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/wdacSupplementalPolicies/{windowsDefenderApplicationControlSupplementalPolicy%2Did}/deviceStatuses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder instantiates a new WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WdacsupplementalpoliciesItemDevicestatusesCountRequestBuilder when successful
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) Count()(*WdacsupplementalpoliciesItemDevicestatusesCountRequestBuilder) {
    return NewWdacsupplementalpoliciesItemDevicestatusesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
// returns a WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusCollectionResponseable), nil
}
// Post create new navigation property to deviceStatuses for deviceAppManagement
// returns a WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, requestConfiguration *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable), nil
}
// ToGetRequestInformation the list of device deployment states for this WindowsDefenderApplicationControl supplemental policy.
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceStatuses for deviceAppManagement
// returns a *RequestInformation when successful
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable, requestConfiguration *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) WithUrl(rawUrl string)(*WdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewWdacsupplementalpoliciesItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
