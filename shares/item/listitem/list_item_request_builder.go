package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0b4ec6fc991be1a5bad593d9f97db9c6abcea4505d6ecaaba08d05caabf2592e "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/fields"
    i38a12227778a54f043ed4c9fc899ccea2f8119513b144d335bbb24ff9feb35b6 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/createlink"
    i3c8545bb4c08289b3ae6c4f2a46d3955a639f2f123c347b15444adbb06feed1a "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/driveitem"
    i51d2330b33ed1489d1581eebff193d7fe83fbb44e60294d8c7c24190a7e28c79 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/versions"
    i722d2b0783dd48ef2c67f1b80b73ae19de29fd8738a39e7f55054268ef0f071e "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/analytics"
    ibd6fe918e90b1b84ad35a09f6c26ed4f4979618638425dd07b61021a5aeef29d "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/activities"
    ie42ad52aa15579e40c5740c2a43c72ab454f4fd52ca4c960af27593651aae334 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ic3476f4a830279b9e89835a349ed86eeeaf5ee2df8edae7f638b271a51d8cace "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/versions/item"
    ie882571704f16fce890ce1823a4378f95130fa7dd10d24c9f8cca3163c9d0344 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/listitem/activities/item"
)

// ListItemRequestBuilder builds and executes requests for operations under \shares\{sharedDriveItem-id}\listItem
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetQueryParameters used to access the underlying listItem
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*ibd6fe918e90b1b84ad35a09f6c26ed4f4979618638425dd07b61021a5aeef29d.ActivitiesRequestBuilder) {
    return ibd6fe918e90b1b84ad35a09f6c26ed4f4979618638425dd07b61021a5aeef29d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*ie882571704f16fce890ce1823a4378f95130fa7dd10d24c9f8cca3163c9d0344.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ie882571704f16fce890ce1823a4378f95130fa7dd10d24c9f8cca3163c9d0344.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i722d2b0783dd48ef2c67f1b80b73ae19de29fd8738a39e7f55054268ef0f071e.AnalyticsRequestBuilder) {
    return i722d2b0783dd48ef2c67f1b80b73ae19de29fd8738a39e7f55054268ef0f071e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation used to access the underlying listItem
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation used to access the underlying listItem
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemRequestBuilder) CreateLink()(*i38a12227778a54f043ed4c9fc899ccea2f8119513b144d335bbb24ff9feb35b6.CreateLinkRequestBuilder) {
    return i38a12227778a54f043ed4c9fc899ccea2f8119513b144d335bbb24ff9feb35b6.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation used to access the underlying listItem
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete used to access the underlying listItem
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
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
func (m *ListItemRequestBuilder) DriveItem()(*i3c8545bb4c08289b3ae6c4f2a46d3955a639f2f123c347b15444adbb06feed1a.DriveItemRequestBuilder) {
    return i3c8545bb4c08289b3ae6c4f2a46d3955a639f2f123c347b15444adbb06feed1a.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i0b4ec6fc991be1a5bad593d9f97db9c6abcea4505d6ecaaba08d05caabf2592e.FieldsRequestBuilder) {
    return i0b4ec6fc991be1a5bad593d9f97db9c6abcea4505d6ecaaba08d05caabf2592e.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get used to access the underlying listItem
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewListItem() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval builds and executes requests for operations under \shares\{sharedDriveItem-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*ie42ad52aa15579e40c5740c2a43c72ab454f4fd52ca4c960af27593651aae334.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ie42ad52aa15579e40c5740c2a43c72ab454f4fd52ca4c960af27593651aae334.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Patch used to access the underlying listItem
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
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
func (m *ListItemRequestBuilder) Versions()(*i51d2330b33ed1489d1581eebff193d7fe83fbb44e60294d8c7c24190a7e28c79.VersionsRequestBuilder) {
    return i51d2330b33ed1489d1581eebff193d7fe83fbb44e60294d8c7c24190a7e28c79.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*ic3476f4a830279b9e89835a349ed86eeeaf5ee2df8edae7f638b271a51d8cace.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ic3476f4a830279b9e89835a349ed86eeeaf5ee2df8edae7f638b271a51d8cace.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
