package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2ae23c159a0e1c54b40f5c4dbc44d533dd75c043d914b63af64ce3d7eb8b08d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/createlink"
    i4321d96761714ba2cc83162cb402b89d4e805e4ecc124d7114441285013dd2d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/activities"
    i8a24f739f347ff43fd36489569e5d2c2207209505cd977be1f9b191f53a20fc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/driveitem"
    i91e074b2b0610bea13c640e2ce112db82e3b4d249aeab535290f280af12e6cb7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/fields"
    iaf529e3e5d745c235f4a8b4ac3af66abf03adeea5b24b734f6341d06495615e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ib303fb970842c7034c8b260984089f9522c1c085877f9d672a778a852ef10990 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/analytics"
    ic8ee5373cd68c2593ffcfc7beb4b07fd6c2892ca6790dc38e1b9bdda6e971af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/versions"
    i14a568448b0792ffb00e14536ea34ad11b3fe7f80a866bf8dbd7c21f82634c34 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/versions/item"
    icec7566431140abfa599d8de5fdcd89d71e43c77e2be18256afcd8c5d4c658f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/activities/item/listitem/activities/item"
)

// Builds and executes requests for operations under \drive\list\activities\{itemActivityOLD-id}\listItem
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
// Get listItem from drive
type ListItemRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
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
func (m *ListItemRequestBuilder) Activities()(*i4321d96761714ba2cc83162cb402b89d4e805e4ecc124d7114441285013dd2d7.ActivitiesRequestBuilder) {
    return i4321d96761714ba2cc83162cb402b89d4e805e4ecc124d7114441285013dd2d7.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.activities.item.listItem.activities.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*icec7566431140abfa599d8de5fdcd89d71e43c77e2be18256afcd8c5d4c658f5.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id1"] = id
    }
    return icec7566431140abfa599d8de5fdcd89d71e43c77e2be18256afcd8c5d4c658f5.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*ib303fb970842c7034c8b260984089f9522c1c085877f9d672a778a852ef10990.AnalyticsRequestBuilder) {
    return ib303fb970842c7034c8b260984089f9522c1c085877f9d672a778a852ef10990.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/activities/{itemActivityOLD_id}/listItem{?select,expand}";
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
// Delete navigation property listItem for drive
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
// Get listItem from drive
// Parameters:
//  - options : Options for the request
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
func (m *ListItemRequestBuilder) CreateLink()(*i2ae23c159a0e1c54b40f5c4dbc44d533dd75c043d914b63af64ce3d7eb8b08d7.CreateLinkRequestBuilder) {
    return i2ae23c159a0e1c54b40f5c4dbc44d533dd75c043d914b63af64ce3d7eb8b08d7.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update the navigation property listItem in drive
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
// Delete navigation property listItem for drive
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
func (m *ListItemRequestBuilder) DriveItem()(*i8a24f739f347ff43fd36489569e5d2c2207209505cd977be1f9b191f53a20fc1.DriveItemRequestBuilder) {
    return i8a24f739f347ff43fd36489569e5d2c2207209505cd977be1f9b191f53a20fc1.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i91e074b2b0610bea13c640e2ce112db82e3b4d249aeab535290f280af12e6cb7.FieldsRequestBuilder) {
    return i91e074b2b0610bea13c640e2ce112db82e3b4d249aeab535290f280af12e6cb7.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get listItem from drive
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
// Builds and executes requests for operations under \drive\list\activities\{itemActivityOLD-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*iaf529e3e5d745c235f4a8b4ac3af66abf03adeea5b24b734f6341d06495615e9.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return iaf529e3e5d745c235f4a8b4ac3af66abf03adeea5b24b734f6341d06495615e9.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Update the navigation property listItem in drive
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
func (m *ListItemRequestBuilder) Versions()(*ic8ee5373cd68c2593ffcfc7beb4b07fd6c2892ca6790dc38e1b9bdda6e971af2.VersionsRequestBuilder) {
    return ic8ee5373cd68c2593ffcfc7beb4b07fd6c2892ca6790dc38e1b9bdda6e971af2.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.activities.item.listItem.versions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) VersionsById(id string)(*i14a568448b0792ffb00e14536ea34ad11b3fe7f80a866bf8dbd7c21f82634c34.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i14a568448b0792ffb00e14536ea34ad11b3fe7f80a866bf8dbd7c21f82634c34.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
