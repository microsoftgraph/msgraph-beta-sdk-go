package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SolutionsRequestBuilder provides operations to manage the solutionsRoot singleton.
type SolutionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SolutionsRequestBuilderGetQueryParameters get solutions
type SolutionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SolutionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SolutionsRequestBuilderGetQueryParameters
}
// SolutionsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SolutionsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BackupRestore provides operations to manage the backupRestore property of the microsoft.graph.solutionsRoot entity.
// returns a *BackuprestoreBackupRestoreRequestBuilder when successful
func (m *SolutionsRequestBuilder) BackupRestore()(*BackuprestoreBackupRestoreRequestBuilder) {
    return NewBackuprestoreBackupRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BookingBusinesses provides operations to manage the bookingBusinesses property of the microsoft.graph.solutionsRoot entity.
// returns a *BookingbusinessesBookingBusinessesRequestBuilder when successful
func (m *SolutionsRequestBuilder) BookingBusinesses()(*BookingbusinessesBookingBusinessesRequestBuilder) {
    return NewBookingbusinessesBookingBusinessesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BookingCurrencies provides operations to manage the bookingCurrencies property of the microsoft.graph.solutionsRoot entity.
// returns a *BookingcurrenciesBookingCurrenciesRequestBuilder when successful
func (m *SolutionsRequestBuilder) BookingCurrencies()(*BookingcurrenciesBookingCurrenciesRequestBuilder) {
    return NewBookingcurrenciesBookingCurrenciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BusinessScenarios provides operations to manage the businessScenarios property of the microsoft.graph.solutionsRoot entity.
// returns a *BusinessscenariosBusinessScenariosRequestBuilder when successful
func (m *SolutionsRequestBuilder) BusinessScenarios()(*BusinessscenariosBusinessScenariosRequestBuilder) {
    return NewBusinessscenariosBusinessScenariosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BusinessScenariosWithUniqueName provides operations to manage the businessScenarios property of the microsoft.graph.solutionsRoot entity.
// returns a *BusinessscenarioswithuniquenameBusinessScenariosWithUniqueNameRequestBuilder when successful
func (m *SolutionsRequestBuilder) BusinessScenariosWithUniqueName(uniqueName *string)(*BusinessscenarioswithuniquenameBusinessScenariosWithUniqueNameRequestBuilder) {
    return NewBusinessscenarioswithuniquenameBusinessScenariosWithUniqueNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, uniqueName)
}
// NewSolutionsRequestBuilderInternal instantiates a new SolutionsRequestBuilder and sets the default values.
func NewSolutionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsRequestBuilder) {
    m := &SolutionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewSolutionsRequestBuilder instantiates a new SolutionsRequestBuilder and sets the default values.
func NewSolutionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SolutionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSolutionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get solutions
// returns a SolutionsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SolutionsRequestBuilder) Get(ctx context.Context, requestConfiguration *SolutionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SolutionsRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSolutionsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SolutionsRootable), nil
}
// Patch update solutions
// returns a SolutionsRootable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SolutionsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SolutionsRootable, requestConfiguration *SolutionsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SolutionsRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSolutionsRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SolutionsRootable), nil
}
// ToGetRequestInformation get solutions
// returns a *RequestInformation when successful
func (m *SolutionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SolutionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update solutions
// returns a *RequestInformation when successful
func (m *SolutionsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SolutionsRootable, requestConfiguration *SolutionsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// VirtualEvents provides operations to manage the virtualEvents property of the microsoft.graph.solutionsRoot entity.
// returns a *VirtualeventsVirtualEventsRequestBuilder when successful
func (m *SolutionsRequestBuilder) VirtualEvents()(*VirtualeventsVirtualEventsRequestBuilder) {
    return NewVirtualeventsVirtualEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SolutionsRequestBuilder when successful
func (m *SolutionsRequestBuilder) WithUrl(rawUrl string)(*SolutionsRequestBuilder) {
    return NewSolutionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
