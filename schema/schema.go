package schema

func getSchema() string {
	return `
schema {
	query: Queries
}

type Queries {
	listUsers: [User]
}

type User {
	id: ID
	displayname: String
	username: String
}
`
}
