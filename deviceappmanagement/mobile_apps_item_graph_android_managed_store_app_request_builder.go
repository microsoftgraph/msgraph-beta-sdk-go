package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder casts the previous resource to androidManagedStoreApp.
type MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetQueryParameters get the item of type microsoft.graph.mobileApp as microsoft.graph.androidManagedStoreApp
type MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetQueryParameters
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) Assignments()(*MobileAppsItemGraphAndroidManagedStoreAppAssignmentsRequestBuilder) {
    return NewMobileAppsItemGraphAndroidManagedStoreAppAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Categories provides operations to manage the categories property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) Categories()(*MobileAppsItemGraphAndroidManagedStoreAppCategoriesRequestBuilder) {
    return NewMobileAppsItemGraphAndroidManagedStoreAppCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilderInternal instantiates a new GraphAndroidManagedStoreAppRequestBuilder and sets the default values.
func NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) {
    m := &MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/graph.androidManagedStoreApp{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilder instantiates a new GraphAndroidManagedStoreAppRequestBuilder and sets the default values.
func NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the item of type microsoft.graph.mobileApp as microsoft.graph.androidManagedStoreApp
func (m *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAppFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppable), nil
}
// Relationships provides operations to manage the relationships property of the microsoft.graph.mobileApp entity.
func (m *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) Relationships()(*MobileAppsItemGraphAndroidManagedStoreAppRelationshipsRequestBuilder) {
    return NewMobileAppsItemGraphAndroidManagedStoreAppRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the item of type microsoft.graph.mobileApp as microsoft.graph.androidManagedStoreApp
func (m *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) WithUrl(rawUrl string)(*MobileAppsItemGraphAndroidManagedStoreAppRequestBuilder) {
    return NewMobileAppsItemGraphAndroidManagedStoreAppRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
