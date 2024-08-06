package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder provides operations to manage the managedDeviceWindowsOSImages property of the microsoft.graph.deviceManagement entity.
type ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetQueryParameters a list of ManagedDeviceWindowsOperatingSystemImages
type ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetQueryParameters
}
// ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderInternal instantiates a new ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder and sets the default values.
func NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) {
    m := &ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDeviceWindowsOSImages/{managedDeviceWindowsOperatingSystemImage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder instantiates a new ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder and sets the default values.
func NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedDeviceWindowsOSImages for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a list of ManagedDeviceWindowsOperatingSystemImages
// returns a ManagedDeviceWindowsOperatingSystemImageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceWindowsOperatingSystemImageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable), nil
}
// GetAllManagedDeviceWindowsOSImages provides operations to call the getAllManagedDeviceWindowsOSImages method.
// returns a *ManagedDeviceWindowsOSImagesItemGetAllManagedDeviceWindowsOSImagesRequestBuilder when successful
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) GetAllManagedDeviceWindowsOSImages()(*ManagedDeviceWindowsOSImagesItemGetAllManagedDeviceWindowsOSImagesRequestBuilder) {
    return NewManagedDeviceWindowsOSImagesItemGetAllManagedDeviceWindowsOSImagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managedDeviceWindowsOSImages in deviceManagement
// returns a ManagedDeviceWindowsOperatingSystemImageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, requestConfiguration *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceWindowsOperatingSystemImageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable), nil
}
// ToDeleteRequestInformation delete navigation property managedDeviceWindowsOSImages for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a list of ManagedDeviceWindowsOperatingSystemImages
// returns a *RequestInformation when successful
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedDeviceWindowsOSImages in deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, requestConfiguration *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder when successful
func (m *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) WithUrl(rawUrl string)(*ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) {
    return NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
