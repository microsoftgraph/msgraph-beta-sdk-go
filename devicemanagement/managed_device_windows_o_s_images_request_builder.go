package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDeviceWindowsOSImagesRequestBuilder provides operations to manage the managedDeviceWindowsOSImages property of the microsoft.graph.deviceManagement entity.
type ManagedDeviceWindowsOSImagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDeviceWindowsOSImagesRequestBuilderGetQueryParameters a list of ManagedDeviceWindowsOperatingSystemImages
type ManagedDeviceWindowsOSImagesRequestBuilderGetQueryParameters struct {
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
// ManagedDeviceWindowsOSImagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceWindowsOSImagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDeviceWindowsOSImagesRequestBuilderGetQueryParameters
}
// ManagedDeviceWindowsOSImagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDeviceWindowsOSImagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedDeviceWindowsOperatingSystemImageId provides operations to manage the managedDeviceWindowsOSImages property of the microsoft.graph.deviceManagement entity.
// returns a *ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder when successful
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) ByManagedDeviceWindowsOperatingSystemImageId(managedDeviceWindowsOperatingSystemImageId string)(*ManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceWindowsOperatingSystemImageId != "" {
        urlTplParams["managedDeviceWindowsOperatingSystemImage%2Did"] = managedDeviceWindowsOperatingSystemImageId
    }
    return NewManagedDeviceWindowsOSImagesManagedDeviceWindowsOperatingSystemImageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedDeviceWindowsOSImagesRequestBuilderInternal instantiates a new ManagedDeviceWindowsOSImagesRequestBuilder and sets the default values.
func NewManagedDeviceWindowsOSImagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceWindowsOSImagesRequestBuilder) {
    m := &ManagedDeviceWindowsOSImagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDeviceWindowsOSImages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedDeviceWindowsOSImagesRequestBuilder instantiates a new ManagedDeviceWindowsOSImagesRequestBuilder and sets the default values.
func NewManagedDeviceWindowsOSImagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDeviceWindowsOSImagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceWindowsOSImagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedDeviceWindowsOSImagesCountRequestBuilder when successful
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) Count()(*ManagedDeviceWindowsOSImagesCountRequestBuilder) {
    return NewManagedDeviceWindowsOSImagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of ManagedDeviceWindowsOperatingSystemImages
// returns a ManagedDeviceWindowsOperatingSystemImageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDeviceWindowsOSImagesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceWindowsOperatingSystemImageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageCollectionResponseable), nil
}
// Post create new navigation property to managedDeviceWindowsOSImages for deviceManagement
// returns a ManagedDeviceWindowsOperatingSystemImageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, requestConfiguration *ManagedDeviceWindowsOSImagesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation a list of ManagedDeviceWindowsOperatingSystemImages
// returns a *RequestInformation when successful
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDeviceWindowsOSImagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managedDeviceWindowsOSImages for deviceManagement
// returns a *RequestInformation when successful
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceWindowsOperatingSystemImageable, requestConfiguration *ManagedDeviceWindowsOSImagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedDeviceWindowsOSImagesRequestBuilder when successful
func (m *ManagedDeviceWindowsOSImagesRequestBuilder) WithUrl(rawUrl string)(*ManagedDeviceWindowsOSImagesRequestBuilder) {
    return NewManagedDeviceWindowsOSImagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
