# The Owl House API

WIP

### First Deliverable

Based on the self-study material and mentorship covered until this deliverable, we suggest you perform the following:
Create an API
Add an endpoint to read from a CSV file
The CSV should have any information, for example:

```
1,bulbasaur
2,ivysaur
3,venusaur
```

- The items in the CSV must have an ID element (int value)
- The endpoint should get information from the CSV by some field
- The result should be displayed as a response
- Clean architecture proposal
- Use best practices
- Handle the Errors

> Note: what has been listed in this deliverable is just for guidance and to help you distribute your workload; you can
> deliver more or fewer items if necessary. However, if you deliver fewer items at this point, you have to cover the
> remaining tasks in the following deliverable.

## Resources

Resources to improve the API.

https://github.com/golang-standards/project-layout

https://www.youtube.com/watch?v=oL6JBUk6tj0

https://refactoring.guru/design-patterns/catalog

## Mockery

```
mockery --name=CharacterRepository --srcpkg=./internal/service --output=./internal/service/mocks
```
