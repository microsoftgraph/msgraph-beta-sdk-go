package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder provides operations to call the hasPayloadLinks method.
type MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal instantiates a new MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    m := &MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mdmWindowsInformationProtectionPolicies/hasPayloadLinks", pathParameters),
    }
    return m
}
// NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder instantiates a new MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder and sets the default values.
func NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action hasPayloadLinks
// Deprecated: This method is obsolete. Use PostAsHasPayloadLinksPostResponse instead.
// returns a MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) Post(ctx context.Context, body MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksResponseable), nil
}
// PostAsHasPayloadLinksPostResponse invoke action hasPayloadLinks
// returns a MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) PostAsHasPayloadLinksPostResponse(ctx context.Context, body MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostResponseable), nil
}
// ToPostRequestInformation invoke action hasPayloadLinks
// returns a *RequestInformation when successful
func (m *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksPostRequestBodyable, requestConfiguration *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) WithUrl(rawUrl string)(*MdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewMdmwindowsinformationprotectionpoliciesHaspayloadlinksHasPayloadLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
