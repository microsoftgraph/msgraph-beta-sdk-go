package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder provides operations to manage the contentVersions property of the microsoft.graph.managedMobileLobApp entity.
type MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetQueryParameters the list of content versions for this app.
type MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderInternal instantiates a new MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) {
    m := &MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedMobileLobApp/contentVersions/{mobileAppContent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder instantiates a new MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContainedApps provides operations to manage the containedApps property of the microsoft.graph.mobileAppContent entity.
// returns a *MobileAppsItemGraphManagedMobileLobAppContentVersionsItemContainedAppsRequestBuilder when successful
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) ContainedApps()(*MobileAppsItemGraphManagedMobileLobAppContentVersionsItemContainedAppsRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppContentVersionsItemContainedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentVersions for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Files provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
// returns a *MobileAppsItemGraphManagedMobileLobAppContentVersionsItemFilesRequestBuilder when successful
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) Files()(*MobileAppsItemGraphManagedMobileLobAppContentVersionsItemFilesRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppContentVersionsItemFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of content versions for this app.
// returns a MobileAppContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentable), nil
}
// Patch update the navigation property contentVersions in deviceAppManagement
// returns a MobileAppContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentable, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentable), nil
}
// ToDeleteRequestInformation delete navigation property contentVersions for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedMobileLobApp/contentVersions/{mobileAppContent%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of content versions for this app.
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property contentVersions in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentable, requestConfiguration *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.managedMobileLobApp/contentVersions/{mobileAppContent%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder when successful
func (m *MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder) {
    return NewMobileAppsItemGraphManagedMobileLobAppContentVersionsMobileAppContentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
