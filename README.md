# Asana Tiny Go

This is a very (!) tiny API wrapper for [Asana](https://asana.com/developers), specifically
to create tasks. It doesn’t do anything else but creat tasks. I might add more functionality
in the future, but there are no promises.

It is also severely limited in its functionality when creating tasks. The only parameters it
supports are

```go
type TaskRequest struct {
  Workspace   int64             // ID of the workspace the tasks should go in
  Name        string            // task name (title)
  Notes       string            // task description (plain text)
  Followers   []*TaskFollower
  Memberships []*TaskMembership // assign task to project and section in that project
}
```

Why is it not just called `asana-create-task`, you ask? Because I can see myself needing and
therefore adding more features to it in the future, but it will never be a feature-complete
client.

## Why is this thing so limited?

I needed a simple way to create tasks programmatically. I looked at
[orijtech/asana](https://github.com/orijtech/asana), but due to the request payload for
task creation being different than the response payload, that library (unfortunately) doesn’t
support setting a project for a task, which was a hard required for my use case.

I googled around a lot and found an [answer deep in the Asana developer forums][forum]
that even confirms that this is a common area of confusion:

  [forum]: https://community.asana.com/t/how-create-task-with-membership-via-api/10481/2

> […] you’re getting caught by a very common area of confusion with our API that I’ve got
> a task to clarify in our documentation “someday hopefully soon”.

So here goes.

## Installation

At this point I honestly don’t think anybody should "just install it" and start using it.
You can, of course, but I’d prefer it if you read the code, understand it, and then copy
the parts you need instead (see LICENSE). It’s really not much.

If not:

```
go get github.com/pagespeed-io/asana-tiny-go
```

Import as usual and refer to as `asana`.

## Usage

You’ll need a [personal access token from Asana](https://asana.com/guide/help/api/api#gl-access-tokens).
Then, create a client:

```go
client, err := asana.New("your-token-here")
```

Then create a `TaskRequest` and use that to create the task:

```go
tr := TaskRequest{
  Workspace: 98765,
  Name:      "Task created with Asana Tiny Go",
  Notes:     "Put some plain text here",
  Followers: []*TaskFollower{}, // no followers. If not set, the owner of the access token will be the follower
  Memberships: []*TaskMembership{
    &TaskMembership{
      Project: 54321,
      Section: 12345,
    },
  },
}

task, err := c.CreateTask(&tr)
```

`Followers` and `Memberships` are optional.


## License

This code is given into the Public Domain using the [Unlicense](http://unlicense.org/).
See the `LICENSE.md` for details.
