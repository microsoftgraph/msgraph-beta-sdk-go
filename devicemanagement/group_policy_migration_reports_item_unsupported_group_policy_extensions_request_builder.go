package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetQueryParameters a list of unsupported group policy extensions inside the Group Policy Object.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetQueryParameters
}
// GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnsupportedGroupPolicyExtensionId provides operations to manage the unsupportedGroupPolicyExtensions property of the microsoft.graph.groupPolicyMigrationReport entity.
// returns a *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder when successful
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) ByUnsupportedGroupPolicyExtensionId(unsupportedGroupPolicyExtensionId string)(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unsupportedGroupPolicyExtensionId != "" {
        urlTplParams["unsupportedGroupPolicyExtension%2Did"] = unsupportedGroupPolicyExtensionId
    }
    return NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderInternal instantiates a new GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) {
    m := &GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/unsupportedGroupPolicyExtensions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder instantiates a new GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsCountRequestBuilder when successful
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) Count()(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsCountRequestBuilder) {
    return NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of unsupported group policy extensions inside the Group Policy Object.
// returns a UnsupportedGroupPolicyExtensionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnsupportedGroupPolicyExtensionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionCollectionResponseable), nil
}
// Post create new navigation property to unsupportedGroupPolicyExtensions for deviceManagement
// returns a UnsupportedGroupPolicyExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation a list of unsupported group policy extensions inside the Group Policy Object.
// returns a *RequestInformation when successful
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to unsupportedGroupPolicyExtensions for deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnsupportedGroupPolicyExtensionable, requestConfiguration *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder when successful
func (m *GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder) {
    return NewGroupPolicyMigrationReportsItemUnsupportedGroupPolicyExtensionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
