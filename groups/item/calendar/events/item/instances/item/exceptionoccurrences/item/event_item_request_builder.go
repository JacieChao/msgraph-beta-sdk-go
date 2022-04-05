package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i15f92df896b379c185f794e9ffab67e334dedbec9e88d7b86bd84d68d577ae5f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i286f1c31f8fd38525ee7d192dffca61b2a6863707a278d98bdd7935381861683 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/forward"
    i35573439ed685eddef23590bc0a1c9fb4f2ea0675471fd4af0af13251f5a7c46 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/accept"
    i36bd95c8b35fc004367c0ffab7c4b0a838fb9f35ac8297d9992fa9c7e6915b1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/cancel"
    i541aa50f9ab36177d08a293614f2097ab3bbac67379aa06d892a8c89f9b9d945 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/decline"
    i5c0221b58587f947d01c8d30e7024054af1cad4538de5d1c62ffb60e9f974153 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    i5d6da724adab0c2686d22abc37fe4b97f4fbfc7f69783572b96133b75798ef15 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions"
    i72787732b27ce6fe3b41b3ec1fd7bad4c374f6c4d97c34101ad274f1c9783b72 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i83564140ebf2e1f802106b61ba96f9a9d9d439ec86f9f2a08676beea206b2ade "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    ia561bc68d074f8a902de97a01b770979657ec027086a50f6155f19e112a56d6a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    icaf70b8c2e4eb4b4ea516ec92c86f703b2e1b0b35c32339faec7df8abf5aa7ba "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/calendar"
    icde6de5a6d7e4df71858b7b4d0a200e3583cbf4f2ec6a50dd173713f3f76c79c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments"
    i188c5cbea292c3ad74ec0251c7d640496b3f2d7301aa020f91325adea45a4427 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i308a595453d756eba27a16accae5d44920f86478ad7c30f3173b004296a51dc3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/extensions/item"
    i632edb3f4dacb5abe9443ab6f3ac4264eff3b76a60b4802a9594b7dc718c4bef "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/attachments/item"
    iba8898c88c80074167c6661c5f80875a66bb01111f106d0509b6aeb775567ac7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from groups
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i35573439ed685eddef23590bc0a1c9fb4f2ea0675471fd4af0af13251f5a7c46.AcceptRequestBuilder) {
    return i35573439ed685eddef23590bc0a1c9fb4f2ea0675471fd4af0af13251f5a7c46.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*icde6de5a6d7e4df71858b7b4d0a200e3583cbf4f2ec6a50dd173713f3f76c79c.AttachmentsRequestBuilder) {
    return icde6de5a6d7e4df71858b7b4d0a200e3583cbf4f2ec6a50dd173713f3f76c79c.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i632edb3f4dacb5abe9443ab6f3ac4264eff3b76a60b4802a9594b7dc718c4bef.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i632edb3f4dacb5abe9443ab6f3ac4264eff3b76a60b4802a9594b7dc718c4bef.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*icaf70b8c2e4eb4b4ea516ec92c86f703b2e1b0b35c32339faec7df8abf5aa7ba.CalendarRequestBuilder) {
    return icaf70b8c2e4eb4b4ea516ec92c86f703b2e1b0b35c32339faec7df8abf5aa7ba.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i36bd95c8b35fc004367c0ffab7c4b0a838fb9f35ac8297d9992fa9c7e6915b1a.CancelRequestBuilder) {
    return i36bd95c8b35fc004367c0ffab7c4b0a838fb9f35ac8297d9992fa9c7e6915b1a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/events/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*i541aa50f9ab36177d08a293614f2097ab3bbac67379aa06d892a8c89f9b9d945.DeclineRequestBuilder) {
    return i541aa50f9ab36177d08a293614f2097ab3bbac67379aa06d892a8c89f9b9d945.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for groups
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
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
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i72787732b27ce6fe3b41b3ec1fd7bad4c374f6c4d97c34101ad274f1c9783b72.DismissReminderRequestBuilder) {
    return i72787732b27ce6fe3b41b3ec1fd7bad4c374f6c4d97c34101ad274f1c9783b72.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i5d6da724adab0c2686d22abc37fe4b97f4fbfc7f69783572b96133b75798ef15.ExtensionsRequestBuilder) {
    return i5d6da724adab0c2686d22abc37fe4b97f4fbfc7f69783572b96133b75798ef15.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i308a595453d756eba27a16accae5d44920f86478ad7c30f3173b004296a51dc3.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i308a595453d756eba27a16accae5d44920f86478ad7c30f3173b004296a51dc3.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i286f1c31f8fd38525ee7d192dffca61b2a6863707a278d98bdd7935381861683.ForwardRequestBuilder) {
    return i286f1c31f8fd38525ee7d192dffca61b2a6863707a278d98bdd7935381861683.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from groups
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i83564140ebf2e1f802106b61ba96f9a9d9d439ec86f9f2a08676beea206b2ade.MultiValueExtendedPropertiesRequestBuilder) {
    return i83564140ebf2e1f802106b61ba96f9a9d9d439ec86f9f2a08676beea206b2ade.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iba8898c88c80074167c6661c5f80875a66bb01111f106d0509b6aeb775567ac7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return iba8898c88c80074167c6661c5f80875a66bb01111f106d0509b6aeb775567ac7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in groups
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ia561bc68d074f8a902de97a01b770979657ec027086a50f6155f19e112a56d6a.SingleValueExtendedPropertiesRequestBuilder) {
    return ia561bc68d074f8a902de97a01b770979657ec027086a50f6155f19e112a56d6a.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i188c5cbea292c3ad74ec0251c7d640496b3f2d7301aa020f91325adea45a4427.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i188c5cbea292c3ad74ec0251c7d640496b3f2d7301aa020f91325adea45a4427.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*i15f92df896b379c185f794e9ffab67e334dedbec9e88d7b86bd84d68d577ae5f.SnoozeReminderRequestBuilder) {
    return i15f92df896b379c185f794e9ffab67e334dedbec9e88d7b86bd84d68d577ae5f.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*i5c0221b58587f947d01c8d30e7024054af1cad4538de5d1c62ffb60e9f974153.TentativelyAcceptRequestBuilder) {
    return i5c0221b58587f947d01c8d30e7024054af1cad4538de5d1c62ffb60e9f974153.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
