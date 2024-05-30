package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder provides operations to manage the managedDeviceEncryptionStates property of the microsoft.graph.deviceManagement entity.
type ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetQueryParameters encryption report for devices in this account
type ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetQueryParameters
}
// ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderInternal instantiates a new ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder and sets the default values.
func NewManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) {
    m := &ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDeviceEncryptionStates/{managedDeviceEncryptionState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder instantiates a new ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder and sets the default values.
func NewManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDeviceEncryptionStates for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get encryption report for devices in this account
// returns a ManagedDeviceEncryptionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceEncryptionStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceEncryptionStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceEncryptionStateable), nil
}
// Patch update the navigation property managedDeviceEncryptionStates in deviceManagement
// returns a ManagedDeviceEncryptionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceEncryptionStateable, requestConfiguration *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceEncryptionStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceEncryptionStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceEncryptionStateable), nil
}
// ToDeleteRequestInformation delete navigation property managedDeviceEncryptionStates for deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation encryption report for devices in this account
// returns a *RequestInformation when successful
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDeviceEncryptionStates in deviceManagement
// returns a *RequestInformation when successful
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceEncryptionStateable, requestConfiguration *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder when successful
func (m *ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) WithUrl(rawUrl string)(*ManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder) {
    return NewManageddeviceencryptionstatesManagedDeviceEncryptionStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
