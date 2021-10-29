package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0cc28725ee9e02cabd2adf77ea9c69ec170f32ec4ee52511e8eddcb2e861df29 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/multivalueextendedproperties"
    i33b54123be3148d600d0cddcbfcbbbb662c7ff52db8a7d37e2dd5a4e8994d064 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/calendarview"
    i77f3a24b6b57772037b003338c2b4e5d8a9fd6b47d8800933b312d72ea34690a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/allowedcalendarsharingroleswithuser"
    i7ce51eb88e0c5dc68784a693fcfda6a4aba4840349f6bee4293e61263d098551 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/events"
    ic2d7b5e3eae7a46aae0380e3a087b31b5000e6ffcfa62ba4d23a8bfaea5999cd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/getschedule"
    iccaa418815ccc4668355da7ed0ea28ccdc048c6c71150b352d3fa08b52b6e0b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/singlevalueextendedproperties"
    icda45a7db0684335193e68f364ef8373bc48fa41fde3f149a4ffec11ebc4c2aa "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/calendarpermissions"
    i26a988227b1128d6b58450eda6f3fed71a462c102f264f67d445446604e37cc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/singlevalueextendedproperties/item"
    i5fdac5e8efc3e7a96f300f80c0a2d1c1afde47805ae01e0d7ebde91aa1954a55 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/multivalueextendedproperties/item"
    ibbad30d0ee1b32e444e020119260f05cf24fdaadcb7fb15e30d1c5aedff2b8bd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/calendarview/item"
    if19567c8855c0e38d0b32211a5b792af17ba9b21afd1df931c47fde5478fa6a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/calendarpermissions/item"
    ifa63a8ed728f8e9e1407b60f546706ea70e32c9780f431ae3e05ae271bbf6fc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/events/item/calendar/events/item"
)

type CalendarRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i77f3a24b6b57772037b003338c2b4e5d8a9fd6b47d8800933b312d72ea34690a.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i77f3a24b6b57772037b003338c2b4e5d8a9fd6b47d8800933b312d72ea34690a.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*icda45a7db0684335193e68f364ef8373bc48fa41fde3f149a4ffec11ebc4c2aa.CalendarPermissionsRequestBuilder) {
    return icda45a7db0684335193e68f364ef8373bc48fa41fde3f149a4ffec11ebc4c2aa.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*if19567c8855c0e38d0b32211a5b792af17ba9b21afd1df931c47fde5478fa6a1.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return if19567c8855c0e38d0b32211a5b792af17ba9b21afd1df931c47fde5478fa6a1.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i33b54123be3148d600d0cddcbfcbbbb662c7ff52db8a7d37e2dd5a4e8994d064.CalendarViewRequestBuilder) {
    return i33b54123be3148d600d0cddcbfcbbbb662c7ff52db8a7d37e2dd5a4e8994d064.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*ibbad30d0ee1b32e444e020119260f05cf24fdaadcb7fb15e30d1c5aedff2b8bd.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ibbad30d0ee1b32e444e020119260f05cf24fdaadcb7fb15e30d1c5aedff2b8bd.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/events/{event_id}/calendar{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CalendarRequestBuilder) CreateGetRequestInformation(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(CalendarRequestBuilderGetQueryParameters)
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
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *CalendarRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CalendarRequestBuilder) Events()(*i7ce51eb88e0c5dc68784a693fcfda6a4aba4840349f6bee4293e61263d098551.EventsRequestBuilder) {
    return i7ce51eb88e0c5dc68784a693fcfda6a4aba4840349f6bee4293e61263d098551.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) EventsById(id string)(*ifa63a8ed728f8e9e1407b60f546706ea70e32c9780f431ae3e05ae271bbf6fc3.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ifa63a8ed728f8e9e1407b60f546706ea70e32c9780f431ae3e05ae271bbf6fc3.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) Get(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCalendar() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*ic2d7b5e3eae7a46aae0380e3a087b31b5000e6ffcfa62ba4d23a8bfaea5999cd.GetScheduleRequestBuilder) {
    return ic2d7b5e3eae7a46aae0380e3a087b31b5000e6ffcfa62ba4d23a8bfaea5999cd.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i0cc28725ee9e02cabd2adf77ea9c69ec170f32ec4ee52511e8eddcb2e861df29.MultiValueExtendedPropertiesRequestBuilder) {
    return i0cc28725ee9e02cabd2adf77ea9c69ec170f32ec4ee52511e8eddcb2e861df29.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i5fdac5e8efc3e7a96f300f80c0a2d1c1afde47805ae01e0d7ebde91aa1954a55.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i5fdac5e8efc3e7a96f300f80c0a2d1c1afde47805ae01e0d7ebde91aa1954a55.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*iccaa418815ccc4668355da7ed0ea28ccdc048c6c71150b352d3fa08b52b6e0b9.SingleValueExtendedPropertiesRequestBuilder) {
    return iccaa418815ccc4668355da7ed0ea28ccdc048c6c71150b352d3fa08b52b6e0b9.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i26a988227b1128d6b58450eda6f3fed71a462c102f264f67d445446604e37cc0.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i26a988227b1128d6b58450eda6f3fed71a462c102f264f67d445446604e37cc0.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}