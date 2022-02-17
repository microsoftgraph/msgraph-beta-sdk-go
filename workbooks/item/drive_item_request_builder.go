package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0146b9e31c3cc4f2ab496bd8b7648bd7422072075b9c3fdb0b6223600352be8a "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/children"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i07fa0b9bcf1ebf1422ecf81f650c09280a537299bd61ea4a889e1779aed5cdb1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/createlink"
    i0d7540bf7da4f5a327b935abc70409898a3fd6c1c9c792496574df321f76f182 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/createuploadsession"
    i12f8dcf6e8ea636350dee797d39120b2eff5eb823b749db83370bd17cde4ee26 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/deltawithtoken"
    i1559af655b2f07d6991e747733c0a5dbe16a79c9110b9cd4e95e89d1f533b436 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/analytics"
    i1dda96dcc6ca1d6bba90ca93c12a40d7bd6714ce98037fc360f32671277baa94 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i1ffc874dc11bc59b8f43e08f999b16d5537ed76c0d32ceacaf84e9d01a9a0a54 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/subscriptions"
    i286a92b61d7508a9e3bd4e1607d4dcca000c4c41be5af19c587bbc43e6d8f2e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/content"
    i3434cc42e81799ee683ff05c68279a9af935557da306a30871d4244ec69a586d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/copy"
    i35894371297ae928bdc491107d0da925520373cf663e0c084b04a7fe5c9e9a44 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook"
    i3a42b727ca0a85dd5841ea0caf5e4ddcc0b92314677b92c30cdff812f9d18de6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/checkin"
    i4403a8b36ffdd31c6449e834ff7354472d5bcaf819a2c59390ccca08a944e4e0 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/thumbnails"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i62c92f7e8ef283c94ff015c3fe5d9d7f09ae77403203082ac6544c4f367d9676 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/follow"
    i652ad18530a0d67db6af9e17a408a0fe46a970bfe04431202d422c3171c97065 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/restore"
    i68453248b6adc2e0a123184b9f0397803b74547d93a538bca62fb45840ec7c39 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities"
    i69d7ab1245b51d446b0639d6cff194b64c1ff6ee87304184b6c17dd90f5c093c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/listitem"
    i710d1b03133218cf42db16a0f513411325c8307a829644968e565a02e56ae7f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/invite"
    i7ecb3bb901e726c93053e2981950ac4cad1568eec1a8ca7051a9125074a2f106 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/searchwithq"
    i8236d13341080e5a3a7f2865e569b3a9db26acaf5e7254cd2dfb7594f29e3b31 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/validatepermission"
    i9e601979c83767d12347ff319a26aa16bd8cde9f1cf30b7a89c1c15121909250 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/preview"
    ia10d6bc417675919b129deccbb4a8ff6d0e423f7280ed162a4c41f59cc0e8de4 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/permissions"
    ib109f1365e8ac7f6b077a01adab1935b4dd6adc390c9e072a8b9a8ef2653e116 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/checkout"
    ibcde7b93c66be3ae4c6307030c9c1fd4912a9617a23339afbb10615f9a735dd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/unfollow"
    iccf725507e0de7b68f548524e89423d3c820bc9d1292124875957bfd3d78a358 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/delta"
    iea15fd2e4983cd4bc7f77c56f9e7d10b714e2d4c76c191417ae77f325254d1fd "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/versions"
    i23745af441bc7e806a6ea374d93773d7f5b4d51f0b23dbfc2e289f9ed360c56f "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/versions/item"
    i8af00cb19bae1d8d547aa285fc2e34cbe5c5c5096ff8de885009b1dcd3611d71 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/permissions/item"
    ic2204a9b8bb31b272b46aa5d8e8e1802c6b262a7377a2425b7aeb9b634548a33 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/subscriptions/item"
    if9d9f1bd841c39df4ae58bdaf1c987671d7244b7fbc793a0e55c2c3bcb60d75b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/thumbnails/item"
)

