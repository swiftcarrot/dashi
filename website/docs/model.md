---
id: model
title: Model
sidebar_label: model
---

| attribute    | postgres  | go            | graphql type |
| ------------ | --------- | ------------- | ------------ |
| integer      | int       | int           | Int          |
| bigint       | bigint    | int64         | Int          |
| string       | text      | string        | String       |
| json         | jsonb     | slices.Map    | Map          |
| uuid         | uuid      | uuid.UUID     | UUID         |
| date         | date      | time.Time     | Time         |
| datetime     | timestamp | time.Time     | Time         |
| text         | text      | string        | String       |
| float        | float     | float64       | Float        |
| bool/boolean | boolean   | bool          | Boolean      |
| strings      | \_text    | slices.String | [String!]    |
| integers     | \_int4    | slices.Int    | [Int!]       |
| uuids        | \_uuid    | slices.UUID   | [UUID!]      |
| floats       | \_float8  | slices.Float  | [Float!]     |

**For All attributes with `nulls` prefix**

- **Go**: nulls.[OriginalType], support string, float, int, time, uuid, bool
- **Postgres**: same type, but column nullable
- **Graphql**: same type but nullable
