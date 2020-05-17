# 3. JSON [see notes](../docs/03-json.md)

- Create a type `Todo` with fields for ID, Title, Completed, URL, and Order
- Rename the `Echo` handler to `ListTodos`
- Create a slice of Todo values with some dummy data.
- Marshal the slice to JSON and write it to the client.
- Use `w.WriteHeader` to explicitly set the response status code.
- Include the Content-Type header so clients understand the response.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
- See what happens when a nil slice is provided.

## File Changes:

```
Modified main.go
```
