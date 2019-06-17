package schema

const schemaSctring = `
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
