package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0c7163e7563398de0c6eedbea150cd9e5c1c21ea59a334143f70cad20c9d43f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/versions"
    i0c83698d287f5489aa1424c8bc6596606eecb5f5c8b1e1825897f591c733b2d1 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/fields"
    i143ab0d56c3186f4003976f3a56daf883a7feddb052227d3b27fd75d467b2d68 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/activities"
    i5839c9cf06f41e2efbf50a0df5ddbee26cfdc7b50e280312410f3af115005dac "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/createlink"
    i5a82ed211510c3d25ad68aeffcc05950d7bebeb9574a332007705eb8e0df6703 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/driveitem"
    iae482eb7e4b9a9f425e355ac636ccaf96866c0e3c8dfee0c3384e6c913ff3eab "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ib6f2dbda4f296007a6efe6f6f01af19139c8c33fe39c7916cd403d8e0f687daf "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/analytics"
    i0b3617f4ee916e25039fd57e02d9dd38fe1279565358209883f26f9f3e2ada9f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/versions/item"
    ia5ef275ced42be64d82fae8b372afadf65eefbea2b4d475df418ae4f56314e4b "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities/item/listitem/activities/item"
)

// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\activities\{itemActivityOLD-id}\listItem
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Get listItem from shares
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *ListItemRequestBuilder) Activities()(*i143ab0d56c3186f4003976f3a56daf883a7feddb052227d3b27fd75d467b2d68.ActivitiesRequestBuilder) {
    return i143ab0d56c3186f4003976f3a56daf883a7feddb052227d3b27fd75d467b2d68.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.activities.item.listItem.activities.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*ia5ef275ced42be64d82fae8b372afadf65eefbea2b4d475df418ae4f56314e4b.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id1"] = id
    }
    return ia5ef275ced42be64d82fae8b372afadf65eefbea2b4d475df418ae4f56314e4b.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*ib6f2dbda4f296007a6efe6f6f01af19139c8c33fe39c7916cd403d8e0f687daf.AnalyticsRequestBuilder) {
    return ib6f2dbda4f296007a6efe6f6f01af19139c8c33fe39c7916cd403d8e0f687daf.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}/list/activities/{itemActivityOLD_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property listItem for shares
// Parameters:
//  - options : Options for the request
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
// Get listItem from shares
// Parameters:
//  - options : Options for the request
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
func (m *ListItemRequestBuilder) CreateLink()(*i5839c9cf06f41e2efbf50a0df5ddbee26cfdc7b50e280312410f3af115005dac.CreateLinkRequestBuilder) {
    return i5839c9cf06f41e2efbf50a0df5ddbee26cfdc7b50e280312410f3af115005dac.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update the navigation property listItem in shares
// Parameters:
//  - options : Options for the request
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
// Delete navigation property listItem for shares
// Parameters:
//  - options : Options for the request
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i5a82ed211510c3d25ad68aeffcc05950d7bebeb9574a332007705eb8e0df6703.DriveItemRequestBuilder) {
    return i5a82ed211510c3d25ad68aeffcc05950d7bebeb9574a332007705eb8e0df6703.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i0c83698d287f5489aa1424c8bc6596606eecb5f5c8b1e1825897f591c733b2d1.FieldsRequestBuilder) {
    return i0c83698d287f5489aa1424c8bc6596606eecb5f5c8b1e1825897f591c733b2d1.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get listItem from shares
// Parameters:
//  - options : Options for the request
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewListItem() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem), nil
}
// Builds and executes requests for operations under \shares\{sharedDriveItem-id}\list\activities\{itemActivityOLD-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*iae482eb7e4b9a9f425e355ac636ccaf96866c0e3c8dfee0c3384e6c913ff3eab.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return iae482eb7e4b9a9f425e355ac636ccaf96866c0e3c8dfee0c3384e6c913ff3eab.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Update the navigation property listItem in shares
// Parameters:
//  - options : Options for the request
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*i0c7163e7563398de0c6eedbea150cd9e5c1c21ea59a334143f70cad20c9d43f7.VersionsRequestBuilder) {
    return i0c7163e7563398de0c6eedbea150cd9e5c1c21ea59a334143f70cad20c9d43f7.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.activities.item.listItem.versions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) VersionsById(id string)(*i0b3617f4ee916e25039fd57e02d9dd38fe1279565358209883f26f9f3e2ada9f.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i0b3617f4ee916e25039fd57e02d9dd38fe1279565358209883f26f9f3e2ada9f.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
