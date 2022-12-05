package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
type DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetQueryParameters a list of unsupported group policy extensions inside the Group Policy Object.
type DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal instantiates a new UnsupportedGroupPolicyExtensionItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/unsupportedGroupPolicyExtensions/{unsupportedGroupPolicyExtension%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder instantiates a new UnsupportedGroupPolicyExtensionItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property unsupportedGroupPolicyExtensions for deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a list of unsupported group policy extensions inside the Group Policy Object.
func (m *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property unsupportedGroupPolicyExtensions in deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property unsupportedGroupPolicyExtensions for deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get a list of unsupported group policy extensions inside the Group Policy Object.
func (m *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable), nil
}
// Patch update the navigation property unsupportedGroupPolicyExtensions in deviceManagement
func (m *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, requestConfiguration *DeviceManagementGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable), nil
}
