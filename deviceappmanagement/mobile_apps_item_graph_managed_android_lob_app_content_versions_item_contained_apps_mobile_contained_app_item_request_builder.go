package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetQueryParameters the collection of contained apps in a MobileLobApp acting as a package.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderInternal instantiates a new MobileContainedAppItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) {
    m := &MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedAndroidLobApp/contentVersions/{mobileAppContent%2Did}/containedApps/{mobileContainedApp%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder instantiates a new MobileContainedAppItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property containedApps for deviceAppManagement
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of contained apps in a MobileLobApp acting as a package.
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileContainedAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable), nil
}
// Patch update the navigation property containedApps in deviceAppManagement
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileContainedAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable), nil
}
// ToDeleteRequestInformation delete navigation property containedApps for deviceAppManagement
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the collection of contained apps in a MobileLobApp acting as a package.
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property containedApps in deviceAppManagement
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileContainedAppable, requestConfiguration *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder) {
    return NewMobileAppsItemGraphManagedAndroidLobAppContentVersionsItemContainedAppsMobileContainedAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
