package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder provides operations to call the validatePermission method.
type FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/validatePermission", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action validatePermission
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action validatePermission
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemValidatepermissionValidatePermissionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
