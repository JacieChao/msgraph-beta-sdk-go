package models
type AutoAdmittedUsersType int

const (
    EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE AutoAdmittedUsersType = iota
    EVERYONE_AUTOADMITTEDUSERSTYPE
)

func (i AutoAdmittedUsersType) String() string {
    return []string{"everyoneInCompany", "everyone"}[i]
}
func ParseAutoAdmittedUsersType(v string) (any, error) {
    result := EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE
    switch v {
        case "everyoneInCompany":
            result = EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE
        case "everyone":
            result = EVERYONE_AUTOADMITTEDUSERSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAutoAdmittedUsersType(values []AutoAdmittedUsersType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AutoAdmittedUsersType) isMultiValue() bool {
    return false
}
