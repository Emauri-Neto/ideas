package pg

type Driver struct{}

func (d Driver) Schema() []string {
	return []string{
		createMigrationTable(),
		createUserTable(),
		createVerificationCodeTable(),
		createSessionTable(),
		createStudyTable(),
	}
}