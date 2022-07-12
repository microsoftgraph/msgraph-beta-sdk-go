package userpreferencepayload

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserPreferencePayloadRequestBuilder provides operations to manage the media for the officeConfiguration entity.
type UserPreferencePayloadRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserPreferencePayloadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserPreferencePayloadRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserPreferencePayloadRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserPreferencePayloadRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserPreferencePayloadRequestBuilderInternal instantiates a new UserPreferencePayloadRequestBuilder and sets the default values.
func NewUserPreferencePayloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserPreferencePayloadRequestBuilder) {
    m := &UserPreferencePayloadRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration/clientConfigurations/{officeClientConfiguration%2Did}/userPreferencePayload";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserPreferencePayloadRequestBuilder instantiates a new UserPreferencePayloadRequestBuilder and sets the default values.
func NewUserPreferencePayloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserPreferencePayloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserPreferencePayloadRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserPreferencePayloadRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePutRequestInformation preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) CreatePutRequestInformation(body []byte)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePutRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePutRequestInformationWithRequestConfiguration preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) CreatePutRequestInformationWithRequestConfiguration(body []byte, requestConfiguration *UserPreferencePayloadRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) Get()([]byte, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *UserPreferencePayloadRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "[]byte", responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
// Put preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) Put(body []byte)(error) {
    return m.PutWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PutWithRequestConfigurationAndResponseHandler preference settings JSON string in binary format, these values can be overridden by the user.
func (m *UserPreferencePayloadRequestBuilder) PutWithRequestConfigurationAndResponseHandler(body []byte, requestConfiguration *UserPreferencePayloadRequestBuilderPutRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePutRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
