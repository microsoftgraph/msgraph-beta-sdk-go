package officeconfiguration

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ClientConfigurationsItemUserPreferencePayloadRequestBuilder provides operations to manage the media for the officeConfiguration entity.
type ClientConfigurationsItemUserPreferencePayloadRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClientConfigurationsItemUserPreferencePayloadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClientConfigurationsItemUserPreferencePayloadRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClientConfigurationsItemUserPreferencePayloadRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClientConfigurationsItemUserPreferencePayloadRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClientConfigurationsItemUserPreferencePayloadRequestBuilderInternal instantiates a new UserPreferencePayloadRequestBuilder and sets the default values.
func NewClientConfigurationsItemUserPreferencePayloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientConfigurationsItemUserPreferencePayloadRequestBuilder) {
    m := &ClientConfigurationsItemUserPreferencePayloadRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/officeConfiguration/clientConfigurations/{officeClientConfiguration%2Did}/userPreferencePayload", pathParameters),
    }
    return m
}
// NewClientConfigurationsItemUserPreferencePayloadRequestBuilder instantiates a new UserPreferencePayloadRequestBuilder and sets the default values.
func NewClientConfigurationsItemUserPreferencePayloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClientConfigurationsItemUserPreferencePayloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClientConfigurationsItemUserPreferencePayloadRequestBuilderInternal(urlParams, requestAdapter)
}
// Get preference settings JSON string in binary format, these values can be overridden by the user.
func (m *ClientConfigurationsItemUserPreferencePayloadRequestBuilder) Get(ctx context.Context, requestConfiguration *ClientConfigurationsItemUserPreferencePayloadRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put preference settings JSON string in binary format, these values can be overridden by the user.
func (m *ClientConfigurationsItemUserPreferencePayloadRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ClientConfigurationsItemUserPreferencePayloadRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToGetRequestInformation preference settings JSON string in binary format, these values can be overridden by the user.
func (m *ClientConfigurationsItemUserPreferencePayloadRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClientConfigurationsItemUserPreferencePayloadRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPutRequestInformation preference settings JSON string in binary format, these values can be overridden by the user.
func (m *ClientConfigurationsItemUserPreferencePayloadRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ClientConfigurationsItemUserPreferencePayloadRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.SetStreamContent(body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
