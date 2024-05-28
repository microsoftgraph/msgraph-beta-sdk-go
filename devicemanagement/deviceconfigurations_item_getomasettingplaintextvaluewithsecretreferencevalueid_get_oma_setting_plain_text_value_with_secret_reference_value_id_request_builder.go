package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder provides operations to call the getOmaSettingPlainTextValue method.
type DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal instantiates a new DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, secretReferenceValueId *string)(*DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    m := &DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/getOmaSettingPlainTextValue(secretReferenceValueId='{secretReferenceValueId}')", pathParameters),
    }
    if secretReferenceValueId != nil {
        m.BaseRequestBuilder.PathParameters["secretReferenceValueId"] = *secretReferenceValueId
    }
    return m
}
// NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder instantiates a new DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOmaSettingPlainTextValue
// Deprecated: This method is obsolete. Use GetAsGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponse instead.
// returns a DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdResponseable), nil
}
// GetAsGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponse invoke function getOmaSettingPlainTextValue
// returns a DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) GetAsGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponse(ctx context.Context, requestConfiguration *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdGetResponseable), nil
}
// ToGetRequestInformation invoke function getOmaSettingPlainTextValue
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder when successful
func (m *DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder) {
    return NewDeviceconfigurationsItemGetomasettingplaintextvaluewithsecretreferencevalueidGetOmaSettingPlainTextValueWithSecretReferenceValueIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
