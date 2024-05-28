package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder provides operations to call the enableLostMode method.
type ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderInternal instantiates a new ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder and sets the default values.
func NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) {
    m := &ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/enableLostMode", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder instantiates a new ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder and sets the default values.
func NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enable lost mode
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) Post(ctx context.Context, body ItemManageddevicesItemEnablelostmodeEnableLostModePostRequestBodyable, requestConfiguration *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation enable lost mode
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemManageddevicesItemEnablelostmodeEnableLostModePostRequestBodyable, requestConfiguration *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder when successful
func (m *ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder) {
    return NewItemManageddevicesItemEnablelostmodeEnableLostModeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
