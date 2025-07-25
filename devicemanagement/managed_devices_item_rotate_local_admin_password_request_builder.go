// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder provides operations to call the rotateLocalAdminPassword method.
type ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal instantiates a new ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder and sets the default values.
func NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    m := &ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/rotateLocalAdminPassword", pathParameters),
    }
    return m
}
// NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilder instantiates a new ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder and sets the default values.
func NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiates a manual rotation for the local admin password on the device
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) Post(ctx context.Context, requestConfiguration *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder when successful
func (m *ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemRotateLocalAdminPasswordRequestBuilder) {
    return NewManagedDevicesItemRotateLocalAdminPasswordRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
