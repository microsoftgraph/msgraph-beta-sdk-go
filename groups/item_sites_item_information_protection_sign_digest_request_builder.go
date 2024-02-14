package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionSignDigestRequestBuilder provides operations to call the signDigest method.
type ItemSitesItemInformationProtectionSignDigestRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionSignDigestRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionSignDigestRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionSignDigestRequestBuilderInternal instantiates a new ItemSitesItemInformationProtectionSignDigestRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionSignDigestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionSignDigestRequestBuilder) {
    m := &ItemSitesItemInformationProtectionSignDigestRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/signDigest", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionSignDigestRequestBuilder instantiates a new ItemSitesItemInformationProtectionSignDigestRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionSignDigestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionSignDigestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionSignDigestRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action signDigest
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a SigningResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionSignDigestRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationProtectionSignDigestPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionSignDigestRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SigningResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSigningResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SigningResultable), nil
}
// ToPostRequestInformation invoke action signDigest
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionSignDigestRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationProtectionSignDigestPostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionSignDigestRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: This API will no longer be accessible, please see microsoft.graph.security.informationProtection APIs. as of 2021-02/Beta_SensitivityLabels
// returns a *ItemSitesItemInformationProtectionSignDigestRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionSignDigestRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionSignDigestRequestBuilder) {
    return NewItemSitesItemInformationProtectionSignDigestRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
