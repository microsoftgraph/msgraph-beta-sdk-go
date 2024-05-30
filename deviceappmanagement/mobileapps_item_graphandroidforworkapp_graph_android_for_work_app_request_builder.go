package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder casts the previous resource to androidForWorkApp.
type MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.androidForWorkApp
type MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphandroidforworkappAssignmentsRequestBuilder when successful
func (m *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) Assignments()(*MobileappsItemGraphandroidforworkappAssignmentsRequestBuilder) {
    return NewMobileappsItemGraphandroidforworkappAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphandroidforworkappCategoriesRequestBuilder when successful
func (m *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) Categories()(*MobileappsItemGraphandroidforworkappCategoriesRequestBuilder) {
    return NewMobileappsItemGraphandroidforworkappCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderInternal instantiates a new MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) {
    m := &MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.androidForWorkApp{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder instantiates a new MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder and sets the default values.
func NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.androidForWorkApp
// returns a AndroidForWorkAppable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidForWorkAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidForWorkAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
// returns a *MobileappsItemGraphandroidforworkappRelationshipsRequestBuilder when successful
func (m *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) Relationships()(*MobileappsItemGraphandroidforworkappRelationshipsRequestBuilder) {
    return NewMobileappsItemGraphandroidforworkappRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.androidForWorkApp
// returns a *RequestInformation when successful
func (m *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder when successful
func (m *MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) WithUrl(rawUrl string)(*MobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder) {
    return NewMobileappsItemGraphandroidforworkappGraphAndroidForWorkAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
