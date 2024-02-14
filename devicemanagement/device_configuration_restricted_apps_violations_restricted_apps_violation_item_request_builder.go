package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder provides operations to manage the deviceConfigurationRestrictedAppsViolations property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetQueryParameters restricted apps violations for this account.
type DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetQueryParameters
}
// DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderInternal instantiates a new DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) {
    m := &DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurationRestrictedAppsViolations/{restrictedAppsViolation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder instantiates a new DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder and sets the default values.
func NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceConfigurationRestrictedAppsViolations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get restricted apps violations for this account.
// returns a RestrictedAppsViolationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property deviceConfigurationRestrictedAppsViolations in deviceManagement
// returns a RestrictedAppsViolationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, requestConfiguration *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property deviceConfigurationRestrictedAppsViolations for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/deviceConfigurationRestrictedAppsViolations/{restrictedAppsViolation%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation restricted apps violations for this account.
// returns a *RequestInformation when successful
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceConfigurationRestrictedAppsViolations in deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestrictedAppsViolationable, requestConfiguration *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/deviceConfigurationRestrictedAppsViolations/{restrictedAppsViolation%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder when successful
func (m *DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) WithUrl(rawUrl string)(*DeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder) {
    return NewDeviceConfigurationRestrictedAppsViolationsRestrictedAppsViolationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
