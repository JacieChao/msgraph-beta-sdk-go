package registration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i30653050b595cb21b15f68a043c33a39a8a59b4e9b628f2f8ce1367de0c76f29 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/registration/customquestions"
    id83f8888f721170d686bae4191da98eefd34bb35b44be13eed9720c503cd8d9a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/registration/registrants"
    i022718a803bf01f48de9971caabac22fb9388c922fc5c47de1d0a1f764c5eb23 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/registration/customquestions/item"
    if37161e0027df4594e8f68714a353b4fbfb428917e9e15ffaa9f02010de1384d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/onlinemeetings/item/registration/registrants/item"
)

// Builds and executes requests for operations under \me\onlineMeetings\{onlineMeeting-id}\registration
type RegistrationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type RegistrationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type RegistrationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RegistrationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
type RegistrationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type RegistrationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingRegistration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new RegistrationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewRegistrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegistrationRequestBuilder) {
    m := &RegistrationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/onlineMeetings/{onlineMeeting_id}/registration{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new RegistrationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewRegistrationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RegistrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRegistrationRequestBuilderInternal(urlParams, requestAdapter)
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) CreateDeleteRequestInformation(options *RegistrationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) CreateGetRequestInformation(options *RegistrationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) CreatePatchRequestInformation(options *RegistrationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *RegistrationRequestBuilder) CustomQuestions()(*i30653050b595cb21b15f68a043c33a39a8a59b4e9b628f2f8ce1367de0c76f29.CustomQuestionsRequestBuilder) {
    return i30653050b595cb21b15f68a043c33a39a8a59b4e9b628f2f8ce1367de0c76f29.NewCustomQuestionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onlineMeetings.item.registration.customQuestions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *RegistrationRequestBuilder) CustomQuestionsById(id string)(*i022718a803bf01f48de9971caabac22fb9388c922fc5c47de1d0a1f764c5eb23.MeetingRegistrationQuestionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingRegistrationQuestion_id"] = id
    }
    return i022718a803bf01f48de9971caabac22fb9388c922fc5c47de1d0a1f764c5eb23.NewMeetingRegistrationQuestionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) Delete(options *RegistrationRequestBuilderDeleteOptions)(error) {
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
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) Get(options *RegistrationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingRegistration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewMeetingRegistration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.MeetingRegistration), nil
}
// The registration that has been enabled for an online meeting. One online meeting can only have one registration enabled.
// Parameters:
//  - options : Options for the request
func (m *RegistrationRequestBuilder) Patch(options *RegistrationRequestBuilderPatchOptions)(error) {
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
func (m *RegistrationRequestBuilder) Registrants()(*id83f8888f721170d686bae4191da98eefd34bb35b44be13eed9720c503cd8d9a.RegistrantsRequestBuilder) {
    return id83f8888f721170d686bae4191da98eefd34bb35b44be13eed9720c503cd8d9a.NewRegistrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.me.onlineMeetings.item.registration.registrants.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *RegistrationRequestBuilder) RegistrantsById(id string)(*if37161e0027df4594e8f68714a353b4fbfb428917e9e15ffaa9f02010de1384d.MeetingRegistrantRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["meetingRegistrant_id"] = id
    }
    return if37161e0027df4594e8f68714a353b4fbfb428917e9e15ffaa9f02010de1384d.NewMeetingRegistrantRequestBuilderInternal(urlTplParams, m.requestAdapter);
}