package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemRemotelockRemoteLockRequestBuilder provides operations to call the remoteLock method.
type ItemManageddevicesItemRemotelockRemoteLockRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemRemotelockRemoteLockRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemRemotelockRemoteLockRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemRemotelockRemoteLockRequestBuilderInternal instantiates a new ItemManageddevicesItemRemotelockRemoteLockRequestBuilder and sets the default values.
func NewItemManageddevicesItemRemotelockRemoteLockRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    m := &ItemManageddevicesItemRemotelockRemoteLockRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/remoteLock", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemRemotelockRemoteLockRequestBuilder instantiates a new ItemManageddevicesItemRemotelockRemoteLockRequestBuilder and sets the default values.
func NewItemManageddevicesItemRemotelockRemoteLockRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemRemotelockRemoteLockRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remote lock
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemRemotelockRemoteLockRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remote lock
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemRemotelockRemoteLockRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemRemotelockRemoteLockRequestBuilder when successful
func (m *ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemRemotelockRemoteLockRequestBuilder) {
    return NewItemManageddevicesItemRemotelockRemoteLockRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
