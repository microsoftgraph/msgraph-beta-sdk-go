package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder provides operations to call the createGooglePlayWebToken method.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    m := &DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/androidManagedStoreAccountEnterpriseSettings/microsoft.graph.createGooglePlayWebToken";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder instantiates a new CreateGooglePlayWebTokenRequestBuilder and sets the default values.
func NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation generates a web token that is used in an embeddable component.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post generates a web token that is used in an embeddable component.
func (m *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenPostRequestBodyable, requestConfiguration *DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementAndroidManagedStoreAccountEnterpriseSettingsCreateGooglePlayWebTokenResponseable), nil
}
