package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder provides operations to manage the managedDeviceMobileAppConfigurationStates property of the microsoft.graph.managedDevice entity.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetQueryParameters managed device mobile app configuration states for this device.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetQueryParameters
}
// ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal instantiates a new ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    m := &ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/managedDeviceMobileAppConfigurationStates/{managedDeviceMobileAppConfigurationState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder instantiates a new ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder and sets the default values.
func NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDeviceMobileAppConfigurationStates for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get managed device mobile app configuration states for this device.
// returns a ManagedDeviceMobileAppConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable), nil
}
// Patch update the navigation property managedDeviceMobileAppConfigurationStates in users
// returns a ManagedDeviceMobileAppConfigurationStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable), nil
}
// ToDeleteRequestInformation delete navigation property managedDeviceMobileAppConfigurationStates for users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation managed device mobile app configuration states for this device.
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDeviceMobileAppConfigurationStates in users
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationStateable, requestConfiguration *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder when successful
func (m *ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder) {
    return NewItemManageddevicesItemManageddevicemobileappconfigurationstatesManagedDeviceMobileAppConfigurationStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
