package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetQueryParameters a list of unsupported group policy extensions inside the Group Policy Object.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetQueryParameters
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal instantiates a new UnsupportedGroupPolicyExtensionItemRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) {
    m := &GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/unsupportedGroupPolicyExtensions/{unsupportedGroupPolicyExtension%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder instantiates a new UnsupportedGroupPolicyExtensionItemRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property unsupportedGroupPolicyExtensions for deviceManagement
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get a list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable), nil
}
// Patch update the navigation property unsupportedGroupPolicyExtensions in deviceManagement
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable), nil
}
// ToDeleteRequestInformation delete navigation property unsupportedGroupPolicyExtensions for deviceManagement
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation a list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property unsupportedGroupPolicyExtensions in deviceManagement
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
