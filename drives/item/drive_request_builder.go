package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/sharedwithme"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities"
    i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root"
    i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following"
    i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/recent"
    i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special"
    ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items"
    ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/searchwithq"
    id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list"
    ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles"
    i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/following/item"
    i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/special/item"
    i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/items/item"
    ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item"
    ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/bundles/item"
)

type DriveRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DriveRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *DriveRequestBuilder) Activities()(*i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f.ActivitiesRequestBuilder) {
    return i57502ad1c70adde13f4f4191c9621344f90ed303c506bcfac62cc6e0f211be2f.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) ActivitiesById(id string)(*ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ib301dcbbaf0fbf1ec225026dfd49aed454924f5c763fba68c72d04e8b817701e.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) Bundles()(*ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718.BundlesRequestBuilder) {
    return ie5ad7fd3da5b255316aa44176c2ac2f76115483b551789e79d599fe1a09af718.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) BundlesById(id string)(*ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ifcd10669d080989622673b9afdc2d4fc13b76538e8a7179f0887b9fb078584da.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/drives/{drive_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDriveRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveRequestBuilder) CreateGetRequestInformation(q func (value *DriveRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DriveRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveRequestBuilder) Following()(*i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111.FollowingRequestBuilder) {
    return i766cace85617fc30182532069c7098a07d81041ac72cd410a073786ca7d45111.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) FollowingById(id string)(*i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i2ad598655381501616b067740dbd6e99ee1ac52e664f55c88fbd4f908cb5ad17.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) Get(q func (value *DriveRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDrive() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive), nil
}
func (m *DriveRequestBuilder) Items()(*ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c.ItemsRequestBuilder) {
    return ia54161ca01b96b4ea12e25e22325b7cb0b3cdd9294192cfbb560f72ea0db851c.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) ItemsById(id string)(*i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i8f091105e59b53a3e7170d648bf47128b9c71e68ca777e009b742a1fec950621.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) List()(*id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af.ListRequestBuilder) {
    return id58c10ecda198f10fa52e1b05c30148577f6b2f028c79657172b0133231e02af.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveRequestBuilder) Recent()(*i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe.RecentRequestBuilder) {
    return i7da819773ad4a87ff5f9f3d30f0fb8b5d88885b8321f8d2383deef979e936cfe.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Root()(*i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512.RootRequestBuilder) {
    return i6215746464d214fce2dc1b69b48ffe1487b12ff3c49cd51dca35bca5dbb61512.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778.SearchWithQRequestBuilder) {
    return ib66a732acc47ab5473e84f72cd54f49f7de1f8a9430fe5f4ba9571cb77231778.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
func (m *DriveRequestBuilder) SharedWithMe()(*i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680.SharedWithMeRequestBuilder) {
    return i1edeb1c46833b3c99f53e69de7b4d55945a2b859ad55580fca564f0c77f14680.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Special()(*i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80.SpecialRequestBuilder) {
    return i80ad3ea091c56b67759525855e8ed8174db7bd7b55988fb87998e2088ca25d80.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) SpecialById(id string)(*i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i30b8e702569d51cb08732e656d751cb44abd73ef6cd97f739a42e0e3b1c6e4c4.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
