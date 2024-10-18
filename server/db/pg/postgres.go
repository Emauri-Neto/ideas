package pg

type Driver struct{}

func (d Driver) Schema() []string {
	return []string{
		createMigrationTable(),
		createUserTable(),
		createWorkspace(),
		createStudyTable(),
		createDiscussionThreadTable(),
		createInvitationTable(),
		createUsersInvitationTable(),
		createUsersStudyTable(),
		createUsersThreadTable(),
	}
}