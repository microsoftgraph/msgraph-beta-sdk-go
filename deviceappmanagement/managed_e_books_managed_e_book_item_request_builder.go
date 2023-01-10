package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedEBooksManagedEBookItemRequestBuilder provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
type ManagedEBooksManagedEBookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters the Managed eBook.
type ManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedEBooksManagedEBookItemRequestBuilderGetQueryParameters
}
// ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Assign()(*ManagedEBooksItemAssignRequestBuilder) {
    return NewManagedEBooksItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Assignments()(*ManagedEBooksItemAssignmentsRequestBuilder) {
    return NewManagedEBooksItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) AssignmentsById(id string)(*ManagedEBooksItemAssignmentsManagedEBookAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment%2Did"] = id
    }
    return NewManagedEBooksItemAssignmentsManagedEBookAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories provides operations to manage the categories property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Categories()(*ManagedEBooksItemCategoriesRequestBuilder) {
    return NewManagedEBooksItemCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById provides operations to manage the categories property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) CategoriesById(id string)(*ManagedEBooksItemCategoriesManagedEBookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookCategory%2Did"] = id
    }
    return NewManagedEBooksItemCategoriesManagedEBookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedEBooksManagedEBookItemRequestBuilderInternal instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBooksManagedEBookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBooksManagedEBookItemRequestBuilder) {
    m := &ManagedEBooksManagedEBookItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedEBooksManagedEBookItemRequestBuilder instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBooksManagedEBookItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBooksManagedEBookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBooksManagedEBookItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// DeviceStates provides operations to manage the deviceStates property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) DeviceStates()(*ManagedEBooksItemDeviceStatesRequestBuilder) {
    return NewManagedEBooksItemDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById provides operations to manage the deviceStates property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) DeviceStatesById(id string)(*ManagedEBooksItemDeviceStatesDeviceInstallStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState%2Did"] = id
    }
    return NewManagedEBooksItemDeviceStatesDeviceInstallStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the Managed eBook.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedEBookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable), nil
}
// InstallSummary provides operations to manage the installSummary property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) InstallSummary()(*ManagedEBooksItemInstallSummaryRequestBuilder) {
    return NewManagedEBooksItemInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedEBookFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable), nil
}
// ToDeleteRequestInformation delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the Managed eBook.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBooksManagedEBookItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *ManagedEBooksManagedEBookItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserStateSummary provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) UserStateSummary()(*ManagedEBooksItemUserStateSummaryRequestBuilder) {
    return NewManagedEBooksItemUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStateSummaryById provides operations to manage the userStateSummary property of the microsoft.graph.managedEBook entity.
func (m *ManagedEBooksManagedEBookItemRequestBuilder) UserStateSummaryById(id string)(*ManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary%2Did"] = id
    }
    return NewManagedEBooksItemUserStateSummaryUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
