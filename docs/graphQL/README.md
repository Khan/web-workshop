# GraphQL

The GraphQL schema in this directory was generated using [openapi-to-graphql](https://github.com/IBM/openapi-to-graphql/tree/master/packages/openapi-to-graphql-cli) from the OpenAPIv3 spec. Unfortunately, it does not support OpenAPIv3 YAML input, so you have to convert to JSON first. [Swagger Editor]() or [yq](https://github.com/mikefarah/yq)


```
npm init -f
npm i -D openapi-to-graphql-cli
npx openapi-to-graphql ./todoOA3.json --save ./todoGQL.yaml
