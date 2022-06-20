package importdeviceidentitylist

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImportDeviceIdentityListRequestBuilder provides operations to call the importDeviceIdentityList method.
type ImportDeviceIdentityListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ImportDeviceIdentityListRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportDeviceIdentityListRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportDeviceIdentityListRequestBuilderInternal instantiates a new ImportDeviceIdentityListRequestBuilder and sets the default values.
func NewImportDeviceIdentityListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportDeviceIdentityListRequestBuilder) {
    m := &ImportDeviceIdentityListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedDeviceIdentities/microsoft.graph.importDeviceIdentityList";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportDeviceIdentityListRequestBuilder instantiates a new ImportDeviceIdentityListRequestBuilder and sets the default values.
func NewImportDeviceIdentityListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportDeviceIdentityListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportDeviceIdentityListRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action importDeviceIdentityList
func (m *ImportDeviceIdentityListRequestBuilder) CreatePostRequestInformation(body ImportDeviceIdentityListPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action importDeviceIdentityList
func (m *ImportDeviceIdentityListRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action importDeviceIdentityList
func (m *ImportDeviceIdentityListRequestBuilder) Post(body ImportDeviceIdentityListPostRequestBodyable)(ImportDeviceIdentityListResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action importDeviceIdentityList
func (m *ImportDeviceIdentityListRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportDeviceIdentityListRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ImportDeviceIdentityListResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateImportDeviceIdentityListResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ImportDeviceIdentityListResponseable), nil
}
