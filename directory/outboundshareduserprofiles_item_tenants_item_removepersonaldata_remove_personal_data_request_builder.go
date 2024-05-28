package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder provides operations to call the removePersonalData method.
type OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderInternal instantiates a new OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) {
    m := &OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile%2DuserId}/tenants/{tenantReference%2DtenantId}/removePersonalData", pathParameters),
    }
    return m
}
// NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder instantiates a new OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder and sets the default values.
func NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a request to remove the personal data for an outboundSharedUserProfile.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tenantreference-removepersonaldata?view=graph-rest-beta
func (m *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) Post(ctx context.Context, requestConfiguration *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation create a request to remove the personal data for an outboundSharedUserProfile.
// returns a *RequestInformation when successful
func (m *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder when successful
func (m *OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) WithUrl(rawUrl string)(*OutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder) {
    return NewOutboundshareduserprofilesItemTenantsItemRemovepersonaldataRemovePersonalDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
