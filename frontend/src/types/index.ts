export interface APIResponse<T> {
    statusCode: number,
    data: T
}

export type User = {
    id: number,
    firstName: string,
    lastName: string,
    username: string,
    image: string,
    email: string,
    password: string,
    phone: string,
    linkTreeUrl: string,
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