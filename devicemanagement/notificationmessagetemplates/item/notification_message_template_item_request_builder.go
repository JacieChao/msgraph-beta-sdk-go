package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i92bdddc6d39f7ceb527fbc3a0c281fdad40f0d03ea69fda531a27f1fb720fdc0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/notificationmessagetemplates/item/sendtestmessage"
    ie131f760e231c41c85ab4a627830ad8e63680d71cfdfeb33b08cb086c57f4e8a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages"
    i8f7ffa70b6e5e6cd3a41114a81ef71a7ef4f81cad3974df901008b8a88e8cf8e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/notificationmessagetemplates/item/localizednotificationmessages/item"
)

// NotificationMessageTemplateItemRequestBuilder builds and executes requests for operations under \deviceManagement\notificationMessageTemplates\{notificationMessageTemplate-id}
type NotificationMessageTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NotificationMessageTemplateItemRequestBuilderDeleteOptions options for Delete
type NotificationMessageTemplateItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotificationMessageTemplateItemRequestBuilderGetOptions options for Get
type NotificationMessageTemplateItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *NotificationMessageTemplateItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NotificationMessageTemplateItemRequestBuilderGetQueryParameters the Notification Message Templates.
type NotificationMessageTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NotificationMessageTemplateItemRequestBuilderPatchOptions options for Patch
type NotificationMessageTemplateItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotificationMessageTemplate;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewNotificationMessageTemplateItemRequestBuilderInternal instantiates a new NotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewNotificationMessageTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotificationMessageTemplateItemRequestBuilder) {
    m := &NotificationMessageTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/notificationMessageTemplates/{notificationMessageTemplate_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewNotificationMessageTemplateItemRequestBuilder instantiates a new NotificationMessageTemplateItemRequestBuilder and sets the default values.
func NewNotificationMessageTemplateItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*NotificationMessageTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewNotificationMessageTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) CreateDeleteRequestInformation(options *NotificationMessageTemplateItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) CreateGetRequestInformation(options *NotificationMessageTemplateItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) CreatePatchRequestInformation(options *NotificationMessageTemplateItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) Delete(options *NotificationMessageTemplateItemRequestBuilderDeleteOptions)(error) {
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
// Get the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) Get(options *NotificationMessageTemplateItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotificationMessageTemplate, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewNotificationMessageTemplate() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NotificationMessageTemplate), nil
}
func (m *NotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessages()(*ie131f760e231c41c85ab4a627830ad8e63680d71cfdfeb33b08cb086c57f4e8a.LocalizedNotificationMessagesRequestBuilder) {
    return ie131f760e231c41c85ab4a627830ad8e63680d71cfdfeb33b08cb086c57f4e8a.NewLocalizedNotificationMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// LocalizedNotificationMessagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.notificationMessageTemplates.item.localizedNotificationMessages.item collection
func (m *NotificationMessageTemplateItemRequestBuilder) LocalizedNotificationMessagesById(id string)(*i8f7ffa70b6e5e6cd3a41114a81ef71a7ef4f81cad3974df901008b8a88e8cf8e.LocalizedNotificationMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["localizedNotificationMessage_id"] = id
    }
    return i8f7ffa70b6e5e6cd3a41114a81ef71a7ef4f81cad3974df901008b8a88e8cf8e.NewLocalizedNotificationMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the Notification Message Templates.
func (m *NotificationMessageTemplateItemRequestBuilder) Patch(options *NotificationMessageTemplateItemRequestBuilderPatchOptions)(error) {
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
func (m *NotificationMessageTemplateItemRequestBuilder) SendTestMessage()(*i92bdddc6d39f7ceb527fbc3a0c281fdad40f0d03ea69fda531a27f1fb720fdc0.SendTestMessageRequestBuilder) {
    return i92bdddc6d39f7ceb527fbc3a0c281fdad40f0d03ea69fda531a27f1fb720fdc0.NewSendTestMessageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}