package drive

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i192e552808d72c666fa89ea5e18d213d8060951b798ab506f46a9f7ffdf83e2a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items"
    i1cde53b5097fe4f2cb7d9aa6a1d6263ca7dc45f359a394f6451d31a82baa66f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/searchwithq"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2f3fa4e2e4547c94b77043edcae61480a38269272b0e62403050580a8364f5c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/sharedwithme"
    i3dba79c08f5528ca1f738cd0ec4d5c0b89913bea342b964174e72bc175dec1d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/recent"
    i4e0de778e26a386528f38fa1e34ccca7e9004fb8a097e4795502e4f91e7913ba "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following"
    i8195f82296fd049d853cdd9b0d96130049f0e63bf935143dc8aed2b4066352f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special"
    i89fcd6e566145dd6fd7ea2956e5aefd9030e1a74e3028ab21e1cac08c5d17ead "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/activities"
    ie75e0ac9d6da32720a3dd716caabe531bea74aa98e60d982dabe2088718e1e8d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root"
    iee45bfc2e2e1ccdf052d488c498a03424a7325bcbac25a672e0946a3251ff078 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list"
    iee7675fea869eb6434c7377c309bb1287126402b996d180fde09843483ab5b0d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles"
    i50d8648f64c4421d77d91aa220edfff9d5ba824309e53cbcc3025b0592285332 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item"
    i5d9fbdca07c6e20b1cc99de5efe021bbc1fb6e6b9aec7b3e289fa7d91f62974d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item"
    i75b317198811cf16709f6ce43e8700a26a56733bc41233a94a23920a7de396f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/activities/item"
    i7ba1848b91a2a586b2d0cf4042a70732acf64fd1f43dec2814281ec5e078edc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item"
    id087d332b06d6a2e3560c5cbdd124c635b06366297ad2046147358062b8678af "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item"
)

// DriveRequestBuilder provides operations to manage the drive singleton.
type DriveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DriveRequestBuilderGetQueryParameters get drive
type DriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DriveRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DriveRequestBuilderGetQueryParameters
}
// DriveRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DriveRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *DriveRequestBuilder) Activities()(*i89fcd6e566145dd6fd7ea2956e5aefd9030e1a74e3028ab21e1cac08c5d17ead.ActivitiesRequestBuilder) {
    return i89fcd6e566145dd6fd7ea2956e5aefd9030e1a74e3028ab21e1cac08c5d17ead.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.activities.item collection
func (m *DriveRequestBuilder) ActivitiesById(id string)(*i75b317198811cf16709f6ce43e8700a26a56733bc41233a94a23920a7de396f5.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i75b317198811cf16709f6ce43e8700a26a56733bc41233a94a23920a7de396f5.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Bundles the bundles property
func (m *DriveRequestBuilder) Bundles()(*iee7675fea869eb6434c7377c309bb1287126402b996d180fde09843483ab5b0d.BundlesRequestBuilder) {
    return iee7675fea869eb6434c7377c309bb1287126402b996d180fde09843483ab5b0d.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.bundles.item collection
func (m *DriveRequestBuilder) BundlesById(id string)(*id087d332b06d6a2e3560c5cbdd124c635b06366297ad2046147358062b8678af.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return id087d332b06d6a2e3560c5cbdd124c635b06366297ad2046147358062b8678af.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveRequestBuilderInternal instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveRequestBuilder instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get drive
func (m *DriveRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get drive
func (m *DriveRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DriveRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update drive
func (m *DriveRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update drive
func (m *DriveRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Following the following property
func (m *DriveRequestBuilder) Following()(*i4e0de778e26a386528f38fa1e34ccca7e9004fb8a097e4795502e4f91e7913ba.FollowingRequestBuilder) {
    return i4e0de778e26a386528f38fa1e34ccca7e9004fb8a097e4795502e4f91e7913ba.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.following.item collection
func (m *DriveRequestBuilder) FollowingById(id string)(*i50d8648f64c4421d77d91aa220edfff9d5ba824309e53cbcc3025b0592285332.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i50d8648f64c4421d77d91aa220edfff9d5ba824309e53cbcc3025b0592285332.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get drive
func (m *DriveRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get drive
func (m *DriveRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DriveRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable), nil
}
// Items the items property
func (m *DriveRequestBuilder) Items()(*i192e552808d72c666fa89ea5e18d213d8060951b798ab506f46a9f7ffdf83e2a.ItemsRequestBuilder) {
    return i192e552808d72c666fa89ea5e18d213d8060951b798ab506f46a9f7ffdf83e2a.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.items.item collection
func (m *DriveRequestBuilder) ItemsById(id string)(*i7ba1848b91a2a586b2d0cf4042a70732acf64fd1f43dec2814281ec5e078edc1.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i7ba1848b91a2a586b2d0cf4042a70732acf64fd1f43dec2814281ec5e078edc1.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// List the list property
func (m *DriveRequestBuilder) List()(*iee45bfc2e2e1ccdf052d488c498a03424a7325bcbac25a672e0946a3251ff078.ListRequestBuilder) {
    return iee45bfc2e2e1ccdf052d488c498a03424a7325bcbac25a672e0946a3251ff078.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update drive
func (m *DriveRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update drive
func (m *DriveRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Driveable, requestConfiguration *DriveRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// Recent provides operations to call the recent method.
func (m *DriveRequestBuilder) Recent()(*i3dba79c08f5528ca1f738cd0ec4d5c0b89913bea342b964174e72bc175dec1d4.RecentRequestBuilder) {
    return i3dba79c08f5528ca1f738cd0ec4d5c0b89913bea342b964174e72bc175dec1d4.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Root the root property
func (m *DriveRequestBuilder) Root()(*ie75e0ac9d6da32720a3dd716caabe531bea74aa98e60d982dabe2088718e1e8d.RootRequestBuilder) {
    return ie75e0ac9d6da32720a3dd716caabe531bea74aa98e60d982dabe2088718e1e8d.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*i1cde53b5097fe4f2cb7d9aa6a1d6263ca7dc45f359a394f6451d31a82baa66f1.SearchWithQRequestBuilder) {
    return i1cde53b5097fe4f2cb7d9aa6a1d6263ca7dc45f359a394f6451d31a82baa66f1.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveRequestBuilder) SharedWithMe()(*i2f3fa4e2e4547c94b77043edcae61480a38269272b0e62403050580a8364f5c9.SharedWithMeRequestBuilder) {
    return i2f3fa4e2e4547c94b77043edcae61480a38269272b0e62403050580a8364f5c9.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Special the special property
func (m *DriveRequestBuilder) Special()(*i8195f82296fd049d853cdd9b0d96130049f0e63bf935143dc8aed2b4066352f6.SpecialRequestBuilder) {
    return i8195f82296fd049d853cdd9b0d96130049f0e63bf935143dc8aed2b4066352f6.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.special.item collection
func (m *DriveRequestBuilder) SpecialById(id string)(*i5d9fbdca07c6e20b1cc99de5efe021bbc1fb6e6b9aec7b3e289fa7d91f62974d.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return i5d9fbdca07c6e20b1cc99de5efe021bbc1fb6e6b9aec7b3e289fa7d91f62974d.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
