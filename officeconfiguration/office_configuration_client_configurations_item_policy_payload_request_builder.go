package officeconfiguration

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder provides operations to manage the media for the officeConfiguration entity.
type OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderInternal instantiates a new PolicyPayloadRequestBuilder and sets the default values.
func NewOfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) {
    m := &OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration/clientConfigurations/{officeClientConfiguration%2Did}/policyPayload";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder instantiates a new PolicyPayloadRequestBuilder and sets the default values.
func NewOfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePutRequestInformation policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) CreatePutRequestInformation(ctx context.Context, body []byte, requestConfiguration *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) Get(ctx context.Context, requestConfiguration *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *OfficeConfigurationClientConfigurationsItemPolicyPayloadRequestBuilderPutRequestConfiguration)(error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
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
