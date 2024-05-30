package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder provides operations to manage the privilegeManagementElevations property of the microsoft.graph.deviceManagement entity.
type PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetQueryParameters the endpoint privilege management elevation event entity contains elevation details.
type PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetQueryParameters struct {
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
// PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetQueryParameters
}
// PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByPrivilegeManagementElevationId provides operations to manage the privilegeManagementElevations property of the microsoft.graph.deviceManagement entity.
// returns a *PrivilegemanagementelevationsPrivilegeManagementElevationItemRequestBuilder when successful
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) ByPrivilegeManagementElevationId(privilegeManagementElevationId string)(*PrivilegemanagementelevationsPrivilegeManagementElevationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if privilegeManagementElevationId != "" {
        urlTplParams["privilegeManagementElevation%2Did"] = privilegeManagementElevationId
    }
    return NewPrivilegemanagementelevationsPrivilegeManagementElevationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderInternal instantiates a new PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder and sets the default values.
func NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) {
    m := &PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/privilegeManagementElevations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder instantiates a new PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder and sets the default values.
func NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *PrivilegemanagementelevationsCountRequestBuilder when successful
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) Count()(*PrivilegemanagementelevationsCountRequestBuilder) {
    return NewPrivilegemanagementelevationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the endpoint privilege management elevation event entity contains elevation details.
// returns a PrivilegeManagementElevationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegeManagementElevationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationCollectionResponseable), nil
}
// Post create new navigation property to privilegeManagementElevations for deviceManagement
// returns a PrivilegeManagementElevationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationable, requestConfiguration *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegeManagementElevationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationable), nil
}
// ToGetRequestInformation the endpoint privilege management elevation event entity contains elevation details.
// returns a *RequestInformation when successful
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to privilegeManagementElevations for deviceManagement
// returns a *RequestInformation when successful
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegeManagementElevationable, requestConfiguration *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder when successful
func (m *PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) WithUrl(rawUrl string)(*PrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder) {
    return NewPrivilegemanagementelevationsPrivilegeManagementElevationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
