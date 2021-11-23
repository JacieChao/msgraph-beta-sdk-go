package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// group 
type Group struct {
    DirectoryObject
    // The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
    acceptedSenders []DirectoryObject;
    // 
    accessType *GroupAccessType;
    // Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    allowExternalSenders *bool;
    // Represents the app roles a group has been granted for an application. Supports $expand.
    appRoleAssignments []AppRoleAssignment;
    // The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only.
    assignedLabels []AssignedLabel;
    // The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
    assignedLicenses []AssignedLicense;
    // Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    autoSubscribeNewMembers *bool;
    // The group's calendar. Read-only.
    calendar *Calendar;
    // The calendar view for the calendar. Read-only.
    calendarView []Event;
    // Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
    classification *string;
    // The group's conversations.
    conversations []Conversation;
    // App ID of the app used to create the group. Can be null for some groups. Returned by default. Read-only. Supports $filter (eq, ne, NOT, in, startsWith).
    createdByAppId *string;
    // Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
    createdOnBehalfOf *DirectoryObject;
    // An optional description for the group. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
    description *string;
    // The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string;
    // The group's default drive. Read-only.
    drive *Drive;
    // The group's drives. Read-only.
    drives []Drive;
    // Endpoints for the group. Read-only. Nullable.
    endpoints []Endpoint;
    // The group's calendar events.
    events []Event;
    // Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of open extensions defined for the group. Read-only. Nullable.
    extensions []Extension;
    // The collection of lifecycle policies for this group. Read-only. Nullable.
    groupLifecyclePolicies []GroupLifecyclePolicy;
    // Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, NOT).
    groupTypes []string;
    // Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
    hasMembersWithLicenseErrors *bool;
    // True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    hideFromAddressLists *bool;
    // True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    hideFromOutlookClients *bool;
    // Identifies the info segments assigned to the group. Returned by default. Supports $filter (eq, NOT, ge, le, startsWith).
    infoCatalogs []string;
    // 
    isArchived *bool;
    // Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, NOT).
    isAssignableToRole *bool;
    // 
    isFavorite *bool;
    // Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    isSubscribedByMail *bool;
    // Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
    licenseProcessingState *LicenseProcessingState;
    // The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
    mail *string;
    // Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, NOT).
    mailEnabled *bool;
    // The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] ' ; : . <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
    mailNickname *string;
    // 
    mdmAppId *string;
    // Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand.
    members []DirectoryObject;
    // The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
    membershipRule *string;
    // Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, NOT, in).
    membershipRuleProcessingState *string;
    // Describes the processing status for rules-based dynamic groups. The property is null for non-rule based dynamic groups or if the dynamic group processing has been paused. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). Read-only.
    membershipRuleProcessingStatus *MembershipRuleProcessingStatus;
    // A list of group members with license errors from this group-based license assignment. Read-only.
    membersWithLicenseErrors []DirectoryObject;
    // Read-only.
    onenote *Onenote;
    // Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
    onPremisesDomainName *string;
    // Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
    onPremisesNetBiosName *string;
    // Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, NOT).
    onPremisesProvisioningErrors []OnPremisesProvisioningError;
    // Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith). Read-only.
    onPremisesSamAccountName *string;
    // Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only.
    onPremisesSecurityIdentifier *string;
    // true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, NOT, in, and eq on null values).
    onPremisesSyncEnabled *bool;
    // 
    organizationId *string;
    // The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand.
    owners []DirectoryObject;
    // The permission that has been granted for a group to a specific application. Supports $expand.
    permissionGrants []ResourceSpecificPermissionGrant;
    // The group's profile photo
    photo *ProfilePhoto;
    // The profile photos owned by the group. Read-only. Nullable.
    photos []ProfilePhoto;
    // Entry-point to Planner resource that might exist for a Unified Group.
    planner *PlannerGroup;
    // The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default.
    preferredDataLocation *string;
    // The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
    preferredLanguage *string;
    // Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
    proxyAddresses []string;
    // The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
    rejectedSenders []DirectoryObject;
    // Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
    renewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Specifies the group behaviors that can be set for a Microsoft 365 group during creation. This can be set only as part of creation (POST). Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled. For more information, see Set Microsoft 365 group behaviors and provisioning options.
    resourceBehaviorOptions []string;
    // Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team. For more information, see Set Microsoft 365 group behaviors and provisioning options.
    resourceProvisioningOptions []string;
    // Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, NOT, in).
    securityEnabled *bool;
    // Security identifier of the group, used in Windows scenarios. Returned by default.
    securityIdentifier *string;
    // Read-only. Nullable.
    settings []DirectorySetting;
    // The list of SharePoint sites in this group. Access the default site with /sites/root.
    sites []Site;
    // 
    team *Team;
    // Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
    theme *string;
    // The group's conversation threads. Nullable.
    threads []ConversationThread;
    // 
    transitiveMemberOf []DirectoryObject;
    // 
    transitiveMembers []DirectoryObject;
    // Count of conversations that have been delivered one or more new posts since the signed-in user's last visit to the group. This property is the same as unseenCount. Returned only on $select.
    unseenConversationsCount *int32;
    // Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
    unseenCount *int32;
    // Count of new posts that have been delivered to the group's conversations since the signed-in user's last visit to the group. Returned only on $select.
    unseenMessagesCount *int32;
    // Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
    visibility *string;
}
// NewGroup instantiates a new group and sets the default values.
func NewGroup()(*Group) {
    m := &Group{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// GetAcceptedSenders gets the acceptedSenders property value. The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
func (m *Group) GetAcceptedSenders()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.acceptedSenders
    }
}
// GetAccessType gets the accessType property value. 
func (m *Group) GetAccessType()(*GroupAccessType) {
    if m == nil {
        return nil
    } else {
        return m.accessType
    }
}
// GetAllowExternalSenders gets the allowExternalSenders property value. Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetAllowExternalSenders()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowExternalSenders
    }
}
// GetAppRoleAssignments gets the appRoleAssignments property value. Represents the app roles a group has been granted for an application. Supports $expand.
func (m *Group) GetAppRoleAssignments()([]AppRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.appRoleAssignments
    }
}
// GetAssignedLabels gets the assignedLabels property value. The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only.
func (m *Group) GetAssignedLabels()([]AssignedLabel) {
    if m == nil {
        return nil
    } else {
        return m.assignedLabels
    }
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
func (m *Group) GetAssignedLicenses()([]AssignedLicense) {
    if m == nil {
        return nil
    } else {
        return m.assignedLicenses
    }
}
// GetAutoSubscribeNewMembers gets the autoSubscribeNewMembers property value. Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetAutoSubscribeNewMembers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoSubscribeNewMembers
    }
}
// GetCalendar gets the calendar property value. The group's calendar. Read-only.
func (m *Group) GetCalendar()(*Calendar) {
    if m == nil {
        return nil
    } else {
        return m.calendar
    }
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Read-only.
func (m *Group) GetCalendarView()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.calendarView
    }
}
// GetClassification gets the classification property value. Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Group) GetClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
// GetConversations gets the conversations property value. The group's conversations.
func (m *Group) GetConversations()([]Conversation) {
    if m == nil {
        return nil
    } else {
        return m.conversations
    }
}
// GetCreatedByAppId gets the createdByAppId property value. App ID of the app used to create the group. Can be null for some groups. Returned by default. Read-only. Supports $filter (eq, ne, NOT, in, startsWith).
func (m *Group) GetCreatedByAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByAppId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCreatedOnBehalfOf gets the createdOnBehalfOf property value. The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
func (m *Group) GetCreatedOnBehalfOf()(*DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.createdOnBehalfOf
    }
}
// GetDescription gets the description property value. An optional description for the group. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
func (m *Group) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Group) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDrive gets the drive property value. The group's default drive. Read-only.
func (m *Group) GetDrive()(*Drive) {
    if m == nil {
        return nil
    } else {
        return m.drive
    }
}
// GetDrives gets the drives property value. The group's drives. Read-only.
func (m *Group) GetDrives()([]Drive) {
    if m == nil {
        return nil
    } else {
        return m.drives
    }
}
// GetEndpoints gets the endpoints property value. Endpoints for the group. Read-only. Nullable.
func (m *Group) GetEndpoints()([]Endpoint) {
    if m == nil {
        return nil
    } else {
        return m.endpoints
    }
}
// GetEvents gets the events property value. The group's calendar events.
func (m *Group) GetEvents()([]Event) {
    if m == nil {
        return nil
    } else {
        return m.events
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the group. Read-only. Nullable.
func (m *Group) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetGroupLifecyclePolicies gets the groupLifecyclePolicies property value. The collection of lifecycle policies for this group. Read-only. Nullable.
func (m *Group) GetGroupLifecyclePolicies()([]GroupLifecyclePolicy) {
    if m == nil {
        return nil
    } else {
        return m.groupLifecyclePolicies
    }
}
// GetGroupTypes gets the groupTypes property value. Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, NOT).
func (m *Group) GetGroupTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupTypes
    }
}
// GetHasMembersWithLicenseErrors gets the hasMembersWithLicenseErrors property value. Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
func (m *Group) GetHasMembersWithLicenseErrors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasMembersWithLicenseErrors
    }
}
// GetHideFromAddressLists gets the hideFromAddressLists property value. True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetHideFromAddressLists()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideFromAddressLists
    }
}
// GetHideFromOutlookClients gets the hideFromOutlookClients property value. True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetHideFromOutlookClients()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hideFromOutlookClients
    }
}
// GetInfoCatalogs gets the infoCatalogs property value. Identifies the info segments assigned to the group. Returned by default. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Group) GetInfoCatalogs()([]string) {
    if m == nil {
        return nil
    } else {
        return m.infoCatalogs
    }
}
// GetIsArchived gets the isArchived property value. 
func (m *Group) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
// GetIsAssignableToRole gets the isAssignableToRole property value. Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, NOT).
func (m *Group) GetIsAssignableToRole()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssignableToRole
    }
}
// GetIsFavorite gets the isFavorite property value. 
func (m *Group) GetIsFavorite()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavorite
    }
}
// GetIsSubscribedByMail gets the isSubscribedByMail property value. Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetIsSubscribedByMail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSubscribedByMail
    }
}
// GetLicenseProcessingState gets the licenseProcessingState property value. Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
func (m *Group) GetLicenseProcessingState()(*LicenseProcessingState) {
    if m == nil {
        return nil
    } else {
        return m.licenseProcessingState
    }
}
// GetMail gets the mail property value. The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
func (m *Group) GetMail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mail
    }
}
// GetMailEnabled gets the mailEnabled property value. Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, NOT).
func (m *Group) GetMailEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailEnabled
    }
}
// GetMailNickname gets the mailNickname property value. The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] ' ; : . <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
func (m *Group) GetMailNickname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailNickname
    }
}
// GetMdmAppId gets the mdmAppId property value. 
func (m *Group) GetMdmAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mdmAppId
    }
}
// GetMemberOf gets the memberOf property value. Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
func (m *Group) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// GetMembers gets the members property value. Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand.
func (m *Group) GetMembers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetMembershipRule gets the membershipRule property value. The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Group) GetMembershipRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRule
    }
}
// GetMembershipRuleProcessingState gets the membershipRuleProcessingState property value. Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, NOT, in).
func (m *Group) GetMembershipRuleProcessingState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleProcessingState
    }
}
// GetMembershipRuleProcessingStatus gets the membershipRuleProcessingStatus property value. Describes the processing status for rules-based dynamic groups. The property is null for non-rule based dynamic groups or if the dynamic group processing has been paused. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). Read-only.
func (m *Group) GetMembershipRuleProcessingStatus()(*MembershipRuleProcessingStatus) {
    if m == nil {
        return nil
    } else {
        return m.membershipRuleProcessingStatus
    }
}
// GetMembersWithLicenseErrors gets the membersWithLicenseErrors property value. A list of group members with license errors from this group-based license assignment. Read-only.
func (m *Group) GetMembersWithLicenseErrors()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.membersWithLicenseErrors
    }
}
// GetOnenote gets the onenote property value. Read-only.
func (m *Group) GetOnenote()(*Onenote) {
    if m == nil {
        return nil
    } else {
        return m.onenote
    }
}
// GetOnPremisesDomainName gets the onPremisesDomainName property value. Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
func (m *Group) GetOnPremisesDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesDomainName
    }
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
func (m *Group) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// GetOnPremisesNetBiosName gets the onPremisesNetBiosName property value. Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
func (m *Group) GetOnPremisesNetBiosName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesNetBiosName
    }
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, NOT).
func (m *Group) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningError) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesProvisioningErrors
    }
}
// GetOnPremisesSamAccountName gets the onPremisesSamAccountName property value. Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith). Read-only.
func (m *Group) GetOnPremisesSamAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSamAccountName
    }
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only.
func (m *Group) GetOnPremisesSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSecurityIdentifier
    }
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, NOT, in, and eq on null values).
func (m *Group) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// GetOrganizationId gets the organizationId property value. 
func (m *Group) GetOrganizationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationId
    }
}
// GetOwners gets the owners property value. The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand.
func (m *Group) GetOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.owners
    }
}
// GetPermissionGrants gets the permissionGrants property value. The permission that has been granted for a group to a specific application. Supports $expand.
func (m *Group) GetPermissionGrants()([]ResourceSpecificPermissionGrant) {
    if m == nil {
        return nil
    } else {
        return m.permissionGrants
    }
}
// GetPhoto gets the photo property value. The group's profile photo
func (m *Group) GetPhoto()(*ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photo
    }
}
// GetPhotos gets the photos property value. The profile photos owned by the group. Read-only. Nullable.
func (m *Group) GetPhotos()([]ProfilePhoto) {
    if m == nil {
        return nil
    } else {
        return m.photos
    }
}
// GetPlanner gets the planner property value. Entry-point to Planner resource that might exist for a Unified Group.
func (m *Group) GetPlanner()(*PlannerGroup) {
    if m == nil {
        return nil
    } else {
        return m.planner
    }
}
// GetPreferredDataLocation gets the preferredDataLocation property value. The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default.
func (m *Group) GetPreferredDataLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredDataLocation
    }
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
func (m *Group) GetPreferredLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguage
    }
}
// GetProxyAddresses gets the proxyAddresses property value. Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Group) GetProxyAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.proxyAddresses
    }
}
// GetRejectedSenders gets the rejectedSenders property value. The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *Group) GetRejectedSenders()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.rejectedSenders
    }
}
// GetRenewedDateTime gets the renewedDateTime property value. Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) GetRenewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.renewedDateTime
    }
}
// GetResourceBehaviorOptions gets the resourceBehaviorOptions property value. Specifies the group behaviors that can be set for a Microsoft 365 group during creation. This can be set only as part of creation (POST). Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled. For more information, see Set Microsoft 365 group behaviors and provisioning options.
func (m *Group) GetResourceBehaviorOptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resourceBehaviorOptions
    }
}
// GetResourceProvisioningOptions gets the resourceProvisioningOptions property value. Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team. For more information, see Set Microsoft 365 group behaviors and provisioning options.
func (m *Group) GetResourceProvisioningOptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.resourceProvisioningOptions
    }
}
// GetSecurityEnabled gets the securityEnabled property value. Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, NOT, in).
func (m *Group) GetSecurityEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityEnabled
    }
}
// GetSecurityIdentifier gets the securityIdentifier property value. Security identifier of the group, used in Windows scenarios. Returned by default.
func (m *Group) GetSecurityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.securityIdentifier
    }
}
// GetSettings gets the settings property value. Read-only. Nullable.
func (m *Group) GetSettings()([]DirectorySetting) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetSites gets the sites property value. The list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *Group) GetSites()([]Site) {
    if m == nil {
        return nil
    } else {
        return m.sites
    }
}
// GetTeam gets the team property value. 
func (m *Group) GetTeam()(*Team) {
    if m == nil {
        return nil
    } else {
        return m.team
    }
}
// GetTheme gets the theme property value. Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
func (m *Group) GetTheme()(*string) {
    if m == nil {
        return nil
    } else {
        return m.theme
    }
}
// GetThreads gets the threads property value. The group's conversation threads. Nullable.
func (m *Group) GetThreads()([]ConversationThread) {
    if m == nil {
        return nil
    } else {
        return m.threads
    }
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. 
func (m *Group) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// GetTransitiveMembers gets the transitiveMembers property value. 
func (m *Group) GetTransitiveMembers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMembers
    }
}
// GetUnseenConversationsCount gets the unseenConversationsCount property value. Count of conversations that have been delivered one or more new posts since the signed-in user's last visit to the group. This property is the same as unseenCount. Returned only on $select.
func (m *Group) GetUnseenConversationsCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unseenConversationsCount
    }
}
// GetUnseenCount gets the unseenCount property value. Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) GetUnseenCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unseenCount
    }
}
// GetUnseenMessagesCount gets the unseenMessagesCount property value. Count of new posts that have been delivered to the group's conversations since the signed-in user's last visit to the group. Returned only on $select.
func (m *Group) GetUnseenMessagesCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unseenMessagesCount
    }
}
// GetVisibility gets the visibility property value. Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
func (m *Group) GetVisibility()(*string) {
    if m == nil {
        return nil
    } else {
        return m.visibility
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Group) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["acceptedSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetAcceptedSenders(res)
        }
        return nil
    }
    res["accessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(GroupAccessType)
            m.SetAccessType(&cast)
        }
        return nil
    }
    res["allowExternalSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalSenders(val)
        }
        return nil
    }
    res["appRoleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAppRoleAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AppRoleAssignment))
            }
            m.SetAppRoleAssignments(res)
        }
        return nil
    }
    res["assignedLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedLabel))
            }
            m.SetAssignedLabels(res)
        }
        return nil
    }
    res["assignedLicenses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAssignedLicense() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicense, len(val))
            for i, v := range val {
                res[i] = *(v.(*AssignedLicense))
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
    res["autoSubscribeNewMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoSubscribeNewMembers(val)
        }
        return nil
    }
    res["calendar"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCalendar() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(*Calendar))
        }
        return nil
    }
    res["calendarView"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*Event))
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val)
        }
        return nil
    }
    res["conversations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Conversation, len(val))
            for i, v := range val {
                res[i] = *(v.(*Conversation))
            }
            m.SetConversations(res)
        }
        return nil
    }
    res["createdByAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByAppId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["createdOnBehalfOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnBehalfOf(val.(*DirectoryObject))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["drive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(*Drive))
        }
        return nil
    }
    res["drives"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDrive() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Drive, len(val))
            for i, v := range val {
                res[i] = *(v.(*Drive))
            }
            m.SetDrives(res)
        }
        return nil
    }
    res["endpoints"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEndpoint() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Endpoint, len(val))
            for i, v := range val {
                res[i] = *(v.(*Endpoint))
            }
            m.SetEndpoints(res)
        }
        return nil
    }
    res["events"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Event, len(val))
            for i, v := range val {
                res[i] = *(v.(*Event))
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["groupLifecyclePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroupLifecyclePolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupLifecyclePolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*GroupLifecyclePolicy))
            }
            m.SetGroupLifecyclePolicies(res)
        }
        return nil
    }
    res["groupTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupTypes(res)
        }
        return nil
    }
    res["hasMembersWithLicenseErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasMembersWithLicenseErrors(val)
        }
        return nil
    }
    res["hideFromAddressLists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideFromAddressLists(val)
        }
        return nil
    }
    res["hideFromOutlookClients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideFromOutlookClients(val)
        }
        return nil
    }
    res["infoCatalogs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInfoCatalogs(res)
        }
        return nil
    }
    res["isArchived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsArchived(val)
        }
        return nil
    }
    res["isAssignableToRole"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssignableToRole(val)
        }
        return nil
    }
    res["isFavorite"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavorite(val)
        }
        return nil
    }
    res["isSubscribedByMail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSubscribedByMail(val)
        }
        return nil
    }
    res["licenseProcessingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLicenseProcessingState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseProcessingState(val.(*LicenseProcessingState))
        }
        return nil
    }
    res["mail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailEnabled(val)
        }
        return nil
    }
    res["mailNickname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["mdmAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmAppId(val)
        }
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["membershipRule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRule(val)
        }
        return nil
    }
    res["membershipRuleProcessingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleProcessingState(val)
        }
        return nil
    }
    res["membershipRuleProcessingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMembershipRuleProcessingStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipRuleProcessingStatus(val.(*MembershipRuleProcessingStatus))
        }
        return nil
    }
    res["membersWithLicenseErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetMembersWithLicenseErrors(res)
        }
        return nil
    }
    res["onenote"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnenote() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnenote(val.(*Onenote))
        }
        return nil
    }
    res["onPremisesDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesDomainName(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesNetBiosName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesNetBiosName(val)
        }
        return nil
    }
    res["onPremisesProvisioningErrors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesProvisioningError() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesProvisioningError, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnPremisesProvisioningError))
            }
            m.SetOnPremisesProvisioningErrors(res)
        }
        return nil
    }
    res["onPremisesSamAccountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSamAccountName(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["organizationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationId(val)
        }
        return nil
    }
    res["owners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["permissionGrants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceSpecificPermissionGrant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceSpecificPermissionGrant, len(val))
            for i, v := range val {
                res[i] = *(v.(*ResourceSpecificPermissionGrant))
            }
            m.SetPermissionGrants(res)
        }
        return nil
    }
    res["photo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoto(val.(*ProfilePhoto))
        }
        return nil
    }
    res["photos"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProfilePhoto() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProfilePhoto, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProfilePhoto))
            }
            m.SetPhotos(res)
        }
        return nil
    }
    res["planner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanner(val.(*PlannerGroup))
        }
        return nil
    }
    res["preferredDataLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredDataLocation(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["proxyAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetProxyAddresses(res)
        }
        return nil
    }
    res["rejectedSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetRejectedSenders(res)
        }
        return nil
    }
    res["renewedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewedDateTime(val)
        }
        return nil
    }
    res["resourceBehaviorOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResourceBehaviorOptions(res)
        }
        return nil
    }
    res["resourceProvisioningOptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetResourceProvisioningOptions(res)
        }
        return nil
    }
    res["securityEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEnabled(val)
        }
        return nil
    }
    res["securityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityIdentifier(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectorySetting() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectorySetting, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectorySetting))
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["sites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSite() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Site, len(val))
            for i, v := range val {
                res[i] = *(v.(*Site))
            }
            m.SetSites(res)
        }
        return nil
    }
    res["team"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeam() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeam(val.(*Team))
        }
        return nil
    }
    res["theme"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val)
        }
        return nil
    }
    res["threads"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationThread() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationThread, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConversationThread))
            }
            m.SetThreads(res)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    res["transitiveMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObject, len(val))
            for i, v := range val {
                res[i] = *(v.(*DirectoryObject))
            }
            m.SetTransitiveMembers(res)
        }
        return nil
    }
    res["unseenConversationsCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnseenConversationsCount(val)
        }
        return nil
    }
    res["unseenCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnseenCount(val)
        }
        return nil
    }
    res["unseenMessagesCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnseenMessagesCount(val)
        }
        return nil
    }
    res["visibility"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
func (m *Group) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Group) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAcceptedSenders()))
        for i, v := range m.GetAcceptedSenders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("acceptedSenders", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessType() != nil {
        cast := m.GetAccessType().String()
        err = writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowExternalSenders", m.GetAllowExternalSenders())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAppRoleAssignments()))
        for i, v := range m.GetAppRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("appRoleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedLabels()))
        for i, v := range m.GetAssignedLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLabels", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoSubscribeNewMembers", m.GetAutoSubscribeNewMembers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classification", m.GetClassification())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConversations()))
        for i, v := range m.GetConversations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("conversations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByAppId", m.GetCreatedByAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDrives()))
        for i, v := range m.GetDrives() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupLifecyclePolicies()))
        for i, v := range m.GetGroupLifecyclePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupLifecyclePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("groupTypes", m.GetGroupTypes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasMembersWithLicenseErrors", m.GetHasMembersWithLicenseErrors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideFromAddressLists", m.GetHideFromAddressLists())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideFromOutlookClients", m.GetHideFromOutlookClients())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("infoCatalogs", m.GetInfoCatalogs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssignableToRole", m.GetIsAssignableToRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavorite", m.GetIsFavorite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSubscribedByMail", m.GetIsSubscribedByMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("licenseProcessingState", m.GetLicenseProcessingState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mailEnabled", m.GetMailEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmAppId", m.GetMdmAppId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("membershipRule", m.GetMembershipRule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("membershipRuleProcessingState", m.GetMembershipRuleProcessingState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("membershipRuleProcessingStatus", m.GetMembershipRuleProcessingStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembersWithLicenseErrors()))
        for i, v := range m.GetMembersWithLicenseErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("membersWithLicenseErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesDomainName", m.GetOnPremisesDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesNetBiosName", m.GetOnPremisesNetBiosName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOnPremisesProvisioningErrors()))
        for i, v := range m.GetOnPremisesProvisioningErrors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSamAccountName", m.GetOnPremisesSamAccountName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationId", m.GetOrganizationId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPermissionGrants()))
        for i, v := range m.GetPermissionGrants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("permissionGrants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhotos()))
        for i, v := range m.GetPhotos() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("photos", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("planner", m.GetPlanner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredDataLocation", m.GetPreferredDataLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRejectedSenders()))
        for i, v := range m.GetRejectedSenders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("rejectedSenders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("renewedDateTime", m.GetRenewedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("resourceBehaviorOptions", m.GetResourceBehaviorOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("resourceProvisioningOptions", m.GetResourceProvisioningOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityEnabled", m.GetSecurityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("securityIdentifier", m.GetSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("team", m.GetTeam())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("theme", m.GetTheme())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreads()))
        for i, v := range m.GetThreads() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("threads", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMembers()))
        for i, v := range m.GetTransitiveMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unseenConversationsCount", m.GetUnseenConversationsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unseenCount", m.GetUnseenCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unseenMessagesCount", m.GetUnseenMessagesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedSenders sets the acceptedSenders property value. The list of users or groups that are allowed to create post's or calendar events in this group. If this list is non-empty then only users or groups listed here are allowed to post.
func (m *Group) SetAcceptedSenders(value []DirectoryObject)() {
    m.acceptedSenders = value
}
// SetAccessType sets the accessType property value. 
func (m *Group) SetAccessType(value *GroupAccessType)() {
    m.accessType = value
}
// SetAllowExternalSenders sets the allowExternalSenders property value. Indicates if people external to the organization can send messages to the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetAllowExternalSenders(value *bool)() {
    m.allowExternalSenders = value
}
// SetAppRoleAssignments sets the appRoleAssignments property value. Represents the app roles a group has been granted for an application. Supports $expand.
func (m *Group) SetAppRoleAssignments(value []AppRoleAssignment)() {
    m.appRoleAssignments = value
}
// SetAssignedLabels sets the assignedLabels property value. The list of sensitivity label pairs (label ID, label name) associated with a Microsoft 365 group. Returned only on $select. Read-only.
func (m *Group) SetAssignedLabels(value []AssignedLabel)() {
    m.assignedLabels = value
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the group. Returned only on $select. Supports $filter (eq).Read-only.
func (m *Group) SetAssignedLicenses(value []AssignedLicense)() {
    m.assignedLicenses = value
}
// SetAutoSubscribeNewMembers sets the autoSubscribeNewMembers property value. Indicates if new members added to the group will be auto-subscribed to receive email notifications. You can set this property in a PATCH request for the group; do not set it in the initial POST request that creates the group. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetAutoSubscribeNewMembers(value *bool)() {
    m.autoSubscribeNewMembers = value
}
// SetCalendar sets the calendar property value. The group's calendar. Read-only.
func (m *Group) SetCalendar(value *Calendar)() {
    m.calendar = value
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Read-only.
func (m *Group) SetCalendarView(value []Event)() {
    m.calendarView = value
}
// SetClassification sets the classification property value. Describes a classification for the group (such as low, medium or high business impact). Valid values for this property are defined by creating a ClassificationList setting value, based on the template definition.Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Group) SetClassification(value *string)() {
    m.classification = value
}
// SetConversations sets the conversations property value. The group's conversations.
func (m *Group) SetConversations(value []Conversation)() {
    m.conversations = value
}
// SetCreatedByAppId sets the createdByAppId property value. App ID of the app used to create the group. Can be null for some groups. Returned by default. Read-only. Supports $filter (eq, ne, NOT, in, startsWith).
func (m *Group) SetCreatedByAppId(value *string)() {
    m.createdByAppId = value
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the group was created. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedOnBehalfOf sets the createdOnBehalfOf property value. The user (or application) that created the group. NOTE: This is not set if the user is an administrator. Read-only.
func (m *Group) SetCreatedOnBehalfOf(value *DirectoryObject)() {
    m.createdOnBehalfOf = value
}
// SetDescription sets the description property value. An optional description for the group. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith) and $search.
func (m *Group) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Group) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. The group's default drive. Read-only.
func (m *Group) SetDrive(value *Drive)() {
    m.drive = value
}
// SetDrives sets the drives property value. The group's drives. Read-only.
func (m *Group) SetDrives(value []Drive)() {
    m.drives = value
}
// SetEndpoints sets the endpoints property value. Endpoints for the group. Read-only. Nullable.
func (m *Group) SetEndpoints(value []Endpoint)() {
    m.endpoints = value
}
// SetEvents sets the events property value. The group's calendar events.
func (m *Group) SetEvents(value []Event)() {
    m.events = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Timestamp of when the group is set to expire. The value cannot be modified and is automatically populated when the group is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the group. Read-only. Nullable.
func (m *Group) SetExtensions(value []Extension)() {
    m.extensions = value
}
// SetGroupLifecyclePolicies sets the groupLifecyclePolicies property value. The collection of lifecycle policies for this group. Read-only. Nullable.
func (m *Group) SetGroupLifecyclePolicies(value []GroupLifecyclePolicy)() {
    m.groupLifecyclePolicies = value
}
// SetGroupTypes sets the groupTypes property value. Specifies the group type and its membership.  If the collection contains Unified, the group is a Microsoft 365 group; otherwise, it's either a security group or distribution group. For details, see groups overview.If the collection includes DynamicMembership, the group has dynamic membership; otherwise, membership is static.  Returned by default. Supports $filter (eq, NOT).
func (m *Group) SetGroupTypes(value []string)() {
    m.groupTypes = value
}
// SetHasMembersWithLicenseErrors sets the hasMembersWithLicenseErrors property value. Indicates whether there are members in this group that have license errors from its group-based license assignment. This property is never returned on a GET operation. You can use it as a $filter argument to get groups that have members with license errors (that is, filter for this property being true). See an example. Supports $filter (eq).
func (m *Group) SetHasMembersWithLicenseErrors(value *bool)() {
    m.hasMembersWithLicenseErrors = value
}
// SetHideFromAddressLists sets the hideFromAddressLists property value. True if the group is not displayed in certain parts of the Outlook UI: the Address Book, address lists for selecting message recipients, and the Browse Groups dialog for searching groups; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetHideFromAddressLists(value *bool)() {
    m.hideFromAddressLists = value
}
// SetHideFromOutlookClients sets the hideFromOutlookClients property value. True if the group is not displayed in Outlook clients, such as Outlook for Windows and Outlook on the web; otherwise, false. Default value is false. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetHideFromOutlookClients(value *bool)() {
    m.hideFromOutlookClients = value
}
// SetInfoCatalogs sets the infoCatalogs property value. Identifies the info segments assigned to the group. Returned by default. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Group) SetInfoCatalogs(value []string)() {
    m.infoCatalogs = value
}
// SetIsArchived sets the isArchived property value. 
func (m *Group) SetIsArchived(value *bool)() {
    m.isArchived = value
}
// SetIsAssignableToRole sets the isAssignableToRole property value. Indicates whether this group can be assigned to an Azure Active Directory role or not. Optional. This property can only be set while creating the group and is immutable. If set to true, the securityEnabled property must also be set to true and the group cannot be a dynamic group (that is, groupTypes cannot contain DynamicMembership). Only callers in Global administrator and Privileged role administrator roles can set this property. The caller must be assigned the RoleManagement.ReadWrite.Directory permission to set this property or update the membership of such groups. For more, see Using a group to manage Azure AD role assignmentsReturned by default. Supports $filter (eq, ne, NOT).
func (m *Group) SetIsAssignableToRole(value *bool)() {
    m.isAssignableToRole = value
}
// SetIsFavorite sets the isFavorite property value. 
func (m *Group) SetIsFavorite(value *bool)() {
    m.isFavorite = value
}
// SetIsSubscribedByMail sets the isSubscribedByMail property value. Indicates whether the signed-in user is subscribed to receive email conversations. Default value is true. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetIsSubscribedByMail(value *bool)() {
    m.isSubscribedByMail = value
}
// SetLicenseProcessingState sets the licenseProcessingState property value. Indicates status of the group license assignment to all members of the group. Default value is false. Read-only. Possible values: QueuedForProcessing, ProcessingInProgress, and ProcessingComplete.Returned only on $select. Read-only.
func (m *Group) SetLicenseProcessingState(value *LicenseProcessingState)() {
    m.licenseProcessingState = value
}
// SetMail sets the mail property value. The SMTP address for the group, for example, 'serviceadmins@contoso.onmicrosoft.com'. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
func (m *Group) SetMail(value *string)() {
    m.mail = value
}
// SetMailEnabled sets the mailEnabled property value. Specifies whether the group is mail-enabled. Required. Returned by default. Supports $filter (eq, ne, NOT).
func (m *Group) SetMailEnabled(value *bool)() {
    m.mailEnabled = value
}
// SetMailNickname sets the mailNickname property value. The mail alias for the group, unique in the organization. Maximum length is 64 characters. This property can contain only characters in the ASCII character set 0 - 127 except the following: @ () / [] ' ; : . <> , SPACE. Required. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
func (m *Group) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetMdmAppId sets the mdmAppId property value. 
func (m *Group) SetMdmAppId(value *string)() {
    m.mdmAppId = value
}
// SetMemberOf sets the memberOf property value. Groups that this group is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand.
func (m *Group) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
// SetMembers sets the members property value. Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), POST (supported for Microsoft 365 groups, security groups and mail-enabled security groups), DELETE (supported for Microsoft 365 groups and security groups). Nullable. Supports $expand.
func (m *Group) SetMembers(value []DirectoryObject)() {
    m.members = value
}
// SetMembershipRule sets the membershipRule property value. The rule that determines members for this group if the group is a dynamic group (groupTypes contains DynamicMembership). For more information about the syntax of the membership rule, see Membership Rules syntax. Returned by default. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Group) SetMembershipRule(value *string)() {
    m.membershipRule = value
}
// SetMembershipRuleProcessingState sets the membershipRuleProcessingState property value. Indicates whether the dynamic membership processing is on or paused. Possible values are On or Paused. Returned by default. Supports $filter (eq, ne, NOT, in).
func (m *Group) SetMembershipRuleProcessingState(value *string)() {
    m.membershipRuleProcessingState = value
}
// SetMembershipRuleProcessingStatus sets the membershipRuleProcessingStatus property value. Describes the processing status for rules-based dynamic groups. The property is null for non-rule based dynamic groups or if the dynamic group processing has been paused. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}). Read-only.
func (m *Group) SetMembershipRuleProcessingStatus(value *MembershipRuleProcessingStatus)() {
    m.membershipRuleProcessingStatus = value
}
// SetMembersWithLicenseErrors sets the membersWithLicenseErrors property value. A list of group members with license errors from this group-based license assignment. Read-only.
func (m *Group) SetMembersWithLicenseErrors(value []DirectoryObject)() {
    m.membersWithLicenseErrors = value
}
// SetOnenote sets the onenote property value. Read-only.
func (m *Group) SetOnenote(value *Onenote)() {
    m.onenote = value
}
// SetOnPremisesDomainName sets the onPremisesDomainName property value. Contains the on-premises domain FQDN, also called dnsDomainName synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
func (m *Group) SetOnPremisesDomainName(value *string)() {
    m.onPremisesDomainName = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. Indicates the last time at which the group was synced with the on-premises directory.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
func (m *Group) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesNetBiosName sets the onPremisesNetBiosName property value. Contains the on-premises netBios name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Read-only.
func (m *Group) SetOnPremisesNetBiosName(value *string)() {
    m.onPremisesNetBiosName = value
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. Errors when using Microsoft synchronization product during provisioning. Returned by default. Supports $filter (eq, NOT).
func (m *Group) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningError)() {
    m.onPremisesProvisioningErrors = value
}
// SetOnPremisesSamAccountName sets the onPremisesSamAccountName property value. Contains the on-premises SAM account name synchronized from the on-premises directory. The property is only populated for customers who are synchronizing their on-premises directory to Azure Active Directory via Azure AD Connect.Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith). Read-only.
func (m *Group) SetOnPremisesSamAccountName(value *string)() {
    m.onPremisesSamAccountName = value
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. Contains the on-premises security identifier (SID) for the group that was synchronized from on-premises to the cloud. Returned by default. Supports $filter on null values. Read-only.
func (m *Group) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this group is synced from an on-premises directory; false if this group was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Returned by default. Read-only. Supports $filter (eq, ne, NOT, in, and eq on null values).
func (m *Group) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetOrganizationId sets the organizationId property value. 
func (m *Group) SetOrganizationId(value *string)() {
    m.organizationId = value
}
// SetOwners sets the owners property value. The owners of the group. The owners are a set of non-admin users who are allowed to modify this object. Limited to 100 owners. Nullable. If this property is not specified when creating a Microsoft 365 group, the calling user is automatically assigned as the group owner. Supports $expand.
func (m *Group) SetOwners(value []DirectoryObject)() {
    m.owners = value
}
// SetPermissionGrants sets the permissionGrants property value. The permission that has been granted for a group to a specific application. Supports $expand.
func (m *Group) SetPermissionGrants(value []ResourceSpecificPermissionGrant)() {
    m.permissionGrants = value
}
// SetPhoto sets the photo property value. The group's profile photo
func (m *Group) SetPhoto(value *ProfilePhoto)() {
    m.photo = value
}
// SetPhotos sets the photos property value. The profile photos owned by the group. Read-only. Nullable.
func (m *Group) SetPhotos(value []ProfilePhoto)() {
    m.photos = value
}
// SetPlanner sets the planner property value. Entry-point to Planner resource that might exist for a Unified Group.
func (m *Group) SetPlanner(value *PlannerGroup)() {
    m.planner = value
}
// SetPreferredDataLocation sets the preferredDataLocation property value. The preferred data location for the Microsoft 365 group. By default, the group inherits the group creator's preferred data location. To set this property, the calling user must be assigned one of the following Azure AD roles:  Global Administrator  User Account Administrator Directory Writer  Exchange Administrator  SharePoint Administrator  For more information about this property, see  OneDrive Online Multi-Geo. Nullable. Returned by default.
func (m *Group) SetPreferredDataLocation(value *string)() {
    m.preferredDataLocation = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for a Microsoft 365 group. Should follow ISO 639-1 Code; for example en-US. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in, startsWith, and eq on null values).
func (m *Group) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetProxyAddresses sets the proxyAddresses property value. Email addresses for the group that direct to the same group mailbox. For example: ['SMTP: bob@contoso.com', 'smtp: bob@sales.contoso.com']. The any operator is required to filter expressions on multi-valued properties. Returned by default. Read-only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Group) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// SetRejectedSenders sets the rejectedSenders property value. The list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *Group) SetRejectedSenders(value []DirectoryObject)() {
    m.rejectedSenders = value
}
// SetRenewedDateTime sets the renewedDateTime property value. Timestamp of when the group was last renewed. This cannot be modified directly and is only updated via the renew service action. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Returned by default. Supports $filter (eq, ne, NOT, ge, le, in). Read-only.
func (m *Group) SetRenewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.renewedDateTime = value
}
// SetResourceBehaviorOptions sets the resourceBehaviorOptions property value. Specifies the group behaviors that can be set for a Microsoft 365 group during creation. This can be set only as part of creation (POST). Possible values are AllowOnlyMembersToPost, HideGroupInOutlook, SubscribeNewGroupMembers, WelcomeEmailDisabled. For more information, see Set Microsoft 365 group behaviors and provisioning options.
func (m *Group) SetResourceBehaviorOptions(value []string)() {
    m.resourceBehaviorOptions = value
}
// SetResourceProvisioningOptions sets the resourceProvisioningOptions property value. Specifies the group resources that are provisioned as part of Microsoft 365 group creation, that are not normally part of default group creation. Possible value is Team. For more information, see Set Microsoft 365 group behaviors and provisioning options.
func (m *Group) SetResourceProvisioningOptions(value []string)() {
    m.resourceProvisioningOptions = value
}
// SetSecurityEnabled sets the securityEnabled property value. Specifies whether the group is a security group. Required. Returned by default. Supports $filter (eq, ne, NOT, in).
func (m *Group) SetSecurityEnabled(value *bool)() {
    m.securityEnabled = value
}
// SetSecurityIdentifier sets the securityIdentifier property value. Security identifier of the group, used in Windows scenarios. Returned by default.
func (m *Group) SetSecurityIdentifier(value *string)() {
    m.securityIdentifier = value
}
// SetSettings sets the settings property value. Read-only. Nullable.
func (m *Group) SetSettings(value []DirectorySetting)() {
    m.settings = value
}
// SetSites sets the sites property value. The list of SharePoint sites in this group. Access the default site with /sites/root.
func (m *Group) SetSites(value []Site)() {
    m.sites = value
}
// SetTeam sets the team property value. 
func (m *Group) SetTeam(value *Team)() {
    m.team = value
}
// SetTheme sets the theme property value. Specifies a Microsoft 365 group's color theme. Possible values are Teal, Purple, Green, Blue, Pink, Orange or Red. Returned by default.
func (m *Group) SetTheme(value *string)() {
    m.theme = value
}
// SetThreads sets the threads property value. The group's conversation threads. Nullable.
func (m *Group) SetThreads(value []ConversationThread)() {
    m.threads = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. 
func (m *Group) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
// SetTransitiveMembers sets the transitiveMembers property value. 
func (m *Group) SetTransitiveMembers(value []DirectoryObject)() {
    m.transitiveMembers = value
}
// SetUnseenConversationsCount sets the unseenConversationsCount property value. Count of conversations that have been delivered one or more new posts since the signed-in user's last visit to the group. This property is the same as unseenCount. Returned only on $select.
func (m *Group) SetUnseenConversationsCount(value *int32)() {
    m.unseenConversationsCount = value
}
// SetUnseenCount sets the unseenCount property value. Count of conversations that have received new posts since the signed-in user last visited the group. Returned only on $select. Supported only on the Get group API (GET /groups/{ID}).
func (m *Group) SetUnseenCount(value *int32)() {
    m.unseenCount = value
}
// SetUnseenMessagesCount sets the unseenMessagesCount property value. Count of new posts that have been delivered to the group's conversations since the signed-in user's last visit to the group. Returned only on $select.
func (m *Group) SetUnseenMessagesCount(value *int32)() {
    m.unseenMessagesCount = value
}
// SetVisibility sets the visibility property value. Specifies the group join policy and group content visibility for groups. Possible values are: Private, Public, or Hiddenmembership. Hiddenmembership can be set only for Microsoft 365 groups, when the groups are created. It can't be updated later. Other values of visibility can be updated after group creation. If visibility value is not specified during group creation on Microsoft Graph, a security group is created as Private by default and Microsoft 365 group is Public. Groups assignable to roles are always Private. See group visibility options to learn more. Returned by default. Nullable.
func (m *Group) SetVisibility(value *string)() {
    m.visibility = value
}