// DriveItemRequestBuilder builds and executes requests for operations under \workbooks\{driveItem-id}
type DriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemRequestBuilderDeleteOptions options for Delete
type DriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetOptions options for Get
type DriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetQueryParameters get entity from workbooks by key
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemRequestBuilderPatchOptions options for Patch
type DriveItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemRequestBuilder) Activities()(*i68453248b6adc2e0a123184b9f0397803b74547d93a538bca62fb45840ec7c39.ActivitiesRequestBuilder) {
    return i68453248b6adc2e0a123184b9f0397803b74547d93a538bca62fb45840ec7c39.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Analytics()(*i1559af655b2f07d6991e747733c0a5dbe16a79c9110b9cd4e95e89d1f533b436.AnalyticsRequestBuilder) {
    return i1559af655b2f07d6991e747733c0a5dbe16a79c9110b9cd4e95e89d1f533b436.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Checkin()(*i3a42b727ca0a85dd5841ea0caf5e4ddcc0b92314677b92c30cdff812f9d18de6.CheckinRequestBuilder) {
    return i3a42b727ca0a85dd5841ea0caf5e4ddcc0b92314677b92c30cdff812f9d18de6.NewCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Checkout()(*ib109f1365e8ac7f6b077a01adab1935b4dd6adc390c9e072a8b9a8ef2653e116.CheckoutRequestBuilder) {
    return ib109f1365e8ac7f6b077a01adab1935b4dd6adc390c9e072a8b9a8ef2653e116.NewCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Children()(*i0146b9e31c3cc4f2ab496bd8b7648bd7422072075b9c3fdb0b6223600352be8a.ChildrenRequestBuilder) {
    return i0146b9e31c3cc4f2ab496bd8b7648bd7422072075b9c3fdb0b6223600352be8a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveItemRequestBuilder) Content()(*i286a92b61d7508a9e3bd4e1607d4dcca000c4c41be5af19c587bbc43e6d8f2e5.ContentRequestBuilder) {
    return i286a92b61d7508a9e3bd4e1607d4dcca000c4c41be5af19c587bbc43e6d8f2e5.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Copy()(*i3434cc42e81799ee683ff05c68279a9af935557da306a30871d4244ec69a586d.CopyRequestBuilder) {
    return i3434cc42e81799ee683ff05c68279a9af935557da306a30871d4244ec69a586d.NewCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete entity from workbooks
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from workbooks by key
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(options *DriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveItemRequestBuilder) CreateLink()(*i07fa0b9bcf1ebf1422ecf81f650c09280a537299bd61ea4a889e1779aed5cdb1.CreateLinkRequestBuilder) {
    return i07fa0b9bcf1ebf1422ecf81f650c09280a537299bd61ea4a889e1779aed5cdb1.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update entity in workbooks
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveItemRequestBuilder) CreateUploadSession()(*i0d7540bf7da4f5a327b935abc70409898a3fd6c1c9c792496574df321f76f182.CreateUploadSessionRequestBuilder) {
    return i0d7540bf7da4f5a327b935abc70409898a3fd6c1c9c792496574df321f76f182.NewCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete entity from workbooks
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Delta builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.delta()
func (m *DriveItemRequestBuilder) Delta()(*iccf725507e0de7b68f548524e89423d3c820bc9d1292124875957bfd3d78a358.DeltaRequestBuilder) {
    return iccf725507e0de7b68f548524e89423d3c820bc9d1292124875957bfd3d78a358.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.delta(token='{token}')
func (m *DriveItemRequestBuilder) DeltaWithToken(token *string)(*i12f8dcf6e8ea636350dee797d39120b2eff5eb823b749db83370bd17cde4ee26.DeltaWithTokenRequestBuilder) {
    return i12f8dcf6e8ea636350dee797d39120b2eff5eb823b749db83370bd17cde4ee26.NewDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
func (m *DriveItemRequestBuilder) Follow()(*i62c92f7e8ef283c94ff015c3fe5d9d7f09ae77403203082ac6544c4f367d9676.FollowRequestBuilder) {
    return i62c92f7e8ef283c94ff015c3fe5d9d7f09ae77403203082ac6544c4f367d9676.NewFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from workbooks by key
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDriveItem() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItem), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
func (m *DriveItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i1dda96dcc6ca1d6bba90ca93c12a40d7bd6714ce98037fc360f32671277baa94.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i1dda96dcc6ca1d6bba90ca93c12a40d7bd6714ce98037fc360f32671277baa94.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
func (m *DriveItemRequestBuilder) Invite()(*i710d1b03133218cf42db16a0f513411325c8307a829644968e565a02e56ae7f6.InviteRequestBuilder) {
    return i710d1b03133218cf42db16a0f513411325c8307a829644968e565a02e56ae7f6.NewInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) ListItem()(*i69d7ab1245b51d446b0639d6cff194b64c1ff6ee87304184b6c17dd90f5c093c.ListItemRequestBuilder) {
    return i69d7ab1245b51d446b0639d6cff194b64c1ff6ee87304184b6c17dd90f5c093c.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in workbooks
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveItemRequestBuilder) Permissions()(*ia10d6bc417675919b129deccbb4a8ff6d0e423f7280ed162a4c41f59cc0e8de4.PermissionsRequestBuilder) {
    return ia10d6bc417675919b129deccbb4a8ff6d0e423f7280ed162a4c41f59cc0e8de4.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.permissions.item collection
func (m *DriveItemRequestBuilder) PermissionsById(id string)(*i8af00cb19bae1d8d547aa285fc2e34cbe5c5c5096ff8de885009b1dcd3611d71.PermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i8af00cb19bae1d8d547aa285fc2e34cbe5c5c5096ff8de885009b1dcd3611d71.NewPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Preview()(*i9e601979c83767d12347ff319a26aa16bd8cde9f1cf30b7a89c1c15121909250.PreviewRequestBuilder) {
    return i9e601979c83767d12347ff319a26aa16bd8cde9f1cf30b7a89c1c15121909250.NewPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Restore()(*i652ad18530a0d67db6af9e17a408a0fe46a970bfe04431202d422c3171c97065.RestoreRequestBuilder) {
    return i652ad18530a0d67db6af9e17a408a0fe46a970bfe04431202d422c3171c97065.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ builds and executes requests for operations under \workbooks\{driveItem-id}\microsoft.graph.search(q='{q}')
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*i7ecb3bb901e726c93053e2981950ac4cad1568eec1a8ca7051a9125074a2f106.SearchWithQRequestBuilder) {
    return i7ecb3bb901e726c93053e2981950ac4cad1568eec1a8ca7051a9125074a2f106.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
func (m *DriveItemRequestBuilder) Subscriptions()(*i1ffc874dc11bc59b8f43e08f999b16d5537ed76c0d32ceacaf84e9d01a9a0a54.SubscriptionsRequestBuilder) {
    return i1ffc874dc11bc59b8f43e08f999b16d5537ed76c0d32ceacaf84e9d01a9a0a54.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.subscriptions.item collection
func (m *DriveItemRequestBuilder) SubscriptionsById(id string)(*ic2204a9b8bb31b272b46aa5d8e8e1802c6b262a7377a2425b7aeb9b634548a33.SubscriptionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return ic2204a9b8bb31b272b46aa5d8e8e1802c6b262a7377a2425b7aeb9b634548a33.NewSubscriptionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Thumbnails()(*i4403a8b36ffdd31c6449e834ff7354472d5bcaf819a2c59390ccca08a944e4e0.ThumbnailsRequestBuilder) {
    return i4403a8b36ffdd31c6449e834ff7354472d5bcaf819a2c59390ccca08a944e4e0.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.thumbnails.item collection
func (m *DriveItemRequestBuilder) ThumbnailsById(id string)(*if9d9f1bd841c39df4ae58bdaf1c987671d7244b7fbc793a0e55c2c3bcb60d75b.ThumbnailSetRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return if9d9f1bd841c39df4ae58bdaf1c987671d7244b7fbc793a0e55c2c3bcb60d75b.NewThumbnailSetRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Unfollow()(*ibcde7b93c66be3ae4c6307030c9c1fd4912a9617a23339afbb10615f9a735dd3.UnfollowRequestBuilder) {
    return ibcde7b93c66be3ae4c6307030c9c1fd4912a9617a23339afbb10615f9a735dd3.NewUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) ValidatePermission()(*i8236d13341080e5a3a7f2865e569b3a9db26acaf5e7254cd2dfb7594f29e3b31.ValidatePermissionRequestBuilder) {
    return i8236d13341080e5a3a7f2865e569b3a9db26acaf5e7254cd2dfb7594f29e3b31.NewValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Versions()(*iea15fd2e4983cd4bc7f77c56f9e7d10b714e2d4c76c191417ae77f325254d1fd.VersionsRequestBuilder) {
    return iea15fd2e4983cd4bc7f77c56f9e7d10b714e2d4c76c191417ae77f325254d1fd.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.versions.item collection
func (m *DriveItemRequestBuilder) VersionsById(id string)(*i23745af441bc7e806a6ea374d93773d7f5b4d51f0b23dbfc2e289f9ed360c56f.DriveItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i23745af441bc7e806a6ea374d93773d7f5b4d51f0b23dbfc2e289f9ed360c56f.NewDriveItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Workbook()(*i35894371297ae928bdc491107d0da925520373cf663e0c084b04a7fe5c9e9a44.WorkbookRequestBuilder) {
    return i35894371297ae928bdc491107d0da925520373cf663e0c084b04a7fe5c9e9a44.NewWorkbookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
