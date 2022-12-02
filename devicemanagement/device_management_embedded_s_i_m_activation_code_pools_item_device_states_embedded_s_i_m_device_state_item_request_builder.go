package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder provides operations to manage the deviceStates property of the microsoft.graph.embeddedSIMActivationCodePool entity.
type DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters navigational property to a list of device states for this pool.
type DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetQueryParameters
}
// DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal instantiates a new EmbeddedSIMDeviceStateItemRequestBuilder and sets the default values.
func NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    m := &DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool%2Did}/deviceStates/{embeddedSIMDeviceState%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder instantiates a new EmbeddedSIMDeviceStateItemRequestBuilder and sets the default values.
func NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceStates for deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation navigational property to a list of device states for this pool.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceStates in deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceStates for deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get navigational property to a list of device states for this pool.
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable), nil
}
// Patch update the navigation property deviceStates in deviceManagement
func (m *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, requestConfiguration *DeviceManagementEmbeddedSIMActivationCodePoolsItemDeviceStatesEmbeddedSIMDeviceStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEmbeddedSIMDeviceStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EmbeddedSIMDeviceStateable), nil
}
