package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4d9d201299918a35fb4e4cf579ccd2a733831ba18bfc3b796f25ff72712c65ff "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagedappprotections/item/apps"
    i4ed8a48e1a91904077b1c7b41a39725b3ff69fd77ff37f36cc5ca55a3771fcd2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagedappprotections/item/assign"
    iaea643d7019f43c457c70b1d3282412e6e2fcac099fb44847cfee5e56e6c4b6c "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagedappprotections/item/targetapps"
    ie2b148f923b331e886164975a27dc60b5fee8486b6ff177093e08d135da109c0 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagedappprotections/item/assignments"
    i04796f9559857226eaa48cd43add7d74f4c5fdb3061437f60850d3556e9ea211 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagedappprotections/item/assignments/item"
    i0cc29f4a22b83e5ca5d2e726aed9e5dd7d1e0b6491976a69c92134e4188679a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/windowsmanagedappprotections/item/apps/item"
)

// WindowsManagedAppProtectionItemRequestBuilder provides operations to manage the windowsManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
type WindowsManagedAppProtectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsManagedAppProtectionItemRequestBuilderGetQueryParameters windows managed app policies.
type WindowsManagedAppProtectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsManagedAppProtectionItemRequestBuilderGetQueryParameters
}
// WindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Apps the apps property
func (m *WindowsManagedAppProtectionItemRequestBuilder) Apps()(*i4d9d201299918a35fb4e4cf579ccd2a733831ba18bfc3b796f25ff72712c65ff.AppsRequestBuilder) {
    return i4d9d201299918a35fb4e4cf579ccd2a733831ba18bfc3b796f25ff72712c65ff.NewAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.windowsManagedAppProtections.item.apps.item collection
func (m *WindowsManagedAppProtectionItemRequestBuilder) AppsById(id string)(*i0cc29f4a22b83e5ca5d2e726aed9e5dd7d1e0b6491976a69c92134e4188679a7.ManagedMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedMobileApp%2Did"] = id
    }
    return i0cc29f4a22b83e5ca5d2e726aed9e5dd7d1e0b6491976a69c92134e4188679a7.NewManagedMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assign the assign property
func (m *WindowsManagedAppProtectionItemRequestBuilder) Assign()(*i4ed8a48e1a91904077b1c7b41a39725b3ff69fd77ff37f36cc5ca55a3771fcd2.AssignRequestBuilder) {
    return i4ed8a48e1a91904077b1c7b41a39725b3ff69fd77ff37f36cc5ca55a3771fcd2.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *WindowsManagedAppProtectionItemRequestBuilder) Assignments()(*ie2b148f923b331e886164975a27dc60b5fee8486b6ff177093e08d135da109c0.AssignmentsRequestBuilder) {
    return ie2b148f923b331e886164975a27dc60b5fee8486b6ff177093e08d135da109c0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.windowsManagedAppProtections.item.assignments.item collection
func (m *WindowsManagedAppProtectionItemRequestBuilder) AssignmentsById(id string)(*i04796f9559857226eaa48cd43add7d74f4c5fdb3061437f60850d3556e9ea211.TargetedManagedAppPolicyAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppPolicyAssignment%2Did"] = id
    }
    return i04796f9559857226eaa48cd43add7d74f4c5fdb3061437f60850d3556e9ea211.NewTargetedManagedAppPolicyAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewWindowsManagedAppProtectionItemRequestBuilderInternal instantiates a new WindowsManagedAppProtectionItemRequestBuilder and sets the default values.
func NewWindowsManagedAppProtectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsManagedAppProtectionItemRequestBuilder) {
    m := &WindowsManagedAppProtectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/windowsManagedAppProtections/{windowsManagedAppProtection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsManagedAppProtectionItemRequestBuilder instantiates a new WindowsManagedAppProtectionItemRequestBuilder and sets the default values.
func NewWindowsManagedAppProtectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsManagedAppProtectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsManagedAppProtectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property windowsManagedAppProtections for deviceAppManagement
func (m *WindowsManagedAppProtectionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property windowsManagedAppProtections for deviceAppManagement
func (m *WindowsManagedAppProtectionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *WindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation windows managed app policies.
func (m *WindowsManagedAppProtectionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration windows managed app policies.
func (m *WindowsManagedAppProtectionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *WindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property windowsManagedAppProtections in deviceAppManagement
func (m *WindowsManagedAppProtectionItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property windowsManagedAppProtections in deviceAppManagement
func (m *WindowsManagedAppProtectionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *WindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property windowsManagedAppProtections for deviceAppManagement
func (m *WindowsManagedAppProtectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsManagedAppProtectionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get windows managed app policies.
func (m *WindowsManagedAppProtectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsManagedAppProtectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsManagedAppProtectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable), nil
}
// Patch update the navigation property windowsManagedAppProtections in deviceAppManagement
func (m *WindowsManagedAppProtectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsManagedAppProtectionable, requestConfiguration *WindowsManagedAppProtectionItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// TargetApps the targetApps property
func (m *WindowsManagedAppProtectionItemRequestBuilder) TargetApps()(*iaea643d7019f43c457c70b1d3282412e6e2fcac099fb44847cfee5e56e6c4b6c.TargetAppsRequestBuilder) {
    return iaea643d7019f43c457c70b1d3282412e6e2fcac099fb44847cfee5e56e6c4b6c.NewTargetAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
