package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder provides operations to call the importDeviceIdentityList method.
type ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderInternal instantiates a new ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder and sets the default values.
func NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) {
    m := &ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedDeviceIdentities/importDeviceIdentityList", pathParameters),
    }
    return m
}
// NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder instantiates a new ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder and sets the default values.
func NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importDeviceIdentityList
// Deprecated: This method is obsolete. Use PostAsImportDeviceIdentityListPostResponse instead.
// returns a ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) Post(ctx context.Context, body ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListResponseable), nil
}
// PostAsImportDeviceIdentityListPostResponse invoke action importDeviceIdentityList
// returns a ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) PostAsImportDeviceIdentityListPostResponse(ctx context.Context, body ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostResponseable), nil
}
// ToPostRequestInformation invoke action importDeviceIdentityList
// returns a *RequestInformation when successful
func (m *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) ToPostRequestInformation(ctx context.Context, body ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListPostRequestBodyable, requestConfiguration *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder when successful
func (m *ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) WithUrl(rawUrl string)(*ImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder) {
    return NewImporteddeviceidentitiesImportdeviceidentitylistImportDeviceIdentityListRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
