package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAssignlicenseAssignLicenseRequestBuilder provides operations to call the assignLicense method.
type ItemAssignlicenseAssignLicenseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAssignlicenseAssignLicenseRequestBuilderInternal instantiates a new ItemAssignlicenseAssignLicenseRequestBuilder and sets the default values.
func NewItemAssignlicenseAssignLicenseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    m := &ItemAssignlicenseAssignLicenseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/assignLicense", pathParameters),
    }
    return m
}
// NewItemAssignlicenseAssignLicenseRequestBuilder instantiates a new ItemAssignlicenseAssignLicenseRequestBuilder and sets the default values.
func NewItemAssignlicenseAssignLicenseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAssignlicenseAssignLicenseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add or remove licenses for the user to enable or disable their use of Microsoft cloud offerings. For example, an organization can have a Microsoft 365 Enterprise E3 subscription with 100 licenses, and this request assigns one of those licenses to a specific user. You can also enable and disable specific plans associated with a subscription. To learn more about subscriptions and licenses, see this Technet article. To get the subscriptions available in the directory, perform a GET subscribedSkus request. 
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-assignlicense?view=graph-rest-beta
func (m *ItemAssignlicenseAssignLicenseRequestBuilder) Post(ctx context.Context, body ItemAssignlicenseAssignLicensePostRequestBodyable, requestConfiguration *ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// ToPostRequestInformation add or remove licenses for the user to enable or disable their use of Microsoft cloud offerings. For example, an organization can have a Microsoft 365 Enterprise E3 subscription with 100 licenses, and this request assigns one of those licenses to a specific user. You can also enable and disable specific plans associated with a subscription. To learn more about subscriptions and licenses, see this Technet article. To get the subscriptions available in the directory, perform a GET subscribedSkus request. 
// returns a *RequestInformation when successful
func (m *ItemAssignlicenseAssignLicenseRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAssignlicenseAssignLicensePostRequestBodyable, requestConfiguration *ItemAssignlicenseAssignLicenseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAssignlicenseAssignLicenseRequestBuilder when successful
func (m *ItemAssignlicenseAssignLicenseRequestBuilder) WithUrl(rawUrl string)(*ItemAssignlicenseAssignLicenseRequestBuilder) {
    return NewItemAssignlicenseAssignLicenseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
