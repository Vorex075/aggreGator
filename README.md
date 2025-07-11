_Note. This project is part of [boot.dev](https://boot.dev) course_

# aggreGator

`aggreGator` is a simple CLI blog aggregator. You can have different users, each following as many RSS feeds as they want.

## Requirements

- Go (1.23 or greater)
- A postgres server

## Installation

Simply run in your terminal.

```Bash
go install github.com/Vorex075/aggreGator
```

The program should be located at `go/bin`. Ensure to add this directory to `$PATH` so that you can use it from anywhere.

## Configuration

You need to create a file in your home directory called `.gatorconfig.json`. This file will contain the configuration json, needed for `aggreGator` to work properly:

- `db_url` - The url connection to the postgres database.
- `current_user_name` - The current logged in user. I recommend you not to set this value at all, since the user should be registered in the database.

## Commands

An extensive list of the avaible commands, and its usage

| command     | description                                                                                                    | usage                                 |
| ----------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------- |
| `register`  | Registers a new user into `aggreGator`                                                                         | `aggreGator register (name)`          |
| `login`     | Changes the active user                                                                                        | `aggreGator login (name)`             |
| `users`     | Lists all users in the database, and marks out the active user                                                 | `aggreGator users`                    |
| `addfeed`   | Adds a new feed into `aggreGator`                                                                              | `aggreGator addfeed (feedName) (url)` |
| `feeds`     | Lists all avaible feeds, and the user that created it                                                          | `aggreGator feeds`                    |
| `follow`    | Follows the specified feed by url                                                                              | `aggreGator follow (url)`             |
| `following` | Lists all feeds that the active user is following                                                              | `aggreGator following`                |
| `unfollow`  | Unfollows a feed by url                                                                                        | `aggreGator unfollow (url)`           |
| `browse`    | Given an optional limit, displays the most recent posts for the active user following feeds. Default limit = 2 | `aggreGator browse [limit]`           |
| `agg`       | Fetches the information from `all feeds` in the database, given a time (ex. 2s)                                | `aggreGator agg (time)`               |
| `bookmark`  | Bookmarks a post for the user to search it later                                                               | `aggreGator bookmark (post_id)`       |

> [!NOTE]
> The `agg` command is the only one that behaves a bit different, since you should run on the background while using `aggreGator`.

There is a command (`reset`) that wipes the entire database. Use it by your own discretion. Usage: `aggreGator reset`
