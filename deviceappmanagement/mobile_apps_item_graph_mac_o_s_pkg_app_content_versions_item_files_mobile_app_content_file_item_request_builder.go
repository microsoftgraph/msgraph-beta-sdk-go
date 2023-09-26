package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder provides operations to manage the files property of the microsoft.graph.mobileAppContent entity.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters the list of files for this app content version.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetQueryParameters
}
// MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Commit provides operations to call the commit method.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) Commit()(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesItemCommitRequestBuilder) {
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesItemCommitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderInternal instantiates a new MobileAppContentFileItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) {
    m := &MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.macOSPkgApp/contentVersions/{mobileAppContent%2Did}/files/{mobileAppContentFile%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder instantiates a new MobileAppContentFileItemRequestBuilder and sets the default values.
func NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property files for deviceAppManagement
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of files for this app content version.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable), nil
}
// Patch update the navigation property files in deviceAppManagement
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppContentFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable), nil
}
// RenewUpload provides operations to call the renewUpload method.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) RenewUpload()(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesItemRenewUploadRequestBuilder) {
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesItemRenewUploadRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property files for deviceAppManagement
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation the list of files for this app content version.
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property files in deviceAppManagement
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppContentFileable, requestConfiguration *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder) {
    return NewMobileAppsItemGraphMacOSPkgAppContentVersionsItemFilesMobileAppContentFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
