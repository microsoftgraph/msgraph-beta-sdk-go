package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
type DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetQueryParameters restricted apps violations for this account.
type DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetQueryParameters struct {
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
// DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetQueryParameters
}
// DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByRestrictedAppsViolationId provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationrestrictedappsviolationsRestrictedAppsViolationItemRequestBuilder when successful
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) ByRestrictedAppsViolationId(restrictedAppsViolationId string)(*DeviceconfigurationrestrictedappsviolationsRestrictedAppsViolationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if restrictedAppsViolationId != "" {
        urlTplParams["restrictedAppsViolation%2Did"] = restrictedAppsViolationId
    }
    return NewDeviceconfigurationrestrictedappsviolationsRestrictedAppsViolationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal instantiates a new DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder and sets the default values.
func NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    m := &DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurationRestrictedAppsViolations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder instantiates a new DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder and sets the default values.
func NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceconfigurationrestrictedappsviolationsCountRequestBuilder when successful
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) Count()(*DeviceconfigurationrestrictedappsviolationsCountRequestBuilder) {
    return NewDeviceconfigurationrestrictedappsviolationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get restricted apps violations for this account.
// returns a RestrictedAppsViolationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRestrictedAppsViolationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationCollectionResponseable), nil
}
// Post create new navigation property to deviceConfigurationRestrictedAppsViolations for deviceManagement
// returns a RestrictedAppsViolationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, requestConfiguration *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateRestrictedAppsViolationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable), nil
}
// ToGetRequestInformation restricted apps violations for this account.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceConfigurationRestrictedAppsViolations for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, requestConfiguration *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder when successful
func (m *DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder) {
    return NewDeviceconfigurationrestrictedappsviolationsDeviceConfigurationRestrictedAppsViolationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
