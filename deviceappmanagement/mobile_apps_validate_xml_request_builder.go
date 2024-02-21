package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsValidateXmlRequestBuilder provides operations to call the validateXml method.
type MobileAppsValidateXmlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsValidateXmlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsValidateXmlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsValidateXmlRequestBuilderInternal instantiates a new MobileAppsValidateXmlRequestBuilder and sets the default values.
func NewMobileAppsValidateXmlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsValidateXmlRequestBuilder) {
    m := &MobileAppsValidateXmlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/validateXml", pathParameters),
    }
    return m
}
// NewMobileAppsValidateXmlRequestBuilder instantiates a new MobileAppsValidateXmlRequestBuilder and sets the default values.
func NewMobileAppsValidateXmlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsValidateXmlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsValidateXmlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateXml
// Deprecated: This method is obsolete. Use PostAsValidateXmlPostResponse instead.
// returns a MobileAppsValidateXmlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsValidateXmlRequestBuilder) Post(ctx context.Context, body MobileAppsValidateXmlPostRequestBodyable, requestConfiguration *MobileAppsValidateXmlRequestBuilderPostRequestConfiguration)(MobileAppsValidateXmlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMobileAppsValidateXmlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MobileAppsValidateXmlResponseable), nil
}
// PostAsValidateXmlPostResponse invoke action validateXml
// returns a MobileAppsValidateXmlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsValidateXmlRequestBuilder) PostAsValidateXmlPostResponse(ctx context.Context, body MobileAppsValidateXmlPostRequestBodyable, requestConfiguration *MobileAppsValidateXmlRequestBuilderPostRequestConfiguration)(MobileAppsValidateXmlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMobileAppsValidateXmlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MobileAppsValidateXmlPostResponseable), nil
}
// ToPostRequestInformation invoke action validateXml
// returns a *RequestInformation when successful
func (m *MobileAppsValidateXmlRequestBuilder) ToPostRequestInformation(ctx context.Context, body MobileAppsValidateXmlPostRequestBodyable, requestConfiguration *MobileAppsValidateXmlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileAppsValidateXmlRequestBuilder when successful
func (m *MobileAppsValidateXmlRequestBuilder) WithUrl(rawUrl string)(*MobileAppsValidateXmlRequestBuilder) {
    return NewMobileAppsValidateXmlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
