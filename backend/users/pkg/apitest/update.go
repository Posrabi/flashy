package apitest

type CreateTestType int

// nolint: revive,  stylecheck
const (
	CreateTestType_USER = iota + 1
)

// Needed for non static test data fields.
func UpdateTestDataAfterCreate(createType CreateTestType) error {
	switch createType {
	case CreateTestType_USER:
		SetUserIDs()
	default:
	}
	return nil
}
