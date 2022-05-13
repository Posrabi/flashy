# flashy - WIP

An app to help you learn English (or just about any languages) better!

Type in your words and sentences into these flashcards, then try to memorize what you just wrote. You will be surprised at how efficient this is.

---

## UI/UX

<img src="flow.gif" width="50%" height="50%"/>

_*I'm not a designer so somehow the cards look like a floppy disk*_ :)

**_The app flow is simple and intuitive:_**

- Sign In.
- Choose how many words you want to learn.
- Learn them.
- Go back to home to see new words appear in your history.

## Features:

- **Auto Login:**

  Log in once and never have to worry about it ever again!

- **History:**

  Keeps a record of every words you have learned.

- **Versus mode - WIP:**

  Allow players to compete head to head thanks to gRPC's amazing bidirectional streams.

- **Rankings - WIP:**

  Knowing how many words your friends have learned to date. Ability to create a mini competition between friends.

## Specs:

- A Go gRPC CRUD server will handle storing all user's info and words record.
- App built on TypeScript React Native and Recoil, used React Query for server state caching.
- Hooked up React Native to gRPC by using Java Native modules.
- Versus mode possible thanks to gRPC bidirectional streams.
- Friends/rankings using Facebook Log In API.
- Used Android's Shared Preferences/ IOS's Keychain to persist log in.

### TODOs:

- Deploy somewhere cheap (preferably locally on a Raspberry-Pi, cloud is too expensive at this level).
- Deploy to Playstore.
- Friends, rankings.
- Client versus mode (gRPC bidirectional streams in Java).
- UI for friends, leaderboards, versus.
- Keep record of all the wrong guesses.
