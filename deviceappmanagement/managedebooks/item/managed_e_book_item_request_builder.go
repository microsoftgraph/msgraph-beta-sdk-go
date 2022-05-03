package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i74a0325b0065a251b01ec5c782cded3d19e332e6d91b4d38e8a2778d6d46efc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/assign"
    i8796b5653a9f0dac43ca660f182d458d71ab4290f3f1a84fbb80d020d2f00738 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/categories"
    ia371fb97d4603b6022b57d8702547b0b74f86539b1fc17374e88f8232f4a6ff6 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/devicestates"
    iae1ce53ec1fe3d26e228e31ca74e284059c70d35ab9331dd10ab28b6e3e5f006 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/assignments"
    iba6a5bdbe4569e6c8f560b2c6f52bc985906914c6b733d7db6a101b9ab423314 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary"
    ifa3abdbb40057ad8079ae0439f61eb6074f9d3512d729656d322eacb98aa3403 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/installsummary"
    i2d88a7f499b3854e7de1e7af5887f042e5cdba6db3b6afe628aa868205b0ae3e "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/categories/item"
    i47d65cb9c49a38dcee79025d300de84f32574e9a653d54c08c771b81f324a256 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/devicestates/item"
    i93a3f61d471b69c96df478aa01820a8a69b1252d396baad2e4d4e1b17f5f21fe "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary/item"
    i9ca2634cfe73b9f5258e0a85ae4be6b9174af8d4172b8471d7cf5054a507dc5b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/managedebooks/item/assignments/item"
)

// ManagedEBookItemRequestBuilder provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
type ManagedEBookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedEBookItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBookItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedEBookItemRequestBuilderGetQueryParameters the Managed eBook.
type ManagedEBookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedEBookItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBookItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedEBookItemRequestBuilderGetQueryParameters
}
// ManagedEBookItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedEBookItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign the assign property
func (m *ManagedEBookItemRequestBuilder) Assign()(*i74a0325b0065a251b01ec5c782cded3d19e332e6d91b4d38e8a2778d6d46efc1.AssignRequestBuilder) {
    return i74a0325b0065a251b01ec5c782cded3d19e332e6d91b4d38e8a2778d6d46efc1.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *ManagedEBookItemRequestBuilder) Assignments()(*iae1ce53ec1fe3d26e228e31ca74e284059c70d35ab9331dd10ab28b6e3e5f006.AssignmentsRequestBuilder) {
    return iae1ce53ec1fe3d26e228e31ca74e284059c70d35ab9331dd10ab28b6e3e5f006.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedEBooks.item.assignments.item collection
func (m *ManagedEBookItemRequestBuilder) AssignmentsById(id string)(*i9ca2634cfe73b9f5258e0a85ae4be6b9174af8d4172b8471d7cf5054a507dc5b.ManagedEBookAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment%2Did"] = id
    }
    return i9ca2634cfe73b9f5258e0a85ae4be6b9174af8d4172b8471d7cf5054a507dc5b.NewManagedEBookAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories the categories property
func (m *ManagedEBookItemRequestBuilder) Categories()(*i8796b5653a9f0dac43ca660f182d458d71ab4290f3f1a84fbb80d020d2f00738.CategoriesRequestBuilder) {
    return i8796b5653a9f0dac43ca660f182d458d71ab4290f3f1a84fbb80d020d2f00738.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedEBooks.item.categories.item collection
func (m *ManagedEBookItemRequestBuilder) CategoriesById(id string)(*i2d88a7f499b3854e7de1e7af5887f042e5cdba6db3b6afe628aa868205b0ae3e.ManagedEBookCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookCategory%2Did"] = id
    }
    return i2d88a7f499b3854e7de1e7af5887f042e5cdba6db3b6afe628aa868205b0ae3e.NewManagedEBookCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedEBookItemRequestBuilderInternal instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBookItemRequestBuilder) {
    m := &ManagedEBookItemRequestBuilder{
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
// NewManagedEBookItemRequestBuilder instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBookItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedEBookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBookItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ManagedEBookItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedEBookItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *ManagedEBookItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ManagedEBookItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property managedEBooks for deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *ManagedEBookItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStates the deviceStates property
func (m *ManagedEBookItemRequestBuilder) DeviceStates()(*ia371fb97d4603b6022b57d8702547b0b74f86539b1fc17374e88f8232f4a6ff6.DeviceStatesRequestBuilder) {
    return ia371fb97d4603b6022b57d8702547b0b74f86539b1fc17374e88f8232f4a6ff6.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedEBooks.item.deviceStates.item collection
func (m *ManagedEBookItemRequestBuilder) DeviceStatesById(id string)(*i47d65cb9c49a38dcee79025d300de84f32574e9a653d54c08c771b81f324a256.DeviceInstallStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState%2Did"] = id
    }
    return i47d65cb9c49a38dcee79025d300de84f32574e9a653d54c08c771b81f324a256.NewDeviceInstallStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// GetWithResponseHandler the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ManagedEBookItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) GetWithResponseHandler(requestConfiguration *ManagedEBookItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedEBookFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable), nil
}
// InstallSummary the installSummary property
func (m *ManagedEBookItemRequestBuilder) InstallSummary()(*ifa3abdbb40057ad8079ae0439f61eb6074f9d3512d729656d322eacb98aa3403.InstallSummaryRequestBuilder) {
    return ifa3abdbb40057ad8079ae0439f61eb6074f9d3512d729656d322eacb98aa3403.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PatchWithResponseHandler update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *ManagedEBookItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property managedEBooks in deviceAppManagement
func (m *ManagedEBookItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedEBookable, requestConfiguration *ManagedEBookItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UserStateSummary the userStateSummary property
func (m *ManagedEBookItemRequestBuilder) UserStateSummary()(*iba6a5bdbe4569e6c8f560b2c6f52bc985906914c6b733d7db6a101b9ab423314.UserStateSummaryRequestBuilder) {
    return iba6a5bdbe4569e6c8f560b2c6f52bc985906914c6b733d7db6a101b9ab423314.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStateSummaryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.managedEBooks.item.userStateSummary.item collection
func (m *ManagedEBookItemRequestBuilder) UserStateSummaryById(id string)(*i93a3f61d471b69c96df478aa01820a8a69b1252d396baad2e4d4e1b17f5f21fe.UserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary%2Did"] = id
    }
    return i93a3f61d471b69c96df478aa01820a8a69b1252d396baad2e4d4e1b17f5f21fe.NewUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
