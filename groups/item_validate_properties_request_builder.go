package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemValidatePropertiesRequestBuilder provides operations to call the validateProperties method.
type ItemValidatePropertiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemValidatePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemValidatePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemValidatePropertiesRequestBuilderInternal instantiates a new ItemValidatePropertiesRequestBuilder and sets the default values.
func NewItemValidatePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemValidatePropertiesRequestBuilder) {
    m := &ItemValidatePropertiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/validateProperties", pathParameters),
    }
    return m
}
// NewItemValidatePropertiesRequestBuilder instantiates a new ItemValidatePropertiesRequestBuilder and sets the default values.
func NewItemValidatePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemValidatePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemValidatePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate if a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use the API to determine if a display name or mail nickname is valid before trying to update a Microsoft 365 group. For validating properties before creating a group, use the validateProperties function for directory objects. The following validations are performed for the display name and mail nickname properties: This API returns with the first failure encountered. If one or more properties fail multiple validations, only the property with the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/group-validateproperties?view=graph-rest-1.0
func (m *ItemValidatePropertiesRequestBuilder) Post(ctx context.Context, body ItemValidatePropertiesPostRequestBodyable, requestConfiguration *ItemValidatePropertiesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation validate if a Microsoft 365 group's display name or mail nickname complies with naming policies. Clients can use the API to determine if a display name or mail nickname is valid before trying to update a Microsoft 365 group. For validating properties before creating a group, use the validateProperties function for directory objects. The following validations are performed for the display name and mail nickname properties: This API returns with the first failure encountered. If one or more properties fail multiple validations, only the property with the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy.
// returns a *RequestInformation when successful
func (m *ItemValidatePropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemValidatePropertiesPostRequestBodyable, requestConfiguration *ItemValidatePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemValidatePropertiesRequestBuilder when successful
func (m *ItemValidatePropertiesRequestBuilder) WithUrl(rawUrl string)(*ItemValidatePropertiesRequestBuilder) {
    return NewItemValidatePropertiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
