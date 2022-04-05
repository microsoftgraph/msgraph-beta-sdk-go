package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i019948c5cc3432cae1aa4d885d81eaee957e9c7ddfea6d161a8a6b12d04d693b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/installsummary"
    i0a3036e7599c27e8daf0ac76114f7781865beca45274e73af8c8395fc2a9665d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses"
    i1e3d61799ac5b1670257d353559a6aa4ec7e545fb82dd83595bf36c04e7e9b1d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/assign"
    i2cc000920189f2c9f846e98b23559d1c3218e590b3b9322853afced5128198d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp"
    i3490d416205b3df42d9b8ee0ff657b79e6262e2c3b79708a8f0fb50f5a1c2499 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses"
    i413f580db54415c14dab7adbde8526f1b741cd16690fd87a8ff30b9108f0ec47 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/updaterelationships"
    i9b6c1e48145f74fcaf231180b0b8a979f941d054e34a28d2e470e01c6a7ee864 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/getrelatedappstateswithuserprincipalnamewithdeviceid"
    iaca75b4f217f7322eba62ee5fd8c052d9514c3394b480376576627995983ab4a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/relationships"
    idd22d981d731ad036f829684abde198fe174082361ffe1a63d853407dca69d18 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/categories"
    ifb2db1cec9c276d6c90084ff24743ae3570fed79bdca00e6ab2191624f49ad72 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/assignments"
    i732f121d128a6358acbcb613d4f9cc7f5a9dd32364d5caff115da0580dadf6af "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/categories/item"
    i7578dc338c9dc393bea0c169238874f0e19fedfffc44418ab516b34a4b525138 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item"
    i8ae606a13a79cd675f0f56386d1cfce72468cf3465afc5d50a1e5bfa0e57e603 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/assignments/item"
    ia6a71ed4e86b560dd63fb4ca626455944edeea774053bb5436fcd463ee9a2cc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/relationships/item"
    ifddf3b1a5eae62e5fad4e7d420246006149401e2aacd31254c493f3bd5f95b80 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item"
)

