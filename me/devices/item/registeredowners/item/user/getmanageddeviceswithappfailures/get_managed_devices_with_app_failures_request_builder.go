package getmanageddeviceswithappfailures

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetManagedDevicesWithAppFailuresRequestBuilder provides operations to call the getManagedDevicesWithAppFailures method.
type GetManagedDevicesWithAppFailuresRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetManagedDevicesWithAppFailuresRequestBuilderInternal instantiates a new GetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedDevicesWithAppFailuresRequestBuilder) {
    m := &GetManagedDevicesWithAppFailuresRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}/microsoft.graph.user/microsoft.graph.getManagedDevicesWithAppFailures()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetManagedDevicesWithAppFailuresRequestBuilder instantiates a new GetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewGetManagedDevicesWithAppFailuresRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedDevicesWithAppFailuresRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) Get()(GetManagedDevicesWithAppFailuresResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetManagedDevicesWithAppFailuresRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetManagedDevicesWithAppFailuresResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagedDevicesWithAppFailuresResponseable), nil
}
