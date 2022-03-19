export type Book = {
	id: string,
	title: string,
	snippet: string,
	genre_id: string,
	pages_cnt: number,
	pub_year: number,
	deleted: boolean,
}

export type Genre = {
	id: string,
	title: string,
	snippet: string,
	deleted: boolean,
}

export type Author = {
	id: string,
	firstname: string,
	lastname: string,
	surname: string,
	snippet: string,
	deleted: boolean,
}

export type AuthorWithBooks = {
	author: Author,
	books_ids: string[],
}

export type BookAuthor = {
	id: string,
	book_id: string,
	author_id: string,
}

export type BookWithAuthors = {
	book: Book,
	authors_ids: string[],
}

export type User = {
	id: string,
	username: string,
	email: string,
	is_admin: boolean,
	deleted: boolean,
}

export type LoginBody = {
	username: string,
	password: string,
}

export type RegisterBody= {
	username: string,
	password: string,
	email: string,
}

export type Link = {
	id: string,
	link: string,
	deleted: boolean,
	book_id: string,
}