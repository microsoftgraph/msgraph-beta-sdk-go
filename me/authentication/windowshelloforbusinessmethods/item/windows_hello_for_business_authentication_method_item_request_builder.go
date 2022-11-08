package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4f9db3ce868389e1f5235eb54ab431b059ccaf6e498c4773259c3da435589ddc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/authentication/windowshelloforbusinessmethods/item/device"
)

// WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder provides operations to manage the windowsHelloForBusinessMethods property of the microsoft.graph.authentication entity.
type WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetQueryParameters represents the Windows Hello for Business authentication method registered to a user for authentication.
type WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal instantiates a new WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder and sets the default values.
func NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    m := &WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder instantiates a new WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder and sets the default values.
func NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsHelloForBusinessAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsHelloForBusinessMethods for me
func (m *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property windowsHelloForBusinessMethods for me
func (m *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Device provides operations to manage the device property of the microsoft.graph.windowsHelloForBusinessAuthenticationMethod entity.
func (m *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) Device()(*i4f9db3ce868389e1f5235eb54ab431b059ccaf6e498c4773259c3da435589ddc.DeviceRequestBuilder) {
    return i4f9db3ce868389e1f5235eb54ab431b059ccaf6e498c4773259c3da435589ddc.NewDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsHelloForBusinessAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsHelloForBusinessAuthenticationMethodable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsHelloForBusinessAuthenticationMethodable), nil
}
