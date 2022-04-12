package gettargetedusersanddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetTargetedUsersAndDevicesRequestBuilder provides operations to call the getTargetedUsersAndDevices method.
type GetTargetedUsersAndDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTargetedUsersAndDevicesRequestBuilderPostOptions options for Post
type GetTargetedUsersAndDevicesRequestBuilderPostOptions struct {
    // 
    Body GetTargetedUsersAndDevicesRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetTargetedUsersAndDevicesRequestBuilderInternal instantiates a new GetTargetedUsersAndDevicesRequestBuilder and sets the default values.
func NewGetTargetedUsersAndDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTargetedUsersAndDevicesRequestBuilder) {
    m := &GetTargetedUsersAndDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/microsoft.graph.getTargetedUsersAndDevices";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTargetedUsersAndDevicesRequestBuilder instantiates a new GetTargetedUsersAndDevicesRequestBuilder and sets the default values.
func NewGetTargetedUsersAndDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTargetedUsersAndDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTargetedUsersAndDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getTargetedUsersAndDevices
func (m *GetTargetedUsersAndDevicesRequestBuilder) CreatePostRequestInformation(options *GetTargetedUsersAndDevicesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Post invoke action getTargetedUsersAndDevices
func (m *GetTargetedUsersAndDevicesRequestBuilder) Post(options *GetTargetedUsersAndDevicesRequestBuilderPostOptions)(GetTargetedUsersAndDevicesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTargetedUsersAndDevicesResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTargetedUsersAndDevicesResponseable), nil
}