// MobileAppItemRequestBuilder provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
type MobileAppItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MobileAppItemRequestBuilderDeleteOptions options for Delete
type MobileAppItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// MobileAppItemRequestBuilderGetOptions options for Get
type MobileAppItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *MobileAppItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// MobileAppItemRequestBuilderGetQueryParameters the mobile apps.
type MobileAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MobileAppItemRequestBuilderPatchOptions options for Patch
type MobileAppItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *MobileAppItemRequestBuilder) Assign()(*i1e3d61799ac5b1670257d353559a6aa4ec7e545fb82dd83595bf36c04e7e9b1d.AssignRequestBuilder) {
    return i1e3d61799ac5b1670257d353559a6aa4ec7e545fb82dd83595bf36c04e7e9b1d.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *MobileAppItemRequestBuilder) Assignments()(*ifb2db1cec9c276d6c90084ff24743ae3570fed79bdca00e6ab2191624f49ad72.AssignmentsRequestBuilder) {
    return ifb2db1cec9c276d6c90084ff24743ae3570fed79bdca00e6ab2191624f49ad72.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item.assignments.item collection
func (m *MobileAppItemRequestBuilder) AssignmentsById(id string)(*i8ae606a13a79cd675f0f56386d1cfce72468cf3465afc5d50a1e5bfa0e57e603.MobileAppAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppAssignment_id"] = id
    }
    return i8ae606a13a79cd675f0f56386d1cfce72468cf3465afc5d50a1e5bfa0e57e603.NewMobileAppAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Categories the categories property
func (m *MobileAppItemRequestBuilder) Categories()(*idd22d981d731ad036f829684abde198fe174082361ffe1a63d853407dca69d18.CategoriesRequestBuilder) {
    return idd22d981d731ad036f829684abde198fe174082361ffe1a63d853407dca69d18.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item.categories.item collection
func (m *MobileAppItemRequestBuilder) CategoriesById(id string)(*i732f121d128a6358acbcb613d4f9cc7f5a9dd32364d5caff115da0580dadf6af.MobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory_id"] = id
    }
    return i732f121d128a6358acbcb613d4f9cc7f5a9dd32364d5caff115da0580dadf6af.NewMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMobileAppItemRequestBuilderInternal instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppItemRequestBuilder) {
    m := &MobileAppItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMobileAppItemRequestBuilder instantiates a new MobileAppItemRequestBuilder and sets the default values.
func NewMobileAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mobileApps for deviceAppManagement
func (m *MobileAppItemRequestBuilder) CreateDeleteRequestInformation(options *MobileAppItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the mobile apps.
func (m *MobileAppItemRequestBuilder) CreateGetRequestInformation(options *MobileAppItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property mobileApps in deviceAppManagement
func (m *MobileAppItemRequestBuilder) CreatePatchRequestInformation(options *MobileAppItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property mobileApps for deviceAppManagement
func (m *MobileAppItemRequestBuilder) Delete(options *MobileAppItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceStatuses the deviceStatuses property
func (m *MobileAppItemRequestBuilder) DeviceStatuses()(*i3490d416205b3df42d9b8ee0ff657b79e6262e2c3b79708a8f0fb50f5a1c2499.DeviceStatusesRequestBuilder) {
    return i3490d416205b3df42d9b8ee0ff657b79e6262e2c3b79708a8f0fb50f5a1c2499.NewDeviceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item.deviceStatuses.item collection
func (m *MobileAppItemRequestBuilder) DeviceStatusesById(id string)(*i7578dc338c9dc393bea0c169238874f0e19fedfffc44418ab516b34a4b525138.MobileAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppInstallStatus_id"] = id
    }
    return i7578dc338c9dc393bea0c169238874f0e19fedfffc44418ab516b34a4b525138.NewMobileAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the mobile apps.
func (m *MobileAppItemRequestBuilder) Get(options *MobileAppItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppable), nil
}
// GetRelatedAppStatesWithUserPrincipalNameWithDeviceId provides operations to call the getRelatedAppStates method.
func (m *MobileAppItemRequestBuilder) GetRelatedAppStatesWithUserPrincipalNameWithDeviceId(deviceId *string, userPrincipalName *string)(*i9b6c1e48145f74fcaf231180b0b8a979f941d054e34a28d2e470e01c6a7ee864.GetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilder) {
    return i9b6c1e48145f74fcaf231180b0b8a979f941d054e34a28d2e470e01c6a7ee864.NewGetRelatedAppStatesWithUserPrincipalNameWithDeviceIdRequestBuilderInternal(m.pathParameters, m.requestAdapter, deviceId, userPrincipalName);
}
// InstallSummary the installSummary property
func (m *MobileAppItemRequestBuilder) InstallSummary()(*i019948c5cc3432cae1aa4d885d81eaee957e9c7ddfea6d161a8a6b12d04d693b.InstallSummaryRequestBuilder) {
    return i019948c5cc3432cae1aa4d885d81eaee957e9c7ddfea6d161a8a6b12d04d693b.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosVppApp the iosVppApp property
func (m *MobileAppItemRequestBuilder) IosVppApp()(*i2cc000920189f2c9f846e98b23559d1c3218e590b3b9322853afced5128198d6.IosVppAppRequestBuilder) {
    return i2cc000920189f2c9f846e98b23559d1c3218e590b3b9322853afced5128198d6.NewIosVppAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property mobileApps in deviceAppManagement
func (m *MobileAppItemRequestBuilder) Patch(options *MobileAppItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Relationships the relationships property
func (m *MobileAppItemRequestBuilder) Relationships()(*iaca75b4f217f7322eba62ee5fd8c052d9514c3394b480376576627995983ab4a.RelationshipsRequestBuilder) {
    return iaca75b4f217f7322eba62ee5fd8c052d9514c3394b480376576627995983ab4a.NewRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationshipsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item.relationships.item collection
func (m *MobileAppItemRequestBuilder) RelationshipsById(id string)(*ia6a71ed4e86b560dd63fb4ca626455944edeea774053bb5436fcd463ee9a2cc2.MobileAppRelationshipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppRelationship_id"] = id
    }
    return ia6a71ed4e86b560dd63fb4ca626455944edeea774053bb5436fcd463ee9a2cc2.NewMobileAppRelationshipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UpdateRelationships the updateRelationships property
func (m *MobileAppItemRequestBuilder) UpdateRelationships()(*i413f580db54415c14dab7adbde8526f1b741cd16690fd87a8ff30b9108f0ec47.UpdateRelationshipsRequestBuilder) {
    return i413f580db54415c14dab7adbde8526f1b741cd16690fd87a8ff30b9108f0ec47.NewUpdateRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatuses the userStatuses property
func (m *MobileAppItemRequestBuilder) UserStatuses()(*i0a3036e7599c27e8daf0ac76114f7781865beca45274e73af8c8395fc2a9665d.UserStatusesRequestBuilder) {
    return i0a3036e7599c27e8daf0ac76114f7781865beca45274e73af8c8395fc2a9665d.NewUserStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStatusesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.mobileApps.item.userStatuses.item collection
func (m *MobileAppItemRequestBuilder) UserStatusesById(id string)(*ifddf3b1a5eae62e5fad4e7d420246006149401e2aacd31254c493f3bd5f95b80.UserAppInstallStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userAppInstallStatus_id"] = id
    }
    return ifddf3b1a5eae62e5fad4e7d420246006149401e2aacd31254c493f3bd5f95b80.NewUserAppInstallStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
