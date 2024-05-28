package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsValidatexmlValidateXmlRequestBuilder provides operations to call the validateXml method.
type MobileappsValidatexmlValidateXmlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsValidatexmlValidateXmlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsValidatexmlValidateXmlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappsValidatexmlValidateXmlRequestBuilderInternal instantiates a new MobileappsValidatexmlValidateXmlRequestBuilder and sets the default values.
func NewMobileappsValidatexmlValidateXmlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsValidatexmlValidateXmlRequestBuilder) {
    m := &MobileappsValidatexmlValidateXmlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/validateXml", pathParameters),
    }
    return m
}
// NewMobileappsValidatexmlValidateXmlRequestBuilder instantiates a new MobileappsValidatexmlValidateXmlRequestBuilder and sets the default values.
func NewMobileappsValidatexmlValidateXmlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsValidatexmlValidateXmlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsValidatexmlValidateXmlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validateXml
// Deprecated: This method is obsolete. Use PostAsValidateXmlPostResponse instead.
// returns a MobileappsValidatexmlValidateXmlResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsValidatexmlValidateXmlRequestBuilder) Post(ctx context.Context, body MobileappsValidatexmlValidateXmlPostRequestBodyable, requestConfiguration *MobileappsValidatexmlValidateXmlRequestBuilderPostRequestConfiguration)(MobileappsValidatexmlValidateXmlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMobileappsValidatexmlValidateXmlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MobileappsValidatexmlValidateXmlResponseable), nil
}
// PostAsValidateXmlPostResponse invoke action validateXml
// returns a MobileappsValidatexmlValidateXmlPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsValidatexmlValidateXmlRequestBuilder) PostAsValidateXmlPostResponse(ctx context.Context, body MobileappsValidatexmlValidateXmlPostRequestBodyable, requestConfiguration *MobileappsValidatexmlValidateXmlRequestBuilderPostRequestConfiguration)(MobileappsValidatexmlValidateXmlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMobileappsValidatexmlValidateXmlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MobileappsValidatexmlValidateXmlPostResponseable), nil
}
// ToPostRequestInformation invoke action validateXml
// returns a *RequestInformation when successful
func (m *MobileappsValidatexmlValidateXmlRequestBuilder) ToPostRequestInformation(ctx context.Context, body MobileappsValidatexmlValidateXmlPostRequestBodyable, requestConfiguration *MobileappsValidatexmlValidateXmlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappsValidatexmlValidateXmlRequestBuilder when successful
func (m *MobileappsValidatexmlValidateXmlRequestBuilder) WithUrl(rawUrl string)(*MobileappsValidatexmlValidateXmlRequestBuilder) {
    return NewMobileappsValidatexmlValidateXmlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
