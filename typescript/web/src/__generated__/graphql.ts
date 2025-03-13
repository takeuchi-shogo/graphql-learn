export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type Node = {
  id: Scalars['ID']['output'];
};

export type Post = Node & {
  __typename?: 'Post';
  id: Scalars['ID']['output'];
  title: Scalars['String']['output'];
  user: User;
};

export type PostConnection = {
  __typename?: 'PostConnection';
  edges: Array<PostEdge>;
  nodes: Array<Post>;
};

export type PostEdge = {
  __typename?: 'PostEdge';
  cursor: Scalars['ID']['output'];
  node: Post;
};

export type Query = {
  __typename?: 'Query';
  posts: PostConnection;
  user: User;
};


export type QueryUserArgs = {
  id: Scalars['ID']['input'];
};

export type User = Node & {
  __typename?: 'User';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type PostFragment = { __typename?: 'Post', id: string, title: string };

export type GetPostsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetPostsQuery = { __typename?: 'Query', posts: { __typename?: 'PostConnection', edges: Array<{ __typename?: 'PostEdge', cursor: string, node: { __typename?: 'Post', id: string, title: string } }> } };
