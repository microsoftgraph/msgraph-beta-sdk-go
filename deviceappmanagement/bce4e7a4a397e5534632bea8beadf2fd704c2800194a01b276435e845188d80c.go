// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder provides operations to call the convertFromMobileAppCatalogPackage method.
type MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderInternal instantiates a new MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder and sets the default values.
func NewMobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, mobileAppCatalogPackageId *string)(*MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) {
    m := &MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/convertFromMobileAppCatalogPackage(mobileAppCatalogPackageId='{mobileAppCatalogPackageId}')", pathParameters),
    }
    if mobileAppCatalogPackageId != nil {
        m.BaseRequestBuilder.PathParameters["mobileAppCatalogPackageId"] = *mobileAppCatalogPackageId
    }
    return m
}
// NewMobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder instantiates a new MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder and sets the default values.
func NewMobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function convertFromMobileAppCatalogPackage
// returns a MobileAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// ToGetRequestInformation invoke function convertFromMobileAppCatalogPackage
// returns a *RequestInformation when successful
func (m *MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder when successful
func (m *MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) WithUrl(rawUrl string)(*MobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder) {
    return NewMobileAppsConvertFromMobileAppCatalogPackageWithMobileAppCatalogPackageIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
