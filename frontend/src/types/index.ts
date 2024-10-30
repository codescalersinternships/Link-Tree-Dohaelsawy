export interface APIResponse<T> {
    statusCode: number,
    data: T
}

export type User = {
    id: number,
    first_name: string,
    last_name: string,
    username: string,
    image: string,
    email: string,
    password: string,
    phone: string,
    link_tree_url: string,
    bio: string,
    token: string,
    createdAt: string,
}

export type Link = {
    id: number,
    name: string,
    url: string,
    userId: number,
    createdAt: string,
    updatedAt: string,
}

export type Error = {
    err: Error
}