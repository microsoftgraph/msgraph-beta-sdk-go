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
// GetManagedDevicesWithAppFailuresRequestBuilderGetOptions options for Get
type GetManagedDevicesWithAppFailuresRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetManagedDevicesWithAppFailuresRequestBuilderInternal instantiates a new GetManagedDevicesWithAppFailuresRequestBuilder and sets the default values.
func NewGetManagedDevicesWithAppFailuresRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetManagedDevicesWithAppFailuresRequestBuilder) {
    m := &GetManagedDevicesWithAppFailuresRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.getManagedDevicesWithAppFailures()";
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
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) CreateGetRequestInformation(options *GetManagedDevicesWithAppFailuresRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get retrieves the list of devices with failed apps
func (m *GetManagedDevicesWithAppFailuresRequestBuilder) Get(options *GetManagedDevicesWithAppFailuresRequestBuilderGetOptions)(GetManagedDevicesWithAppFailuresResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetManagedDevicesWithAppFailuresResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetManagedDevicesWithAppFailuresResponseable), nil
}
