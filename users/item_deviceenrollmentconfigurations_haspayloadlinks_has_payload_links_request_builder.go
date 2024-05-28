package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderInternal instantiates a new ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder and sets the default values.
func NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) {
    m := &ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder instantiates a new ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder and sets the default values.
func NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use PostAsHasPayloadLinksPostResponse instead.
// returns a ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) Post(ctx context.Context, body ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
// returns a ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
// returns a *RequestInformation when successful
func (m *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
