package schema

const schemaSctring = `
schema {
	query: Queries
}

type Queries {
	listUsers: [User]
}

type User {
	displayname: String
	username: String
	image: String
}
`
