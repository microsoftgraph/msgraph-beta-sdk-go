package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder provides operations to call the rotateLocalAdminPassword method.
type ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderInternal instantiates a new ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder and sets the default values.
func NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) {
    m := &ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/rotateLocalAdminPassword", pathParameters),
    }
    return m
}
// NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder instantiates a new ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder and sets the default values.
func NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiates a manual rotation for the local admin password on the device
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation initiates a manual rotation for the local admin password on the device
// returns a *RequestInformation when successful
func (m *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder when successful
func (m *ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder) {
    return NewItemManageddevicesItemRotatelocaladminpasswordRotateLocalAdminPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
