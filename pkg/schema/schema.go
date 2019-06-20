package schema

const schemaSctring = `
schema {
	query: Queries
	mutation: Mutation
}

type Queries {
	listUsers: [User]
}

type Mutation {
	saveUser(user: UserInput!): User
}

type User {
	displayname: String
	username: String
}

input UserInput {
	displayname: String
	username: String
}
`
