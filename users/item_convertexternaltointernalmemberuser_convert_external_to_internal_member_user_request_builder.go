package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder provides operations to call the convertExternalToInternalMemberUser method.
type ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderInternal instantiates a new ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder and sets the default values.
func NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) {
    m := &ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/convertExternalToInternalMemberUser", pathParameters),
    }
    return m
}
// NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder instantiates a new ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder and sets the default values.
func NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Post convert an externally authenticated user into an internal user. The user is able to sign into the host tenant as an internal user and access resources as a member. For more information about this conversion, see Convert external users to internal users.
// returns a ConversionUserDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-convertexternaltointernalmemberuser?view=graph-rest-beta
func (m *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) Post(ctx context.Context, body ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserPostRequestBodyable, requestConfiguration *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversionUserDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateConversionUserDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ConversionUserDetailsable), nil
}
// ToPostRequestInformation convert an externally authenticated user into an internal user. The user is able to sign into the host tenant as an internal user and access resources as a member. For more information about this conversion, see Convert external users to internal users.
// returns a *RequestInformation when successful
func (m *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserPostRequestBodyable, requestConfiguration *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder when successful
func (m *ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) WithUrl(rawUrl string)(*ItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder) {
    return NewItemConvertexternaltointernalmemberuserConvertExternalToInternalMemberUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
