package getmanageddeviceswithfailedorpendingapps

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetManagedDevicesWithFailedOrPendingAppsRequestBuilder provides operations to call the getManagedDevicesWithFailedOrPendingApps method.
type GetManagedDevicesWithFailedOrPendingAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal instantiates a new GetManagedDevicesWithFailedOrPendingAppsRequestBuilder and sets the default values.
func NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    m := &GetManagedDevicesWithFailedOrPendingAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.getManagedDevicesWithFailedOrPendingApps()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilder instantiates a new GetManagedDevicesWithFailedOrPendingAppsRequestBuilder and sets the default values.
func NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagedDevicesWithFailedOrPendingAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation retrieves the list of devices with failed or pending apps
func (m *GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieves the list of devices with failed or pending apps
func (m *GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get retrieves the list of devices with failed or pending apps
func (m *GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) Get()(GetManagedDevicesWithFailedOrPendingAppsResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler retrieves the list of devices with failed or pending apps
func (m *GetManagedDevicesWithFailedOrPendingAppsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetManagedDevicesWithFailedOrPendingAppsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetManagedDevicesWithFailedOrPendingAppsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagedDevicesWithFailedOrPendingAppsResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(GetManagedDevicesWithFailedOrPendingAppsResponseable), nil
}
